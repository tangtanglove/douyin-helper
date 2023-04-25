package main

import (
	"github.com/quarkcms/douyin-helper/model"
	"github.com/quarkcms/douyin-helper/resource"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 自动构建本地数据库
func DBMigrate() {

	// 迁移数据
	db.Client.AutoMigrate(
		&model.Video{},
	)

	// 数据填充
	(&model.Video{}).Seeder()
}

func main() {

	// 注册服务
	providers := []interface{}{
		&resource.Video{},
	}

	// 数据库配置信息
	dsn := "./data.db"

	// 加载后台服务
	providers = append(providers, admin.Providers...)

	// 配置资源
	config := &builder.Config{
		AppKey:    "abcdefg",
		Providers: providers,
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 实例化对象
	b := builder.New(config)

	// 静态文件
	b.Static("/", "./website")

	// 自动构建数据库、拉取静态文件
	install.Handle()

	// 自动构建本地数据库
	DBMigrate()

	// 后台中间件
	b.Use(middleware.Handle)

	// 重定向到后台管理
	b.GET("/", func(ctx *builder.Context) error {
		return ctx.Redirect(301, "/admin/")
	})

	// 启动服务
	b.Run(":3000")
}
