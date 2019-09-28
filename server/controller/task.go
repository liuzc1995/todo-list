package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/542213314/frame-sample/server/vm"
	"go.mongodb.org/mongo-driver/bson"
	"golanger.com/log"
)

//任务添加
func taskAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		m := bson.M{
			"status":  0,
			"message": "",
		}
		taskName := r.Form.Get("name")
		if err := vm.AddTask(taskName); err != nil {
			m["message"] = "Error in add task"
			log.Debug("add task error:", err)
		} else {
			m["message"] = "添加成功"
			m["status"] = 1
		}
		ret, _ := json.Marshal(m)
		w.Write(ret)
		return
	}
}

//任务状态更新
func taskUpdateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		m := bson.M{
			"status":  0,
			"message": "",
		}
		ids := r.Form.Get("id")
		statuss := r.Form.Get("status")
		id, _ := strconv.Atoi(ids)
		status, _ := strconv.Atoi(statuss)
		if err := vm.UpdateTaskByIDStatus(id, status); err != nil {
			m["message"] = "Error in update task"
			log.Debug("update task error:", err)
		} else {
			m["message"] = "更改成功"
			m["status"] = 1
		}
		ret, _ := json.Marshal(m)
		w.Write(ret)
		return
	}
}

//删除事务
func taskDeleteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		m := bson.M{
			"status":  0,
			"message": "",
		}
		ids := r.Form.Get("id")
		id, _ := strconv.Atoi(ids)
		if err := vm.DeleteTaskByID(id); err != nil {
			m["message"] = "Error in delete task"
			log.Debug("delete task error:", err)
		} else {
			m["message"] = "更改成功"
			m["status"] = 1
		}
		ret, _ := json.Marshal(m)
		w.Write(ret)
		return
	}
}
