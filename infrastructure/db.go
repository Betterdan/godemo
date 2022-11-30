package infrastructure

import (
	appConfig "demo/configs"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type (
	MyDb *gorm.DB
)

func LoadMysql()  {
	MyDb := initMysql("mysql")
	defer MyDb.Close()
}

func initMysql(configName string) *gorm.DB {
	curConfig := appConfig.VTool
	dsn :=fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		curConfig.Get(configName+".username"),
		curConfig.Get(configName+".password"),
		curConfig.Get(configName+".host"),
		curConfig.GetUint32(configName+".port"),
		curConfig.Get(configName+".dbname"),
		"10s",
	)
	db,err :=gorm.Open("mysql",dsn)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("mysql connect success")
	}
	return db
}