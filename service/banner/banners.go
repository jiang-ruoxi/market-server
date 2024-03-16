package banner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/banner"
	bannerReq "github.com/flipped-aurora/gin-vue-admin/server/model/banner/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BannersService struct {
}

// CreateBanners 创建zm_banner表记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannersService *BannersService) CreateBanners(banners *banner.Banners) (err error) {
	//banners.Image = "https://static.58haha.com/" + banners.Image
	err = global.MustGetGlobalDBByDBName("market").Create(banners).Error
	return err
}

// DeleteBanners 删除zm_banner表记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannersService *BannersService)DeleteBanners(banners banner.Banners) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&banners).Error
	return err
}

// DeleteBannersByIds 批量删除zm_banner表记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannersService *BannersService)DeleteBannersByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]banner.Banners{},"id in ?",ids.Ids).Error
	return err
}

// UpdateBanners 更新zm_banner表记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannersService *BannersService)UpdateBanners(banners banner.Banners) (err error) {
	//banners.Image = "https://static.58haha.com/" + banners.Image
	err = global.MustGetGlobalDBByDBName("market").Save(&banners).Error
	return err
}

// GetBanners 根据id获取zm_banner表记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannersService *BannersService)GetBanners(id uint) (banners banner.Banners, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&banners).Error
	return
}

// GetBannersInfoList 分页获取zm_banner表记录
// Author [piexlmax](https://github.com/piexlmax)
func (bannersService *BannersService)GetBannersInfoList(info bannerReq.BannersSearch) (list []banner.Banners, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&banner.Banners{})
    var bannerss []banner.Banners
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
	
	err = db.Find(&bannerss).Error
	return  bannerss, total, err
}
