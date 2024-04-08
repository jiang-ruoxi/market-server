package banner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/banner"
	bannerReq "github.com/flipped-aurora/gin-vue-admin/server/model/banner/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"strconv"
)

type BannersService struct {
}

// CreateBanners 创建zm_banner表记录
func (bannersService *BannersService) CreateBanners(banners *banner.Banners) (err error) {
	err = global.MustGetGlobalDBByDBName("api").Create(banners).Error
	return err
}

// DeleteBanners 删除zm_banner表记录
func (bannersService *BannersService) DeleteBanners(banners banner.Banners) (err error) {
	var s banner.Banners
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id=?", banners.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteBannersByIds 批量删除zm_banner表记录
func (bannersService *BannersService) DeleteBannersByIds(ids request.IdsReq) (err error) {
	var s banner.Banners
	err = global.MustGetGlobalDBByDBName("api").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&banner.Banners{IsDeleted: 1}).Error
	return err
}

// UpdateBanners 更新zm_banner表记录
func (bannersService *BannersService) UpdateBanners(banners banner.Banners) (err error) {
	err = global.MustGetGlobalDBByDBName("api").Where("id=?", banners.ID).Save(&banners).Error
	return err
}

// GetBanners 根据id获取zm_banner表记录
func (bannersService *BannersService) GetBanners(id int) (bannerInfo banner.BannersResponse, err error) {
	var banner banner.Banners
	err = global.MustGetGlobalDBByDBName("api").Where("id = ?", id).First(&banner).Error
	bannerInfo.ID = banner.ID
	bannerInfo.Name = banner.Name
	bannerInfo.Image = banner.Image

	statusBool := false
	if banner.Status != 0 {
		statusBool = true
	}

	bannerInfo.Status = &statusBool
	bannerInfo.Type = strconv.Itoa(banner.Type)
	return
}

// GetBannersInfoList 分页获取zm_banner表记录
func (bannersService *BannersService) GetBannersInfoList(info bannerReq.BannersSearch) (list []banner.Banners, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("api").Model(&banner.Banners{}).Debug().Where("is_deleted = 0")
	var bannerss []banner.Banners
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&bannerss).Error
	return bannerss, total, err
}
