package main

import (
	"log"
	"meetingserver/router"
	"meetingserver/util"

	"github.com/gin-gonic/gin"
)

func main() {

	//读取配制文件
	cfg, err := util.ParseConfig("./config/app.json")
	if err != nil {
		log.Print(err.Error())
		panic(err)
		return
	}

	//初始化数据库
	_,err = util.InitDataBase(cfg)
	if err != nil {
		log.Print(err.Error())
		panic(err)
		return
	}

	app := gin.Default()


	appRouter := new(router.AppRouter)
	appRouter.UserRouter(app)
	appRouter.MeetingRouter(app)

	app.Run(cfg.AppHost + ":" + cfg.AppPort)
}


