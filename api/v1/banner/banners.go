package banner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/banner"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    bannerReq "github.com/flipped-aurora/gin-vue-admin/server/model/banner/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type BannersApi struct {
}

var bannersService = service.ServiceGroupApp.BannerServiceGroup.BannersService


// CreateBanners 创建zm_banner表
// @Tags Banners
// @Summary 创建zm_banner表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body banner.Banners true "创建zm_banner表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /banners/createBanners [post]
func (bannersApi *BannersApi) CreateBanners(c *gin.Context) {
	var banners banner.Banners
	err := c.ShouldBindJSON(&banners)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Image":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(banners, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := bannersService.CreateBanners(&banners); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBanners 删除zm_banner表
// @Tags Banners
// @Summary 删除zm_banner表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body banner.Banners true "删除zm_banner表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /banners/deleteBanners [delete]
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
// @Tags Banners
// @Summary 批量删除zm_banner表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zm_banner表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /banners/deleteBannersByIds [delete]
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
// @Tags Banners
// @Summary 更新zm_banner表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body banner.Banners true "更新zm_banner表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /banners/updateBanners [put]
func (bannersApi *BannersApi) UpdateBanners(c *gin.Context) {
	var banners banner.Banners
	err := c.ShouldBindJSON(&banners)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Image":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(banners, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := bannersService.UpdateBanners(banners); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBanners 用id查询zm_banner表
// @Tags Banners
// @Summary 用id查询zm_banner表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query banner.Banners true "用id查询zm_banner表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /banners/findBanners [get]
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
// @Tags Banners
// @Summary 分页获取zm_banner表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bannerReq.BannersSearch true "分页获取zm_banner表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /banners/getBannersList [get]
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
