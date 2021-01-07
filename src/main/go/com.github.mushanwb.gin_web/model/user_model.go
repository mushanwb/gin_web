package model

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/logger"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/pkg"
)

type User struct {
	BaseModel

	Name     string `gorm:"type:varchar(64); not null; unique" json:"name"`
	Phone    string `gorm:"type:varchar(11); not null; unique" json:"phone"`
	Password string `gorm:"type:varchar(255); not null;" json:"-"`
}

func (user *User) CreateUser() error {
	result := pkg.DefaultDB.Create(&user)
	if err := result.Error; err != nil {
		logger.Logger.Info(err)
		return err
	}
	return nil
}
