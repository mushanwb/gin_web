package com_github_mushanwb_gin_web_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestRoute(t *testing.T) {
	baseURL := "http://localhost:8080"

	resp, err := http.Get(baseURL + "/ping")
	assert.NoError(t, err, "请求发生错误")

	body, err := ioutil.ReadAll(resp.Body)

	var dataMap map[string]interface{}
	_ = json.Unmarshal(body, &dataMap)

	defer resp.Body.Close()

	// 2. 检测 —— 是否无错误且 200
	assert.Equal(t, 200, resp.StatusCode, "应返回状态码 200")
	assert.Equal(t, "pong", dataMap["message"], "应返值为 pong")

}
