package main

import (
	"github.com/gin-gonic/gin"
	"github.com/isHuangXin/tiktok-backend/service"
)

func main() {
	go service.RunMessageServer()
	r := gin.Default()

	initRouter(r)

	// Listen and serve on 0.0.0.0:8080 ("localhost:8080")
	r.Run()
}
