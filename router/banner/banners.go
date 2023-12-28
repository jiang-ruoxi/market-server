package banner

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BannersRouter struct {
}

// InitBannersRouter 初始化 zm_banner表 路由信息
func (s *BannersRouter) InitBannersRouter(Router *gin.RouterGroup) {
	bannersRouter := Router.Group("banners").Use(middleware.OperationRecord())
	bannersRouterWithoutRecord := Router.Group("banners")
	var bannersApi = v1.ApiGroupApp.BannerApiGroup.BannersApi
	{
		bannersRouter.POST("createBanners", bannersApi.CreateBanners)   // 新建zm_banner表
		bannersRouter.DELETE("deleteBanners", bannersApi.DeleteBanners) // 删除zm_banner表
		bannersRouter.DELETE("deleteBannersByIds", bannersApi.DeleteBannersByIds) // 批量删除zm_banner表
		bannersRouter.PUT("updateBanners", bannersApi.UpdateBanners)    // 更新zm_banner表
	}
	{
		bannersRouterWithoutRecord.GET("findBanners", bannersApi.FindBanners)        // 根据ID获取zm_banner表
		bannersRouterWithoutRecord.GET("getBannersList", bannersApi.GetBannersList)  // 获取zm_banner表列表
	}
}
