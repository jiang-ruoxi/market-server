// 自动生成模板BadWords
package bad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// bad 结构体  BadWords
type BadWords struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:敏感词;size:255;"`  //敏感词 
}


// TableName bad BadWords自定义表名 zm_bad_words
func (BadWords) TableName() string {
  return "zm_bad_words"
}

