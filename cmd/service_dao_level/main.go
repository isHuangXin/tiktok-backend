package main

import (
	pbfavorite "github.com/isHuangxin/tiktok-backend/api/rpc_controller_service/favorite/route"
	initialization "github.com/isHuangxin/tiktok-backend/init"
	"github.com/isHuangxin/tiktok-backend/internal/dao"
	"github.com/isHuangxin/tiktok-backend/internal/service"
	"github.com/isHuangxin/tiktok-backend/internal/utils/cronUtils"
	"github.com/isHuangxin/tiktok-backend/internal/utils/jwt"
	"github.com/isHuangxin/tiktok-backend/internal/utils/logger"
	"google.golang.org/grpc"
	"net"
)

const (
	port = ":50051"
)

func initAll() {
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

func main() {
	initAll()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		logger.GlobalLogger.Fatal().Err(err)
	}
	s := grpc.NewServer()
	pbfavorite.RegisterFavoriteInfoServer(s, service.GetFavoriteServiceInstance())
	if err = s.Serve(lis); err != nil {
		panic(err)
	}
}
