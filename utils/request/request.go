package request

import (
	"net/http"
	"strings"

	"github.com/guonaihong/gout"
	"github.com/quarkcms/douyin-helper/utils/sign"
)

var Cookies []*http.Cookie

type Query struct {
	Key   string
	Value string
}

// 请求结构体
type Request struct {
	Debug       bool
	Url         string
	Method      string
	Header      map[string]string
	Querys      []*Query
	QueryString string
	MsToken     string
	XBogus      string
}

// 创建对象
func New() *Request {
	p := &Request{}

	// 初始化MsToken
	p.MsToken, _ = sign.New().GetMsToken(128)

	return p
}

// 设置Debug模式
func (p *Request) SetDebug(debug bool) *Request {
	p.Debug = debug

	return p
}

// 获取通用头部
func (p *Request) SetHeader(header map[string]string) *Request {
	p.Header = header

	return p
}

// 获取通用Query参数
func (p *Request) SetQuery(querys []*Query) *Request {
	p.Querys = querys

	return p
}

// 将query转换字符串
func (p *Request) QueryToString() *Request {
	queryString := ""
	for _, v := range p.Querys {
		queryString = queryString + v.Key + "=" + v.Value + "&"
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
