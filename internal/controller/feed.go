package controller

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/isHuangxin/tiktok-backend/api"
	"github.com/isHuangxin/tiktok-backend/internal/service"
	"github.com/isHuangxin/tiktok-backend/internal/utils/constants"
	"github.com/isHuangxin/tiktok-backend/internal/utils/jwt"
	"github.com/isHuangxin/tiktok-backend/internal/utils/logger"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	api.Response
	VideoList []api.Video `json:"video_list,omitempty"`
	NextTime  int64       `json:"next_time,omitempty"`
}

// Feed 推送视频流
func Feed(c context.Context, ctx *app.RequestContext) {
	token := ctx.Query("token")
	var userId int64
	var err error
	if token != "" {
		jwt.JwtMiddleware.MiddlewareFunc()(c, ctx)
		userId, err = jwt.GetUserId(c, ctx)
		if err != nil {
			logger.GlobalLogger.Printf("Time = %v, can't get user from Token", time.Now())
			if errors.Is(constants.InvalidTokenErr, err) {
				ctx.JSON(consts.StatusOK, api.Response{
					StatusCode: int32(api.TokenInvalidErr),
					StatusMsg:  api.ErrorCodeToMsg[api.TokenInvalidErr],
				})
			} else {
				ctx.JSON(consts.StatusOK, api.Response{
					StatusCode: int32(api.InnerDataBaseErr),
					StatusMsg:  api.ErrorCodeToMsg[api.InnerDataBaseErr],
				})
			}
			return
		}
	}
	latestTimeStr := ctx.Query("latest_time")
	logger.GlobalLogger.Printf("latestTime = %v", latestTimeStr)
	var latestTime time.Time
	if latestTimeStr == "" {
		latestTime = time.Now()
	} else {
		latestTimeInt, err := strconv.ParseInt(latestTimeStr, 10, 64)
		if err != nil {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.InputFormatCheckErr),
				StatusMsg:  api.ErrorCodeToMsg[api.InputFormatCheckErr],
			})
			return
		}
		latestTime = time.Unix(latestTimeInt, 0)
	}

	nextTime, videoList, err := service.GetFeedServiceInstance().Feed(userId, latestTime)
	if err != nil {
		if errors.Is(constants.RecordNotExistErr, err) {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.RecordNotExistErr),
				StatusMsg:  api.ErrorCodeToMsg[api.RecordNotExistErr],
			})
		} else if errors.Is(constants.NoVideoErr, err) {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.NoVideoErr),
				StatusMsg:  api.ErrorCodeToMsg[api.NoVideoErr],
			})
		} else {
			ctx.JSON(consts.StatusOK, api.Response{
				StatusCode: int32(api.InnerDataBaseErr),
				StatusMsg:  api.ErrorCodeToMsg[api.InnerDataBaseErr],
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, FeedResponse{
		Response:  api.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}
