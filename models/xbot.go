package models

import (
	db "farmer/autocs/database"
	"fmt"
	"github.com/deepzz0/go-com/time"
	"farmer/autocs/config"
	"github.com/jinzhu/gorm"
)
type (
	Xbot struct {
		Id int    `json:"id"`
		Title	string    `json:"title"`
		ReplyType	int    `json:"reply_type"`
		ReplyImg	string    `json:"reply_img"`
		ReplyText	string        `json:"reply_text"`
		Views	int    `json:"views"`
		DeletedAt	time.Now    `json:"deleted_at"`
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
		DeletedAt	time.Now    `json:"deleted_at"`
	}

	QaWeb struct {
		Id int
		Title	string
		Keywords string
		Content	string
		ReplyType	int
		ReplyImg	string
		ReplyText	string
		Views	int
		CreateDate	string
		DeletedAt	time.Now
	}
)

func (Qa) TableName() string  {
	return "xmks_qa";
}
func (QaWeb) TableName() string  {
	return "xmks_qa";
}
func (Xbot) TableName() string  {
	return "xmks_qa";
}

func GetListByKeyword(keyword string)(x []Xbot, err error) {
	x = []Xbot{}
	err = db.GetORM().Where("keywords like ?","%"+ keyword +"%").Order("id desc").Find(&x).Limit(5).Error
	if err != nil {
		fmt.Println(err)
	}else{
		db.GetORM().Model(&x).Where("keywords like ?","%"+ keyword +"%").UpdateColumn("views", gorm.Expr("views + ?", 1))
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

	//带分页查询
	pg1 := fmcfg.Config.GetInt("app.perPageNum")
	var pg2 int
	if page < 2{
		pg2 = 0
	}else{
		pg2 = (page - 1)*pg1
	}

	total = 0
	db.GetORM().Find(&x).Count(&total)
	err = db.GetORM().Order("id desc").Limit(pg1).Offset(pg2).Find(&x).Error

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

func FaqDelDo(id int) error {
	var qa Qa
	err := db.GetORM().Where("id = ?",id).Delete(&qa).Error
	return err
}


func (* Qa)FaqUpdate(qa Qa)(Qa,error){
	err := db.GetORM().Table("xmks_qa").Where("id = ?",qa.Id).Update(&qa).Error
	return qa,err
}

func GetKeywordsViews(keyword string) int {
	var qa QaWeb
	type Result struct {
		Total int
	}
	var r Result
	db.GetORM().Model(&qa).Select("sum(views) as total").Where("keywords like ?","%"+ keyword +"%").Scan(&r)
	return r.Total
}