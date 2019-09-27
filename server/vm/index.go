package vm

import (
	"github.com/542213314/frame-sample/server/model"
)

//页面呈现内容
type IndexViewModel struct {
	BaseViewModel
	TaskFinish []model.Task
	TaskBeing  []model.Task
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	v := IndexViewModel{}
	tb, _ := model.GetTaskByStatus(0) //进行中事务
	tf, _ := model.GetTaskByStatus(1) //已完成事务
	v.TaskBeing = *tb
	v.TaskFinish = *tf
	v.SetTitle("TodoList")
	return v
}

//添加事务
func AddTask(taskName string) error {
	return model.AddTask(taskName)
}

//更新事务状态
func UpdateTaskByIDStatus(id, status int) error {
	return model.UpdateTaskStatus(id, status)
}
