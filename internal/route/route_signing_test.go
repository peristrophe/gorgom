package route

import (
	"encoding/json"
	"fmt"
	"gorgom/internal/controller"
	"gorgom/internal/entity"
	"gorgom/internal/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const UNAUTH_EXPECT_BODY = `{
    "error": "Unauthorized"
}`

func TestRoute_SignUp(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	userID, err := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	if err != nil {
		panic(err)
	}
	request := struct {
		Email    string
		Password string
	}{Email: "hoge@example.com", Password: "hogehoge"}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	userStub := entity.User{
		ID:        userID,
		Email:     request.Email,
		Name:      "hoge",
		Location:  "tokyo",
		CreatedAt: date,
		UpdatedAt: date,
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().CreateUser(request.Email, request.Password).Return(&userStub, nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	reader := strings.NewReader(string(requestBytes))
	req, _ := http.NewRequest("POST", "/api/v1/signup", reader)
	r.ServeHTTP(w, req)

	var response struct {
		UserID string `json:"userID"`
		Token  string `json:"token"`
	}
	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		panic(err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "5b4ccb43-81ab-4357-8591-95b42d42e339", response.UserID)
	assert.Equal(t, "string", fmt.Sprintf("%T", response.Token))
	assert.NotZero(t, response.Token)

	cookies := w.Result().Cookies()
	assert.NotEmpty(t, len(cookies))
	assert.Equal(t, "token", cookies[0].Name)
	assert.Equal(t, response.Token, cookies[0].Value)
}

func TestRoute_SignIn(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	userID, err := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	if err != nil {
		panic(err)
	}
	salt, err := uuid.Parse("d4ec27ec-2878-4deb-af7a-b2da6ac75e5e")
	if err != nil {
		panic(err)
	}
	request := struct {
		Email    string
		Password string
	}{Email: "hoge@example.com", Password: "hogehoge"}
	requestBytes, err := json.Marshal(request)
	if err != nil {
		panic(err)
	}
	userStub := entity.User{
		ID:        userID,
		Email:     request.Email,
		Salt:      salt,
		Name:      "hoge",
		Location:  "tokyo",
		CreatedAt: date,
		UpdatedAt: date,
	}
	if err = userStub.SetPassword(request.Password); err != nil {
		panic(nil)
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByEmail(request.Email).Return(&userStub, nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	reader := strings.NewReader(string(requestBytes))
	req, _ := http.NewRequest("POST", "/api/v1/signin", reader)
	r.ServeHTTP(w, req)

	var response struct {
		UserID string `json:"userID"`
		Token  string `json:"token"`
	}
	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		panic(err)
	}
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "5b4ccb43-81ab-4357-8591-95b42d42e339", response.UserID)
	assert.Equal(t, "string", fmt.Sprintf("%T", response.Token))
	assert.NotZero(t, response.Token)

	cookies := w.Result().Cookies()
	assert.NotEmpty(t, len(cookies))
	assert.Equal(t, "token", cookies[0].Name)
	assert.Equal(t, response.Token, cookies[0].Value)
}
