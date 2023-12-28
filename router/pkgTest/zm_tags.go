package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ZmTagsRouter struct {
}

// InitZmTagsRouter 初始化 zmTags表 路由信息
func (s *ZmTagsRouter) InitZmTagsRouter(Router *gin.RouterGroup) {
	zmTagsRouter := Router.Group("zmTags").Use(middleware.OperationRecord())
	zmTagsRouterWithoutRecord := Router.Group("zmTags")
	var zmTagsApi = v1.ApiGroupApp.PkgTestApiGroup.ZmTagsApi
	{
		zmTagsRouter.POST("createZmTags", zmTagsApi.CreateZmTags)   // 新建zmTags表
		zmTagsRouter.DELETE("deleteZmTags", zmTagsApi.DeleteZmTags) // 删除zmTags表
		zmTagsRouter.DELETE("deleteZmTagsByIds", zmTagsApi.DeleteZmTagsByIds) // 批量删除zmTags表
		zmTagsRouter.PUT("updateZmTags", zmTagsApi.UpdateZmTags)    // 更新zmTags表
	}
	{
		zmTagsRouterWithoutRecord.GET("findZmTags", zmTagsApi.FindZmTags)        // 根据ID获取zmTags表
		zmTagsRouterWithoutRecord.GET("getZmTagsList", zmTagsApi.GetZmTagsList)  // 获取zmTags表列表
	}
}
