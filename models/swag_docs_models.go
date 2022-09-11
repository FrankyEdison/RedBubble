package models

import (
	"RedBubble/common/responseCode"
	"time"
)

/**
swag生成接口文档时，使用这里的模型来表示响应数据的类型（与整个项目的运行无关，单纯是swagger文档使用）
*/

// ResponseAllCategory
type ResponseAllCategory struct {
	Code responseCode.ResCode // 业务响应状态码
	Msg  string               // 提示信息
	Data []*docAllCategory    // 数据
}

// ResponseCategory
type ResponseCategory struct {
	Code responseCode.ResCode // 业务响应状态码
	Msg  string               // 提示信息
	Data *docCategory         // 数据
}

type docAllCategory struct {
	CategoryName string
	Introduction string
}
type docCategory struct {
	ID           uint
	CreatedAt    time.Time
	CategoryName string
	Introduction string
}
