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

func GetList()(x []QaWeb, err error) {
	x = []QaWeb{}
	err = db.GetORM().Table("xmks_qa").Order("id desc").Find(&x).Limit(10).Error
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