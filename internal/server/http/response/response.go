package response

import (
	"log"
	"net/http"

	"github.com/gaogao-asia/golang-template/pkg/errs"
	"github.com/gin-gonic/gin"
)

type ResponseBody struct {
	Data  interface{}        `json:"data,omitempty"`
	Error *ErrorResponseBody `json:"error,omitempty"`
}

type ErrorResponseBody struct {
	Code    int    `json:"code,omitempty"`
	MsgCode string `json:"msg_code,omitempty"`
	Message string `json:"msg,omitempty"`
}

func GeneralError(c *gin.Context, err error) {
	var body ResponseBody

	if er, ok := err.(errs.ClientError); ok {
		log.Printf("%s %v\n", errs.CLIENT_ERROR, er)

		body = ResponseBody{
			Error: &ErrorResponseBody{
				Code:    er.Code,
				MsgCode: er.MsgCode,
			},
		}
		filterClientErrorHTTPCode(c, er, body)
		return
	}

	log.Printf("%s %+v\n", errs.SERVER_ERROR, err)

	internalErr := errs.ErrInternalServer
	if er, ok := err.(errs.ServerError); ok {
		internalErr = er
	}

	body = ResponseBody{
		Error: &ErrorResponseBody{
			Code:    internalErr.Code,
			MsgCode: internalErr.MsgCode,
		},
	}

	c.JSON(http.StatusInternalServerError, body)
}

// sendToHTTPCode -.
func filterClientErrorHTTPCode(c *gin.Context, err errs.ClientError, body ResponseBody) {
	switch err {
	case errs.ErrBadRequest,
		errs.ErrCreateNewAccountRequestInvalid,
		errs.ErrCreateNewAccountRequestRoleInvalid:

		c.JSON(http.StatusBadRequest, body)
	case errs.ErrUnAuthorized,
		errs.ErrAccessTokenExpired,
		errs.ErrAccessTokenNotExist:

		c.JSON(http.StatusUnauthorized, body)
	case errs.ErrForbidden,
		errs.ErrUserNotAllow,
		errs.ErrUserCanNotAccess:

		c.JSON(http.StatusForbidden, body)
	case errs.ErrNotFound,
		errs.ErrUserNotExist:

		c.JSON(http.StatusNotFound, body)
	default:
		c.JSON(http.StatusBadRequest, body)
	}
}
