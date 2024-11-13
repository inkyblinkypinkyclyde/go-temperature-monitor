package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RouterTestSuite struct {
	suite.Suite
	DefaultReceiver
	router *gin.Engine
}

func TestRouterTestSuite(t *testing.T) {
	suite.Run(t, new(RouterTestSuite))
}

func (s *RouterTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
}

func (s *RouterTestSuite) TearDownSuite() {
	// delete test files
}

func (s *RouterTestSuite) TestLogRequestHappyPath() {
	s.router.POST("/logtemperature", s.DefaultReceiver.LogTemperature)

	requestBody := map[string]string{
		"origin":  "localhost",
		"message": "message1",
	}
	body, _ := json.Marshal(requestBody)
	res := performRequest(s.router, "POST", "/logtemperature", body)

	assert.Equal(s.T(), http.StatusOK, res.Code)
}
func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
