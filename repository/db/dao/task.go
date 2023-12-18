package dao

import (
	"Todolist/repository/db/model"
	"context"
	"gorm.io/gorm"
)

type TaskDao struct {
	*gorm.DB
}

func NewTaskDao(c context.Context) *TaskDao {
	if c == nil {
		c = context.Background()
	}
	return &TaskDao{NewDBClient(c)}
}

// CreateTask
// 创建一个新的任务
func (s *TaskDao) CreateTask(task *model.TaskModel) error {
	return s.Model(&model.TaskModel{}).Create(&task).Error
}
