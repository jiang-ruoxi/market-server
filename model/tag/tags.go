// 自动生成模板Tags
package tag

// zm_tags表 结构体  Tags
type Tags struct {
	ID        int    `gorm:"primarykey" json:"ID"`                                                         // 主键ID
	Name      string `json:"name" form:"name" gorm:"column:name;comment:icon名称;size:256;"`                 //icon名称
	Icon      string `json:"icon" form:"icon" gorm:"column:icon;comment:icon链接;size:1024;"`                //icon链接
	Status    *bool  `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;"`                //状态,1启用,0禁用
	IsDeleted int    `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zm_tags表 Tags自定义表名 zm_tags
func (Tags) TableName() string {
	return "zm_tags"
}
