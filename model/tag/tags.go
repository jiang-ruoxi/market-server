// 自动生成模板Tags
package tag

import (
	"gorm.io/gorm"
	"time"
)

// zm_tags表 结构体  Tags
type Tags struct {
	ID        int            `gorm:"primarykey" json:"ID"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdatedAt time.Time      // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`                                                // 删除时间
	Name      string         `json:"name" form:"name" gorm:"column:name;comment:icon名称;size:256;"`  //icon名称
	Icon      string         `json:"icon" form:"icon" gorm:"column:icon;comment:icon链接;size:1024;"` //icon链接
	Status    *bool          `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;"` //状态,1启用,0禁用
}

// TableName zm_tags表 Tags自定义表名 zm_tags
func (Tags) TableName() string {
	return "zm_tags"
}
