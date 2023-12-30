package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/order"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type OrdersApi struct {
}

var ordersService = service.ServiceGroupApp.OrderServiceGroup.OrdersService


// CreateOrders 创建zmOrder表
// @Tags Orders
// @Summary 创建zmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.Orders true "创建zmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /orders/createOrders [post]
func (ordersApi *OrdersApi) CreateOrders(c *gin.Context) {
	var orders order.Orders
	err := c.ShouldBindJSON(&orders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "UserId":{utils.NotEmpty()},
        "OrderId":{utils.NotEmpty()},
        "Type":{utils.NotEmpty()},
        "CPrice":{utils.NotEmpty()},
        "Number":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
        "PayTime":{utils.NotEmpty()},
    }
	if err := utils.Verify(orders, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := ordersService.CreateOrders(&orders); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrders 删除zmOrder表
// @Tags Orders
// @Summary 删除zmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.Orders true "删除zmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /orders/deleteOrders [delete]
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
// @Tags Orders
// @Summary 批量删除zmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /orders/deleteOrdersByIds [delete]
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

// UpdateOrders 更新zmOrder表
// @Tags Orders
// @Summary 更新zmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.Orders true "更新zmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /orders/updateOrders [put]
func (ordersApi *OrdersApi) UpdateOrders(c *gin.Context) {
	var orders order.Orders
	err := c.ShouldBindJSON(&orders)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "UserId":{utils.NotEmpty()},
          "OrderId":{utils.NotEmpty()},
          "Type":{utils.NotEmpty()},
          "CPrice":{utils.NotEmpty()},
          "Number":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
          "PayTime":{utils.NotEmpty()},
      }
    if err := utils.Verify(orders, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := ordersService.UpdateOrders(orders); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrders 用id查询zmOrder表
// @Tags Orders
// @Summary 用id查询zmOrder表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query order.Orders true "用id查询zmOrder表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /orders/findOrders [get]
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
// @Tags Orders
// @Summary 分页获取zmOrder表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderReq.OrdersSearch true "分页获取zmOrder表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /orders/getOrdersList [get]
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
