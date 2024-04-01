package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/member"
	memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tag"
)

type MembersService struct {
}

// DeleteMembers 删除zmUser表记录
func (membersService *MembersService) DeleteMembers(members member.Members) (err error) {
	var s member.Members
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id=?", members.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteMembersByIds 批量删除zmUser表记录
func (membersService *MembersService) DeleteMembersByIds(ids request.IdsReq) (err error) {
	var s member.Members
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&member.Members{IsDeleted: 1}).Error
	return err
}

// GetMembers 根据id获取zmUser表记录
func (membersService *MembersService) GetMembers(id int) (members member.Members, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&members).Error

	var tagInfo tag.Tags
	db1 := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{})
	db1.Where("id=?", members.TagId).First(&tagInfo)

	members.TagName = tagInfo.Name
	return
}

// GetMembersInfoList 分页获取zmUser表记录
func (membersService *MembersService) GetMembersInfoList(info memberReq.MembersSearch) (list []member.Members, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&member.Members{}).Where("is_deleted = 0")
	var memberss []member.Members
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(" and created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&memberss).Error
	return memberss, total, err
}
