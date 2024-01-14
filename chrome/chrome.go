package chrome

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"

	"github.com/lifei6671/chrome-douyin/config"
)

var pattern = regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
var (
	ErrParseURLFail = errors.New("parse douyin url fail")
)

type Chrome struct {
	opts    *options
	browser *rod.Browser
	cookies map[string][]*proto.NetworkCookieParam
}

func New(opts ...Option) *Chrome {
	opt := &options{
		Timeout: time.Minute,
	}
	for _, o := range opts {
		o(opt)
	}

	c := &Chrome{
		opts:    opt,
		cookies: make(map[string][]*proto.NetworkCookieParam),
	}
	r, err := c.initBrowser()
	if err != nil {
		panic(err)
	}
	c.browser = r
	if len(opt.Cookies) > 0 {
		for _, cookie := range opt.Cookies {
			if len(cookie.Cookie) == 0 {
				continue
			}
			var cookies []*proto.NetworkCookie
			header := http.Header{}
			header.Add("Cookie", cookie.Cookie)
			request := http.Request{Header: header}

			for _, rCookie := range request.Cookies() {
				cookies = append(cookies, &proto.NetworkCookie{
					Name:    rCookie.Name,
					Value:   rCookie.Value,
					Domain:  cookie.Domain,
					Path:    "/",
					Expires: proto.TimeSinceEpoch(time.Now().Add(7 * 24 * time.Hour).Unix()),
					Secure:  true,
				})
			}

			c.cookies[cookie.Name] = proto.CookiesToParams(cookies)
		}
	}
	return c
}

func (c *Chrome) initBrowser() (*rod.Browser, error) {
	var r *rod.Browser
	if strings.HasPrefix(c.opts.ServiceURL, "ws://") {
		r = rod.New().ControlURL(c.opts.ServiceURL)
	} else {
		l := launcher.MustNewManaged(c.opts.ServiceURL)
		l.Set("disable-gpu").Delete("disable-gpu")
		l.Headless(false).XVFB("--server-num=5", "--server-args=-screen 0 1600x900x16")
		l.Set("no-sandbox").
			Set("disable-setuid-sandbox").
			Set("disable-gpu").
			Set("disable-dev-shm-usage"). // 在这里已经设置为 disable 了
			Set("unlimited-storage").
			Set("disable-accelerated-2d-canvas").
			Set("full-memory-crash-report")

		client, err := l.Client()

		if err != nil {
			return nil, err
		}
		r = rod.New().Client(client)
	}
	if err := r.Connect(); err != nil {
		return nil, err
	}
	return r, nil
}

