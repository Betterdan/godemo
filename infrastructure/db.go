package infrastructure

import (

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type (
	MyDb *gorm.DB
)

func loadMysql(mysqlConfig string)  {

}