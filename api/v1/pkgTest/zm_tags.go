package pkgTest

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    pkgTestReq "github.com/flipped-aurora/gin-vue-admin/server/model/pkgTest/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type ZmTagsApi struct {
}

var zmTagsService = service.ServiceGroupApp.PkgTestServiceGroup.ZmTagsService


// CreateZmTags 创建zmTags表
// @Tags ZmTags
// @Summary 创建zmTags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.ZmTags true "创建zmTags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /zmTags/createZmTags [post]
func (zmTagsApi *ZmTagsApi) CreateZmTags(c *gin.Context) {
	var zmTags pkgTest.ZmTags
	err := c.ShouldBindJSON(&zmTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Icon":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(zmTags, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := zmTagsService.CreateZmTags(&zmTags); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteZmTags 删除zmTags表
// @Tags ZmTags
// @Summary 删除zmTags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.ZmTags true "删除zmTags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /zmTags/deleteZmTags [delete]
func (zmTagsApi *ZmTagsApi) DeleteZmTags(c *gin.Context) {
	var zmTags pkgTest.ZmTags
	err := c.ShouldBindJSON(&zmTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zmTagsService.DeleteZmTags(zmTags); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteZmTagsByIds 批量删除zmTags表
// @Tags ZmTags
// @Summary 批量删除zmTags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zmTags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /zmTags/deleteZmTagsByIds [delete]
func (zmTagsApi *ZmTagsApi) DeleteZmTagsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := zmTagsService.DeleteZmTagsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateZmTags 更新zmTags表
// @Tags ZmTags
// @Summary 更新zmTags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body pkgTest.ZmTags true "更新zmTags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /zmTags/updateZmTags [put]
func (zmTagsApi *ZmTagsApi) UpdateZmTags(c *gin.Context) {
	var zmTags pkgTest.ZmTags
	err := c.ShouldBindJSON(&zmTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Icon":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(zmTags, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := zmTagsService.UpdateZmTags(zmTags); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindZmTags 用id查询zmTags表
// @Tags ZmTags
// @Summary 用id查询zmTags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTest.ZmTags true "用id查询zmTags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /zmTags/findZmTags [get]
func (zmTagsApi *ZmTagsApi) FindZmTags(c *gin.Context) {
	var zmTags pkgTest.ZmTags
	err := c.ShouldBindQuery(&zmTags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rezmTags, err := zmTagsService.GetZmTags(zmTags.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rezmTags": rezmTags}, c)
	}
}

// GetZmTagsList 分页获取zmTags表列表
// @Tags ZmTags
// @Summary 分页获取zmTags表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query pkgTestReq.ZmTagsSearch true "分页获取zmTags表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /zmTags/getZmTagsList [get]
func (zmTagsApi *ZmTagsApi) GetZmTagsList(c *gin.Context) {
	var pageInfo pkgTestReq.ZmTagsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := zmTagsService.GetZmTagsInfoList(pageInfo); err != nil {
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
