// 自动生成模板Pays
package pay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zmPay表 结构体  Pays
type Pays struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:名称;size:256;"`  //名称 
      CPrice  *int `json:"cPrice" form:"cPrice" gorm:"column:c_price;comment:现价;size:10;"`  //现价 
      OPrice  *int `json:"oPrice" form:"oPrice" gorm:"column:o_price;comment:原价;size:10;"`  //原价 
      Number  *int `json:"number" form:"number" gorm:"column:number;comment:有效天数;size:10;"`  //有效天数 
      NumberExt  *int `json:"numberExt" form:"numberExt" gorm:"column:number_ext;comment:赠送天数;size:10;"`  //赠送天数 
      Type  *bool `json:"type" form:"type" gorm:"column:type;comment:类型,1付费,2积分;"`  //类型,1付费,2积分 
}


// TableName zmPay表 Pays自定义表名 zm_pay
func (Pays) TableName() string {
  return "zm_pay"
}

