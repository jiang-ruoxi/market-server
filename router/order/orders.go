package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type OrdersRouter struct {
}

// InitOrdersRouter 初始化 zmOrder表 路由信息
func (s *OrdersRouter) InitOrdersRouter(Router *gin.RouterGroup) {
	ordersRouter := Router.Group("orders").Use(middleware.OperationRecord())
	ordersRouterWithoutRecord := Router.Group("orders")
	var ordersApi = v1.ApiGroupApp.OrderApiGroup.OrdersApi
	{
		ordersRouter.DELETE("deleteOrders", ordersApi.DeleteOrders)           // 删除zmOrder表
		ordersRouter.DELETE("deleteOrdersByIds", ordersApi.DeleteOrdersByIds) // 批量删除zmOrder表
	}
	{
		ordersRouter.GET("refundOrders", ordersApi.RefundOrders)             // 退费zm_tags表
		ordersRouterWithoutRecord.GET("findOrders", ordersApi.FindOrders)       // 根据ID获取zmOrder表
		ordersRouterWithoutRecord.GET("getOrdersList", ordersApi.GetOrdersList) // 获取zmOrder表列表
	}
}
