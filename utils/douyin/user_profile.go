package douyin

import "github.com/quarkcms/douyin-helper/utils/request"

// 解析获取用户信息返回的Body数据
type UserProfileRspBody struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
	Data    string `json:"data"`
}

// 解析获取用户信息返回的Header数据
type UserProfileRspHeader struct {
	Sid  string `header:"sid"`
	Time int    `header:"time"`
}

// 结构体
type UserProfile struct {
	Debug       bool
	Headers     interface{}
	QueryString string
}

// 获取用户信息
func (p *Douyin) GetUserProfile() error {

	// 获取查询参数
	querys := p.queryParser([]*request.Query{
		{
			Key:   "publish_video_strategy_type",
			Value: "2",
		},
		{
			Key:   "source",
			Value: "channel_pc_web",
		},
	})

	// 获取头部参数
	header := p.headerParser()

	// 发送请求
	result := request.New().
		GET(mainDomain + userProfileSelf).
		SetDebug(p.debug).
		SetHeader(header).
		SetQuery(querys).
		Do()

	return result
}
