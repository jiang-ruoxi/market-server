package order

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"io/ioutil"
	"net/http"
	"time"
)

type OrdersService struct {
}

//RefundOrders 退费操作
func (ordersService *OrdersService) RefundOrders(id int) (code int, message string) {
	//var s order.Orders
	fmt.Printf("退费的id:%#v \n", id)

	//根据订单id查询订单信息
	orderInfo, _ := ordersService.GetOrders(id)
	if orderInfo.ID > 0 {
		//判断有效期-根据类型 todo
		//此处直接发起退款操作
		code, message = ordersService.sendHttpPostData(*orderInfo.OrderId)
		return
	}
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
	TimeLocation, _ := time.LoadLocation("Asia/Shanghai") //指定时区
	dateTime := time.Unix(orders.RefundTime, 0).In(TimeLocation).Format("2006-01-02 15:04:05")
	orders.RefundTimeStr = dateTime
	return
}

// GetOrdersInfoList 分页获取zmOrder表记录
func (ordersService *OrdersService) GetOrdersInfoList(info orderReq.OrdersSearch) (list []order.Orders, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&order.Orders{}).Where("is_deleted = 0").Debug()
	var orderss []order.Orders
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.UserId != nil && *info.UserId > 0 {
		db = db.Where(" user_id =?", *info.UserId)
	}
	if info.Status != nil && *info.Status > -5 {
		db = db.Where(" status =?", *info.Status)
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(" pay_time BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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

//sendHttpPostData 发起请求
func (ordersService *OrdersService) sendHttpPostData(orderId int) (code int, message string) {
	// 要发送的数据
	postBody, _ := json.Marshal(map[string]int{
		"order_id": orderId,
	})

	// 将数据转换为字节序列
	requestBody := bytes.NewBuffer(postBody)

	// 创建请求
	refundUrl := global.GVA_CONFIG.WxPay.Refund
	req, err := http.NewRequest("POST", refundUrl, requestBody)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 设置请求头信息，指定内容类型为JSON
	req.Header.Set("Content-Type", "application/json")

	// 执行请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	type RefundsGenerated struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"data"`
	}
	var refund RefundsGenerated
	json.Unmarshal(body, &refund)

	fmt.Printf("%#v \n", refund)
	fmt.Printf("%#v \n", refund.Data.Code)
	fmt.Printf("%#v \n", refund.Data.Message)
	//{"code":10000,"msg":"success","data":{"code":200,"message":"success"}}
	fmt.Println(string(body))

	return refund.Data.Code, refund.Data.Message
}
