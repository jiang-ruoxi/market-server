package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
)

type OrdersService struct {
}


// DeleteOrders 删除zmOrder表记录
func (ordersService *OrdersService)DeleteOrders(orders order.Orders) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&orders).Error
	return err
}

// DeleteOrdersByIds 批量删除zmOrder表记录
func (ordersService *OrdersService)DeleteOrdersByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]order.Orders{},"id in ?",ids.Ids).Error
	return err
}

// GetOrders 根据id获取zmOrder表记录
func (ordersService *OrdersService)GetOrders(id uint) (orders order.Orders, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&orders).Error
	return
}

// GetOrdersInfoList 分页获取zmOrder表记录
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
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }
	
	err = db.Order("id desc").Find(&orderss).Error
	return  orderss, total, err
}
