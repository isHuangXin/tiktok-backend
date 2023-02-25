package main

import (
	"github.com/RaymondCode/simple-demo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	go service.RunMessageServer()
	r := gin.Default()

	initRouter(r)

	// Listen and serve on 0.0.0.0:8080 ("localhost:8080")
	r.Run()
}
