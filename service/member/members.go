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
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id=?", members.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteMembersByIds 批量删除zmUser表记录
func (membersService *MembersService) DeleteMembersByIds(ids request.IdsReq) (err error) {
	var s member.Members
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&member.Members{IsDeleted: 1}).Error
	return err
}

// GetMembers 根据id获取zmUser表记录
func (membersService *MembersService) GetMembers(id int) (members member.Members, err error) {
	err = global.MustGetGlobalDBByDBName("api").Where("id = ?", id).First(&members).Error

	var tagInfo tag.Tags
	db1 := global.MustGetGlobalDBByDBName("api").Model(&tag.Tags{})
	db1.Where("id=?", members.TagId).First(&tagInfo)

	members.TagName = tagInfo.Name
	return
}

// GetMembersInfoList 分页获取zmUser表记录
func (membersService *MembersService) GetMembersInfoList(info memberReq.MembersSearch) (list []member.Members, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("api").Model(&member.Members{}).Where("is_deleted = 0")
	var memberss []member.Members
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != "" {
		db = db.Where("user_id =?", info.UserId)
	}
	if info.Mobile != "" {
		db = db.Where("mobile =?", info.Mobile)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(" created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&memberss).Error

	for idx,_:= range memberss{
		if memberss[idx].HeadUrl == "" {
			memberss[idx].HeadUrl = "https://oss.58haha.com/style/market/default.png"
		}
		if memberss[idx].NickName == "" {
			memberss[idx].NickName = "-"
		}
	}
	return memberss, total, err
}
