package model

type Task struct {
	ID       int    `gorm:"primary_key"`
	TaskName string `gorm:"type:varchar(30)"`
	Status   int64  `gorm:"type:int(10)"`
}
