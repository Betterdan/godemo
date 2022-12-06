package boot

import (
	"demo/common"
	"demo/configs"
	"demo/infrastructure"
	"demo/infrastructure/logger"

	//"demo/infrastructure/logger"
	"demo/router"
	"fmt"
	"log"
	"net/http"
	"time"
)

func Init() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("程序启动失败！err:", err)
		}
	}()
	path,err := common.GetCurrentPath()
	if err != nil {
		fmt.Println("加载配置失败!",err)
	}

	/**加载日志**/
	logger.LoadMyLog()


	/** 加载配置 */
	configs.Init(path+"/configs")

	/** 加载数据库  */
	infrastructure.LoadMysql()

	/** 加载路由 */
	r := router.LoadRouter()

	s := &http.Server{
		Addr:           ":8787",
		Handler:        r,
		ReadTimeout:    1000 * time.Second,
		WriteTimeout:   1000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
