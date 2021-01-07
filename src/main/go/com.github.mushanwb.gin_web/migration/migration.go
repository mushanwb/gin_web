package migration

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/model"
	"gorm.io/gorm"
)

func Migration(db *gorm.DB) {

	// 自动迁移
	db.AutoMigrate(
		&model.User{},
	)

}
