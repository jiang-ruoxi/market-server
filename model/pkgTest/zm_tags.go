// 自动生成模板ZmTags
package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zmTags表 结构体  ZmTags
type ZmTags struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:icon名称;size:256;"`  //icon名称 
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:icon链接;size:1024;"`  //icon链接 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;"`  //状态,1启用,0禁用 
      AddTime  *int `json:"addTime" form:"addTime" gorm:"column:add_time;comment:添加时间;size:19;"`  //添加时间 
}


// TableName zmTags表 ZmTags自定义表名 zm_tags
func (ZmTags) TableName() string {
  return "zm_tags"
}

