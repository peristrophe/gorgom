package route

import (
	"encoding/json"
	"gorgom/internal/mock"
	"gorgom/internal/util"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func prepareRouter(
	mc *gomock.Controller,
	stub func(*gin.Context),
	handlerName string,
) *gin.Engine {
	otherStub := func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "NG"})
	}

	mockAppCtrl := mock.NewMockController(mc)
	mockRecorder := mockAppCtrl.EXPECT()

	tv := reflect.TypeOf(mockRecorder)
	rv := reflect.ValueOf(mockRecorder)

	for i := 0; i < tv.NumMethod(); i++ {
		tm := tv.Method(i)
		rm := rv.Method(i)
		rets := rm.Call([]reflect.Value{})
		ret := rets[0].MethodByName("Return")

		// mockAppCtrl.EXPECT().Handler().Return(func(*gin.Context)) ...
		if tm.Name == handlerName {
			ret.Call([]reflect.Value{reflect.ValueOf(stub)})
		} else {
			ret.Call([]reflect.Value{reflect.ValueOf(otherStub)})
		}
	}

	appRoute := NewRoute(mockAppCtrl)
	r := appRoute.Setup()
	return r
}

func routeTester(
	t *testing.T,
	request *http.Request,
	handlerName string,
	handlerStub func(*gin.Context),
	withToken bool,
) *httptest.ResponseRecorder {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	r := prepareRouter(mockCtrl, handlerStub, handlerName)

	w := httptest.NewRecorder()

	if withToken {
		userID, _ := uuid.Parse("5b4ccb43-81ab-4357-8591-95b42d42e339")
		token := util.NewJWT(userID.String())
		if token == nil {
			t.Error("NewJWT failed.")
		}
		request.AddCookie(&http.Cookie{Name: "token", Value: string(*token)})
	}

	r.ServeHTTP(w, request)
	return w
}

func TestRoute_SignUp(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true

		var postParams struct{ Email, Password string }
		c.BindJSON(&postParams)
		c.JSON(
			http.StatusOK,
			gin.H{
				"message":  "OK",
				"email":    postParams.Email,
				"password": postParams.Password,
			},
		)
	}

	requestBodyObj := struct{ Email, Password string }{Email: "hoge@example.com", Password: "hogehoge"}
	requestBodyBytes, _ := json.Marshal(requestBodyObj)
	reader := strings.NewReader(string(requestBodyBytes))
	request, _ := http.NewRequest("POST", "/api/v1/signup", reader)

	w := routeTester(t, request, "SignUp", handlerStub, false)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"email":"hoge@example.com","message":"OK","password":"hogehoge"}`, w.Body.String())
}

func TestRoute_SignIn(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true

		var postParams struct{ Email, Password string }
		c.BindJSON(&postParams)
		c.JSON(
			http.StatusOK,
			gin.H{
				"message":  "OK",
				"email":    postParams.Email,
				"password": postParams.Password,
			},
		)
	}

	requestBodyObj := struct{ Email, Password string }{Email: "hoge@example.com", Password: "hogehoge"}
	requestBodyBytes, _ := json.Marshal(requestBodyObj)
	reader := strings.NewReader(string(requestBodyBytes))
	request, _ := http.NewRequest("POST", "/api/v1/signin", reader)

	w := routeTester(t, request, "SignIn", handlerStub, false)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"email":"hoge@example.com","message":"OK","password":"hogehoge"}`, w.Body.String())
}

func TestRoute_MyPage(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
			},
		)
	}

	request, _ := http.NewRequest("GET", "/api/v1/users/mypage", nil)

	w := routeTester(t, request, "MyPage", handlerStub, true)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"OK"}`, w.Body.String())
}

func TestRoute_Groups(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
			},
		)
	}

	request, _ := http.NewRequest("GET", "/api/v1/groups", nil)

	w := routeTester(t, request, "Groups", handlerStub, true)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"OK"}`, w.Body.String())
}

func TestRoute_GroupDetail(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
				"groupID": c.Param("groupID"),
			},
		)
	}

	request, _ := http.NewRequest("GET", "/api/v1/groups/ABCDEFG", nil)

	w := routeTester(t, request, "GroupDetail", handlerStub, true)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"groupID":"ABCDEFG","message":"OK"}`, w.Body.String())
}

func TestRoute_Boards(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
			},
		)
	}

	request, _ := http.NewRequest("GET", "/api/v1/boards", nil)

	w := routeTester(t, request, "Boards", handlerStub, true)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"message":"OK"}`, w.Body.String())
}

func TestRoute_BoardDetail(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
				"boardID": c.Param("boardID"),
			},
		)
	}

	request, _ := http.NewRequest("GET", "/api/v1/boards/ABCDEFG", nil)

	w := routeTester(t, request, "BoardDetail", handlerStub, true)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"boardID":"ABCDEFG","message":"OK"}`, w.Body.String())
}

func TestRoute_CardDetail(t *testing.T) {
	stubCalled := false
	handlerStub := func(c *gin.Context) {
		stubCalled = true
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "OK",
				"cardID":  c.Param("cardID"),
			},
		)
	}

	request, _ := http.NewRequest("GET", "/api/v1/cards/ABCDEFG", nil)

	w := routeTester(t, request, "CardDetail", handlerStub, true)
	assert.True(t, stubCalled)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"cardID":"ABCDEFG","message":"OK"}`, w.Body.String())
}
