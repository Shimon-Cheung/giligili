package serializer

import "giligili/model"

// 视频的序列化器，用来组织返回的data数据结构的，
//这个需要和models进行区分，这个是用来控制展示哪些字段
type Video struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Info      string `json:"info"`
	CreatedAt int64  `json:"created_at"`
}

func BuildVideo(item model.Video) Video {
	return Video{
		ID:        item.ID,
		Title:     item.Title,
		Info:      item.Info,
		CreatedAt: item.CreatedAt.Unix(),
	}
}
