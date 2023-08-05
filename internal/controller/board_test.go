package controller

import (
	"fmt"
	"gorgom/internal/entity"
	"gorgom/internal/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const (
	BOARD_DETAIL_EXPECT_BODY = `{
    "id": "cc6ede1a-c2dc-43e3-a992-ffd8a610be92",
    "title": "foo",
    "ownerGroupId": "40f0e6f9-cc36-49aa-9c73-856c34bcc915",
    "boxes": null,
    "definedTags": null,
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`
	BOARDS_EXPECT_BODY = `[
    {
        "id": "cc6ede1a-c2dc-43e3-a992-ffd8a610be92",
        "title": "foo",
        "ownerGroupId": "40f0e6f9-cc36-49aa-9c73-856c34bcc915",
        "boxes": null,
        "definedTags": null,
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    },
    {
        "id": "0f9bb0a6-acd5-4fa9-9f0a-40fc10a5a380",
        "title": "bar",
        "ownerGroupId": "40f0e6f9-cc36-49aa-9c73-856c34bcc915",
        "boxes": null,
        "definedTags": null,
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    }
]`
)

func TestController_BoardDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}

	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")
	boardStub := entity.Board{
		ID:           boardID,
		Title:        "foo",
		OwnerGroupID: groupID,
		CreatedAt:    date,
		UpdatedAt:    date,
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().GetBoardByID(boardID).Return(&boardStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards/:boardID", appCtrl.BoardDetail())

	endpoint := fmt.Sprintf("/api/v1/boards/%s", boardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, BOARD_DETAIL_EXPECT_BODY, w.Body.String())
}

func TestController_BoardDetail_AuthError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userID, _ := uuid.Parse(TESTING_USER_ID)
	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(nil, fmt.Errorf("Unauthorized."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards/:boardID", appCtrl.BoardDetail())

	endpoint := fmt.Sprintf("/api/v1/boards/%s", boardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Unauthorized.\"\n}", w.Body.String())
}

func TestController_BoardDetail_FetchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}
	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().GetBoardByID(boardID).Return(nil, fmt.Errorf("Board fetch failed."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards/:boardID", appCtrl.BoardDetail())

	endpoint := fmt.Sprintf("/api/v1/boards/%s", boardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Board fetch failed.\"\n}", w.Body.String())
}

func TestController_BoardDetail_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		Groups:    []entity.Group{},
		CreatedAt: date,
		UpdatedAt: date,
	}

	boardID, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")
	boardStub := entity.Board{
		ID:           boardID,
		Title:        "foo",
		OwnerGroupID: groupID,
		CreatedAt:    date,
		UpdatedAt:    date,
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().GetBoardByID(boardID).Return(&boardStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards/:boardID", appCtrl.BoardDetail())

	endpoint := fmt.Sprintf("/api/v1/boards/%s", boardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Board not found.\"\n}", w.Body.String())
}

func TestController_Boards(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}

	boardID0, _ := uuid.Parse("cc6ede1a-c2dc-43e3-a992-ffd8a610be92")
	boardID1, _ := uuid.Parse("0f9bb0a6-acd5-4fa9-9f0a-40fc10a5a380")
	boardsStub := []entity.Board{
		{
			ID:           boardID0,
			Title:        "foo",
			OwnerGroupID: groupID,
			CreatedAt:    date,
			UpdatedAt:    date,
		},
		{
			ID:           boardID1,
			Title:        "bar",
			OwnerGroupID: groupID,
			CreatedAt:    date,
			UpdatedAt:    date,
		},
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().ListBoardsByGroupID(groupID).Return(boardsStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards", appCtrl.Boards())

	request, _ := http.NewRequest("GET", "/api/v1/boards", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, BOARDS_EXPECT_BODY, w.Body.String())
}

func TestController_Boards_AuthError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userID, _ := uuid.Parse(TESTING_USER_ID)

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(nil, fmt.Errorf("Unauthorized."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards", appCtrl.Boards())

	request, _ := http.NewRequest("GET", "/api/v1/boards", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Unauthorized.\"\n}", w.Body.String())
}

func TestController_Boards_FetchError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().ListBoardsByGroupID(groupID).Return(nil, fmt.Errorf("Boards fetch failed."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/boards", appCtrl.Boards())

	request, _ := http.NewRequest("GET", "/api/v1/boards", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Boards fetch failed.\"\n}", w.Body.String())
}
