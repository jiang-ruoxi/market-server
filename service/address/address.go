package address

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/address"
	addressReq "github.com/flipped-aurora/gin-vue-admin/server/model/address/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type AddressService struct {
}

// CreateAddress 创建zmAddress表记录
func (zmAddressService *AddressService) CreateAddress(zmAddress *address.Address) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(zmAddress).Error
	return err
}

// DeleteAddress 删除zmAddress表记录
func (zmAddressService *AddressService) DeleteAddress(zmAddress address.Address) (err error) {
	var s address.Address
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id=?", zmAddress.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteAddressByIds 批量删除zmAddress表记录
func (zmAddressService *AddressService) DeleteAddressByIds(ids request.IdsReq) (err error) {
	var s address.Address
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&address.Address{IsDeleted: 1}).Error
	return err
}

// UpdateAddress 更新zmAddress表记录
func (zmAddressService *AddressService) UpdateAddress(zmAddress address.Address) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Save(&zmAddress).Error
	return err
}

// GetAddress 根据id获取zmAddress表记录
func (zmAddressService *AddressService) GetAddress(id int) (zmAddress address.Address, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&zmAddress).Error
	return
}

// GetAddressInfoList 分页获取zmAddress表记录
func (zmAddressService *AddressService) GetAddressInfoList(info addressReq.AddressSearch) (list []address.Address, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&address.Address{}).Where("is_deleted=0")
	var zmAddresss []address.Address
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&zmAddresss).Error
	return zmAddresss, total, err
}

// GetAddressAllList
func (zmAddressService *AddressService) GetAddressAllList() (list []address.Address, err error) {
	db := global.MustGetGlobalDBByDBName("market").Model(&address.Address{}).Debug().Where("is_deleted= 0 and parent_id = 0")
	err = db.Find(&list).Error
	return list, err
}
