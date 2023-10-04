package requestbind

import (
	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/gaogao-asia/golang-template/pkg/utils"
	"github.com/gin-gonic/gin"
)

// BindJson bind json param to struct
//
// Example: /api/v1/users
//
// Body payload: {"phone": "123456789"}
func BindJson[T any](c *gin.Context) (*T, error) {
	var req = new(T)
	err := c.ShouldBindJSON(req)
	if err != nil {
		return nil, errs.ErrBadRequest.WithErr(err)
	}

	paramMap := make(map[string]interface{})
	bytes, _ := utils.JsonMarshal(req)
	_ = utils.JsonUnmarshal(bytes, &paramMap)
	if _, ok := paramMap["password"]; ok {
		paramMap["password"] = "********"
	}

	return req, nil
}

// BindFormOrQuery bind form or query param to struct
//
// Example: /api/v1/users?phone=123456789
func BindFormOrQuery[T any](c *gin.Context) (*T, error) {
	var req = new(T)
	err := c.ShouldBind(req)
	if err != nil {
		return nil, errs.ErrBadRequest.WithErr(err)
	}

	paramMap := make(map[string]interface{})
	bytes, _ := utils.JsonMarshal(req)
	_ = utils.JsonUnmarshal(bytes, &paramMap)

	return req, nil
}

// BindPathParam bind path param to struct
//
// Example: /api/v1/users/:id
func BindPathParam[T any](c *gin.Context) (*T, error) {
	var req = new(T)
	err := c.ShouldBindUri(req)
	if err != nil {
		return nil, errs.ErrBadRequest.WithErr(err)
	}

	paramMap := make(map[string]interface{})
	bytes, _ := utils.JsonMarshal(req)
	_ = utils.JsonUnmarshal(bytes, &paramMap)

	return req, nil
}
