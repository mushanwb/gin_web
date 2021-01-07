package controller

import (
	"gin_web/src/test/go/com.github.mushanwb.gin_web/base"
	"gin_web/src/test/go/com.github.mushanwb.gin_web/data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterData(t *testing.T) {
	w, _ := base.Response("POST", "/register", data.RegisterData)
	assert.Equal(t, 200, w.Code)
}

func TestRegisterBadData(t *testing.T) {
	w, _ := base.Response("POST", "/register", data.RegisterBadData)
	assert.Equal(t, 400, w.Code)
}
