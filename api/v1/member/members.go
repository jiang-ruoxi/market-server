package member

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/member"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    memberReq "github.com/flipped-aurora/gin-vue-admin/server/model/member/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type MembersApi struct {
}

var membersService = service.ServiceGroupApp.MemberServiceGroup.MembersService


// CreateMembers 创建zmUser表
// @Tags Members
// @Summary 创建zmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Members true "创建zmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /members/createMembers [post]
func (membersApi *MembersApi) CreateMembers(c *gin.Context) {
	var members member.Members
	err := c.ShouldBindJSON(&members)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "OpenId":{utils.NotEmpty()},
        "NickName":{utils.NotEmpty()},
        "HeadUrl":{utils.NotEmpty()},
        "Mobile":{utils.NotEmpty()},
        "ParentId":{utils.NotEmpty()},
    }
	if err := utils.Verify(members, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := membersService.CreateMembers(&members); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMembers 删除zmUser表
// @Tags Members
// @Summary 删除zmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Members true "删除zmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /members/deleteMembers [delete]
func (membersApi *MembersApi) DeleteMembers(c *gin.Context) {
	var members member.Members
	err := c.ShouldBindJSON(&members)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := membersService.DeleteMembers(members); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMembersByIds 批量删除zmUser表
// @Tags Members
// @Summary 批量删除zmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /members/deleteMembersByIds [delete]
func (membersApi *MembersApi) DeleteMembersByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := membersService.DeleteMembersByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMembers 更新zmUser表
// @Tags Members
// @Summary 更新zmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body member.Members true "更新zmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /members/updateMembers [put]
func (membersApi *MembersApi) UpdateMembers(c *gin.Context) {
	var members member.Members
	err := c.ShouldBindJSON(&members)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "OpenId":{utils.NotEmpty()},
          "NickName":{utils.NotEmpty()},
          "HeadUrl":{utils.NotEmpty()},
          "Mobile":{utils.NotEmpty()},
          "ParentId":{utils.NotEmpty()},
      }
    if err := utils.Verify(members, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := membersService.UpdateMembers(members); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMembers 用id查询zmUser表
// @Tags Members
// @Summary 用id查询zmUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query member.Members true "用id查询zmUser表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /members/findMembers [get]
func (membersApi *MembersApi) FindMembers(c *gin.Context) {
	var members member.Members
	err := c.ShouldBindQuery(&members)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if remembers, err := membersService.GetMembers(members.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remembers": remembers}, c)
	}
}

// GetMembersList 分页获取zmUser表列表
// @Tags Members
// @Summary 分页获取zmUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query memberReq.MembersSearch true "分页获取zmUser表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /members/getMembersList [get]
func (membersApi *MembersApi) GetMembersList(c *gin.Context) {
	var pageInfo memberReq.MembersSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := membersService.GetMembersInfoList(pageInfo); err != nil {
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
