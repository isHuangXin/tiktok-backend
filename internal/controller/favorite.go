package controller

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/isHuangxin/tiktok-backend/api"
	"net/http"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c context.Context, ctx *app.RequestContext) {
	token := ctx.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		ctx.JSON(http.StatusOK, api.Response{StatusCode: 0})
	} else {
		ctx.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
