package web

import (
	"embed"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"

	"github.com/lifei6671/chrome-douyin/chrome"
	"github.com/lifei6671/chrome-douyin/config"
)

var _chrome *chrome.Chrome

//go:embed static
var FS embed.FS

var _cache = cache.New(24*time.Hour, 10*time.Minute)

func Run(addr string, ch func() *chrome.Chrome) error {
	_chrome = ch()
	r := gin.Default()

	templ := template.Must(template.New("").ParseFS(FS, "static/templates/*.gohtml"))
	r.SetHTMLTemplate(templ)

	fe, _ := fs.Sub(FS, "static")
	r.StaticFS("/static", http.FS(fe))

	r.GET("/api", DouYin)
	r.GET("/", Application)
	r.POST("/", Application)

	return r.Run(addr)
}

// DouYin API 接口
func DouYin(c *gin.Context) {
	//如果开启了token校验，需要判断是否传递了token
	if config.Default.Authorization.Enable {
		username, password, ok := c.Request.BasicAuth()
		if !ok || config.Default.Authorization.Username != username || config.Default.Authorization.Password != password {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
	v := c.Query("url")
	t := c.Query("type")
	log.Println("douyin share url:", v)
	if len(v) == 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "param error"})
		return
	}
	key := v
	if t == "raw" {
		key = v + "_raw"
	}
	if m, ok := _cache.Get(key); ok {
		if t == "raw" {
			c.String(http.StatusOK, m.(string))
		} else {
			c.JSON(http.StatusOK, chrome.FromDouYinDetail(m.(*chrome.DouYinDetail)))
		}

		return
	}
	m, raw, err := _chrome.DouYinDetail(c.Request.Context(), v)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"msg": err.Error()})
		return
	}
	if m.StatusCode == 0 {
		_cache.Set(v, m, cache.DefaultExpiration)
		_cache.Set(v+"_raw", raw, cache.DefaultExpiration)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"msg": "video not found"})
		return
	}
	if t == "raw" {
		c.String(http.StatusOK, raw)
		return
	}
	c.JSON(http.StatusOK, chrome.FromDouYinDetail(m))
}
