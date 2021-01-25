package base

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/pkg"
	"gin_web/src/test/go/com.github.mushanwb.gin_web/data"
)

func DB() {
	mysqlConfig := data.MysqlConfig
	pkg.ConnectDefaultDB(mysqlConfig)
}
