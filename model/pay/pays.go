// 自动生成模板Pays
package pay

// zmPay表 结构体  Pays
type Pays struct {
	ID        int      `gorm:"primarykey" json:"ID"`                                                      // 主键ID
	Name      string   `json:"name" form:"name" gorm:"column:name;comment:名称;size:256;"`                  //名称
	CPrice    *float64 `json:"cPrice" form:"cPrice" gorm:"column:c_price;comment:现价;size:10;"`            //现价
	OPrice    *float64 `json:"oPrice" form:"oPrice" gorm:"column:o_price;comment:原价;size:10;"`            //原价
	Number    *int     `json:"number" form:"number" gorm:"column:number;comment:有效天数;size:10;"`           //有效天数
	NumberExt *int     `json:"numberExt" form:"numberExt" gorm:"column:number_ext;comment:赠送天数;size:10;"` //赠送天数
	Sort      *int     `json:"sort" form:"sort" gorm:"column:sort;comment:排序;size:10;"`
	Checked   *bool    `json:"checked" form:"checked" gorm:"column:checked;comment:默认状态;size:10;"`           //默认状态
	Status    *bool    `json:"status" form:"status" gorm:"column:status;comment:启用状态;size:10;"`              //启用状态
	Type      string   `json:"type" form:"type" gorm:"column:type;type:enum(256);comment:类型,1付费,2积分;"`       //类型,1付费,2积分
	IsDeleted int      `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zmPay表 Pays自定义表名 zm_pay
func (Pays) TableName() string {
	return "zm_pay"
}
