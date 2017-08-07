package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
	qa "farmer/autocs/models"
	"strconv"
	"html/template"
)

func GetQaInfo(c *gin.Context)  {
	var msg string
	nid := c.Param("id")
	id, _ := strconv.Atoi(nid)
	r, err := qa.GetInfoById(id)
	if err != nil {
		msg = "记录不存在"
	}else{
		msg = r.Title
	}
	con := template.HTML(r.Content)
	c.HTML(http.StatusOK,"qa/info.html", gin.H{"msg": msg,"data":r,"con":con})

}