package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/isHuangxin/tiktok-backend/internal/controller"
	"github.com/isHuangxin/tiktok-backend/internal/utils/jwt"
)

func InitRouter(hertz *server.Hertz) {
	// public directory is used to serve static resources
	hertz.Static("/static", "./public")

	// 用户注册与登录
	hertz.POST("/douyin/user/register", controller.Register)
	hertz.POST("/douyin/user/login", jwt.JwtMiddleware.LoginHandler)

	// 鉴权
	auth := hertz.Group("/douyin", jwt.JwtMiddleware.MiddlewareFunc())

	// basic apis
	auth.GET("/feed/", controller.Feed)
	auth.GET("/user/", controller.UserInfo)
	auth.POST("/publish/action/", controller.Publish)
	auth.GET("/publish/list/", controller.PublishList)

	// extra apis - I
	auth.POST("/favorite/action/", controller.FavoriteAction)
	auth.GET("/favorite/list/", controller.FavoriteList)
	auth.POST("/comment/action/", controller.CommentAction)
	auth.GET("/comment/list/", controller.CommentList)

	// extra apis - II
	auth.POST("/relation/action/", controller.RelationAction)
	auth.GET("/relation/follow/list/", controller.FollowList)
	auth.GET("/relation/follower/list/", controller.FollowerList)
	auth.GET("/relation/friend/list/", controller.FriendList)
	auth.GET("/message/chat/", controller.MessageChat)
	auth.POST("/message/action/", controller.MessageAction)

}
