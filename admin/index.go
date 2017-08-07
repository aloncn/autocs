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