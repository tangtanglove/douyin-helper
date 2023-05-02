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
	debug      bool
	QueryValue *QueryValue
	Header     *Header
}

func New() *Douyin {
	p := &Douyin{}

	p.InitHeader()
	p.InitQuery()

	return p
}

// 设置Debug模式
func (p *Douyin) Debug(debug bool) *Douyin {
	p.debug = debug

	return p
}

// 初始化通用头部
func (p *Douyin) InitHeader() *Douyin {
	p.Header = &Header{
		Accept:          "application/json, text/plain, */*",
		AcceptEncoding:  "gzip, deflate, br",
		AcceptLanguage:  "zh-CN,zh;q=0.9",
		CacheControl:    "no-cache",
		Pragma:          "no-cache",
		Referer:         "https://www.douyin.com/user/self",
		SecChUa:         "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
		SecChUaMobile:   "?0",
		SecChUaPlatform: "\"Windows\"",
		SecFetchDest:    "empty",
		SecFetchMode:    "cors",
		SecFetchSite:    "same-origin",
		UserAgent:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		Host:            "www.douyin.com",
		Connection:      "keep-alive",
	}

	return p
}

// 初始化通用Query参数
func (p *Douyin) InitQuery() *Douyin {
	p.QueryValue = &QueryValue{
		DevicePlatform:  "webapp",
		Aid:             "6383",
		Channel:         "channel_pc_web",
		PcClientType:    "1",
		VersionCode:     "170400",
		VersionName:     "17.4.0",
		CookieEnabled:   "true",
		ScreenWidth:     "1536",
		ScreenHeight:    "864",
		BrowserLanguage: "zh-CN",
		BrowserPlatform: "Win32",
		BrowserName:     "Chrome",
		BrowserVersion:  "110.0.0.0",
		BrowserOnline:   "true",
		EngineName:      "Blink",
		EngineVersion:   "110.0.0.0",
		OsName:          "Windows",
		OsVersion:       "10",
		CpuCoreNum:      "8",
		DeviceMemory:    "8",
		Platform:        "PC",
		Downlink:        "10",
		EffectiveType:   "4g",
		RoundTripTime:   "50",
		Webid:           "7225873996603786752",
	}

	return p
}

// 登录获取用户信息
func (p *Douyin) Login() *Douyin {
	// create chrome instance
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

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 150*time.Second)
	defer cancel()

	// navigate to a page, wait for an element, click
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

// 将query转换字符串
func (p *Douyin) queryParser(apiQuerys []*request.Query) []*request.Query {
	querys := []*request.Query{
		{
			Key:   "device_platform",
			Value: p.QueryValue.DevicePlatform,
		},
		{
			Key:   "aid",
			Value: p.QueryValue.Aid,
		},
		{
			Key:   "channel",
			Value: p.QueryValue.Channel,
		},
	}

	envQuerys := []*request.Query{
		{
			Key:   "pc_client_type",
			Value: p.QueryValue.PcClientType,
		},
		{
			Key:   "version_code",
			Value: p.QueryValue.VersionCode,
		},
		{
			Key:   "version_name",
			Value: p.QueryValue.VersionName,
		},
		{
			Key:   "cookie_enabled",
			Value: p.QueryValue.CookieEnabled,
		},
		{
			Key:   "screen_width",
			Value: p.QueryValue.ScreenWidth,
		},
		{
			Key:   "screen_height",
			Value: p.QueryValue.ScreenHeight,
		},
		{
			Key:   "browser_language",
			Value: p.QueryValue.BrowserLanguage,
		},
		{
			Key:   "browser_platform",
			Value: p.QueryValue.BrowserPlatform,
		},

		{
			Key:   "browser_name",
			Value: p.QueryValue.BrowserName,
		},
		{
			Key:   "browser_version",
			Value: p.QueryValue.BrowserVersion,
		},
		{
			Key:   "browser_online",
			Value: p.QueryValue.BrowserOnline,
		},
		{
			Key:   "engine_name",
			Value: p.QueryValue.EngineName,
		},
		{
			Key:   "engine_version",
			Value: p.QueryValue.EngineVersion,
		},
		{
			Key:   "os_name",
			Value: p.QueryValue.OsName,
		},
		{
			Key:   "os_version",
			Value: p.QueryValue.OsVersion,
		},
		{
			Key:   "cpu_core_num",
			Value: p.QueryValue.CpuCoreNum,
		},
		{
			Key:   "device_memory",
			Value: p.QueryValue.DeviceMemory,
		},
		{
			Key:   "platform",
			Value: p.QueryValue.Platform,
		},
		{
			Key:   "downlink",
			Value: p.QueryValue.Downlink,
		},
		{
			Key:   "effective_type",
			Value: p.QueryValue.EffectiveType,
		},
		{
			Key:   "round_trip_time",
			Value: p.QueryValue.RoundTripTime,
		},
		{
			Key:   "webid",
			Value: p.QueryValue.Webid,
		},
	}

	querys = append(querys, apiQuerys...)
	querys = append(querys, envQuerys...)

	return querys
}

// 通用的头部
func (p *Douyin) headerParser() map[string]string {
	return map[string]string{
		"accept":             p.Header.Accept,
		"accept-language":    p.Header.AcceptLanguage,
		"cache-control":      p.Header.CacheControl,
		"cookie":             p.Header.Cookie,
		"pragma":             p.Header.Pragma,
		"referer":            p.Header.Referer,
		"sec-ch-ua":          p.Header.SecChUa,
		"sec-ch-ua-mobile":   p.Header.SecChUaMobile,
		"sec-ch-ua-platform": p.Header.SecChUaPlatform,
		"sec-fetch-dest":     p.Header.SecFetchDest,
		"sec-fetch-mode":     p.Header.SecFetchMode,
		"sec-fetch-site":     p.Header.SecFetchSite,
		"user-agent":         p.Header.UserAgent,
		"Host":               p.Header.Host,
		"Connection":         p.Header.Connection,
	}
}
