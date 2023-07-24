package route

import (
	"fmt"
	"gorgom/internal/controller"
	"gorgom/internal/entity"
	"gorgom/internal/mock"
	"gorgom/internal/util"
	"net/http"
	"net/http/httptest"
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
	BOARDS_EXPECT_BODY = `[
    {
        "id": "9a42b6f1-237d-11ee-8a00-0242ac383802",
        "title": "hoge",
        "ownerGroupId": "4e4d3517-237e-11ee-b7fd-0242ac383802",
        "boxes": [],
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    }
]`
)

func TestRoute_BoardDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boardID, _ := uuid.Parse("9a42b6f1-237d-11ee-8a00-0242ac383802")
	groupID, _ := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	boardStub := entity.Board{
		ID:           boardID,
		Title:        "hoge",
		OwnerGroupID: groupID,
		Boxes:        []entity.Box{},
		CreatedAt:    date,
		UpdatedAt:    date,
	}
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockRepo.EXPECT().GetBoardByID(boardID).Return(&boardStub, nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT(userID.String())
	if token == nil {
		panic(fmt.Errorf("NewJWT failed."))
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
	assert.Equal(t, UNAUTH_EXPECT_BODY, w.Body.String())
}

func TestRoute_Boards(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boardID, _ := uuid.Parse("9a42b6f1-237d-11ee-8a00-0242ac383802")
	groupID, _ := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	board := entity.Board{
		ID:           boardID,
		Title:        "hoge",
		OwnerGroupID: groupID,
		Boxes:        []entity.Box{},
		CreatedAt:    date,
		UpdatedAt:    date,
	}
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}
	var boardsStub []entity.Board
	boardsStub = append(boardsStub, board)

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockRepo.EXPECT().ListBoardsByGroupID(groupID).Return(boardsStub, nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT(userID.String())
	if token == nil {
		panic(fmt.Errorf("NewJWT failed."))
	}

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/boards", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, BOARDS_EXPECT_BODY, w.Body.String())
}

func TestRoute_Boards_UnauthorizedError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockRepository(mockCtrl)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/boards", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, UNAUTH_EXPECT_BODY, w.Body.String())
}
