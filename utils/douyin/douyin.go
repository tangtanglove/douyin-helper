package douyin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/quarkcms/douyin-helper/utils/request"
)

const (
	mainDomain = "https://www.douyin.com"
	liveDomain = "https://live.douyin.com"
)

// 主域名POST接口
const (
	userSetSettings            = "/aweme/v1/web/user/set/settings/"
	commitFollowUser           = "/aweme/v1/web/commit/follow/user/"
	commitItemDigg             = "/aweme/v1/web/commit/item/digg/"
	commentPublish             = "/aweme/v1/web/comment/publish"
	commentDigg                = "/aweme/v1/web/comment/digg"
	commentDelete              = "/aweme/v1/web/comment/delete"
	commitDislikeItem          = "/aweme/v1/web/commit/dislike/item/"
	userBlock                  = "/aweme/v1/web/user/block/"
	familiarFeed               = "/aweme/v1/web/familiar/feed/"
	noticeDel                  = "/aweme/v1/web/notice/del/"
	qrcodeInfo                 = "/aweme/v1/fancy/qrcode/info/"
	danmakuPost_v2             = "/aweme/v1/web/danmaku/post_v2/"
	danmakuDelete_v2           = "/aweme/v1/web/danmaku/delete_v2/"
	danmakuConfSet             = "/aweme/v1/web/danmaku/conf/set/"
	privacyBatchBuildImage     = "/aweme/v1/web/privacy/batch_build_image/"
	batchConvertImage          = "/aweme/v1/web/privacy/batch_convert_image/"
	historyWrite               = "/aweme/v1/web/history/write/"
	historyClear               = "/aweme/v1/web/history/clear/"
	lvideoSubmitHistory        = "/aweme/v1/web/lvideo/submit/history/"
	formatSubmitHistory        = "/aweme/v1/web/format/submit/history/"
	lvideoClearHistory         = "/aweme/v1/web/lvideo/clear/history/"
	awemeCollect               = "/aweme/v1/web/aweme/collect/"
	writeImpression            = "/aweme/v1/web/write/impression/"
	homepageImpression         = "/aweme/v1/web/homepage/impression/"
	removeFollower             = "/aweme/v1/web/remove/follower/"
	commitFollowRequestApprove = "/aweme/v1/web/commit/follow/request/approve/"
	adReportFeedback           = "/aweme/v1/web/ad/report/feedback/"
	appointWrite               = "/aweme/v1/web/appoint/write/"
	imUserActiveConfigSet      = "/aweme/v1/web/im/user/active/config/set"
)

// 主域名GET接口
const (
	creatorSchoolCollect = "/web/api/creator/school/collect/"
	userProfileSelf      = "/aweme/v1/web/user/profile/self/"
	recommendUserDislike = "/aweme/v1/web/recommend/user/dislike/"
	danmakuGet_v2        = "/aweme/v1/web/danmaku/get_v2/"
	danmakuDigg          = "/aweme/v1/web/danmaku/digg/"
	ocpcWrite            = "/aweme/v1/web/ocpc/write/"
	danmakuConfGet       = "/aweme/v1/web/danmaku/conf/get/"
)

type Douyin struct {
	debug bool
}

func New() *Douyin {
	return &Douyin{}
}

// 设置Debug模式
func (p *Douyin) Debug(debug bool) *Douyin {
	p.debug = debug

	return p
}

// 用户登录
func (p *Douyin) Login() *Douyin {

	// 创建Chrome实例
	ctx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(
			chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)

	ctx, cancel := chromedp.NewContext(
		ctx,
	)
	defer cancel()

	// 设置超时
	ctx, cancel = context.WithTimeout(ctx, 150*time.Second)
	defer cancel()

	// 打开页面
	err := chromedp.Run(ctx,
		chromedp.Navigate(mainDomain),
		chromedp.WaitVisible(`#douyin-header > div.oJArD0aS > header > div > div > div.iqHX00br > div > div > div > div:nth-child(6) > div > a`),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// 获取cookies
			cookies, err := network.GetCookies().Do(ctx)
			if err != nil {
				return err
			}

			request.Cookies = []*http.Cookie{}
			for _, v := range cookies {
				request.Cookies = append(request.Cookies, &http.Cookie{
					Name:     v.Name,
					Value:    v.Value,
					Path:     v.Path,
					Domain:   v.Domain,
					HttpOnly: v.HTTPOnly,
					Secure:   v.Secure,
				})
			}

			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	return p
}

// 获取用户信息
func (p *Douyin) GetUserProfile() error {

	// 请求URL
	url := mainDomain + userProfileSelf

	// 查询参数
	query := map[string]string{
		"publish_video_strategy_type": "2",
		"source":                      "source",
	}

	// 发送请求
	result := request.New().
		GET(url).
		SetDebug(p.debug).
		SetQuery(query).
		Do()

	return result
}
