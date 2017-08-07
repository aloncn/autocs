package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"net/http"
	"html/template"
)

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			if strings.Compare(value, "") != 0 {
				c.Next()
				return
			}
		}
		c.HTML(http.StatusUnauthorized, "public/error.html", gin.H{
			"error": "请先登录！",
			"msg":   template.HTML("你尚未登录或者登录已过期，请 <a href='/login'><button class=\"btn btn-primary btn-labeled ion-log-in\">重新登录</button></a>"),
			"url":   "/login",
		})
		c.Abort()
		return
	}
}