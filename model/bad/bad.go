// 自动生成模板BadWords
package bad

import "time"

// bad 结构体  BadWords
type BadWords struct {
	ID        int       `gorm:"primarykey" json:"ID"`                                      // 主键ID
	Name      string    `json:"name" form:"name" gorm:"column:name;comment:敏感词;size:255;"` //敏感词
	CreatedAt time.Time // 创建时间
	IsDeleted int       `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName bad BadWords自定义表名 zm_bad_words
func (BadWords) TableName() string {
	return "zm_bad_words"
}
