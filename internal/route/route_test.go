package route

import (
	"encoding/json"
	"fmt"
	"gorgom/internal/controller"
	"gorgom/internal/entity"
	"gorgom/internal/mock"
	"gorgom/internal/util"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	BOARD_DETAIL_EXPECT_BODY = `{
    "id": "9a42b6f1-237d-11ee-8a00-0242ac383802",
    "title": "hoge",
    "ownerGroupId": "4e4d3517-237e-11ee-b7fd-0242ac383802",
    "boxes": [],
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`
	BOARD_DETAIL_UNAUTH_EXPECT_BODY = `{
    "error": "Unauthorized"
}`
)

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

func TestRoute_BoardDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boardID, err := uuid.Parse("9a42b6f1-237d-11ee-8a00-0242ac383802")
	if err != nil {
		panic(err)
	}
	groupID, err := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	if err != nil {
		panic(err)
	}
	boardStub := entity.Board{
		ID:           boardID,
		Title:        "hoge",
		OwnerGroupID: groupID,
		Boxes:        []entity.Box{},
		CreatedAt:    date,
		UpdatedAt:    date,
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().BoardByID(boardID).Return(&boardStub)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT("foo")
	if err != nil {
		panic(err)
	}

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	endpoint := fmt.Sprintf("/api/v1/boards/%s", boardID.String())
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, BOARD_DETAIL_EXPECT_BODY, w.Body.String())
}

func TestRoute_BoardDetail_UnauthorizedError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockRepository(mockCtrl)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/boards/foobarbuz", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, BOARD_DETAIL_UNAUTH_EXPECT_BODY, w.Body.String())
}
