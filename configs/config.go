package configs

import (
	enum2 "demo/common/enum"
	"github.com/spf13/viper"
	"log"
)

var (
	VTool *viper.Viper
)

func Init(path string)  {
	VTool = viper.New()
	VTool.AddConfigPath(path)
	VTool.SetConfigName(enum2.CONFIG_FILE_NAME)
	VTool.SetConfigType("ini")
	VTool.ReadInConfig()
	VTool.WatchConfig()
	log.Println("加载的配置项:",VTool.AllSettings())
	log.Println("app common configs:",VTool.Get("common"))
}

func test(){

}