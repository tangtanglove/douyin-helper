package sign

import (
	"crypto/rand"
	_ "embed"
	"regexp"

	"github.com/dop251/goja"
	"github.com/guonaihong/gout"
	"github.com/pkg/errors"
)

//go:embed sign.js
var script string

// 结构体
type Sign struct{}

// 创建对象
func New() *Sign {
	return &Sign{}
}

// 创建msToken，默认128位
func (p *Sign) GetMsToken(length int) (string, error) {
	const characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	token := make([]byte, length)
	for i, b := range randomBytes {
		token[i] = characters[int(b)%len(characters)]
	}

	return string(token), nil
}

// 创建ttwid
func (p *Sign) GetTtwid() (string, error) {
	// 发送请求
	resp, err := gout.
		POST("https://ttwid.bytedance.com/ttwid/union/register/").
		SetJSON(gout.H{
			"aid":           1768,
			"union":         true,
			"needFid":       false,
			"region":        "cn",
			"cbUrlProtocol": "https",
			"service":       "www.ixigua.com",
			"migrate_info":  map[string]string{"ticket": "", "source": "node"},
		}).Response()

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	cookie := resp.Header.Get("Set-Cookie")
	re := regexp.MustCompile(`ttwid=([^;]+)`)
	if match := re.FindStringSubmatch(cookie); match != nil {
		return match[1], nil
	}

	return "", errors.New("douyin ttwid request failed")
}

// 生成X-Bogus
func (p *Sign) GetXBogus(queryString string, userAgent string) (string, error) {

	// 初始化JavaScripts运行时
	vm := goja.New()

	// 加载抖音签名脚本
	_, _ = vm.RunString(script)

	// 签名
	sign, err := vm.RunString("sign('" + queryString + "', '" + userAgent + "')")
	if err != nil {
		return "", err
	}

	return sign.String(), nil
}
