package search

import (
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/builder/template/adminresource/searches"
	"github.com/quarkcms/quark-go/pkg/component/admin/form/fields/selectfield"
	"gorm.io/gorm"
)

type JobLogStatus struct {
	searches.Select
}

// 初始化
func (p *JobLogStatus) Init() *JobLogStatus {
	p.ParentInit()
	p.Name = "状态"

	return p
}

// 执行查询
func (p *JobLogStatus) Apply(ctx *builder.Context, query *gorm.DB, value interface{}) *gorm.DB {
	return query.Where("job_logs.status = ?", value)
}

// 属性
func (p *JobLogStatus) Options(ctx *builder.Context) interface{} {

	return []*selectfield.Option{
		p.Option(0, "失败"),
		p.Option(1, "成功"),
	}
}
