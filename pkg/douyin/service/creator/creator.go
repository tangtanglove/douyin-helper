package creator

import "github.com/quarkcms/douyin-helper/pkg/douyin/request"

// 主域名
const (
	domain = "https://creator.douyin.com"
)

// 接口列表
const (
	mediaAwemePost = "/web/api/media/aweme/post/"
)

type Creator struct {
	debug bool
}

func New() *Creator {
	return &Creator{}
}

// 设置Debug模式
func (p *Creator) Debug(debug bool) *Creator {
	p.debug = debug

	return p
}

// 获取作品列表
func (p *Creator) GetMediaAwemePost() error {

	// 请求URL
	url := domain + mediaAwemePost

	// 查询参数
	query := map[string]string{
		"status":     "0",
		"count":      "12",
		"max_cursor": "0",
	}

	// 发送请求
	result := request.New().
		GET(url).
		SetDebug(p.debug).
		SetQuery(query).
		Do()

	return result
}
