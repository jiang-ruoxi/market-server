package pay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PaysRouter struct {
}

// InitPaysRouter 初始化 zmPay表 路由信息
func (s *PaysRouter) InitPaysRouter(Router *gin.RouterGroup) {
	paysRouter := Router.Group("pays").Use(middleware.OperationRecord())
	paysRouterWithoutRecord := Router.Group("pays")
	var paysApi = v1.ApiGroupApp.PayApiGroup.PaysApi
	{
		paysRouter.POST("createPays", paysApi.CreatePays)   // 新建zmPay表
		paysRouter.DELETE("deletePays", paysApi.DeletePays) // 删除zmPay表
		paysRouter.DELETE("deletePaysByIds", paysApi.DeletePaysByIds) // 批量删除zmPay表
		paysRouter.PUT("updatePays", paysApi.UpdatePays)    // 更新zmPay表
	}
	{
		paysRouterWithoutRecord.GET("findPays", paysApi.FindPays)        // 根据ID获取zmPay表
		paysRouterWithoutRecord.GET("getPaysList", paysApi.GetPaysList)  // 获取zmPay表列表
	}
}
