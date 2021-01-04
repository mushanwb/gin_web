package base

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/bootstrap"
	"gin_web/src/test/go/com.github.mushanwb.gin_web/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	mysqlConfig := data.MysqlConfig
	_, err := bootstrap.ConnectDB(mysqlConfig)
	assert.Equal(t, nil, err)
}
