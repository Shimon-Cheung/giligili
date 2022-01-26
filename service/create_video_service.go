package service

import (
	"giligili/model"
	"giligili/serializer"
)

// CreateVideoService 用来接收上传视频传参的结构体，进行数据的校验，以及具体响应方法
type CreateVideoService struct {
	Title string `json:"title" binding:"required,min=5,max=30"`
	Info  string `json:"info" binding:"required,min=0,max=300"`
}

func (service *CreateVideoService) Create() serializer.Response {
	video := model.Video{
		Title: service.Title,
		Info:  service.Info,
	}
	err := model.DB.Create(&video).Error
	if err != nil {
		return serializer.Response{
			Code:  50001,
			Data:  nil,
			Msg:   "视频保存失败",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code:  0,
		Data:  serializer.BuildVideo(video),
		Msg:   "",
		Error: "",
	}

}
