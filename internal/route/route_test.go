package route

import (
	"fmt"
	"gorgom/internal/controller"
	"gorgom/internal/entity"
	"gorgom/internal/mock"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

const BOARD_DETAIL_EXPECT_BODY = `{
    "id": "9a42b6f1-237d-11ee-8a00-0242ac383802",
    "title": "hoge",
    "ownerGroupId": "4e4d3517-237e-11ee-b7fd-0242ac383802",
    "boxes": [],
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`

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
	board := entity.Board{
		ID:           boardID,
		Title:        "hoge",
		OwnerGroupID: groupID,
		Boxes:        []entity.Box{},
		CreatedAt:    date,
		UpdatedAt:    date,
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().BoardByID(boardID).Return(&board)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	endpoint := fmt.Sprintf("/api/v1/boards/%s", boardID.String())
	req, _ := http.NewRequest("GET", endpoint, nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, BOARD_DETAIL_EXPECT_BODY, w.Body.String())
}
