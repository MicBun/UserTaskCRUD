package models

// Task
type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	UserID      uint   `json:"user_id"`
	TaskName    string `json:"task_name"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	User        User   `json:"-"`
}
