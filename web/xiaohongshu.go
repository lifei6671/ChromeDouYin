package web

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"github.com/lifei6671/chrome-douyin/chrome"
	"github.com/lifei6671/chrome-douyin/config"
	"github.com/patrickmn/go-cache"
	"io"
	"net/http"
	"strings"
)

// XiaoHongShu 小红书解析
func XiaoHongShu(c *gin.Context) {
	//如果开启了token校验，需要判断是否传递了token
	if config.Default.Authorization.Enable || c.Request.RequestURI != "/" {
		username, password, ok := c.Request.BasicAuth()
		if !ok || config.Default.Authorization.Username != username || config.Default.Authorization.Password != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
	v := c.Query("url")
	if len(v) == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "param error"})
		return
	}
	h := md5.New()
	_, _ = io.WriteString(h, v)
	key := hex.EncodeToString(h.Sum(nil))
	if m, ok := _cache.Get(key); ok {
		c.JSON(http.StatusOK, chrome.FromXiaoHongShu(m.(*chrome.XiaoHongShuDetail)))
		return
	}
	m, raw, err := _chrome.XiaoHongShuDetail(c.Request.Context(), v)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	_cache.Set(v, m, cache.DefaultExpiration)
	_cache.Set(v+"_raw", raw, cache.DefaultExpiration)
	c.JSON(http.StatusOK, chrome.FromXiaoHongShu(m))
}

// XiaoHongShuShortcuts 小红书捷径
func XiaoHongShuShortcuts(c *gin.Context) {
	v := c.PostForm("url")
	if len(v) == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "param error"})
		return
	}
	h := md5.New()
	_, _ = io.WriteString(h, v)
	key := hex.EncodeToString(h.Sum(nil))
	if m, ok := _cache.Get(key); ok {
		c.JSON(http.StatusOK, chrome.FromXiaoHongShu(m.(*chrome.XiaoHongShuDetail)))
		return
	}
	m, raw, err := _chrome.XiaoHongShuDetail(c.Request.Context(), v)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	_cache.Set(v, m, cache.DefaultExpiration)
	_cache.Set(v+"_raw", raw, cache.DefaultExpiration)
	c.JSON(http.StatusOK, chrome.FromXiaoHongShu(m))
}

type XiaoHongShuAuthorization struct {
	Version     string   `json:"version"`
	Api         string   `json:"api"`
	Token       string   `json:"token"`
	SuccessMenu []string `json:"success_Menu"`
	FailMenu    []string `json:"fail_Menu"`
}

// XiaoHongShuAuthorizationShortcuts 小红书快捷指令请求地址
func XiaoHongShuAuthorizationShortcuts(c *gin.Context) {
	host := config.Default.SiteURL
	if len(host) == 0 {
		proto := c.Request.Header.Get("X-Forwarded-Proto")
		if proto == "https" {
			host = strings.TrimSuffix("https://"+c.Request.Host, ":443")
		} else {
			host = strings.TrimSuffix("http://"+c.Request.Host, ":80")
		}
	}
	ret := &XiaoHongShuAuthorization{
		Version: "1",
		Api:     strings.TrimSuffix(host, "/") + "/api/xhs/shortcuts",
		Token:   "",
		SuccessMenu: []string{"返回小红书",
			"打开相册",
		}, FailMenu: []string{"返回小红书"},
	}

	c.JSON(http.StatusOK, ret)
}
