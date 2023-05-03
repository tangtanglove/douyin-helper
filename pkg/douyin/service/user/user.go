package user

import "github.com/quarkcms/douyin-helper/pkg/douyin/request"

const (
	domain = "https://www.douyin.com"
)

// 接口列表
const (
	userProfileSelf = "/aweme/v1/web/user/profile/self/"
)

// 用户平台
type User struct {
	debug bool
}

func New() *User {
	return &User{}
}

// 设置Debug模式
func (p *User) Debug(debug bool) *User {
	p.debug = debug

	return p
}

// 获取用户信息
func (p *User) GetUserProfile() (rsp map[string]interface{}, err error) {

	// 请求URL
	url := domain + userProfileSelf

	// 查询参数
	query := map[string]string{
		"publish_video_strategy_type": "2",
		"source":                      "source",
	}

	// 发送请求
	return request.New().
		GET(url).
		SetDebug(p.debug).
		SetQuery(query).
		Do()
}
