package tag

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tag"
	tagReq "github.com/flipped-aurora/gin-vue-admin/server/model/tag/request"
)

type TagsService struct {
}

// CreateTags 创建zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService) CreateTags(tags *tag.Tags) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(tags).Error
	return err
}

// DeleteTags 删除zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService)DeleteTags(tags tag.Tags) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&tags).Error
	return err
}

// DeleteTagsByIds 批量删除zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService)DeleteTagsByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]tag.Tags{},"id in ?",ids.Ids).Error
	return err
}

// UpdateTags 更新zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService)UpdateTags(tags tag.Tags) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Save(&tags).Error
	return err
}

// GetTags 根据id获取zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService)GetTags(id int) (tags tag.Tags, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&tags).Error
	return
}

// GetTagsInfoList 分页获取zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService)GetTagsInfoList(info tagReq.TagsSearch) (list []tag.Tags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{})
    var tagss []tag.Tags
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
	
	err = db.Find(&tagss).Error
	return  tagss, total, err
}

// GetTagsInfoList 分页获取zm_tags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tagsService *TagsService)GetTagsInfoListAll() (list []tag.Tags, err error) {
	db := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{})
	var tagss []tag.Tags

	err = db.Find(&tagss).Error
	return  tagss, err
}