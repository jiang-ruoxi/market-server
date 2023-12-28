// 自动生成模板Banners
package banner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zm_banner表 结构体  Banners
type Banners struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:banner名称;size:256;"`  //banner名称 
      Image  string `json:"image" form:"image" gorm:"column:image;comment:banner链接;size:1024;"`  //banner链接 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;"`  //状态,1启用,0禁用 
}


// TableName zm_banner表 Banners自定义表名 zm_banner
func (Banners) TableName() string {
  return "zm_banner"
}

