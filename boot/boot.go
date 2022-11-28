package boot

import "log"

func Init() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("程序启动失败！err:", err)
		}
	}()



}
