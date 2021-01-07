package service

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/model"
)

func CreateUser(user *model.User) error {
	err := user.CreateUser()
	return err
}
