package apis

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	. "farmer/autocs/models"
	"strconv"
	"github.com/mssola/user_agent"
)

type Reback struct {
	Status int    `json:"status"`
	Msg string `json:"msg"`
	Data interface{}	`json:"data"`
}
func IndexApi(c *gin.Context) {
	ua := user_agent.New(c.Request.UserAgent())		//获取用户UA
	var tpl string
	if ua.Mobile() {
		tpl = "wap/index.html"
	}else{
		tpl = "web/index.html"
	}
	c.HTML(http.StatusOK, tpl,gin.H{"msg":"Hello World!","title":"熊猫快收自助客服系统"})
}

func ChatApi(c *gin.Context)  {
	c.HTML(http.StatusOK, "web/chat.html",nil)
}

func AddPersonApi(c *gin.Context) {
	firstName := c.Request.FormValue("first_name")
	lastName := c.Request.FormValue("last_name")

	p := Person{FirstName: firstName, LastName: lastName}

	ra, err := p.AddPerson()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func GetPersonsApi(c *gin.Context) {
	var R Reback
	p := Person{}

	ra, err := p.GetPersons()
	if err != nil {
		log.Fatalln(err)
		R.Msg = "暂无数据"
	}else{
		R.Data = ra
	}
	fmt.Println(ra)
	//c.String(http.StatusOK, "It works")
	c.JSON(http.StatusOK, R)
}

func GetPersonApi(c *gin.Context) {
	var R Reback
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	r := GetUserInfoById(id)
	/*if err != nil {
		log.Fatalln(r)
		R.Msg = "用户不存在"
	}else{
		R.Data = r
	}*/
	R.Data = r
	c.JSON(http.StatusOK, R)
}

func DelPersonApi(c *gin.Context)  {
	cid := c.Param("id")
	id, _ := strconv.Atoi(cid)
	p := Person{Id:id}
	pid, _ := p.DelPerson()
	msg := fmt.Sprintf("Query successful %d", pid)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
		"data":pid,
	})
}