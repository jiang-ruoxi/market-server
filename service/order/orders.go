package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
)

type OrdersService struct {
}

// CreateOrders 创建zm_order表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService) CreateOrders(orders *order.Orders) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(orders).Error
	return err
}

// DeleteOrders 删除zm_order表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService)DeleteOrders(orders order.Orders) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&orders).Error
	return err
}

// DeleteOrdersByIds 批量删除zm_order表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService)DeleteOrdersByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]order.Orders{},"id in ?",ids.Ids).Error
	return err
}

// UpdateOrders 更新zm_order表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService)UpdateOrders(orders order.Orders) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Save(&orders).Error
	return err
}

// GetOrders 根据id获取zm_order表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService)GetOrders(id uint) (orders order.Orders, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&orders).Error
	return
}

// GetOrdersInfoList 分页获取zm_order表记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordersService *OrdersService)GetOrdersInfoList(info orderReq.OrdersSearch) (list []order.Orders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&order.Orders{})
    var orderss []order.Orders
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
    if info.OrderId != nil {
        db = db.Where("order_id = ?",info.OrderId)
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
    if info.Status != nil {
        db = db.Where("status = ?",info.Status)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Find(&orderss).Error
	return  orderss, total, err
}
