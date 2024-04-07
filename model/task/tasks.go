// 自动生成模板Tasks
package task

import "time"

// zmTask表 结构体  Tasks
type Tasks struct {
	ID        int       `gorm:"primarykey" json:"ID"`                                                        // 主键ID
	Title     string    `json:"title" form:"title" gorm:"column:title;comment:名称;size:256;"`                 //名称
	Desc      string    `json:"desc" form:"desc" gorm:"column:desc;comment:任务描述;size:1024;"`                 //任务描述
	TagId     int       `json:"tagId" form:"tagId" gorm:"column:tag_id;comment:类型id;size:10;"`               //类型id
	UserId    int64     `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"`            //用户id
	Status    int       `json:"status" form:"status" gorm:"column:status;comment:状态,2已完成,1招聘中,0待审核;"`        //状态,2已完成,1招聘中,0待审核
	Address   string    `json:"address" form:"address" gorm:"-"`        //工作地址
	AddressId int       `json:"address_id" form:"address_id" gorm:"column:address_id;comment:地址id;size:10;"` //address_id
	Mobile    string    `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;size:1024;"`            //手机号
	CreatedAt time.Time // 创建时间
	TagName   string    `json:"tag_name" gorm:"-"`
	IsDeleted int       `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
	AddTime   int64     `json:"add_time" form:"add_time" gorm:"column:add_time;comment:添加时间时间戳;"`             //添加时间时间戳
}

// TableName zmTask表 Tasks自定义表名 zm_task
func (Tasks) TableName() string {
	return "zm_task"
}
