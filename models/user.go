package models

import (
	"UserSimpleCRUD/utils/token"
	"gorm.io/gorm"
)

// User
type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Salary   int    `json:"salary"`
	Address  string `json:"address"`
	Task     []Task `json:"-"`
}

type UserLogin interface {
	LoginCheck(db *gorm.DB) (string, error)
}

func (user User) LoginCheck(db *gorm.DB) (string, error) {
	u := User{ID: 0}
	if user.Username != "admin" && user.Password != "admin" {
		err := db.Model(u).Where("username = ? AND password = ?", user.Username, user.Password).Take(&u).Error
		if err != nil {
			return "", err
		}
	}

	tokenGenerated, err := token.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}

	return tokenGenerated, nil
}
