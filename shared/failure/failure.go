package failure

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

var (
	ErrBadRequest       = errors.New("bad request")
	ErrInternalError    = errors.New("internal server error")
	ErrUnimplemented    = errors.New("unimplemented")
	ErrForbidden        = errors.New("forbidden")
	ErrNotFound         = errors.New("not found")
	ErrUnauthorized     = errors.New("unauthorized")
	ErrConflict         = errors.New("conflict")
	ErrFailedDependency = errors.New("failed dependency")
)

const (
	HTTPStatusCodePrefixBadRequest = "Bad Request: "
	HTTPStatusCodePrefixForbidden  = "Forbidden: "
	HTTPStatusCodePrefixNotFound   = "Not Found: "
	HTTPStatusCodePrefixConflict   = "Conflict: "
	HTTPStatusCodePrefixInternal   = "Internal Server Error: "
)

type Failure struct {
	Code      int    `json:"code"`
	ErrorCode string `json:"errorCode"`
	Message   string `json:"message"`
}

func (e *Failure) Error() string {
	return fmt.Sprintf("%s: %s", http.StatusText(e.Code), e.Message)
}

func Init(code int, err error) *Failure {
	if err != nil {
		return &Failure{
			Code:    code,
			Message: err.Error(),
		}
	}
	return nil
}

func New(code int, err error) error {
	if err != nil {
		errCode, _ := ParseErrorCode(err.Error())
		return &Failure{
			Code:      code,
			Message:   err.Error(),
			ErrorCode: errCode,
		}
	}
	return nil
}

func BadRequest(err error) error {
	if err != nil {
		errCode, _ := ParseErrorCode(err.Error())
		return &Failure{
			Code:      http.StatusBadRequest,
			ErrorCode: errCode,
			Message:   err.Error(),
		}
	}
	return nil
}

func BadRequestFromStringf(msg string, messageParams ...any) error {
	errCode, _ := ParseErrorCode(msg)
	return &Failure{
		Code:      http.StatusBadRequest,
		ErrorCode: errCode,
		Message:   fmt.Sprintf(msg, messageParams...),
	}
}

func BadRequestFromString(msg string) error {
	errCode, _ := ParseErrorCode(msg)
	return &Failure{
		Code:      http.StatusBadRequest,
		Message:   msg,
		ErrorCode: errCode,
	}
}

func NotFoundFromString(msg string) error {
	errCode, _ := ParseErrorCode(msg)
	return &Failure{
		Code:      http.StatusNotFound,
		Message:   msg,
		ErrorCode: errCode,
	}
}

func Unauthorized(msg string) error {
	errCode, _ := ParseErrorCode(msg)
	return &Failure{
		Code:      http.StatusUnauthorized,
		Message:   msg,
		ErrorCode: errCode,
	}
}

func UnprocessableEntity(err error) error {
	if err != nil {
		errCode, _ := ParseErrorCode(err.Error())
		return &Failure{
			Code:      http.StatusUnprocessableEntity,
			Message:   err.Error(),
			ErrorCode: errCode,
		}
	}
	return nil
}

func InternalError(err error) error {
	if err != nil {
		errCode, _ := ParseErrorCode(err.Error())

		return &Failure{
			Code:      http.StatusInternalServerError,
			Message:   err.Error(),
			ErrorCode: errCode,
		}
	}
	return nil
}

func InternalErrorFromString(err string, val ...any) error {
	return InternalError(errors.New(fmt.Sprintf(err, val...)))
}

func GetCode(err error) int {
	if f, ok := err.(*Failure); ok {
		return f.Code
	}
	return http.StatusInternalServerError
}

func GetErrorCode(err error) string {
	if f, ok := err.(*Failure); ok {
		return f.ErrorCode
	}
	return ""
}

func ParseErrorCode(c string) (string, bool) {
	index := strings.Index(c, ":")
	if index == -1 {
		return "", false
	}

	return c[:index], true
}
