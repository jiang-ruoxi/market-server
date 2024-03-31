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
	err = global.MustGetGlobalDBByDBName("market").Create(pays).Error
	return err
}

// DeletePays 删除zmPay表记录
func (paysService *PaysService) DeletePays(pays pay.Pays) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&pays).Error
	return err
}

// DeletePaysByIds 批量删除zmPay表记录
func (paysService *PaysService) DeletePaysByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]pay.Pays{}, "id in ?", ids.Ids).Error
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
	err = global.MustGetGlobalDBByDBName("market").Save(&pays).Error
	return err
}

// GetPays 根据id获取zmPay表记录
func (paysService *PaysService) GetPays(id uint) (pays pay.Pays, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&pays).Error
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
	db := global.MustGetGlobalDBByDBName("market").Model(&pay.Pays{})
	var payss []pay.Pays
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
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
