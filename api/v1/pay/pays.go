package pay

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/pay"
	payReq "github.com/flipped-aurora/gin-vue-admin/server/model/pay/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PaysApi struct {
}

var paysService = service.ServiceGroupApp.PayServiceGroup.PaysService

// CreatePays 创建zmPay表
func (paysApi *PaysApi) CreatePays(c *gin.Context) {
	var pays pay.Pays
	err := c.ShouldBindJSON(&pays)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":    {utils.NotEmpty()},
		"CPrice":  {utils.NotEmpty()},
		"Number":  {utils.NotEmpty()},
		"Sort":    {utils.NotEmpty()},
		"Checked": {utils.NotEmpty()},
		"Status":  {utils.NotEmpty()},
		"Type":    {utils.NotEmpty()},
	}
	if err := utils.Verify(pays, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := paysService.CreatePays(&pays); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePays 删除zmPay表
func (paysApi *PaysApi) DeletePays(c *gin.Context) {
	var pays pay.Pays
	err := c.ShouldBindJSON(&pays)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := paysService.DeletePays(pays); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePaysByIds 批量删除zmPay表
func (paysApi *PaysApi) DeletePaysByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := paysService.DeletePaysByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePays 更新zmPay表
func (paysApi *PaysApi) UpdatePays(c *gin.Context) {
	var pays pay.Pays
	err := c.ShouldBindJSON(&pays)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	verify := utils.Rules{
		"Name":    {utils.NotEmpty()},
		"CPrice":  {utils.NotEmpty()},
		"Number":  {utils.NotEmpty()},
		"Sort":    {utils.NotEmpty()},
		"Checked": {utils.NotEmpty()},
		"Status":  {utils.NotEmpty()},
		"Type":    {utils.NotEmpty()},
	}
	if err := utils.Verify(pays, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := paysService.UpdatePays(pays); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPays 用id查询zmPay表
func (paysApi *PaysApi) FindPays(c *gin.Context) {
	var pays pay.Pays
	err := c.ShouldBindQuery(&pays)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repays, err := paysService.GetPays(pays.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repays": repays}, c)
	}
}

// GetPaysList 分页获取zmPay表列表
func (paysApi *PaysApi) GetPaysList(c *gin.Context) {
	var pageInfo payReq.PaysSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := paysService.GetPaysInfoList(pageInfo); err != nil {
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
