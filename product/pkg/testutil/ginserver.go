package testutil

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

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

func GetGinTestContextWithBody(bodyJsonString string) (c *gin.Context, w *httptest.ResponseRecorder) {
	// create a GIN context for test
	w = httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		Body:   io.NopCloser(strings.NewReader(bodyJsonString)),
	}

	return ctx, w
}
