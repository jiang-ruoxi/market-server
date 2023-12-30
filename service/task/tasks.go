package task

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/tag"
	"github.com/flipped-aurora/gin-vue-admin/server/model/task"
	taskReq "github.com/flipped-aurora/gin-vue-admin/server/model/task/request"
)

type TasksService struct {
}

// CreateTasks 创建zmTask表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tasksService *TasksService) CreateTasks(tasks *task.Tasks) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Create(tasks).Error
	return err
}

// DeleteTasks 删除zmTask表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tasksService *TasksService) DeleteTasks(tasks task.Tasks) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&tasks).Error
	return err
}

// DeleteTasksByIds 批量删除zmTask表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tasksService *TasksService) DeleteTasksByIds(ids request.IdsReq) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Delete(&[]task.Tasks{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateTasks 更新zmTask表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tasksService *TasksService) UpdateTasks(tasks task.Tasks) (err error) {
	err = global.MustGetGlobalDBByDBName("market").Debug().Save(&tasks).Error
	return err
}

// GetTasks 根据id获取zmTask表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tasksService *TasksService) GetTasks(id uint) (tasks task.Tasks, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&tasks).Error
	return
}

// GetTasksInfoList 分页获取zmTask表记录
// Author [piexlmax](https://github.com/piexlmax)
func (tasksService *TasksService) GetTasksInfoList(info taskReq.TasksSearch) (list []task.Tasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&task.Tasks{})
	var taskss []task.Tasks
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&taskss).Error

	var tagList []tag.Tags
	db1 := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{})
	db1.Find(&tagList)
	for idx, _ := range taskss {
		for tagIdx, _ := range tagList {
			if tagList[tagIdx].ID == taskss[idx].TagId {
				taskss[idx].TagName = tagList[tagIdx].Name
			}
		}

	}
	return taskss, total, err
}
