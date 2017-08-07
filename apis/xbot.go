package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/huichen/sego"
	"fmt"
	qa "farmer/autocs/models"
	"farmer/autocs/common"
	"strconv"
	"github.com/gorilla/websocket"
	"encoding/json"
	"time"
	"farmer/autocs/config"
)


func CutWordsApi(c *gin.Context)  {
	text := c.PostForm("question")	//接受请求参数

	// 载入词典
	var segmenter sego.Segmenter
	segmenter.LoadDictionary("./data/dictionary.txt")

	// 分词
	//text := "中国云计算很好很强大，今天是个好日子，哇哈哈，有木有优惠？优惠怎么加盟？你们在哪里，201701年优惠政策"
	//fmt.Printf("输入：%s \n",text)
	segments := segmenter.Segment([]byte(text))
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	//fmt.Println(sego.SegmentsToString(segments, false))

	keywords := sego.SegmentsToString(segments, false)

	fmt.Println(keywords)
	/**
	根据匹配关键词查询是否有问答记录
	放在 分词匹配[sego.SegmentsToString]中更高效，这里为了不去对依赖进行修改
	 */
	all := []qa.Xbot{}
	var ids []int
	for _, v := range keywords{
		list, _ := qa.GetListByKeyword(v)
		if len(list) > 0 {
			for _, vv := range list{
				//检查是否已经存在
				if !common.CheckRepeat(ids, vv.Id) {
					all = append(all, vv)
					ids = append(ids, vv.Id)
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{"keywords" : keywords, "list":all})
}



func GetQaInfoApi(c *gin.Context)  {
	var msg string
	var data interface{}

	nid := c.Param("id")
	id, _ := strconv.Atoi(nid)
	r, err := qa.GetInfoById(id)
	if err != nil {
		msg = "记录不存在"
	}

	data = r
	c.JSON(http.StatusOK, gin.H{"msg": msg,"data":data})

}


func GetCutWordsByKey(c string) []qa.Xbot {

	var segmenter sego.Segmenter

	segmenter.LoadDictionary("./data/dictionary.txt")
	segments := segmenter.Segment([]byte(c))
	keywords := sego.SegmentsToString(segments, false)
	all := []qa.Xbot{}
	var ids []int
	for _, v := range keywords{
		list, _ := qa.GetListByKeyword(v)
		if len(list) > 0 {
			for _, vv := range list{
				//检查是否已经存在
				if !common.CheckRepeat(ids, vv.Id) {
					all = append(all, vv)
					ids = append(ids, vv.Id)
				}
			}
		}
	}
	return all
}


var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ChatDemoApi(w http.ResponseWriter, r *http.Request){
	type Mine struct{
		Avatar	string    `json:"avatar"`
		Content string    `json:"content"`
		Id string    `json:"id"`
		Mine bool    `json:"mine"`
		Username string    `json:"username"`
	}
	type To struct {
		Avatar	string    `json:"avatar"`
		Id	string    `json:"id"`
		Name	string    `json:"name"`
		Sign	string    `json:"sign"`
		Type	string    `json:"type"`
	}
	type R struct {
		Type string    `json:"type"`
		Data	struct {
			Mine    `json:"mine"`
			To		`json:"to"`
		}	`json:"data"`
	}
	type Re struct {
		Username	string    	  `json:"username"`	//消息来源用户名
		Avatar		string        `json:"avatar"`	//消息来源用户头像
		Id			string        `json:"id"`	//消息的来源ID（如果是私聊，则是用户id，如果是群聊，则是群组id）
		Type 		string        `json:"type"`	//聊天窗口来源类型，从发送消息传递的to里面获取
		Content 	string        `json:"content"`	//消息内容
		Mine		bool          `json:"mine"`	//是否我发送的消息，如果为true，则会显示在右方
		Fromid		string        `json:"fromid"`	//消息的发送者id（比如群组中的某个消息发送者）
		Timestamp 	int64         `json:"timestamp"`	//服务端动态时间戳
	}


	js := R{}
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}
	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		text := string(msg)
		json.Unmarshal([]byte(text),&js)

		//fmt.Println(js.Data.Content)

		cut := GetCutWordsByKey(js.Data.Content)
		fmt.Println(cut)
		result := ""
		url := fmcfg.Config.GetString("app.url") + ":" + fmcfg.Config.GetString("app.port") + "/faq/"
		if len(cut) > 0 {
			for _,v := range cut {
				result += "👉 a(http://" + url + strconv.Itoa(v.Id) + ")["+ v.Title +"] \n\r"
			}
		}else {
			result = "聊点什么吧，仅限 '云计算'、'优惠政策'相关，不然我回答不上来😢";
		}
		ret := Re{Username:js.Data.To.Name,Avatar:js.Data.To.Avatar,Id:js.Data.To.Id,Type:js.Data.To.Type,Content:result,Mine:false,Fromid:js.Data.Mine.Id,Timestamp:time.Now().Unix()}

		rr, _ := json.Marshal(ret)




		conn.WriteMessage(t, rr)
	}
}
func WssApi(w http.ResponseWriter, r *http.Request) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		text := string(msg)

		var segmenter sego.Segmenter
		segmenter.LoadDictionary("./data/dictionary.txt")
		segments := segmenter.Segment([]byte(text))
		keywords := sego.SegmentsToString(segments, false)

		fmt.Println(keywords)
		all := []qa.Xbot{}
		var ids []int
		for _, v := range keywords{
			list, _ := qa.GetListByKeyword(v)
			if len(list) > 0 {
				for _, vv := range list{
					//检查是否已经存在
					if !common.CheckRepeat(ids, vv.Id) {
						all = append(all, vv)
						ids = append(ids, vv.Id)
					}
				}
			}
		}
		json_str,_ := json.Marshal(all)
		msg = []byte(json_str)

		conn.WriteMessage(t, msg)
	}
}