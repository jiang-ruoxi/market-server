package address

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AddressRouter struct {
}

// InitAddressRouter 初始化 zmAddress表 路由信息
func (s *AddressRouter) InitAddressRouter(Router *gin.RouterGroup) {
	zmAddressRouter := Router.Group("zmAddress").Use(middleware.OperationRecord())
	zmAddressRouterWithoutRecord := Router.Group("zmAddress")
	var zmAddressApi = v1.ApiGroupApp.AddressApiGroup.AddressApi
	{
		zmAddressRouter.POST("createAddress", zmAddressApi.CreateAddress)   // 新建zmAddress表
		zmAddressRouter.DELETE("deleteAddress", zmAddressApi.DeleteAddress) // 删除zmAddress表
		zmAddressRouter.DELETE("deleteAddressByIds", zmAddressApi.DeleteAddressByIds) // 批量删除zmAddress表
		zmAddressRouter.PUT("updateAddress", zmAddressApi.UpdateAddress)    // 更新zmAddress表
	}
	{
		zmAddressRouterWithoutRecord.GET("findAddress", zmAddressApi.FindAddress)        // 根据ID获取zmAddress表
		zmAddressRouterWithoutRecord.GET("getAddressList", zmAddressApi.GetAddressList)  // 获取zmAddress表列表
		zmAddressRouterWithoutRecord.GET("getAddressAllList", zmAddressApi.GetAddressAllList)  // 获取zmAddress表列表
		zmAddressRouterWithoutRecord.GET("getAddressChildList", zmAddressApi.GetAddressChildList)  // 获取zmAddress表列表
	}
}
