package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrdersApi struct {
}

var ordersService = service.ServiceGroupApp.OrderServiceGroup.OrdersService

// RefundOrders 退费操作
func (ordersApi *OrdersApi) RefundOrders(c *gin.Context) {
	var orders order.Orders
	err := c.ShouldBindQuery(&orders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorders, err := ordersService.RefundOrders(orders.ID); err != nil {
		global.GVA_LOG.Error("退费失败!", zap.Error(err))
		response.FailWithMessage("退费失败", c)
	} else {
		response.OkWithData(gin.H{"reorders": reorders}, c)
	}
}

// DeleteOrders 删除zmOrder表
func (ordersApi *OrdersApi) DeleteOrders(c *gin.Context) {
	var orders order.Orders
	err := c.ShouldBindJSON(&orders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ordersService.DeleteOrders(orders); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrdersByIds 批量删除zmOrder表
func (ordersApi *OrdersApi) DeleteOrdersByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ordersService.DeleteOrdersByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindOrders 用id查询zmOrder表
func (ordersApi *OrdersApi) FindOrders(c *gin.Context) {
	var orders order.Orders
	err := c.ShouldBindQuery(&orders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reorders, err := ordersService.GetOrders(orders.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reorders": reorders}, c)
	}
}

// GetOrdersList 分页获取zmOrder表列表
func (ordersApi *OrdersApi) GetOrdersList(c *gin.Context) {
	var pageInfo orderReq.OrdersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ordersService.GetOrdersInfoList(pageInfo); err != nil {
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
