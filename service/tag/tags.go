package tag

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/banner"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tag"
	tagReq "github.com/flipped-aurora/gin-vue-admin/server/model/tag/request"
)

type TagsService struct {
}

// CreateTags 创建zm_tags表记录
func (tagsService *TagsService) CreateTags(tags *tag.Tags) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(tags).Error
	return err
}

// DeleteTags 删除zm_tags表记录
func (tagsService *TagsService) DeleteTags(tags tag.Tags) (err error) {
	var s tag.Tags
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id=?", tags.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteTagsByIds 批量删除zm_tags表记录
func (tagsService *TagsService) DeleteTagsByIds(ids request.IdsReq) (err error) {
	var s tag.Tags
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&banner.Banners{IsDeleted: 1}).Error
	return err
}

// UpdateTags 更新zm_tags表记录
func (tagsService *TagsService) UpdateTags(tags tag.Tags) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Save(&tags).Error
	return err
}

// GetTags 根据id获取zm_tags表记录
func (tagsService *TagsService) GetTags(id int) (tags tag.Tags, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&tags).Error
	return
}

// GetTagsInfoList 分页获取zm_tags表记录
func (tagsService *TagsService) GetTagsInfoList(info tagReq.TagsSearch) (list []tag.Tags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{}).Where("is_deleted = 0")
	var tagss []tag.Tags
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tagss).Error
	return tagss, total, err
}

// GetTagsInfoList 分页获取zm_tags表记录
func (tagsService *TagsService) GetTagsInfoListAll() (list []tag.Tags, err error) {
	db := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{})
	var tagss []tag.Tags

	err = db.Find(&tagss).Error
	return tagss, err
}
