package database
import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"farmer/autocs/config"
	"fmt"
)

//var MyDb *sql.DB
var Gorm  map[string]*gorm.DB
func init() {
	Gorm = make(map[string]*gorm.DB)
}

// 初始化Gorm
func NewDB(dbname string) {
	var orm *gorm.DB
	var err error

	dbHost := fmcfg.Config.GetString(dbname + ".dbHost")
	dbName := fmcfg.Config.GetString(dbname + ".dbName")
	dbUser := fmcfg.Config.GetString(dbname + ".dbUser")
	dbPasswd := fmcfg.Config.GetString(dbname + ".dbPasswd")
	dbPort := fmcfg.Config.GetString(dbname + ".dbPort")
	dbType := fmcfg.Config.GetString(dbname + ".dbType")
	dbOpen := fmcfg.Config.GetInt(dbname + ".dbOpen")
	dbIdle := fmcfg.Config.GetInt(dbname + ".dbIdle")
	orm, err = gorm.Open(dbType, dbUser + ":" + dbPasswd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local")
	//开启sql调试模式
	orm.LogMode(true)

	if err != nil {
		fmt.Println(err)
		panic("数据库连接异常")
	}
	//defer Gorm[dbname].Close()
	//最大打开连接数
	orm.DB().SetMaxOpenConns(dbOpen)
	//连接池的空闲数大小
	orm.DB().SetMaxIdleConns(dbIdle)
	Gorm[dbname] = orm
	//defer Gorm[dbname].Close()
}

// 通过名称获取Gorm实例
func GetORMByName(dbname string) *gorm.DB {

	return Gorm[dbname]
}

// 获取默认的Gorm实例
func GetORM() *gorm.DB {
	return Gorm["dbDefault"]
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
