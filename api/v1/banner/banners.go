package banner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/banner"
	bannerReq "github.com/flipped-aurora/gin-vue-admin/server/model/banner/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type BannersApi struct {
}

var bannersService = service.ServiceGroupApp.BannerServiceGroup.BannersService

// CreateBanners 创建zm_banner表
func (bannersApi *BannersApi) CreateBanners(c *gin.Context) {
	var request banner.BannersRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Image":  {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
	}
	if err := utils.Verify(request, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	statusInt := 0
	if *request.Status == true {
		statusInt = 1
	}
	var banners banner.Banners
	banners.Name = request.Name
	banners.Image = request.Image
	banners.Status = statusInt
	typeInt, _ := strconv.Atoi(request.Type)
	banners.Type = typeInt
	if err := bannersService.CreateBanners(&banners); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBanners 删除zm_banner表
func (bannersApi *BannersApi) DeleteBanners(c *gin.Context) {
	var banners banner.Banners
	err := c.ShouldBindJSON(&banners)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bannersService.DeleteBanners(banners); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBannersByIds 批量删除zm_banner表
func (bannersApi *BannersApi) DeleteBannersByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := bannersService.DeleteBannersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBanners 更新zm_banner表
func (bannersApi *BannersApi) UpdateBanners(c *gin.Context) {
	var request banner.BannersRequest
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":   {utils.NotEmpty()},
		"Image":  {utils.NotEmpty()},
		"Status": {utils.NotEmpty()},
		"Type":   {utils.NotEmpty()},
	}
	if err := utils.Verify(request, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	statusInt := 0
	if *request.Status == true {
		statusInt = 1
	}
	var banners banner.Banners
	banners.Name = request.Name
	banners.Image = request.Image
	banners.Status = statusInt
	typeInt, _ := strconv.Atoi(request.Type)
	banners.Type = typeInt
	banners.ID = request.ID
	if err := bannersService.UpdateBanners(banners); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBanners 用id查询zm_banner表
func (bannersApi *BannersApi) FindBanners(c *gin.Context) {
	var banners banner.Banners
	err := c.ShouldBindQuery(&banners)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebanners, err := bannersService.GetBanners(banners.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebanners": rebanners}, c)
	}
}

// GetBannersList 分页获取zm_banner表列表
func (bannersApi *BannersApi) GetBannersList(c *gin.Context) {
	var pageInfo bannerReq.BannersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := bannersService.GetBannersInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
