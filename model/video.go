package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/pkg/app/model"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"gorm.io/gorm"
)

// 调度器模型
type Video struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	DouyinId  string         `json:"douyin_id" gorm:"size:200;not null"`
	Name      string         `json:"name" gorm:"size:200;not null"`
	Status    int            `json:"status" gorm:"size:4;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Video) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(18) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 18, Name: "抖音管理", GuardName: "admin", Icon: "icon-comment", Type: "default", Pid: 0, Sort: 0, Path: "/douyin", Show: 1, Status: 1},
		{Id: 19, Name: "视频列表", GuardName: "admin", Icon: "", Type: "engine", Pid: 18, Sort: 0, Path: "/api/admin/video/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
