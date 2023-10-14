package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Application APP入口
func Application(c *gin.Context) {
	if c.Request.Method == "GET" {
		c.HTML(http.StatusOK, "index.gohtml", gin.H{})
		return
	}

}
