// 自动生成模板Orders
package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
	
)

// zm_order表 结构体  Orders
type Orders struct {
      global.GVA_MODEL
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"`  //用户id 
      OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单;size:19;"`  //订单 
      Type  *bool `json:"type" form:"type" gorm:"column:type;comment:类型,1普通会员,2优选工匠,3积分兑换;"`  //类型,1普通会员,2优选工匠,3积分兑换 
      CPrice  *int `json:"cPrice" form:"cPrice" gorm:"column:c_price;comment:现价;size:10;"`  //现价 
      OPrice  *int `json:"oPrice" form:"oPrice" gorm:"column:o_price;comment:原价;size:10;"`  //原价 
      Number  *int `json:"number" form:"number" gorm:"column:number;comment:有效天数;size:10;"`  //有效天数 
      NumberExt  *int `json:"numberExt" form:"numberExt" gorm:"column:number_ext;comment:赠送天数;size:10;"`  //赠送天数 
      Status  *bool `json:"status" form:"status" gorm:"column:status;comment:支付状态,1支付完成,0待支付;"`  //支付状态,1支付完成,0待支付 
}


// TableName zm_order表 Orders自定义表名 zm_order
func (Orders) TableName() string {
  return "zm_order"
}

