// 自动生成模板Address
package address

import (
	"time"
)

// zmAddress表 结构体  Address
type Address struct {
	ID        int       `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	Name      string    `json:"name" form:"name" gorm:"column:name;comment:城市;size:255;"` //城市
	Sort      int       `json:"sort" form:"sort" gorm:"column:sort;comment:排序,倒序;"`
	IsDeleted int       `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zmAddress表 Address自定义表名 zm_address
func (Address) TableName() string {
	return "zm_address"
}
