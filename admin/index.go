package admin

import (
	"github.com/gin-gonic/gin"
	"farmer/autocs/models"
	"fmt"
	"strconv"
	"farmer/autocs/common"
	"os"
	"bufio"
	"io"
	"strings"
	"github.com/labstack/gommon/log"
	"html/template"
	"farmer/autocs/config"
)

func QaList(c *gin.Context) {
	var msg string
	var pageNum int
	pageNum, _ = strconv.Atoi(c.Query("p"))
	list,total, err := models.GetList(pageNum)
	if err != nil {
		msg = "暂无数据"
	}
	fmt.Println(total)
	for i, v := range list{
		list[i].Content = common.TrimHtml(v.Content)
	}


	var pages models.Page = models.NewPage(pageNum, fmcfg.Config.GetInt("app.perPageNum"), total, "/admin/faq")
	page := pages.Show()
	c.HTML(200,"admin/faq.html", gin.H{"msg": msg,"list":list,"Mid":"faq", "Page":template.HTML(page)})
}
func QaAdd(c *gin.Context) {
	c.HTML(200,"admin/faq_add.html", gin.H{"Mid":"faq"})
}
func QaAddDo(c *gin.Context) {
	faq := models.Qa{}
	faq.Title = c.PostForm("title")
	faq.ReplyType, _ = strconv.Atoi(c.PostForm("reply_type"))
	faq.ReplyImg = c.PostForm("reply_img")
	faq.ReplyText = c.PostForm("reply_text")
	faq.Content = c.PostForm("content")
	faq.Keywords = c.PostForm("keywords")


	if _, err := faq.FaqAdd(faq);err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{"code":1,"msg":"新增失败"})
		return
	}

	//更新FAQ字典
	go checkDic(faq.Keywords)

	c.JSON(200, gin.H{"msg": "添加成功","code":0,"url":"/admin/faq"})
}


func QaDelAction(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	if err := models.FaqDelDo(id); err != nil{
		log.Fatal(err)
		c.JSON(200, gin.H{"code":1,"msg":"删除失败"})
	}else{
		c.JSON(200, gin.H{"msg": "删除成功","code":0,"url":"/admin/faq"})
	}
}
func QaEditAction(c *gin.Context){
	type R struct {
		Content interface{}
		Id		int
		Title	string
		Keywords string
		ReplyText string
		ReplyType int
		ReplyImg string
		Mid string
	}
	cid, _ := strconv.Atoi(c.Param("id"))
	data, _ := models.GetInfo(cid)
	r := R{Id:data.Id,Title:data.Title, ReplyType:data.ReplyType,ReplyImg:data.ReplyImg,ReplyText:data.ReplyText,Content:template.HTML(data.Content), Keywords:data.Keywords, Mid:"faq"}
	fmt.Println(r)
	c.HTML(200,"admin/qa_edit.html",r)
}

func QaUpdateAction(c *gin.Context) {
	faq := models.Qa{}
	faq.Id, _ = strconv.Atoi(c.PostForm("id"))
	if faq.Id < 1 {
		fmt.Println(111)
		c.JSON(200, gin.H{"code":1,"msg":"更新失败"})
		return
	}
	faq.Title = c.PostForm("title")
	faq.ReplyType, _ = strconv.Atoi(c.PostForm("reply_type"))
	faq.ReplyImg = c.PostForm("reply_img")
	faq.ReplyText = c.PostForm("reply_text")
	faq.Content = c.PostForm("content")
	faq.Keywords = c.PostForm("keywords")

	if _, err := faq.FaqUpdate(faq);err != nil {
		log.Fatal(err)
		c.JSON(200, gin.H{"code":1,"msg":"更新失败"})
		return
	}

	//更新FAQ字典
	//go checkDic(faq.Keywords)

	c.JSON(200, gin.H{"msg": "更新成功","code":0,"url":"/admin/faq"})
}
func checkDic(keywords string)  {
	fmt.Println(keywords)
	words := common.StrToSlice(keywords)
	fmt.Println(words)

	for _, v := range words{
		common.UpdateDic(v)
	}
}

func UpDic(c *gin.Context) {
	str := c.Query("str")
	fmt.Println(str)

	go common.UpdateDic(str)

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
		//aw.Weight, _= strconv.Atoi(thisLine[1])
		aw.Weight = models.GetKeywordsViews(aw.Name)
		aw.Cs = common.RandGetArray(cor)
		ds = append(ds, aw)
		if err != nil || io.EOF == err {
			break
		}else{
			allwords = append(allwords, line)
		}

	}
	c.HTML(200,"admin/keywords.html", gin.H{"list":ds,"Mid":"dic"})
}