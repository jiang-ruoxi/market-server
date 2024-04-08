package pay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pay"
	payReq "github.com/flipped-aurora/gin-vue-admin/server/model/pay/request"
)

type PaysService struct {
}

// CreatePays 创建zmPay表记录
func (paysService *PaysService) CreatePays(pays *pay.Pays) (err error) {
	cPrice := pays.CPrice
	realCPrice := *cPrice * 100
	pays.CPrice = &realCPrice
	oPrice := pays.OPrice
	realOPrice := *oPrice * 100
	pays.OPrice = &realOPrice
	err = global.MustGetGlobalDBByDBName("api").Create(pays).Error
	return err
}

// DeletePays 删除zmPay表记录
func (paysService *PaysService) DeletePays(pays pay.Pays) (err error) {
	var s pay.Pays
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id=?", pays.ID).Update("is_deleted", 1).Error
	return err
}

// DeletePaysByIds 批量删除zmPay表记录
func (paysService *PaysService) DeletePaysByIds(ids request.IdsReq) (err error) {
	var s pay.Pays
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&pay.Pays{IsDeleted: 1}).Error
	return err
}

// UpdatePays 更新zmPay表记录
func (paysService *PaysService) UpdatePays(pays pay.Pays) (err error) {
	cPrice := pays.CPrice
	realCPrice := *cPrice * 100
	pays.CPrice = &realCPrice
	oPrice := pays.OPrice
	realOPrice := *oPrice * 100
	pays.OPrice = &realOPrice
	err = global.MustGetGlobalDBByDBName("api").Save(&pays).Error
	return err
}

// GetPays 根据id获取zmPay表记录
func (paysService *PaysService) GetPays(id int) (pays pay.Pays, err error) {
	err = global.MustGetGlobalDBByDBName("api").Where("id = ?", id).First(&pays).Error
	payCPrice := *pays.CPrice / 100
	payOPrice := *pays.OPrice / 100
	pays.CPrice = &payCPrice
	pays.OPrice = &payOPrice
	return
}

// GetPaysInfoList 分页获取zmPay表记录
func (paysService *PaysService) GetPaysInfoList(info payReq.PaysSearch) (list []pay.Pays, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("api").Model(&pay.Pays{}).Where("is_deleted = 0")
	var payss []pay.Pays
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&payss).Error

	for idx, _ := range payss {
		var payCPrice float64
		var payOPrice float64
		payCPrice = *payss[idx].CPrice / 100
		payOPrice = *payss[idx].OPrice / 100
		payss[idx].CPrice = &payCPrice
		payss[idx].OPrice = &payOPrice
	}
	return payss, total, err
}
