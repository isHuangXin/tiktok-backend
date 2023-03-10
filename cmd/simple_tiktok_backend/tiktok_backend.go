package main

import (
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
	initialization "github.com/isHuangxin/tiktok-backend/init"
	"github.com/isHuangxin/tiktok-backend/init/router"
	"github.com/isHuangxin/tiktok-backend/internal/dao"
	"github.com/isHuangxin/tiktok-backend/internal/utils/cronUtils"
	"github.com/isHuangxin/tiktok-backend/internal/utils/jwt"
	"github.com/isHuangxin/tiktok-backend/internal/utils/logger"
)

// initAll 初始化所有的部分
func initAll() {
	// Init basic operators
	initialization.InitConfig()
	initialization.InitDB()
	initialization.InitOSS()
	initialization.InitRDB()
	initialization.InitKafkaServer()
	initialization.InitKafkaClient()

	// Init Utils
	logger.InitLogger(initialization.LogConf)
	jwt.InitJwt()
	cronUtils.InitCron()

	// Init lower Levels
	dao.DaoInitialization()
}

// 用于单机的极简版抖音后端程序
func main() {
	initAll()
	hServer := server.Default(server.WithHostPorts(fmt.Sprintf("127.0.0.1:%s", initialization.Port)))
	router.InitRouter(hServer)
	hServer.Spin()
}
