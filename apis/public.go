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
	upload_path := "/Users/farmer/igo/src/farmer/autocs/upload"
	out, err := os.Create(upload_path + "/" +filename)
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
