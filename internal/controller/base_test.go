package controller

import (
	"gorgom/internal/middleware"
	"gorgom/internal/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const TESTING_USER_ID = "5b4ccb43-81ab-4357-8591-95b42d42e339"

func middlewareStub(c *gin.Context) {
	c.AddParam(middleware.CONTEXT_PARAM_KEY_AUTH_USER_ID, TESTING_USER_ID)
	c.Next()
}

func TestController_MissingAuthUserID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.GET("/api/v1/users/mypage", appCtrl.MyPage())

	request, _ := http.NewRequest("GET", "/api/v1/users/mypage", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
}
