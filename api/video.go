package api

import (
	"giligili/serializer"
	"giligili/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CreateVideo 上传视频路由函数
func CreateVideo(c *gin.Context) {
	videoService := service.CreateVideoService{}
	if err := c.ShouldBindJSON(&videoService); err == nil {
		res := videoService.Create()
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, serializer.Response{
			Code:  1,
			Data:  nil,
			Msg:   "",
			Error: err.Error(),
		})
	}

}
