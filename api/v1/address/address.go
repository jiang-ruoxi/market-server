package address

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/address"
	addressReq "github.com/flipped-aurora/gin-vue-admin/server/model/address/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AddressApi struct {
}

var zmAddressService = service.ServiceGroupApp.AddressServiceGroup.AddressService


// CreateAddress 创建zmAddress表
func (zmAddressApi *AddressApi) CreateAddress(c *gin.Context) {
	var zmAddress address.Address
	err := c.ShouldBindJSON(&zmAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zmAddressService.CreateAddress(&zmAddress); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAddress 删除zmAddress表
func (zmAddressApi *AddressApi) DeleteAddress(c *gin.Context) {
	var zmAddress address.Address
	err := c.ShouldBindJSON(&zmAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zmAddressService.DeleteAddress(zmAddress); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAddressByIds 批量删除zmAddress表
func (zmAddressApi *AddressApi) DeleteAddressByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zmAddressService.DeleteAddressByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAddress 更新zmAddress表
func (zmAddressApi *AddressApi) UpdateAddress(c *gin.Context) {
	var zmAddress address.Address
	err := c.ShouldBindJSON(&zmAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zmAddressService.UpdateAddress(zmAddress); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAddress 用id查询zmAddress表
func (zmAddressApi *AddressApi) FindAddress(c *gin.Context) {
	var zmAddress address.Address
	err := c.ShouldBindQuery(&zmAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rezmAddress, err := zmAddressService.GetAddress(zmAddress.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezmAddress": rezmAddress}, c)
	}
}

// GetAddressList 分页获取zmAddress表列表
func (zmAddressApi *AddressApi) GetAddressList(c *gin.Context) {
	var pageInfo addressReq.AddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zmAddressService.GetAddressInfoList(pageInfo); err != nil {
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


func (zmAddressApi *AddressApi) GetAddressAllList(c *gin.Context) {
	if list, err := zmAddressService.GetAddressAllList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
		}, "获取成功", c)
	}
}

func (zmAddressApi *AddressApi) GetAddressChildList(c *gin.Context) {
	if list, err := zmAddressService.GetAddressChildList(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
		}, "获取成功", c)
	}
}