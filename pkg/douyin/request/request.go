package request

import (
	"net/http"
	"strings"

	"github.com/guonaihong/gout"
	"github.com/quarkcms/douyin-helper/pkg/douyin/sign"
)

// 全局Cookies
var Cookies []*http.Cookie

// Header模拟数据
var HeaderMock = map[string]string{
	"accept":             "application/json, text/plain, */*",
	"accept-language":    "zh-CN,zh;q=0.9",
	"cache-control":      "no-cache",
	"pragma":             "no-cache",
	"referer":            "https://www.douyin.com/",
	"sec-ch-ua":          "\"Chromium\";v=\"110\", \"Not A(Brand\";v=\"24\", \"Google Chrome\";v=\"110\"",
	"sec-ch-ua-mobile":   "?0",
	"sec-ch-ua-platform": "\"Windows\"",
	"sec-fetch-dest":     "empty",
	"sec-fetch-mode":     "cors",
	"sec-fetch-site":     "same-origin",
	"user-agent":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36",
	"Host":               "www.douyin.com",
	"Connection":         "keep-alive",
}

// Query模拟数据
var QueryMock = map[string]string{
	"device_platform":  "webapp",
	"aid":              "6383",
	"channel":          "channel_pc_web",
	"pc_client_type":   "1",
	"version_code":     "170400",
	"version_name":     "17.4.0",
	"cookie_enabled":   "true",
	"screen_width":     "1536",
	"screen_height":    "864",
	"browser_language": "zh-CN",
	"browser_platform": "Win32",
	"browser_name":     "Chrome",
	"browser_version":  "110.0.0.0",
	"browser_online":   "true",
	"engine_name":      "Blink",
	"engine_version":   "110.0.0.0",
	"os_name":          "Windows",
	"os_version":       "10",
	"cpu_core_num":     "8",
	"device_memory":    "8",
	"platform":         "PC",
	"downlink":         "10",
	"effective_type":   "4g",
	"round_trip_time":  "50",
	"webid":            "7225873996603786752",
}

// 请求结构体
type Request struct {
	Debug       bool
	Url         string
	Method      string
	Query       map[string]string
	Header      map[string]string
	QueryString string
	MsToken     string
	XBogus      string
}

// 创建对象
func New() *Request {
	p := &Request{}

	// 初始化MsToken
	p.MsToken, _ = sign.New().GetMsToken(128)

	// 初始化Header
	p.Header = HeaderMock

	// 初始化Query
	p.Query = QueryMock

	return p
}

// 设置Debug模式
func (p *Request) SetDebug(debug bool) *Request {
	p.Debug = debug

	return p
}

// 设置Header参数
func (p *Request) SetHeader(header map[string]string) *Request {
	for k, v := range header {
		p.Header[k] = v
	}

	return p
}

// 设置Query参数
func (p *Request) SetQuery(query map[string]string) *Request {
	for k, v := range query {
		p.Query[k] = v
	}

	return p
}

// 将Query参数转换字符串
func (p *Request) QueryToString() *Request {
	queryString := ""
	for k, v := range p.Query {
		queryString = queryString + k + "=" + v + "&"
	}
	p.QueryString = strings.Trim(queryString, "&")

	return p
}

// 参数签名
func (p *Request) SignParams() *Request {
	queryString := p.QueryString + "&msToken=" + p.MsToken
	XBogus, _ := sign.New().GetXBogus(queryString, p.Header["user-agent"])
	p.QueryString = queryString + "&X-Bogus=" + XBogus

	return p
}

// GET请求
func (p *Request) GET(url string) *Request {
	p.Method = "GET"
	p.Url = url

	return p
}

// POST请求
func (p *Request) POST(url string) *Request {
	p.Method = "POST"
	p.Url = url

	return p
}

// 发送请求
func (p *Request) Do() error {
	var err error

	// 将query转换字符串
	p.QueryToString()

	// 对参数进行签名
	p.SignParams()

	switch p.Method {
	case "GET":
		err = gout.
			GET(p.Url + "?" + p.QueryString).
			Debug(p.Debug).
			SetHeader(p.Header).
			SetCookies(Cookies...).
			Do()
	case "POST":
		err = gout.
			POST(p.Url + "?" + p.QueryString).
			Debug(p.Debug).
			SetHeader(p.Header).
			SetCookies(Cookies...).
			Do()
	}

	return err
}
