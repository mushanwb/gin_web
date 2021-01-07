package base

import (
	"encoding/json"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/bootstrap"
	"gin_web/src/main/go/com.github.mushanwb.gin_web/logger"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
)

var bodyMap map[string]interface{}

func Response(method string, uri string, body interface{}) (*httptest.ResponseRecorder, map[string]interface{}) {

	r := bootstrap.SetupRoute()

	w := httptest.NewRecorder()

	jsonBody, _ := json.Marshal(body)

	reader := strings.NewReader(string(jsonBody))

	req, _ := http.NewRequest(method, uri, reader)

	req.Header.Set("Content-Type", "application/json;charset=UTF-8")

	r.ServeHTTP(w, req)

	responseBody, _ := ioutil.ReadAll(w.Body)
	_ = json.Unmarshal(responseBody, &bodyMap)

	logger.Logger.Info(bodyMap)

	return w, bodyMap
}
