package controller

import (
	"encoding/json"
	"fmt"
	"gorgom/internal/entity"
	"gorgom/internal/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	SIGNUP_SIGNIN_EXPECT_BODY_HEAD = `{
    "userID": "5b4ccb43-81ab-4357-8591-95b42d42e339",
    "token": "`
	MYPAGE_EXPECT_BODY = `{
    "id": "",
}`
)

func TestController_SignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	email := "hoge@example.com"
	password := "hogehoge"

	userID, _ := uuid.Parse(TESTING_USER_ID)
	userStub := entity.User{
		ID:        userID,
		Email:     email,
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	userStub.SetPassword(password)

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().CreateUser(email, password).Return(&userStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.POST("/api/v1/signup", appCtrl.SignUp())

	requestBodyObj := struct{ Email, Password string }{Email: email, Password: password}
	requestBodyBytes, _ := json.Marshal(requestBodyObj)
	reader := strings.NewReader(string(requestBodyBytes))

	request, _ := http.NewRequest("POST", "/api/v1/signup", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), SIGNUP_SIGNIN_EXPECT_BODY_HEAD))
}

func TestController_SignUp_DeserializeError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	email := "hoge@example.com"
	password := "hogehoge"

	userID, _ := uuid.Parse(TESTING_USER_ID)
	userStub := entity.User{
		ID:        userID,
		Email:     email,
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	userStub.SetPassword(password)

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.POST("/api/v1/signup", appCtrl.SignUp())

	request, _ := http.NewRequest("POST", "/api/v1/signup", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 400, w.Code)
}

func TestController_SignUp_CreateError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	email := "hoge@example.com"
	password := "hogehoge"

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().CreateUser(email, password).Return(nil, fmt.Errorf("User create failed."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.POST("/api/v1/signup", appCtrl.SignUp())

	requestBodyObj := struct{ Email, Password string }{Email: email, Password: password}
	requestBodyBytes, _ := json.Marshal(requestBodyObj)
	reader := strings.NewReader(string(requestBodyBytes))

	request, _ := http.NewRequest("POST", "/api/v1/signup", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
}

func TestController_SignIn(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	email := "hoge@example.com"
	password := "hogehoge"

	userID, _ := uuid.Parse(TESTING_USER_ID)
	salt, _ := uuid.Parse("762c8a63-c186-445f-a9ae-288ce4d8cd27")
	userStub := entity.User{
		ID:        userID,
		Email:     email,
		Salt:      salt,
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	userStub.SetPassword(password)

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByEmail(email).Return(&userStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.POST("/api/v1/signin", appCtrl.SignIn())

	requestBodyObj := struct{ Email, Password string }{Email: email, Password: password}
	requestBodyBytes, _ := json.Marshal(requestBodyObj)
	reader := strings.NewReader(string(requestBodyBytes))

	request, _ := http.NewRequest("POST", "/api/v1/signin", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.True(t, strings.HasPrefix(w.Body.String(), SIGNUP_SIGNIN_EXPECT_BODY_HEAD))
}

func TestController_SignIn_DeserializeError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.POST("/api/v1/signin", appCtrl.SignIn())

	request, _ := http.NewRequest("POST", "/api/v1/signin", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 400, w.Code)
}

func TestController_SignIn_FetchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	email := "hoge@example.com"
	password := "hogehoge"

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByEmail(email).Return(nil, fmt.Errorf("User fetch failed."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.POST("/api/v1/signin", appCtrl.SignIn())

	requestBodyObj := struct{ Email, Password string }{Email: email, Password: password}
	requestBodyBytes, _ := json.Marshal(requestBodyObj)
	reader := strings.NewReader(string(requestBodyBytes))

	request, _ := http.NewRequest("POST", "/api/v1/signin", reader)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
}

func TestController_MyPage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	salt, _ := uuid.Parse("762c8a63-c186-445f-a9ae-288ce4d8cd27")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Salt:      salt,
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	userStub.SetPassword("hogehoge")

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/users/mypage", appCtrl.MyPage())

	request, _ := http.NewRequest("GET", "/api/v1/users/mypage", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
}

func TestController_MyPage_AuthError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userID, _ := uuid.Parse(TESTING_USER_ID)

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(nil, fmt.Errorf("Unauthorized."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/users/mypage", appCtrl.MyPage())

	request, _ := http.NewRequest("GET", "/api/v1/users/mypage", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
}
