package models

import (
	//"database/sql"
	db "farmer/autocs/database"
	"fmt"
)

type Person struct {
	Id        int    `json:"id" form:"id"`
	Username  string    `json:"username" from:"username"`
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
}

func (Person) TableName() string {
	return "test_user"
}
func (p *Person) AddPerson() (id int64, err error) {
	/*rs, err := db.MyDb.Exec("INSERT INTO test_user(first_name, last_name) VALUES (?, ?)", p.FirstName, p.LastName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()*/
	return
}

func (p *Person) GetPersons() (persons []Person, err error) {
	persons = make([]Person, 0)
	/*rows, err := db.MyDb.Query("SELECT id, username, first_name, last_name FROM test_user limit 10")
	defer rows.Close()

	if err != nil {
		return
	}

	for rows.Next() {
		var person Person
		rows.Scan(&person.Id, &person.Username, &person.FirstName, &person.LastName)
		persons = append(persons, person)
	}
	if err = rows.Err(); err != nil {
		return
	}*/
	return
}

func GetPerson(cid int) (p Person, err error) {
	//var person PersonQ
	/*err = db.MyDb.QueryRow("SELECT id,username,first_name,last_name FROM test_user WHERE id=?", cid).Scan(&person.Id, &person.Username, &person.FirstName, &person.LastName)
	if err != nil {
		fmt.Println(err)
		return
	}
	return person,nil
	//fmt.Println(err)

	id := cid
	var person Person
	err := db.MyDb.QueryRow("SELECT id, username, first_name, last_name FROM test_user WHERE id=?", id).Scan(&person)

	//err = db.MyDb.QueryRow("SELECT id, username, first_name, last_name FROM test_user WHERE id = ?", "1").Scan(&p)
	if err != nil {
		return err
	}*/
	return
}

func (p *Person) DelPerson()(id int64,err error)  {
	/*rs, err := db.MyDb.Exec("DELETE FROM test_user WHERE id=?", p.Id)
	if err != nil {
		log.Fatalln(err)
	}
	id, err = rs.RowsAffected()
	if err != nil {
		log.Fatalln(err)
	}*/
	return
}

func GetUserInfoById(cid int)(person Person)  {
	var u Person

	db.GetORM().Where("id = ?",cid).First(&u)
	fmt.Println(u)
	return u
}