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

const MYPAGE_EXPECT_BODY = `{
    "id": "5b4ccb43-81ab-4357-8591-95b42d42e339",
    "email": "hoge@example.com",
    "name": "",
    "birthday": "0001-01-01T00:00:00Z",
    "location": "",
    "status": 0,
    "groups": [
        {
            "id": "4e4d3517-237e-11ee-b7fd-0242ac383802",
            "name": "",
            "ownerId": "00000000-0000-0000-0000-000000000000",
            "members": null,
            "roles": null,
            "createdAt": "0001-01-01T00:00:00Z",
            "updatedAt": "0001-01-01T00:00:00Z"
        }
    ],
    "roles": [
        {
            "id": "5f411621-ddd3-4568-a9cb-a6d4e54f6ade",
            "name": "",
            "groupId": "4e4d3517-237e-11ee-b7fd-0242ac383802",
            "createdAt": "0001-01-01T00:00:00Z",
            "updatedAt": "0001-01-01T00:00:00Z"
        }
    ],
    "createdAt": "2023-07-16T00:00:00Z",
    "updatedAt": "2023-07-16T00:00:00Z"
}`

func TestUser_MyPage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	groupID, _ := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	roleID, _ := uuid.Parse("5f411621-ddd3-4568-a9cb-a6d4e54f6ade")
	userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	userStub := entity.User{
		ID:        userID,
		Email:     "hoge@example.com",
		Groups:    []entity.Group{{ID: groupID}},
		Roles:     []entity.Role{{ID: roleID, GroupID: groupID}},
		CreatedAt: date,
		UpdatedAt: date,
	}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByID(userID).Return(&userStub, nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT(userID.String())
	if token == nil {
		panic(fmt.Errorf("NewJWT failed."))
	}

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/users/mypage", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, MYPAGE_EXPECT_BODY, w.Body.String())
}
