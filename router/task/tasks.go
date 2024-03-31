package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TasksRouter struct {
}

// InitTasksRouter 初始化 zmTask表 路由信息
func (s *TasksRouter) InitTasksRouter(Router *gin.RouterGroup) {
	tasksRouter := Router.Group("tasks").Use(middleware.OperationRecord())
	tasksRouterWithoutRecord := Router.Group("tasks")
	var tasksApi = v1.ApiGroupApp.TaskApiGroup.TasksApi
	{
		tasksRouter.DELETE("deleteTasks", tasksApi.DeleteTasks) // 删除zmTask表
		tasksRouter.DELETE("deleteTasksByIds", tasksApi.DeleteTasksByIds) // 批量删除zmTask表
	}
	{
		tasksRouterWithoutRecord.GET("findTasks", tasksApi.FindTasks)        // 根据ID获取zmTask表
		tasksRouterWithoutRecord.GET("getTasksList", tasksApi.GetTasksList)  // 获取zmTask表列表
	}
}
