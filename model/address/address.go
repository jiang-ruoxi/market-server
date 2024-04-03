// 自动生成模板Address
package address

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zmAddress表 结构体  Address
type Address struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:城市;size:255;"`  //城市 
}


// TableName zmAddress表 Address自定义表名 zm_address
func (Address) TableName() string {
  return "zm_address"
}

