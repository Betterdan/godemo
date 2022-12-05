package configs

import (
	enum2 "demo/common/enum"
	"fmt"
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
	err := VTool.ReadInConfig()
	if err != nil {
		fmt.Println("读取配置失败",err)
	}
	VTool.WatchConfig()
	log.Println("加载的配置项:",VTool.AllSettings())
	log.Println("app common configs:",VTool.Get("common"))
}

func test(){

}