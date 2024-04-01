package bad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bad"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    badReq "github.com/flipped-aurora/gin-vue-admin/server/model/bad/request"
)

type BadWordsService struct {
}

// CreateBadWords 创建bad记录
// Author [piexlmax](https://github.com/piexlmax)
func (badWordsService *BadWordsService) CreateBadWords(badWords *bad.BadWords) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(badWords).Error
	return err
}

// DeleteBadWords 删除bad记录
// Author [piexlmax](https://github.com/piexlmax)
func (badWordsService *BadWordsService)DeleteBadWords(badWords bad.BadWords) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&badWords).Error
	return err
}

// DeleteBadWordsByIds 批量删除bad记录
// Author [piexlmax](https://github.com/piexlmax)
func (badWordsService *BadWordsService)DeleteBadWordsByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]bad.BadWords{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBadWords 更新bad记录
// Author [piexlmax](https://github.com/piexlmax)
func (badWordsService *BadWordsService)UpdateBadWords(badWords bad.BadWords) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Save(&badWords).Error
	return err
}

// GetBadWords 根据id获取bad记录
// Author [piexlmax](https://github.com/piexlmax)
func (badWordsService *BadWordsService)GetBadWords(id uint) (badWords bad.BadWords, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&badWords).Error
	return
}

// GetBadWordsInfoList 分页获取bad记录
// Author [piexlmax](https://github.com/piexlmax)
func (badWordsService *BadWordsService)GetBadWordsInfoList(info badReq.BadWordsSearch) (list []bad.BadWords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&bad.BadWords{})
    var badWordss []bad.BadWords
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
	
	err = db.Find(&badWordss).Error
	return  badWordss, total, err
}
