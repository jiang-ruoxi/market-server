// 自动生成模板Orders
package order

import (
	"time"
)

// zmOrder表 结构体  Orders
type Orders struct {
	ID            int       `gorm:"primarykey" json:"ID"`                                                          // 主键ID
	UserId        *int      `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;"`                      //用户id
	OrderId       *int      `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单;"`                     //订单
	Name          string    `json:"name" form:"name" gorm:"column:name;comment:订单名称"`                              //订单名称
	PaymentNumber string    `json:"paymentNumber" form:"paymentNumber" gorm:"column:payment_number;comment:支付流水号"` //支付流水号
	Type          string    `json:"type" form:"type" gorm:"column:type;type:enum();comment:类型,1普通会员,2优选工匠,3积分兑换;"` //类型,1普通会员,2优选工匠,3积分兑换
	CPrice        *float64  `json:"cPrice" form:"cPrice" gorm:"column:c_price;comment:现价;"`                        //现价
	OPrice        *float64  `json:"oPrice" form:"oPrice" gorm:"column:o_price;comment:原价;"`                        //原价
	Number        *int      `json:"number" form:"number" gorm:"column:number;comment:有效天数;"`                       //有效天数
	NumberExt     *int      `json:"numberExt" form:"numberExt" gorm:"column:number_ext;comment:赠送天数;"`             //赠送天数
	Status        *int      `json:"status" form:"status" gorm:"column:status;comment:支付状态,1支付完成,0待支付;"`            //支付状态,1支付完成,0待支付
	PayTime       string    `json:"payTime" form:"payTime" gorm:"column:pay_time;comment:支付时间;size:19;"`           //支付时间
	CreatedAt     time.Time // 创建时间
	IsDeleted     int       `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zmOrder表 Orders自定义表名 zm_order
func (Orders) TableName() string {
	return "zm_order"
}
