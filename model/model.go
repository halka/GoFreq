//Package model データモデルを定義している
package model

import (
	"github.com/jinzhu/gorm"
)

//User ユーザ
type User struct {
	gorm.Model
	Username string
	Password string
}

//Freq 周波数
type Freq struct {
	gorm.Model
	Freq     string `json:"freq"`
	Name     string `json:"name"`
	Location string `json:"location"`
	User     User   `gorm:"foreignkey:ID"`
}
