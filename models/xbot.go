package models

import (
	db "farmer/autocs/database"
	"fmt"
	"github.com/deepzz0/go-com/time"
)
type (
	Xbot struct {
		Id int    `json:"id"`
		Title	string    `json:"title"`
		ReplyType	int    `json:"reply_type"`
		ReplyImg	string    `json:"reply_img"`
		ReplyText	string        `json:"reply_text"`
	}
	QaList struct {
		Id int
		Title	string
	}

	Qa struct {
		Id int    			   `json:"id"`
		Title		string     `json:"title"`
		ReplyType	int        `json:"reply_type"`
		ReplyImg	string     `json:"reply_img"`
		ReplyText	string    `json:"reply_text"`
		Content		string     `json:"content"`
		Keywords 	string     `json:"keywords"`
		Views		int    	   `json:"views"`
		CreateDate	time.Now   `json:"create_date"`
	}

	QaWeb struct {
		Id int
		Title	string
		Content	string
		ReplyType	int
		ReplyImg	string
		ReplyText	string
		Views	int
		CreateDate	string
	}
)

func (Qa) TableName() string  {
	return "xmks_qa";
}

func GetListByKeyword(keyword string)(x []Xbot, err error) {
	x = []Xbot{}
	err = db.GetORM().Table("xmks_qa").Where("keywords like ?","%"+ keyword +"%").Order("id desc").Find(&x).Limit(5).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}


func GetInfoById(id int)(x Qa, err error) {
	x = Qa{}
	err = db.GetORM().Where("id = ?",id).Last(&x).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetList(page int)(x []QaWeb,total int, err error) {
	x = []QaWeb{}
	pg1 := 10
	var pg2 int
	if page < 2{
		pg2 = 0
	}else{
		pg2 = (page - 1)*pg1
	}

	total = 0
	db.GetORM().Table("xmks_qa").Find(&x).Count(&total)
	err = db.GetORM().Table("xmks_qa").Order("id desc").Limit(pg1).Offset(pg2).Find(&x).Error

	if err != nil {
		fmt.Println(err)
	}
	return
}

func GetInfo(id int)(x QaWeb, err error) {
	x = QaWeb{}
	err = db.GetORM().Table("xmks_qa").Where("id = ?",id).Last(&x).Error
	if err != nil {
		fmt.Println(err)
	}
	return
}

func (* Qa)FaqAdd(qa Qa)(Qa,error){
	err := db.GetORM().Create(&qa).Error
	return qa,err
}