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

const CARD_DETAIL_EXPECT_BODY = `{
    "id": "250013d6-6298-4572-932f-ab46dbab0b2c",
    "title": "foo",
    "description": "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
    "boxId": "00000000-0000-0000-0000-000000000000",
    "tags": null,
    "comments": null,
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`

func TestController_CardDetail(t *testing.T) {
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

	cardID, _ := uuid.Parse("250013d6-6298-4572-932f-ab46dbab0b2c")
	cardStub := entity.Card{
		ID:          cardID,
		Title:       "foo",
		Description: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Box:         entity.Box{Board: entity.Board{OwnerGroupID: groupID}},
		CreatedAt:   date,
		UpdatedAt:   date,
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().GetCardByID(cardID).Return(&cardStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/cards/:cardID", appCtrl.CardDetail())

	endpoint := fmt.Sprintf("/api/v1/cards/%s", cardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, CARD_DETAIL_EXPECT_BODY, w.Body.String())
}

func TestController_CardDetail_AuthError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userID, _ := uuid.Parse(TESTING_USER_ID)
	cardID, _ := uuid.Parse("250013d6-6298-4572-932f-ab46dbab0b2c")

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(nil, fmt.Errorf("Unauthorized."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/cards/:cardID", appCtrl.CardDetail())

	endpoint := fmt.Sprintf("/api/v1/cards/%s", cardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Unauthorized.\"\n}", w.Body.String())
}

func TestController_CardDetail_FetchError(t *testing.T) {
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

	cardID, _ := uuid.Parse("250013d6-6298-4572-932f-ab46dbab0b2c")

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().GetCardByID(cardID).Return(nil, fmt.Errorf("Card fetch failed."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/cards/:cardID", appCtrl.CardDetail())

	endpoint := fmt.Sprintf("/api/v1/cards/%s", cardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Card fetch failed.\"\n}", w.Body.String())
}

func TestController_CardDetail_NotFound(t *testing.T) {
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

	cardID, _ := uuid.Parse("250013d6-6298-4572-932f-ab46dbab0b2c")
	cardStub := entity.Card{
		ID:          cardID,
		Title:       "foo",
		Description: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		Box:         entity.Box{Board: entity.Board{OwnerGroupID: groupID}},
		CreatedAt:   date,
		UpdatedAt:   date,
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	mockAppRepo.EXPECT().GetCardByID(cardID).Return(&cardStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/cards/:cardID", appCtrl.CardDetail())

	endpoint := fmt.Sprintf("/api/v1/cards/%s", cardID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Card not found.\"\n}", w.Body.String())
}
