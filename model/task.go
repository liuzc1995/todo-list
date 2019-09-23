package model

type Task struct {
	ID       int    `gorm:"primary_key"`
	TaskName string `gorm:"type:varchar(30)"`
	Status   int64  `gorm:"type:int(10)"`
}

func GetTask(status int64) (*Task, error) {
	var task Task
	if err := db.Where("status=", status).Find(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
