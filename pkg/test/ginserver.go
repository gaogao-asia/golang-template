package test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetGinTestContext() (c *gin.Context, w *httptest.ResponseRecorder) {
	// create a GIN context for test
	w = httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	return ctx, w
}
