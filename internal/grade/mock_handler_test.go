package grade

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetGradeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{}
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/grade/:studentId", handler.GetGradeHandler)

	req, _ := http.NewRequest("GET", "/grade/6509650269", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSubmitGradeHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := &MockService{}
	handler := NewHandler(mockService)
	router := gin.Default()
	router.POST("/grade", handler.SubmitGradeHandler)

	body := `{"student_id":"6609650269", "homework":80, "midterm":70, "final":90}`
	req, _ := http.NewRequest("POST", "/grade", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestSubmitGradeHandler_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	handler := NewHandler(&MockService{})
	router := gin.Default()
	router.POST("/grade", handler.SubmitGradeHandler)

	body := `{"student_id":"6609650269"`
	req, _ := http.NewRequest("POST", "/grade", strings.NewReader(body))
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSubmitGradeHandler_ServiceError(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{ShouldReturnError: true}
	handler := NewHandler(mockService)
	router := gin.Default()
	router.POST("/grade", handler.SubmitGradeHandler)

	body := `{"student_id":"6609650269", "homework":80, "midterm":70, "final":90}`
	req, _ := http.NewRequest("POST", "/grade", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestGetGradeHandler_NotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockService{ShouldReturnError: true}
	handler := NewHandler(mockService)

	router := gin.Default()
	router.GET("/grade/:studentId", handler.GetGradeHandler)

	req, _ := http.NewRequest("GET", "/grade/non-existent-id", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
