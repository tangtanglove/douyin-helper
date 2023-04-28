package action

import (
	"github.com/quarkcms/douyin-helper/utils/douyin"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/actions"
	"gorm.io/gorm"
)

type Sync struct {
	actions.Action
}

// 初始化
func (p *Sync) Init() *Sync {
	// 初始化父结构
	p.ParentInit()

	// 行为名称
	p.Name = "同步数据"

	// 执行成功后刷新的组件
	p.Reload = "table"

	// 是否具有loading，当action 的作用类型为ajax,submit时有效
	p.WithLoading = true

	// 设置展示位置
	p.SetOnlyOnIndex(true)

	// 行为类型
	p.ActionType = "ajax"

	return p
}

// 执行行为句柄
func (p *Sync) Handle(ctx *builder.Context, query *gorm.DB) error {

	douyin.New().GetUserInfo()

	return ctx.JSONOk("操作成功")
}
