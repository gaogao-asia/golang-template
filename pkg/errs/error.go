package errs

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	SERVER_ERROR = "SERVER_ERROR"
	CLIENT_ERROR = "CLIENT_ERROR"
)

var (
	ErrBadRequest                         = makeClientError(400000, "BAD_REQUEST", nil)
	ErrCreateNewAccountRequestInvalid     = makeClientError(400001, "CREATE_ACCOUNT_REQUEST_INVALID", nil)
	ErrCreateNewAccountRequestRoleInvalid = makeClientError(400001, "CREATE_ACCOUNT_REQUEST_ROLE_INVALID", nil)

	ErrUnAuthorized        = makeClientError(401000, "UNAUTHORIZED", nil)
	ErrAccessTokenExpired  = makeClientError(401001, "ACCESS_TOKEN_EXPIRED", nil)
	ErrAccessTokenNotExist = makeClientError(401002, "ACCESS_TOKEN_NOT_EXIST", nil)

	ErrForbidden        = makeClientError(403000, "FORBIDDEN", nil)
	ErrUserNotAllow     = makeClientError(403001, "USER_NOT_ALLOW", nil)
	ErrUserCanNotAccess = makeClientError(403002, "USER_CAN_NOT_ACCESSS", nil)

	ErrNotFound     = makeClientError(404000, "NOT_FOUND", nil)
	ErrUserNotExist = makeClientError(404001, "USER_NOT_EXIST", nil)

	ErrInternalServer = makeServerError(500000, "INTERNAL_SERVER_ERROR", nil)
	ErrDBFailed       = makeServerError(500003, "DB_ERROR", nil)
)

type ClientError struct {
	err
}

type err struct {
	Code      int
	MsgCode   string
	OriginErr error
}

func (e ClientError) Error() string {
	if e.OriginErr == nil {
		return fmt.Sprintf("error_code: %d, msg_code: %s", e.Code, e.MsgCode)
	}

	return fmt.Sprintf("error_code: %d, msg_code: %s, origin_err:%s", e.Code, e.MsgCode, e.OriginErr.Error())
}

func IsClientError(err error) bool {
	_, ok := err.(ClientError)
	return ok
}

func makeClientError(code int, msgCode string, er error) ClientError {
	return ClientError{err: err{
		Code:      code,
		MsgCode:   msgCode,
		OriginErr: er,
	}}
}

func (e ClientError) WithErr(err error) ClientError {
	e.OriginErr = err
	return e
}

type ServerError struct {
	err
}

func (e ServerError) Error() string {
	if e.OriginErr == nil {
		return fmt.Sprintf("error_code: %d, msg_code: %s", e.Code, e.MsgCode)
	}

	return fmt.Sprintf("error_code: %d, msg_code: %s, origin_err:%s", e.Code, e.MsgCode, e.OriginErr.Error())
}

func IsServerError(err error) bool {
	_, ok := err.(ServerError)
	return ok
}

func makeServerError(code int, msgCode string, er error) ServerError {
	return ServerError{err: err{
		Code:      code,
		MsgCode:   msgCode,
		OriginErr: er,
	}}
}

func (e ServerError) WithErr(err error) ServerError {
	e.OriginErr = err
	return e
}

func FilterErrNotFoundOrInternal(err error, newErr error) error {
	if err == ErrNotFound || err == gorm.ErrRecordNotFound {
		return newErr
	}

	return ErrInternalServer.WithErr(err)
}

func IsRecordNotFound(err error) bool {
	return err == ErrNotFound || err == gorm.ErrRecordNotFound
}
