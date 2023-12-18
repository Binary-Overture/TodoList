package service

import (
	"Todolist/pkg/ctl"
	"Todolist/pkg/util"
	"Todolist/repository/db/dao"
	"Todolist/repository/db/model"
	"Todolist/types"
	"context"
	"sync"
	"time"
)

var TaskSrvIns *TaskSrv
var TaskSrcOnce sync.Once

type TaskSrv struct {
}

func GetTaskSrv() *TaskSrv {
	TaskSrcOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (s *TaskSrv) CreateTask(c context.Context, req *types.CreateTaskReq) (resp any, err error) {
	u, err := ctl.GetUserInfo(c)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	user, err := dao.NewUserDao(c).FindUserByUserId(u.Id)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	task := &model.TaskModel{
		User:      *user,
		Uid:       user.ID,
		Title:     req.Title,
		Content:   req.Content,
		Status:    0,
		StartTime: time.Now().Unix(),
	}
	err = dao.NewTaskDao(c).CreateTask(task)
	if err != nil {
		util.LogrusObj.Info(err)
		return
	}
	return ctl.RespSuccess(), nil
}
