package admin

import (
	"github.com/gin-gonic/gin"
	qa "farmer/autocs/models"
	"fmt"
	"strconv"
	"farmer/autocs/common"
	"html/template"
	"os"
	"bufio"
	"io"
	"strings"
	"github.com/labstack/gommon/log"
)

func QaList(c *gin.Context) {
	var msg string
	list, err := qa.GetList()
	if err != nil {
		msg = "暂无数据"
	}
	fmt.Println(list)
	for i, v := range list{
		list[i].Content = common.TrimHtml(v.Content)
	}
	c.HTML(200,"admin/faq.html", gin.H{"msg": msg,"list":list,"mid":"faq"})
}

func QaAdd(c *gin.Context) {
	c.HTML(200,"admin/faq_add.html", gin.H{"mid":"faq"})
}
func QaAddDo(c *gin.Context) {
	faq := qa.Qa{}
	faq.Title = c.PostForm("title")
	faq.ReplyType, _ = strconv.Atoi(c.PostForm("reply_type"))
	faq.ReplyImg = c.PostForm("reply_img")
	faq.ReplyText = c.PostForm("reply_text")
	faq.Content = c.PostForm("content")
	faq.Keywords = c.PostForm("keywords")


	if _, err := faq.FaqAdd(faq);err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{"code":1,"msg":"新增失败"})
		c.Abort()
	}


	c.JSON(200, gin.H{"msg": "添加成功","code":0,"url":"/admin/faq"})
}

func QaInfo (c *gin.Context) {
	type r struct {
		Content interface{}
		Title	string
	}
	rr := r{}
	cid, _ := strconv.Atoi(c.Param("id"))
	data, _ := qa.GetInfo(cid)
	rr.Title = data.Title
	rr.Content = template.HTML(data.Content)
	c.HTML(200,"qa/info.html", rr)
}

func add(left int, right int) int{
	return left + right
}
func GetAllWords(c *gin.Context)  {
	var Dics []string
	type D struct {
		Name string
		Weight int
		Cs string
	}
	var ds []D
	f, err := os.Open("./data/dictionary.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	allwords := Dics
	aw := D{}

	cor := []string{"purple", "info", "pink", "success"}

	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		thisLine := strings.Fields(line)

		aw.Name = thisLine[0]
		aw.Weight, _= strconv.Atoi(thisLine[1])
		aw.Cs = common.RandGetArray(cor)
		ds = append(ds, aw)
		if err != nil || io.EOF == err {
			break
		}else{
			allwords = append(allwords, line)
		}

	}
	c.HTML(200,"admin/keywords.html", gin.H{"list":ds,"mid":"dic"})
}