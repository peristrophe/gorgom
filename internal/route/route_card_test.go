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

const CARD_DETAIL_EXPECT_BODY = `{
    "id": "0d756135-968c-403f-a6af-d7e69ac96769",
    "title": "hoge.fuga.piyo",
    "description": "foo bar buz",
    "boxId": "5bde985f-3ecd-46e7-95ca-257dff67d3e4",
    "tags": null,
    "comments": null,
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`

func TestRoute_CardDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	boardID, _ := uuid.Parse("9a42b6f1-237d-11ee-8a00-0242ac383802")
	boxID, _ := uuid.Parse("5bde985f-3ecd-46e7-95ca-257dff67d3e4")
	cardID, _ := uuid.Parse("0d756135-968c-403f-a6af-d7e69ac96769")
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
	box := entity.Box{
		ID:        boxID,
		Title:     "hoge.fuga",
		BoardID:   boardID,
		Board:     board,
		Cards:     []entity.Card{},
		CreatedAt: date,
		UpdatedAt: date,
	}
	cardStub := entity.Card{
		ID:          cardID,
		Title:       "hoge.fuga.piyo",
		Description: "foo bar buz",
		BoxID:       boxID,
		Box:         box,
		CreatedAt:   date,
		UpdatedAt:   date,
	}
	box.Cards = append(box.Cards, cardStub)
	board.Boxes = append(board.Boxes, box)
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Groups:    []entity.Group{{ID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockRepo.EXPECT().GetCardByID(boardID).Return(&cardStub, nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT(userID.String())
	if token == nil {
		panic(fmt.Errorf("NewJWT failed."))
	}

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	endpoint := fmt.Sprintf("/api/v1/cards/%s", boardID.String())
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, CARD_DETAIL_EXPECT_BODY, w.Body.String())
}

func TestRoute_CardDetail_UnauthorizedError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := mock.NewMockRepository(mockCtrl)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/cards/foobarbuz", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.Equal(t, UNAUTH_EXPECT_BODY, w.Body.String())
}
