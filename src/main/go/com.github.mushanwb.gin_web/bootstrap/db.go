package bootstrap

import (
	u "gin_web/src/main/go/com.github.mushanwb.gin_web/Utils"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/pkg"
)

func SetupDB() {
	defaultConfig := u.Get("database.default")
	pkg.ConnectDefaultDB(defaultConfig.(map[string]interface{}))
}
