package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
)

type MembersService struct {
}

// CreateMembers 创建zmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (membersService *MembersService) CreateMembers(members *member.Members) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(members).Error
	return err
}

// DeleteMembers 删除zmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (membersService *MembersService)DeleteMembers(members member.Members) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&members).Error
	return err
}

// DeleteMembersByIds 批量删除zmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (membersService *MembersService)DeleteMembersByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]member.Members{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMembers 更新zmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (membersService *MembersService)UpdateMembers(members member.Members) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Save(&members).Error
	return err
}

// GetMembers 根据id获取zmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (membersService *MembersService)GetMembers(id uint) (members member.Members, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&members).Error
	return
}

// GetMembersInfoList 分页获取zmUser表记录
// Author [piexlmax](https://github.com/piexlmax)
func (membersService *MembersService)GetMembersInfoList(info memberReq.MembersSearch) (list []member.Members, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&member.Members{})
    var memberss []member.Members
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&memberss).Error
	return  memberss, total, err
}
