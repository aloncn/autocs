package apis

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"strings"
	"net/http"
	"farmer/autocs/common"
	"os"
	"io"
	"log"
	"farmer/autocs/config"
	"html/template"
	"strconv"
	"farmer/autocs/models"
)

func PublicLoginApi(c *gin.Context)  {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Printf("Username: %s",username)

	if strings.Compare(username, "admin") != 0{
		c.JSON(http.StatusOK, gin.H{"msg": "用户名错误","code":1})
	}else if strings.Compare(password, "admin999") != 0{
		c.JSON(http.StatusOK, gin.H{"msg": "密码错误","code":2})
	} else {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    common.GetMd5String(username),
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		//这里可以直接跳转到登录后的页面，也可以返回json，由前端去处理
		c.JSON(http.StatusOK, gin.H{"msg": "登陆成功","code":0,"url":"/admin"})
		//c.Redirect(http.StatusFound, "/admin")
	}
}

//上传图片
func UploadApi(c *gin.Context)  {
	//fmt.Println(c)
	file, header , err := c.Request.FormFile("cimage")
	filename := header.Filename
	// 创建临时接收文件
	out, err := os.Create(fmcfg.Config.GetString("basePath") + "upload/" +filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// Copy数据
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	pics := []string{"/upload/" + filename}
	c.JSON(http.StatusOK,gin.H{"errno":0, "data":pics})

}

func QaInfo (c *gin.Context) {
	type R struct {
		Content interface{}
		Title	string
		ReplyText string
		ReplyType int
		ReplyImg string
	}
	cid, _ := strconv.Atoi(c.Param("id"))
	data, _ := models.GetInfo(cid)
	r := R{Title:data.Title, ReplyType:data.ReplyType,ReplyImg:data.ReplyImg,ReplyText:data.ReplyText,Content:template.HTML(data.Content)}
	//rr.Content = template.HTML(data.Content)
	c.HTML(200,"qa/info.html", r)
}