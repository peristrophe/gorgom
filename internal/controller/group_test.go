package controller

import (
	"fmt"
	"gorgom/internal/entity"
	"gorgom/internal/helper"
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
	GROUP_DETAIL_EXPECT_BODY = `{
    "id": "40f0e6f9-cc36-49aa-9c73-856c34bcc915",
    "name": "hoge family",
    "ownerId": "5b4ccb43-81ab-4357-8591-95b42d42e339",
    "memberNum": 1,
    "members": [
        {
            "id": "5b4ccb43-81ab-4357-8591-95b42d42e339",
            "email": "hoge@example.com",
            "name": "hoge",
            "birthday": "0001-01-01T00:00:00Z",
            "location": "",
            "status": 0,
            "groups": null,
            "roles": null,
            "createdAt": "2023-07-16T00:00:00Z",
            "updatedAt": "2023-07-16T00:00:00Z"
        }
    ],
    "definedRoles": null,
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`
	GROUPS_EXPECT_BODY = `[
    {
        "id": "40f0e6f9-cc36-49aa-9c73-856c34bcc915",
        "name": "hoge family",
        "ownerId": "5b4ccb43-81ab-4357-8591-95b42d42e339",
        "memberNum": 1,
        "members": null,
        "definedRoles": null,
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    },
    {
        "id": "ce81e3d5-9809-4ecb-b53f-a3d05543153e",
        "name": "others family",
        "ownerId": "9da5acd7-4e23-484b-a936-565ddfd09293",
        "memberNum": 2,
        "members": null,
        "definedRoles": null,
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    }
]`
)

func TestController_GroupDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	var userInGroup entity.User
	helper.DeepCopy(&userStub, &userInGroup)
	userStub.Groups = []entity.Group{
		{
			ID:        groupID,
			Name:      "hoge family",
			OwnerID:   userID,
			MemberNum: 1,
			Members:   []entity.User{userInGroup},
			CreatedAt: date,
			UpdatedAt: date,
		},
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/groups/:groupID", appCtrl.GroupDetail())

	endpoint := fmt.Sprintf("/api/v1/groups/%s", groupID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, GROUP_DETAIL_EXPECT_BODY, w.Body.String())
}

func TestController_AuthError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(nil, fmt.Errorf("Unauthorized."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/groups/:groupID", appCtrl.GroupDetail())

	endpoint := fmt.Sprintf("/api/v1/groups/%s", groupID.String())
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Unauthorized.\"\n}", w.Body.String())
}

func TestController_NotFound(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID, _ := uuid.Parse(TESTING_USER_ID)
	groupID, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	var userInGroup entity.User
	helper.DeepCopy(&userStub, &userInGroup)
	userStub.Groups = []entity.Group{
		{
			ID:        groupID,
			Name:      "hoge family",
			OwnerID:   userID,
			MemberNum: 1,
			Members:   []entity.User{userInGroup},
			CreatedAt: date,
			UpdatedAt: date,
		},
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/groups/:groupID", appCtrl.GroupDetail())

	requestGroupID := "40f0e6f9-cc36-49aa-9c73-856c34bcc91c"
	endpoint := fmt.Sprintf("/api/v1/groups/%s", requestGroupID)
	request, _ := http.NewRequest("GET", endpoint, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 404, w.Code)
	expectBody := fmt.Sprintf("{\n    \"error\": \"Not found group: %s\"\n}", requestGroupID)
	assert.Equal(t, expectBody, w.Body.String())
}

func TestController_Groups(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)

	userID0, _ := uuid.Parse(TESTING_USER_ID)
	userID1, _ := uuid.Parse("9da5acd7-4e23-484b-a936-565ddfd09293")
	groupID0, _ := uuid.Parse("40f0e6f9-cc36-49aa-9c73-856c34bcc915")
	groupID1, _ := uuid.Parse("ce81e3d5-9809-4ecb-b53f-a3d05543153e")
	userStub := entity.User{
		ID:        userID0,
		Email:     "hoge@example.com",
		Name:      "hoge",
		CreatedAt: date,
		UpdatedAt: date,
	}
	var userInGroup entity.User
	helper.DeepCopy(&userStub, &userInGroup)
	userStub.Groups = []entity.Group{
		{
			ID:        groupID0,
			Name:      "hoge family",
			OwnerID:   userID0,
			MemberNum: 1,
			Members:   []entity.User{userInGroup},
			CreatedAt: date,
			UpdatedAt: date,
		},
		{
			ID:        groupID1,
			Name:      "others family",
			OwnerID:   userID1,
			MemberNum: 2,
			Members: []entity.User{
				{
					ID:        userID1,
					Email:     "fuga@example.com",
					Name:      "fuga",
					CreatedAt: date,
					UpdatedAt: date,
				},
				userInGroup,
			},
			CreatedAt: date,
			UpdatedAt: date,
		},
	}

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID0).Return(&userStub, nil)
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/groups", appCtrl.Groups())

	request, _ := http.NewRequest("GET", "/api/v1/groups", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, GROUPS_EXPECT_BODY, w.Body.String())
}

func TestController_Groups_AuthError(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	userID0, _ := uuid.Parse(TESTING_USER_ID)

	mockAppRepo := mock.NewMockRepository(mockCtrl)
	mockAppRepo.EXPECT().GetUserByID(userID0).Return(nil, fmt.Errorf("Unauthorized."))
	appCtrl := NewController(mockAppRepo)

	r := gin.Default()
	r.Use(middlewareStub)
	r.GET("/api/v1/groups", appCtrl.Groups())

	request, _ := http.NewRequest("GET", "/api/v1/groups", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)

	assert.Equal(t, 500, w.Code)
	assert.Equal(t, "{\n    \"error\": \"Unauthorized.\"\n}", w.Body.String())
}
