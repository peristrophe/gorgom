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
	GROUP_DETAIL_EXPECT_BODY = `{
    "id": "4e4d3517-237e-11ee-b7fd-0242ac383802",
    "name": "",
    "ownerId": "5b4ccb43-81ab-4357-8591-95b42d42e339",
    "memberNum": 1,
    "members": [
        {
            "id": "5b4ccb43-81ab-4357-8591-95b42d42e339",
            "email": "hoge@example.com",
            "name": "",
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
        "id": "4e4d3517-237e-11ee-b7fd-0242ac383802",
        "name": "",
        "ownerId": "5b4ccb43-81ab-4357-8591-95b42d42e339",
        "memberNum": 1,
        "members": null,
        "definedRoles": null,
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    },
    {
        "id": "0c3bc2ed-032a-4e9a-861e-a890d3ac9263",
        "name": "",
        "ownerId": "a26427d0-b1d4-489f-8484-19513ac5e732",
        "memberNum": 2,
        "members": null,
        "definedRoles": null,
        "createdAt": "2023-07-16T00:00:00Z",
        "updatedAt": "2023-07-16T00:00:00Z"
    }
]`
)

func TestRoute_GroupDetail(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	userID0, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	userID1, _ := uuid.Parse("a26427d0-b1d4-489f-8484-19513ac5e732")
	groupID0, _ := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	groupID1, _ := uuid.Parse("0c3bc2ed-032a-4e9a-861e-a890d3ac9263")
	groups := []entity.Group{
		{
			ID:        groupID0,
			OwnerID:   userID0,
			CreatedAt: date,
			UpdatedAt: date,
		},
		{
			ID:        groupID1,
			OwnerID:   userID1,
			CreatedAt: date,
			UpdatedAt: date,
		},
	}

	users := []entity.User{
		{
			ID:        userID0,
			Email:     "hoge@example.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
		{
			ID:        userID1,
			Email:     "fuga@example.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
	}

	var us []entity.User
	util.DeepCopy(&users, &us)
	groups[0].Members = []entity.User{us[0]}
	groups[1].Members = us

	// It is not automatically calculated because mock 'gorm'.
	groups[0].MemberNum = 1
	groups[1].MemberNum = 2
	users[0].Groups = groups
	users[1].Groups = []entity.Group{groups[1]}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByID(userID0).Return(&users[0], nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT(userID0.String())
	if token == nil {
		panic(fmt.Errorf("NewJWT failed."))
	}

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	endpoint := fmt.Sprintf("/api/v1/groups/%s", groupID0.String())
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, GROUP_DETAIL_EXPECT_BODY, w.Body.String())
}

func TestRoute_Groups(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	date := time.Date(2023, 7, 16, 0, 0, 0, 0, time.UTC)
	userID0, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
	userID1, _ := uuid.Parse("a26427d0-b1d4-489f-8484-19513ac5e732")
	groupID0, _ := uuid.Parse("4e4d3517-237e-11ee-b7fd-0242ac383802")
	groupID1, _ := uuid.Parse("0c3bc2ed-032a-4e9a-861e-a890d3ac9263")
	groups := []entity.Group{
		{ID: groupID0, OwnerID: userID0, CreatedAt: date, UpdatedAt: date},
		{ID: groupID1, OwnerID: userID1, CreatedAt: date, UpdatedAt: date},
	}

	users := []entity.User{
		{
			ID:        userID0,
			Email:     "hoge@example.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
		{
			ID:        userID1,
			Email:     "fuga@example.com",
			CreatedAt: date,
			UpdatedAt: date,
		},
	}

	// It is not automatically calculated because mock 'gorm'.
	groups[0].MemberNum = 1
	groups[1].MemberNum = 2
	users[0].Groups = groups
	users[1].Groups = []entity.Group{groups[1]}

	mockRepo := mock.NewMockRepository(mockCtrl)
	mockRepo.EXPECT().GetUserByID(userID0).Return(&users[0], nil)
	appCtrl := controller.NewController(mockRepo)
	appRoute := NewRoute(appCtrl)

	token := util.NewJWT(userID0.String())
	if token == nil {
		panic(fmt.Errorf("NewJWT failed."))
	}

	r := appRoute.Setup()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/groups", nil)
	req.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, GROUPS_EXPECT_BODY, w.Body.String())
}
