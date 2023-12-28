package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrdersRouter struct {
}

// InitOrdersRouter 初始化 zm_order表 路由信息
func (s *OrdersRouter) InitOrdersRouter(Router *gin.RouterGroup) {
	ordersRouter := Router.Group("orders").Use(middleware.OperationRecord())
	ordersRouterWithoutRecord := Router.Group("orders")
	var ordersApi = v1.ApiGroupApp.OrderApiGroup.OrdersApi
	{
		ordersRouter.POST("createOrders", ordersApi.CreateOrders)   // 新建zm_order表
		ordersRouter.DELETE("deleteOrders", ordersApi.DeleteOrders) // 删除zm_order表
		ordersRouter.DELETE("deleteOrdersByIds", ordersApi.DeleteOrdersByIds) // 批量删除zm_order表
		ordersRouter.PUT("updateOrders", ordersApi.UpdateOrders)    // 更新zm_order表
	}
	{
		ordersRouterWithoutRecord.GET("findOrders", ordersApi.FindOrders)        // 根据ID获取zm_order表
		ordersRouterWithoutRecord.GET("getOrdersList", ordersApi.GetOrdersList)  // 获取zm_order表列表
	}
}
