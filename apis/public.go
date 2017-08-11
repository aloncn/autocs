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
	"time"
	"path"
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
	fmt.Println(file)
	filename := header.Filename
	fmt.Println(filename)

	var filenameWithSuffix string
	filenameWithSuffix = path.Base(filename) //获取文件名带后缀
	fmt.Println("filenameWithSuffix =", filenameWithSuffix)
	var fileSuffix string
	fileSuffix = path.Ext(filenameWithSuffix) //获取文件后缀
	fmt.Println("fileSuffix =", fileSuffix)

	//var filenameOnly string
	//filenameOnly = strings.TrimSuffix(filenameWithSuffix, fileSuffix)//获取文件名
	//fmt.Println("filenameOnly =", filenameOnly)

	//存储文件的路径
	fp1 := fmcfg.Config.GetString("app.uploadPath") + "/" + time.Now().Format("2006-01-02")

	//绝对路径
	fp := fmcfg.Config.GetString("app.basePath") + fp1

	if !common.IsExist(fp) {
		if os.Mkdir(fp,0755) != nil {
			c.JSON(http.StatusOK,gin.H{"errno":1, "msg":"创建目录失败"})
			return
		}
	}
	//重命名文件
	newName := strconv.FormatInt(time.Now().Unix(),10) + fileSuffix
	// 创建临时接收文件
	out, err := os.Create(fp + "/" + newName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	// Copy数据
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	pics := []string{fp1 + "/" + newName}
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