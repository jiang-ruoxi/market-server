package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TasksApi struct {
}

var tasksService = service.ServiceGroupApp.TaskServiceGroup.TasksService

// DeleteTasks 删除zmTask表
func (tasksApi *TasksApi) DeleteTasks(c *gin.Context) {
	var tasks task.Tasks
	err := c.ShouldBindJSON(&tasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tasksService.DeleteTasks(tasks); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteTasksByIds 批量删除zmTask表
func (tasksApi *TasksApi) DeleteTasksByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := tasksService.DeleteTasksByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// FindTasks 用id查询zmTask表
func (tasksApi *TasksApi) FindTasks(c *gin.Context) {
	var tasks task.Tasks
	err := c.ShouldBindQuery(&tasks)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retasks, err := tasksService.GetTasks(tasks.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retasks": retasks}, c)
	}
}

// GetTasksList 分页获取zmTask表列表
func (tasksApi *TasksApi) GetTasksList(c *gin.Context) {
	var pageInfo taskReq.TasksSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := tasksService.GetTasksInfoList(pageInfo); err != nil {
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
