package douyin

import (
	"github.com/guonaihong/gout"
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

type Cookie struct {
	Key   string
	Value string
}

type Douyin struct {
	Cookie *Cookie
}

// 解析获取用户信息返回的Body数据
type UserRspBody struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
	Data    string `json:"data"`
}

// 解析获取用户信息返回的Header数据
type UserRspHeader struct {
	Sid  string `header:"sid"`
	Time int    `header:"time"`
}

func New() *Douyin {
	return &Douyin{}
}

// 通用的头部
func header() map[string]interface{} {
	return gout.H{
		"accept":                            " application/json, text/plain, */*",
		"accept-language":                   " zh-CN,zh;q=0.9",
		"bd-ticket-guard-client-cert":       " LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUNGVENDQWJ1Z0F3SUJBZ0lWQU9NUFhOL1NTVWFwNzZSR2hoZFpHbmNPSDFtT01Bb0dDQ3FHU000OUJBTUMKTURFeEN6QUpCZ05WQkFZVEFrTk9NU0l3SUFZRFZRUUREQmwwYVdOclpYUmZaM1ZoY21SZlkyRmZaV05rYzJGZgpNalUyTUI0WERUSXpNRFF5TlRBMk5ETTFORm9YRFRNek1EUXlOVEUwTkRNMU5Gb3dKekVMTUFrR0ExVUVCaE1DClEwNHhHREFXQmdOVkJBTU1EMkprWDNScFkydGxkRjluZFdGeVpEQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDkKQXdFSEEwSUFCRThGZW9ka05XM1h1Q2ZpUTRrNWc5ekZSNFRyMlJwakZtcUNiZUVNYUY5VHU4REhYTm51OERuNQpMR3JuNkxkQThVSVhWbGtrZkpzWEZkbVJXRzZRc3Bhamdia3dnYll3RGdZRFZSMFBBUUgvQkFRREFnV2dNREVHCkExVWRKUVFxTUNnR0NDc0dBUVVGQndNQkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3TUdDQ3NHQVFVRkJ3TUUKTUNrR0ExVWREZ1FpQkNDbFBUei9IMXAwS2NBVlluelpjLzE4NGM4Tjh1UGppV3QxdmlDQVhVRnNYekFyQmdOVgpIU01FSkRBaWdDQXlwV2Zxam1SSUVvM01UazFBZTNNVW0wZHRVM3FrMFlEWGVaU1hleUpIZ3pBWkJnTlZIUkVFCkVqQVFnZzUzZDNjdVpHOTFlV2x1TG1OdmJUQUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpQWkxK0JsOXFrSlo3Z1QKUno1b042alZXMnFyR20rSFA3N0FDT0JPUVNPNGVRSWhBSlFEV3RjMC9sekJUbzRkUjVjRVpJdGhEWVRaeDFmVworZTlzakpQUVUvZGsKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=",
		"bd-ticket-guard-client-data":       " eyJ0c19zaWduIjoidHMuMS41MTEyYzBmZjkxMDQyOGEzN2Q0MzQyZTI4NDQxOWQ4ZDkyMzJhNDMyNDdhYWJiZDFjMzM5OGQ3Yzg4ZjZjOGRmYzRmYmU4N2QyMzE5Y2YwNTMxODYyNGNlZGExNDkxMWNhNDA2ZGVkYmViZWRkYjJlMzBmY2U4ZDRmYTAyNTc1ZCIsInJlcV9jb250ZW50IjoidGlja2V0LHBhdGgsdGltZXN0YW1wIiwicmVxX3NpZ24iOiJNRVlDSVFDMk5wUTZOZURFc0YrMmVOSWtJRmltNmIyTnBFZEZBbGo3TUdPdzZMYktsd0loQU1vZ0I4amNaaSsrbmlPNkdZMUtWZ3d3L3poNGpZYnNsWGNmTGMvcXRVd20iLCJ0aW1lc3RhbXAiOjE2ODI2NTg4NDd9",
		"bd-ticket-guard-iteration-version": " 1",
		"bd-ticket-guard-version":           " 2",
		"cache-control":                     " no-cache",
		"cookie":                            " ttcid=d7863790afe24f00af972153515895eb34; passport_csrf_token=a86bf5be57a0af1af777ead381701560; passport_csrf_token_default=a86bf5be57a0af1af777ead381701560; s_v_web_id=verify_lgvwgpa2_crJUpTiW_Sf1c_4kta_8jbK_SUj6d3Uywqo4; passport_assist_user=Cj1WqiEodmILygka5e0MKBgiEL6mwA7UEWVcwtvKm5C4HoDfiplKQGuDhUCA-5ptGd-XWMB6tkzC5rPb89DTGkgKPOGH9_6KR6lDeVXbY7qEC_cJzrSANlZ3o315EwsceLmv25awV09goK9ThDPxZoeRceWo27M9vZTgERAwuBCktq8NGImv1lQiAQNwza34; n_mh=TZyGh1CAhTx0BO-JVitmjv8kUBPcdxCfi5P1KDYSzps; sso_uid_tt=93bef7ddf39cc4052bf859668c458199; sso_uid_tt_ss=93bef7ddf39cc4052bf859668c458199; toutiao_sso_user=e023c5fa4e28683ff967b9c431a487b8; toutiao_sso_user_ss=e023c5fa4e28683ff967b9c431a487b8; sid_ucp_sso_v1=1.0.0-KDk0NjRmZDg2NGM1ZjFkMmNlM2I1MGRmNmNiZmIwZTM2YmQ2MTlmZWQKHQjRgLnUjAMQqe2dogYY7zEgDDDggaDfBTgGQPQHGgJsZiIgZTAyM2M1ZmE0ZTI4NjgzZmY5NjdiOWM0MzFhNDg3Yjg; ssid_ucp_sso_v1=1.0.0-KDk0NjRmZDg2NGM1ZjFkMmNlM2I1MGRmNmNiZmIwZTM2YmQ2MTlmZWQKHQjRgLnUjAMQqe2dogYY7zEgDDDggaDfBTgGQPQHGgJsZiIgZTAyM2M1ZmE0ZTI4NjgzZmY5NjdiOWM0MzFhNDg3Yjg; passport_auth_status=9cf87f1225db76ede2a406c145eb8b47%2C; passport_auth_status_ss=9cf87f1225db76ede2a406c145eb8b47%2C; uid_tt=51560f1274cf303431e0fb8521b47630; uid_tt_ss=51560f1274cf303431e0fb8521b47630; sid_tt=a1f6020926d9a1174eeab2e395ed78a6; sessionid=a1f6020926d9a1174eeab2e395ed78a6; sessionid_ss=a1f6020926d9a1174eeab2e395ed78a6; _tea_utm_cache_2018=undefined; publish_badge_show_info=%220%2C0%2C0%2C1682405067179%22; LOGIN_STATUS=1; store-region=cn-he; store-region-src=uid; sid_guard=a1f6020926d9a1174eeab2e395ed78a6%7C1682405090%7C5183946%7CSat%2C+24-Jun-2023+06%3A43%3A56+GMT; sid_ucp_v1=1.0.0-KDFiYTk3NDc4MDc4N2EyZjI3MDUyMTg1YzZjM2Q3YzM0YTZlZDgwNjUKGQjRgLnUjAMQ4u2dogYY7zEgDDgGQPQHSAQaAmxxIiBhMWY2MDIwOTI2ZDlhMTE3NGVlYWIyZTM5NWVkNzhhNg; ssid_ucp_v1=1.0.0-KDFiYTk3NDc4MDc4N2EyZjI3MDUyMTg1YzZjM2Q3YzM0YTZlZDgwNjUKGQjRgLnUjAMQ4u2dogYY7zEgDDgGQPQHSAQaAmxxIiBhMWY2MDIwOTI2ZDlhMTE3NGVlYWIyZTM5NWVkNzhhNg; d_ticket=91faccacc5cb307444639ee60d733d13696d7; download_guide=%223%2F20230425%22; pwa2=%220%7C1%22; my_rd=1; douyin.com; strategyABtestKey=%221682648028.606%22; csrf_session_id=56f4604d4df219a719f5f216eab27fc1; bd_ticket_guard_client_data=eyJiZC10aWNrZXQtZ3VhcmQtdmVyc2lvbiI6MiwiYmQtdGlja2V0LWd1YXJkLWl0ZXJhdGlvbi12ZXJzaW9uIjoxLCJiZC10aWNrZXQtZ3VhcmQtY2xpZW50LWNlcnQiOiItLS0tLUJFR0lOIENFUlRJRklDQVRFLS0tLS1cbk1JSUNGVENDQWJ1Z0F3SUJBZ0lWQU9NUFhOL1NTVWFwNzZSR2hoZFpHbmNPSDFtT01Bb0dDQ3FHU000OUJBTUNcbk1ERXhDekFKQmdOVkJBWVRBa05PTVNJd0lBWURWUVFEREJsMGFXTnJaWFJmWjNWaGNtUmZZMkZmWldOa2MyRmZcbk1qVTJNQjRYRFRJek1EUXlOVEEyTkRNMU5Gb1hEVE16TURReU5URTBORE0xTkZvd0p6RUxNQWtHQTFVRUJoTUNcblEwNHhHREFXQmdOVkJBTU1EMkprWDNScFkydGxkRjluZFdGeVpEQlpNQk1HQnlxR1NNNDlBZ0VHQ0NxR1NNNDlcbkF3RUhBMElBQkU4RmVvZGtOVzNYdUNmaVE0azVnOXpGUjRUcjJScGpGbXFDYmVFTWFGOVR1OERIWE5udThEbjVcbkxHcm42TGRBOFVJWFZsa2tmSnNYRmRtUldHNlFzcGFqZ2Jrd2diWXdEZ1lEVlIwUEFRSC9CQVFEQWdXZ01ERUdcbkExVWRKUVFxTUNnR0NDc0dBUVVGQndNQkJnZ3JCZ0VGQlFjREFnWUlLd1lCQlFVSEF3TUdDQ3NHQVFVRkJ3TUVcbk1Da0dBMVVkRGdRaUJDQ2xQVHovSDFwMEtjQVZZbnpaYy8xODRjOE44dVBqaVd0MXZpQ0FYVUZzWHpBckJnTlZcbkhTTUVKREFpZ0NBeXBXZnFqbVJJRW8zTVRrMUFlM01VbTBkdFUzcWswWURYZVpTWGV5SkhnekFaQmdOVkhSRUVcbkVqQVFnZzUzZDNjdVpHOTFlV2x1TG1OdmJUQUtCZ2dxaGtqT1BRUURBZ05JQURCRkFpQWkxK0JsOXFrSlo3Z1RcblJ6NW9ONmpWVzJxckdtK0hQNzdBQ09CT1FTTzRlUUloQUpRRFd0YzAvbHpCVG80ZFI1Y0VaSXRoRFlUWngxZldcbitlOXNqSlBRVS9ka1xuLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLVxuIn0=; _tea_utm_cache_1243=undefined; MONITOR_WEB_ID=dbbfc8a3-ad5a-4cac-b7b4-8e1d962b2a86; ttwid=1%7CToKYWF2xpOiDC-WxWcl7mtK1cL2c8eVvBjmRGf6iMV0%7C1682648875%7Cedeb66f19fedfb72fe81a4ffe4c2ef585733f6e10abec9751a53d8bdf78797c1; VIDEO_FILTER_MEMO_SELECT=%7B%22expireTime%22%3A1683253712060%2C%22type%22%3A1%7D; odin_tt=a7110e30944ce9efef6b45c9ca5b627830908bcc296ed16b25db268cb0d7a5a924328905af69aedc611ab7ec8246985a; FOLLOW_NUMBER_YELLOW_POINT_INFO=%22MS4wLjABAAAAT4RzIMDTdiZnonQn1DgSuGBWfhf77u0M0St8kIgzsAA%2F1682697600000%2F0%2F0%2F1682657579066%22; msToken=8uQGAlywEk6TOqCFvlFN17_Zee3UupB1p1GBwYIJXWRgt8JVLGTLekniw4kCIn4EdSPc7BF9WK4UDDwfAVPxBRxMbWPTiZTDzskYHpGLOyUBTysTsgcI2wcE0kfaLA==; __ac_nonce=0644b5423000a7aca3b31; __ac_signature=_02B4Z6wo00f01xLfEhQAAIDDkt3oVut6hEcS.xaAAKEGpHpUsQTNS2MFClzMu97ISqWoZupH1CKAzH9peDRH2Ta1H3tFKOvEzRD3dAIlAiURRohP9PIN0WRVCXFRWqCvaNrasVQjT2xjGy4w65; FOLLOW_LIVE_POINT_INFO=%22MS4wLjABAAAAT4RzIMDTdiZnonQn1DgSuGBWfhf77u0M0St8kIgzsAA%2F1682697600000%2F0%2F1682658342085%2F0%22; home_can_add_dy_2_desktop=%221%22; passport_fe_beating_status=true; tt_scid=tq7g1w3a3a78hoPckuvnN0ZXZlh8iUSH86fqpJqR2tFQRQFtl4L6pH6ZtHFvXsmJd1ec; msToken=Kcd7RtbVOIAUgazq88P6ATcL8Pv3ednsRHzx4k4scO_xdgSetmPNVrJsfbxllNTVOnKXcDqZyGNUxuSeshNOBtHso_kWGzQKmdoHF6ORti2_d5cqA1Ysbn6kQSf8_g==",
		"pragma":                            " no-cache",
		"referer":                           " https://www.douyin.com/user/self",
		"sec-ch-ua":                         " \"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
		"sec-ch-ua-mobile":                  " ?0",
		"sec-ch-ua-platform":                " \"Windows\"",
		"sec-fetch-dest":                    " empty",
		"sec-fetch-mode":                    " cors",
		"sec-fetch-site":                    " same-origin",
		"user-agent":                        " Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
		"Host":                              "www.douyin.com",
		"Connection":                        "keep-alive",
	}
}

