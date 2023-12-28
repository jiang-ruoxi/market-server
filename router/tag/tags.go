package tag

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TagsRouter struct {
}

// InitTagsRouter 初始化 zm_tags表 路由信息
func (s *TagsRouter) InitTagsRouter(Router *gin.RouterGroup) {
	tagsRouter := Router.Group("tags").Use(middleware.OperationRecord())
	tagsRouterWithoutRecord := Router.Group("tags")
	var tagsApi = v1.ApiGroupApp.TagApiGroup.TagsApi
	{
		tagsRouter.POST("createTags", tagsApi.CreateTags)   // 新建zm_tags表
		tagsRouter.DELETE("deleteTags", tagsApi.DeleteTags) // 删除zm_tags表
		tagsRouter.DELETE("deleteTagsByIds", tagsApi.DeleteTagsByIds) // 批量删除zm_tags表
		tagsRouter.PUT("updateTags", tagsApi.UpdateTags)    // 更新zm_tags表
	}
	{
		tagsRouterWithoutRecord.GET("findTags", tagsApi.FindTags)        // 根据ID获取zm_tags表
		tagsRouterWithoutRecord.GET("getTagsList", tagsApi.GetTagsList)  // 获取zm_tags表列表
	}
}
