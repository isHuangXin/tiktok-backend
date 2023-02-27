package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/isHuangXin/tiktok-backend/api"
	"net/http"
	"path/filepath"
)

type VideoListResponse struct {
	api.Response
	VideoList []api.Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, api.Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	fileName := fmt.Sprintln("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/", fileName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, api.Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, api.Response{
		StatusCode: 0,
		StatusMsg:  fileName + "uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: api.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
