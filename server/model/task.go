package model

import (
	"time"
)

type Task struct {
	ID         int       `gorm:"primary_key;AUTO_INCREMENT"`
	TaskName   string    `gorm:"type:varchar(30)"`
	Status     int64     `gorm:"DEFAULT:0"`
	CreateTime time.Time `sql:"DEFAULT:current_timestamp"`
	UpdateTime time.Time `sql:"DEFAULT:current_timestamp"`
}

//通过状态获取事务
func GetTaskByStatus(status int64) (*[]Task, error) {
	var task []Task
	if err := db.Where("status=?", status).Order("create_time desc,update_time").Find(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

//通过ID获取事务
func GetTaskByID(id int) (*Task, error) {
	var task Task
	if err := db.Where("id=", id).Find(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

//添加事务
func AddTask(taskName string) error {
	task := Task{TaskName: taskName}
	return db.Create(&task).Error
}

//根据ID更新
func UpdateTaskByTaskID(id int, contents map[string]interface{}) error {
	item, err := GetTaskByID(id)
	if err != nil {
		return err
	}
	return db.Model(item).Update(contents).Error
}

//更新事务状态
func UpdateTaskStatus(id, status int) error {
	contents := map[string]interface{}{"status": status}
	return UpdateTaskByTaskID(id, contents)
}

//删除事务
func DeleteTaskByID(id int) error {
	item, err := GetTaskByID(id)
	if err != nil {
		return err
	}
	return db.Delete(item).Error
}
