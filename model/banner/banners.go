// 自动生成模板Banners
package banner

type BannersRequest struct {
	ID     int    `json:"ID"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status *bool  `json:"status"`
	Type   string `json:"type"`
}

type BannersResponse struct {
	ID     int    `json:"ID"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	Status *bool  `json:"status"`
	Type   string `json:"type"`
}

// zm_banner表 结构体  Banners
type Banners struct {
	ID        int    `gorm:"primarykey" json:"ID"`                                                         // 主键ID
	Name      string `json:"name" form:"name" gorm:"column:name;comment:banner名称;size:256;"`               //banner名称
	Image     string `json:"image" form:"image" gorm:"column:image;comment:banner链接;size:1024;"`           //banner链接
	Status    int    `json:"status" form:"status" gorm:"column:status;comment:状态,1启用,0禁用;"`                //状态,1启用,0禁用
	Type      int    `json:"type" form:"type" gorm:"column:type;comment:类型:1小程序,2APP,3小程序弹层,4APP弹层;"`      //状类型:1小程序,2APP,3小程序弹层,4APP弹层
	IsDeleted int    `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zm_banner表 Banners自定义表名 zm_banner
func (Banners) TableName() string {
	return "zm_banner"
}
