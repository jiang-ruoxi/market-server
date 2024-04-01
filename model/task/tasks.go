// 自动生成模板Tasks
package task

import "time"

// zmTask表 结构体  Tasks
type Tasks struct {
	ID        int       `gorm:"primarykey" json:"ID"`                                                 // 主键ID
	Title     string    `json:"title" form:"title" gorm:"column:title;comment:名称;size:256;"`          //名称
	Desc      string    `json:"desc" form:"desc" gorm:"column:desc;comment:任务描述;size:1024;"`          //任务描述
	TagId     int       `json:"tagId" form:"tagId" gorm:"column:tag_id;comment:类型id;size:10;"`        //类型id
	UserId    string    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"`     //用户id
	Status    string    `json:"status" form:"status" gorm:"column:status;comment:状态,2已完成,1招聘中,0待审核;"` //状态,2已完成,1招聘中,0待审核
	Address   string    `json:"address" form:"address" gorm:"column:address;comment:工作地址;size:1024;"` //工作地址
	CreatedAt time.Time // 创建时间
	TagName   string    `json:"tag_name" gorm:"-"`
	IsDeleted int       `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zmTask表 Tasks自定义表名 zm_task
func (Tasks) TableName() string {
	return "zm_task"
}
