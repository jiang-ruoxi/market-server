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

// DeleteTasks 删除zmTask表记录
func (tasksService *TasksService) DeleteTasks(tasks task.Tasks) (err error) {
	var s task.Tasks
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id=?", tasks.ID).Update("is_deleted", 1).Error
	return err
}

// DeleteTasksByIds 批量删除zmTask表记录
func (tasksService *TasksService) DeleteTasksByIds(ids request.IdsReq) (err error) {
	var s task.Tasks
	err = global.MustGetGlobalDBByDBName("market").Model(&s).Debug().Where("id IN ?", ids.Ids).Updates(&task.Tasks{IsDeleted: 1}).Error
	return err
}

// GetTasks 根据id获取zmTask表记录
func (tasksService *TasksService) GetTasks(id int) (tasks task.Tasks, err error) {
	err = global.MustGetGlobalDBByDBName("market").Where("id = ?", id).First(&tasks).Error

	var tagInfo tag.Tags
	db1 := global.MustGetGlobalDBByDBName("market").Model(&tag.Tags{})
	db1.Where("id=?", tasks.TagId).First(&tagInfo)

	tasks.TagName = tagInfo.Name
	return
}

// GetTasksInfoList 分页获取zmTask表记录
func (tasksService *TasksService) GetTasksInfoList(info taskReq.TasksSearch) (list []task.Tasks, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("market").Model(&task.Tasks{}).Where("is_deleted = 0")
	var taskss []task.Tasks
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where(" and created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Order("id desc").Find(&taskss).Error

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
