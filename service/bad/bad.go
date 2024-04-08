package bad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bad"
	badReq "github.com/flipped-aurora/gin-vue-admin/server/model/bad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BadWordsService struct {
}

// CreateBadWords 创建bad记录
func (badWordsService *BadWordsService) CreateBadWords(badWords *bad.BadWords) (err error) {
	err = global.MustGetGlobalDBByDBName("api").Create(badWords).Error
	return err
}

// DeleteBadWords 删除bad记录
func (badWordsService *BadWordsService)DeleteBadWords(badWords bad.BadWords) (err error) {
	var s bad.BadWords
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id=?", badWords.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteBadWordsByIds 批量删除bad记录
func (badWordsService *BadWordsService)DeleteBadWordsByIds(ids request.IdsReq) (err error) {
	var s bad.BadWords
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&bad.BadWords{IsDeleted: 1}).Error
	return err
}

// UpdateBadWords 更新bad记录
func (badWordsService *BadWordsService)UpdateBadWords(badWords bad.BadWords) (err error) {
	err = global.MustGetGlobalDBByDBName("api").Save(&badWords).Error
	return err
}

// GetBadWords 根据id获取bad记录
func (badWordsService *BadWordsService)GetBadWords(id int) (badWords bad.BadWords, err error) {
	err = global.MustGetGlobalDBByDBName("api").Where("id = ?", id).First(&badWords).Error
	return
}

// GetBadWordsInfoList 分页获取bad记录
func (badWordsService *BadWordsService)GetBadWordsInfoList(info badReq.BadWordsSearch) (list []bad.BadWords, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("api").Model(&bad.BadWords{}).Where("is_deleted = 0")
    var badWordss []bad.BadWords
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
