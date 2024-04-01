// 自动生成模板Members
package member

import "time"

// zmUser表 结构体  Members
type Members struct {
	ID          int       `gorm:"primarykey" json:"ID"`                                                              // 主键ID
	UserId      string    `json:"userId" form:"userId" gorm:"column:user_id;comment:用户UserId;size:64;"`              //用户UserId
	OpenId      string    `json:"openId" form:"openId" gorm:"column:open_id;comment:用户OpendId;size:64;"`             //用户OpendId
	NickName    string    `json:"nickName" form:"nickName" gorm:"column:nick_name;comment:用户昵称;size:64;"`            //用户昵称
	HeadUrl     string    `json:"headUrl" form:"headUrl" gorm:"column:head_url;comment:用户头像;size:255;"`              //用户头像
	RealName    string    `json:"realName" form:"realName" gorm:"column:real_name;comment:用户姓名;size:255;"`           //用户姓名
	Mobile      string    `json:"mobile" form:"mobile" gorm:"column:mobile;comment:手机号;"`                            //手机号
	IsBest      *bool     `json:"isBest" form:"isBest" gorm:"column:is_best;comment:优选工匠,1是,0否;"`                    //优选工匠,1是,0否
	BestLimit   *int      `json:"bestLimit" form:"bestLimit" gorm:"column:best_limit;comment:优选工匠截止日期;size:10;"`     //优选工匠截止日期
	TagId       *int      `json:"tagId" form:"tagId" gorm:"column:tag_id;comment:主营类型;size:10;"`                     //主营类型
	ParentId    *int      `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:邀请人;size:19;"`             //邀请人
	IsMember    *bool     `json:"isMember" form:"isMember" gorm:"column:is_member;comment:是否为会员,1是,0否;"`             //是否为会员,1是,0否
	MemberLimit *int      `json:"memberLimit" form:"memberLimit" gorm:"column:member_limit;comment:会员截止日期;size:10;"` //会员截止日期
	TagName     string    `json:"tag_name" gorm:"-"`
	CreatedAt   time.Time // 创建时间
	IsDeleted   int       `json:"is_deleted" form:"is_deleted" gorm:"column:is_deleted;comment:是否删除,1已删除,0正常;"` //是否删除,1已删除,0正常
}

// TableName zmUser表 Members自定义表名 zm_user
func (Members) TableName() string {
	return "zm_user"
}