// 通用的请求
func query() []string {
	return []string{
		"device_platform", "webapp",
		"aid", "6383",
		"channel", "channel_pc_web",
		"publish_video_strategy_type", "2",
		"source", "channel_pc_web",
		"pc_client_type", "1",
		"version_code", "170400",
		"version_name", "17.4.0",
		"cookie_enabled", "true",
		"screen_width", "1536",
		"screen_height", "864",
		"browser_language", "zh-CN",
		"browser_platform", "Win32",
		"browser_name", "Chrome",
		"browser_version", "110.0.0.0",
		"browser_online", "true",
		"engine_name", "Blink",
		"engine_version", "110.0.0.0",
		"os_name", "Windows",
		"os_version", "10",
		"cpu_core_num", "8",
		"device_memory", "8",
		"platform", "PC",
		"downlink", "10",
		"effective_type", "4g",
		"round_trip_time", "50",
		"webid", "7225873996603786752",
		"msToken", "Kcd7RtbVOIAUgazq88P6ATcL8Pv3ednsRHzx4k4scO_xdgSetmPNVrJsfbxllNTVOnKXcDqZyGNUxuSeshNOBtHso_kWGzQKmdoHF6ORti2_d5cqA1Ysbn6kQSf8_g==",
		"X-Bogus", "DFSzswVuSFsANJPIteCwfe9WX7jh",
	}
}

// 获取用户信息
func (p *Douyin) GetUserInfo() error {
	err := gout.
		GET("https://www.douyin.com/aweme/v1/web/user/profile/self/").
		Debug(true).
		SetHeader(header()).
		SetQuery(query()).
		Do()

	return err
}
