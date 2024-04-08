package bad

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bad"
	badReq "github.com/flipped-aurora/gin-vue-admin/server/model/bad/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BadWordsApi struct {
}

var badWordsService = service.ServiceGroupApp.BadServiceGroup.BadWordsService


// CreateBadWords 创建bad
func (badWordsApi *BadWordsApi) CreateBadWords(c *gin.Context) {
	var badWords bad.BadWords
	err := c.ShouldBindJSON(&badWords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    verify := utils.Rules{
        "Name":{utils.NotEmpty()},
    }
	if err := utils.Verify(badWords, verify); err != nil {
    		response.FailWithMessage(err.Error(), c)
    		return
    	}
	if err := badWordsService.CreateBadWords(&badWords); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBadWords 删除bad
func (badWordsApi *BadWordsApi) DeleteBadWords(c *gin.Context) {
	var badWords bad.BadWords
	err := c.ShouldBindJSON(&badWords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := badWordsService.DeleteBadWords(badWords); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBadWordsByIds 批量删除bad
func (badWordsApi *BadWordsApi) DeleteBadWordsByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := badWordsService.DeleteBadWordsByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBadWords 更新bad
func (badWordsApi *BadWordsApi) UpdateBadWords(c *gin.Context) {
	var badWords bad.BadWords
	err := c.ShouldBindJSON(&badWords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
      verify := utils.Rules{
          "Name":{utils.NotEmpty()},
      }
    if err := utils.Verify(badWords, verify); err != nil {
      	response.FailWithMessage(err.Error(), c)
      	return
     }
	if err := badWordsService.UpdateBadWords(badWords); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBadWords 用id查询bad
func (badWordsApi *BadWordsApi) FindBadWords(c *gin.Context) {
	var badWords bad.BadWords
	err := c.ShouldBindQuery(&badWords)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rebadWords, err := badWordsService.GetBadWords(badWords.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rebadWords": rebadWords}, c)
	}
}

// GetBadWordsList 分页获取bad列表
func (badWordsApi *BadWordsApi) GetBadWordsList(c *gin.Context) {
	var pageInfo badReq.BadWordsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := badWordsService.GetBadWordsInfoList(pageInfo); err != nil {
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
