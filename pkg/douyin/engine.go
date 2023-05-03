package douyin

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/quarkcms/douyin-helper/pkg/douyin/request"
	"github.com/quarkcms/douyin-helper/pkg/douyin/service/creator"
	"github.com/quarkcms/douyin-helper/pkg/douyin/service/user"
)

const (
	domain = "https://www.douyin.com"
)

type Engine struct {
	debug bool
}

func New() *Engine {
	return &Engine{}
}

// 设置Debug模式
func (p *Engine) Debug(debug bool) *Engine {
	p.debug = debug

	return p
}

// 用户登录
func (p *Engine) Login() *Engine {

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
		chromedp.Navigate(domain),
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

// 获取用户服务实例
func (p *Engine) UserService() *user.User {

	return user.New().Debug(p.debug)
}

// 获取创作者服务实例
func (p *Engine) CreatorService() *creator.Creator {

	return creator.New().Debug(p.debug)
}
