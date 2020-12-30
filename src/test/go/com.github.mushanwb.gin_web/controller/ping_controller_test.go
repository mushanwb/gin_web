package controller

import (
	"encoding/json"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/bootstrap"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	r := bootstrap.SetupRoute()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	var bodyMap map[string]interface{}
	_ = json.Unmarshal(body, &bodyMap)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", bodyMap["message"])

}
