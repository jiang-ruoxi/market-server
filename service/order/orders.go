package order

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
)

type OrdersService struct {
}

//RefundOrders 退费操作
func (ordersService *OrdersService) RefundOrders(id int) (orders order.Orders, err error) {
	//var s order.Orders
	fmt.Printf("退费的id:%#v \n", id)
	//这里发起http请求，请求api的退费接口
	return
}

// DeleteOrders 删除zmOrder表记录
func (ordersService *OrdersService) DeleteOrders(orders order.Orders) (err error) {
	var s order.Orders
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id=?", orders.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteOrdersByIds 批量删除zmOrder表记录
func (ordersService *OrdersService) DeleteOrdersByIds(ids request.IdsReq) (err error) {
	var s order.Orders
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&order.Orders{IsDeleted: 1}).Error
	return err
}

// GetOrders 根据id获取zmOrder表记录
func (ordersService *OrdersService) GetOrders(id int) (orders order.Orders, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&orders).Error

	payCPrice := *orders.CPrice / 100
	payOPrice := *orders.OPrice / 100
	orders.CPrice = &payCPrice
	orders.OPrice = &payOPrice
	return
}

// GetOrdersInfoList 分页获取zmOrder表记录
func (ordersService *OrdersService) GetOrdersInfoList(info orderReq.OrdersSearch) (list []order.Orders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&order.Orders{}).Where("is_deleted = 0")
	var orderss []order.Orders
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(" and created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&orderss).Error

	for idx, _ := range orderss {
		var payCPrice float64
		var payOPrice float64
		payCPrice = *orderss[idx].CPrice / 100
		payOPrice = *orderss[idx].OPrice / 100
		orderss[idx].CPrice = &payCPrice
		orderss[idx].OPrice = &payOPrice
	}

	return orderss, total, err
}
