package base

import (
	"gin_web/src/main/go/com.github.mushanwb.gin_web/bootstrap"
	_ "gin_web/src/main/go/com.github.mushanwb.gin_web/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDB(t *testing.T) {
	_, err := bootstrap.ConnectDB()
	assert.Equal(t, nil, err)
}
