package bad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BadWordsRouter struct {
}

// InitBadWordsRouter 初始化 bad 路由信息
func (s *BadWordsRouter) InitBadWordsRouter(Router *gin.RouterGroup) {
	badWordsRouter := Router.Group("badWords").Use(middleware.OperationRecord())
	badWordsRouterWithoutRecord := Router.Group("badWords")
	var badWordsApi = v1.ApiGroupApp.BadApiGroup.BadWordsApi
	{
		badWordsRouter.POST("createBadWords", badWordsApi.CreateBadWords)   // 新建bad
		badWordsRouter.DELETE("deleteBadWords", badWordsApi.DeleteBadWords) // 删除bad
		badWordsRouter.DELETE("deleteBadWordsByIds", badWordsApi.DeleteBadWordsByIds) // 批量删除bad
		badWordsRouter.PUT("updateBadWords", badWordsApi.UpdateBadWords)    // 更新bad
	}
	{
		badWordsRouterWithoutRecord.GET("findBadWords", badWordsApi.FindBadWords)        // 根据ID获取bad
		badWordsRouterWithoutRecord.GET("getBadWordsList", badWordsApi.GetBadWordsList)  // 获取bad列表
	}
}
