package base

import (
	_ "gin_web/src/main/go/com.github.mushanwb.gin_web/logger"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/pkg"
	"gin_web/src/test/go/com.github.mushanwb.gin_web/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	mysqlConfig := data.MysqlConfig
	_, err := pkg.ConnectDB(mysqlConfig)
	assert.Equal(t, nil, err)
}
