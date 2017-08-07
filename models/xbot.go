package models

import (
	db "farmer/autocs/database"
	"fmt"
)
type (
	Xbot struct {
		Id int    `json:"id"`
		Title	string    `json:"title"`
	}
	QaList struct {
		Id int
		Title	string
	}

	Qa struct {
		Id int    `json:"id"`
		Title	string    `json:"title"`
		Content	string        `json:"content"`
		Views	int    `json:"views"`
		CreateDate	string    `json:"create_date"`
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
	err = db.GetORM().Table("xmks_qa").Where("keywords like ?","%"+ keyword +"%").Find(&x).Error
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
	err = db.GetORM().Table("xmks_qa").Find(&x).Error
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