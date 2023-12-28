package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
)

type ZmTagsService struct {
}

// CreateZmTags 创建zmTags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zmTagsService *ZmTagsService) CreateZmTags(zmTags *pkgTest.ZmTags) (err error) {
	err = global.GVA_DB.Create(zmTags).Error
	return err
}

// DeleteZmTags 删除zmTags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zmTagsService *ZmTagsService)DeleteZmTags(zmTags pkgTest.ZmTags) (err error) {
	err = global.GVA_DB.Delete(&zmTags).Error
	return err
}

// DeleteZmTagsByIds 批量删除zmTags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zmTagsService *ZmTagsService)DeleteZmTagsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]pkgTest.ZmTags{},"id in ?",ids.Ids).Error
	return err
}

// UpdateZmTags 更新zmTags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zmTagsService *ZmTagsService)UpdateZmTags(zmTags pkgTest.ZmTags) (err error) {
	err = global.GVA_DB.Save(&zmTags).Error
	return err
}

// GetZmTags 根据id获取zmTags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zmTagsService *ZmTagsService)GetZmTags(id uint) (zmTags pkgTest.ZmTags, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&zmTags).Error
	return
}

// GetZmTagsInfoList 分页获取zmTags表记录
// Author [piexlmax](https://github.com/piexlmax)
func (zmTagsService *ZmTagsService)GetZmTagsInfoList(info pkgTestReq.ZmTagsSearch) (list []pkgTest.ZmTags, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&pkgTest.ZmTags{})
    var zmTagss []pkgTest.ZmTags
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
	
	err = db.Find(&zmTagss).Error
	return  zmTagss, total, err
}
