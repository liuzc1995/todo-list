package vm

import "github.com/542213314/frame-sample/server/model"

//页面呈现内容
type IndexViewModel struct {
	BaseViewModel
	TaskFinish *model.Task
	TaskDoing  *model.Task
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	v := IndexViewModel{}
	// td, _ := model.GetTaskByStatus(0) //进行中事务
	// tf, _ := model.GetTaskByStatus(1) //已完成事务
	// v.TaskDoing = td
	// v.TaskFinish = tf
	// log.Debug(tf)
	v.SetTitle("TodoList")
	return v
}
