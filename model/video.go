package model

import "gorm.io/gorm"

// 数据库模型结构体，用来创建表和查询数据使用
type Video struct {
	gorm.Model
	Title string
	Info  string
}
