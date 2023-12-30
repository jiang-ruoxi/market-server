package tag

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tag"
	tagReq "github.com/flipped-aurora/gin-vue-admin/server/model/tag/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TagsApi struct {
}

var tagsService = service.ServiceGroupApp.TagServiceGroup.TagsService


// CreateTags 创建zm_tags表
// @Tags Tags
// @Summary 创建zm_tags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tag.Tags true "创建zm_tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /tags/createTags [post]
func (tagsApi *TagsApi) CreateTags(c *gin.Context) {
	var tags tag.Tags
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
        "Icon":{utils.NotEmpty()},
        "Status":{utils.NotEmpty()},
    }
	if err := utils.Verify(tags, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := tagsService.CreateTags(&tags); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTags 删除zm_tags表
// @Tags Tags
// @Summary 删除zm_tags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tag.Tags true "删除zm_tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /tags/deleteTags [delete]
func (tagsApi *TagsApi) DeleteTags(c *gin.Context) {
	var tags tag.Tags
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tagsService.DeleteTags(tags); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTagsByIds 批量删除zm_tags表
// @Tags Tags
// @Summary 批量删除zm_tags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除zm_tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /tags/deleteTagsByIds [delete]
func (tagsApi *TagsApi) DeleteTagsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tagsService.DeleteTagsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTags 更新zm_tags表
// @Tags Tags
// @Summary 更新zm_tags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body tag.Tags true "更新zm_tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /tags/updateTags [put]
func (tagsApi *TagsApi) UpdateTags(c *gin.Context) {
	var tags tag.Tags
	err := c.ShouldBindJSON(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
          "Icon":{utils.NotEmpty()},
          "Status":{utils.NotEmpty()},
      }
    if err := utils.Verify(tags, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := tagsService.UpdateTags(tags); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTags 用id查询zm_tags表
// @Tags Tags
// @Summary 用id查询zm_tags表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tag.Tags true "用id查询zm_tags表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /tags/findTags [get]
func (tagsApi *TagsApi) FindTags(c *gin.Context) {
	var tags tag.Tags
	err := c.ShouldBindQuery(&tags)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retags, err := tagsService.GetTags(tags.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retags": retags}, c)
	}
}

// GetTagsList 分页获取zm_tags表列表
// @Tags Tags
// @Summary 分页获取zm_tags表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tagReq.TagsSearch true "分页获取zm_tags表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/getTagsList [get]
func (tagsApi *TagsApi) GetTagsList(c *gin.Context) {
	var pageInfo tagReq.TagsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tagsService.GetTagsInfoList(pageInfo); err != nil {
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

// GetTagsListAll 分页获取zm_tags表列表
// @Tags Tags
// @Summary 分页获取zm_tags表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query tagReq.TagsSearch true "分页获取zm_tags表列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /tags/getTagsListAll [get]
func (tagsApi *TagsApi) GetTagsListAll(c *gin.Context) {
	if list, err := tagsService.GetTagsInfoListAll(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
		}, "获取成功", c)
	}
}