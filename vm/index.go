package vm

import (
	"gggo/model"

	"golanger.com/log"
)

//页面呈现内容
type IndexViewModel struct {
	BaseViewModel
	TaskFinish *model.Task
	TaskDoing  *model.Task
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	v := IndexViewModel{}
	td, _ := model.GetTask(0) //进行中事务
	tf, _ := model.GetTask(1) //已完成事务
	v.TaskDoing = td
	v.TaskFinish = tf
	log.Debug(tf)
	v.SetTitle("TodoList")
	return v
}