func (c *Chrome) DouYinDetail(ctx context.Context, b string) (*DouYinDetail, string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*20)
	defer cancel()
	url := pattern.FindString(b)
	if len(url) == 0 {
		return nil, "", ErrParseURLFail
	}
	if !strings.Contains(url, "v.douyin.com") {
		return nil, "", ErrParseURLFail
	}
	log.Println("douyin-url:", url)
	if cookie, ok := c.cookies["douyin"]; ok {
		err := c.browser.SetCookies(cookie)
		if err != nil {
			log.Println(err)
			return nil, "", err
		}
	}
	page, err := stealth.Page(c.browser)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "use of closed network connection") {
			c.browser, err = c.initBrowser()
			if err != nil {
				log.Println(err)
			} else {
				page, err = stealth.Page(c.browser)
			}
		}
		if err != nil {
			return nil, "", err
		}
	}
	page = page.Context(ctx)

	defer page.MustClose()

	var videoType MediaType
	var awemeId string
	ch := make(chan string)

	router := page.HijackRequests()
	defer router.MustStop()
	_ = router.MustAdd("", func(ctx *rod.Hijack) {
		//图片的链接
		if strings.HasPrefix(ctx.Request.URL().String(), "https://www.douyin.com/note/") {
			awemeId = strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(ctx.Request.URL().Path, "/note/"), "/"), "/")
			log.Println(b, awemeId)
		}
		//视频的链接
		if strings.HasPrefix(ctx.Request.URL().String(), "https://www.douyin.com/video/") {
			awemeId = strings.TrimSuffix(strings.TrimPrefix(strings.TrimPrefix(ctx.Request.URL().Path, "/video/"), "/"), "/")
			log.Println(b, awemeId)
		}

		if strings.Contains(ctx.Request.URL().Host, "mcs.zijieapi.com") {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		//拦截字体
		if ctx.Request.Type() == proto.NetworkResourceTypeFont {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		//拦截图片
		if ctx.Request.Type() == proto.NetworkResourceTypeImage {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		//拦截视频
		if ctx.Request.Type() == proto.NetworkResourceTypeMedia {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		//拦截所有WS请求
		//if ctx.Request.Type() == proto.NetworkResourceTypeWebSocket {
		//	ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
		//	return
		//}
		if videoType == 0 {
			if strings.HasPrefix(ctx.Request.URL().Path, config.Default.DouYin.Aweme.Video) {
				videoType = MediaTypeVideo
			}
			if strings.HasPrefix(ctx.Request.URL().Path, config.Default.DouYin.Aweme.Image) {
				videoType = MediaTypeImage
			}
			if videoType > 0 {
				if bodyErr := ctx.LoadResponse(http.DefaultClient, true); bodyErr != nil {
					log.Printf("[ERROR] hijack fail: errmsg[%v] url[%s]", bodyErr, ctx.Request.URL().String())
				}
				log.Printf("video_type[%s] status_code[%d] video_url[%s]", videoType, ctx.Response.Payload().ResponseCode, ctx.Request.URL().String())
				ch <- ctx.Response.Body()
				close(ch)
				return
			}
		}
		log.Println(awemeId, ctx.Request.Type(), ctx.Request.Method(), ctx.Request.URL().String())
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})

	go router.Run()

	err = page.Navigate(url)
	if err != nil {
		log.Println(err)
		return nil, "", err
	}
	var rawData string
	select {
	case rawData = <-ch:
		break
	case <-ctx.Done():
		return nil, "", ctx.Err()
	}
	var data DouYinDetail

	if videoType == MediaTypeVideo {
		err = json.Unmarshal([]byte(rawData), &data)
		if err != nil {
			log.Printf("[ERROR] json unmarshal fail: errmsg[%v] body[%s]", err, rawData)
			return nil, "", err
		}
	} else {
		var images DouYinImages

		err = json.Unmarshal([]byte(rawData), &images)
		if err != nil {
			log.Printf("[ERROR] json unmarshal fail: errmsg[%v] body[%s]", err, rawData)
			return nil, "", err
		}
		for _, aweme := range images.AwemeList {
			if aweme.AwemeId == awemeId {
				data.AwemeDetail = aweme
				break
			}
		}
		data.StatusCode = images.StatusCode
	}
	data.ShareURL = url

	return &data, rawData, nil
}

func (c *Chrome) XiaoHongShuDetail(ctx context.Context, b string) (*XiaoHongShuDetail, string, error) {
	url := pattern.FindString(b)
	if len(url) == 0 {
		return nil, "", ErrParseURLFail
	}
	if !strings.Contains(url, "xhslink.com") {
		return nil, "", ErrParseURLFail
	}
	if cookie, ok := c.cookies["xiaohongshu"]; ok {
		err := c.browser.SetCookies(cookie)
		if err != nil {
			log.Println(err)
			return nil, "", err
		}
	}
	page, err := stealth.Page(c.browser)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "use of closed network connection") {
			c.browser, err = c.initBrowser()
			if err != nil {
				log.Println(err)
			} else {
				page, err = stealth.Page(c.browser)
			}
		}
		if err != nil {
			return nil, "", err
		}
	}
	page = page.Context(ctx)

	defer page.MustClose()

	router := page.HijackRequests()
	defer router.MustStop()

	//匹配所有XHR请求
	_ = router.Add("", proto.NetworkResourceTypeXHR, func(ctx *rod.Hijack) {
		ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
	})
	//拦截所有fetch类型请求
	_ = router.Add("*", proto.NetworkResourceTypeFetch, func(ctx *rod.Hijack) {
		if ctx.Request.Type() == proto.NetworkResourceTypeFetch {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})
	//拦截所有字体请求
	_ = router.Add("*", proto.NetworkResourceTypeFont, func(ctx *rod.Hijack) {
		if ctx.Request.Type() == proto.NetworkResourceTypeFont {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})
	//拦截所有图片的请求
	_ = router.Add("*", proto.NetworkResourceTypeImage, func(ctx *rod.Hijack) {
		if ctx.Request.Type() == proto.NetworkResourceTypeImage {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})
	//拦截所有样式表的请求
	_ = router.Add("*", proto.NetworkResourceTypeStylesheet, func(ctx *rod.Hijack) {
		if ctx.Request.Type() == proto.NetworkResourceTypeStylesheet {
			ctx.Response.Fail(proto.NetworkErrorReasonBlockedByClient)
			return
		}
		ctx.ContinueRequest(&proto.FetchContinueRequest{})
	})

	go router.Run()

	err = page.Navigate(url)
	if err != nil {
		log.Printf("小红书解析失败: url[%s] errmsg[%s]", url, err)
		return nil, "", err
	}
	// 监听 Document 请求
	page.MustWaitLoad()
	eval, err := page.Eval(`() => window.__INITIAL_STATE__`)
	if err != nil {
		log.Printf("小红书解析失败: url[%s] errmsg[%s]", url, err)
		return nil, "", err
	}
	firstNoteId := eval.Value.Get("note").Get("firstNoteId").Get("_value").Val()

	val, err := eval.Value.Get("note").Get("noteDetailMap").Get(firstNoteId.(string)).Get("note").MarshalJSON()
	if err != nil {
		log.Printf("小红书解析失败: url[%s] errmsg[%s]", url, err)
		return nil, "", err
	}
	var data XiaoHongShuDetail

	err = json.Unmarshal(val, &data)
	if err != nil {
		log.Printf("小红书解析失败: url[%s] errmsg[%s]", url, err)
		return nil, "", err
	}

	return &data, string(val), nil
}
