package vm

import (
	"fmt"
	"time"

	"github.com/542213314/frame-sample/server/model"
)

//页面呈现内容
type IndexViewModel struct {
	BaseViewModel
	TaskFinish []model.Task
	TaskBeing  []model.Task
	TaskTime   interface{}
}

type IndexViewModelOp struct{}

func (IndexViewModelOp) GetVM() IndexViewModel {
	v := IndexViewModel{}
	tb, _ := model.GetTaskByStatus(0) //进行中事务
	tf, _ := model.GetTaskByStatus(1) //已完成事务
	v.TaskBeing = *tb
	v.TaskFinish = *tf
	v.TaskTime = func(tm time.Time) string {

		year := tm.Year()
		month := fmt.Sprintf("%d", tm.Month())
		tm.Month()
		if len(month) < 2 {
			month = fmt.Sprintf("%d%d", 0, tm.Month())
		}
		day := ParseIntTime(tm.Day())
		hour := ParseIntTime(tm.Hour())

		minute := ParseIntTime(tm.Minute())
		return fmt.Sprintf("%d/%v/%v %v:%v", year, month, day, hour, minute)
	}
	v.SetTitle("TodoList")
	return v
}

func ParseIntTime(n int) string {
	if n < 10 {
		return fmt.Sprintf("%d%d", 0, n)
	}
	return fmt.Sprintf("%d", n)
}

//添加事务
func AddTask(taskName string) error {
	return model.AddTask(taskName)
}

//更新事务状态
func UpdateTaskByIDStatus(id, status int) error {
	return model.UpdateTaskStatus(id, status)
}

//删除事务
func DeleteTaskByID(id int) error {
	return model.DeleteTaskByID(id)
}
