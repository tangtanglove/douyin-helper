package creator

import (
	"strconv"

	"github.com/quarkcms/douyin-helper/pkg/douyin/request"
)

// 主域名
const (
	domain = "https://creator.douyin.com"
)

// 接口列表
const (
	mediaAwemePost = "/web/api/media/aweme/post/"
)

// 创作服务平台
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

// 获取作品列表, status（状态）:0全部，1已发布，2审核中，3未通过；count（读取数量）:12；maxCursor（下一页游标）:0
func (p *Creator) GetMediaAwemePost(status int, count int, maxCursor int) (rsp map[string]interface{}, err error) {

	// 请求URL
	url := domain + mediaAwemePost

	// 查询参数
	query := map[string]string{
		"status":     strconv.Itoa(status),
		"count":      strconv.Itoa(count),
		"max_cursor": strconv.Itoa(maxCursor),
	}

	// 发送请求
	return request.New().
		GET(url).
		SetDebug(p.debug).
		SetQuery(query).
		Do()
}
