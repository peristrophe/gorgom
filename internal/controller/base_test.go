package controller

import (
	"gorgom/internal/middleware"

	"github.com/gin-gonic/gin"
)

const TESTING_USER_ID = "5b4ccb43-81ab-4357-8591-95b42d42e339"

func middlewareStub(c *gin.Context) {
	c.AddParam(middleware.CONTEXT_PARAM_KEY_AUTH_USER_ID, TESTING_USER_ID)
	c.Next()
}
