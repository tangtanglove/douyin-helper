package resource

import (
	"github.com/quarkcms/douyin-helper/action"
	"github.com/quarkcms/douyin-helper/model"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/actions"
	"github.com/quarkcms/quark-go/pkg/app/handler/admin/searches"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource"
)

type Video struct {
	adminresource.Template
}

// 初始化
func (p *Video) Init() interface{} {

	// 初始化模板
	p.TemplateInit()

	// 标题
	p.Title = "视频"

	// 模型
	p.Model = &model.Video{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Video) Fields(ctx *builder.Context) []interface{} {
	field := &adminresource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "名称").OnlyOnIndex(),

		field.Datetime("created_at", "创建时间").OnlyOnIndex(),
	}
}

// 搜索
func (p *Video) Searches(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&searches.Input{}).Init("name", "名称"),
	}
}

// 行为
func (p *Video) Actions(ctx *builder.Context) []interface{} {

	return []interface{}{
		(&action.Login{}).Init(),
		(&action.Sync{}).Init(),
		(&actions.Delete{}).Init("批量删除"),
		(&actions.Delete{}).Init("删除"),
	}
}
