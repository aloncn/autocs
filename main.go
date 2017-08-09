package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
	"github.com/fvbock/endless"
	db "farmer/autocs/database"
	rt "farmer/autocs/router"
	"farmer/autocs/config"
)

func main()  {
	//初始化配置文件
	fmcfg.NewConfig("conf","sys")
	//fmcfg.NewConfig("/Users/farmer/igo/src/farmer/autocs/conf","sys")
	db.NewDB("dbDefault")	//默认数据库
	//NewDB("db_r")		//读库
	//NewDB("db_w")		//写库


	bmx := rt.InitRouter()	//初始化路由

	err := endless.ListenAndServe(":" + fmcfg.Config.GetString("app.port"),bmx)
	if err != nil{
		log.Println(err)
	}
	log.Println("Server stopped")
	os.Exit(0)
}
