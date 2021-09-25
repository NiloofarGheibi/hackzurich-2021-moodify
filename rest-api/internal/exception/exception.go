package exception

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var (
	ErrPermission             = New(http.StatusForbidden, "not enough permission")
	ErrMissingInputParameters = New(http.StatusBadRequest, "missing input")
	ErrInvalidInput           = New(http.StatusBadRequest, "invalid input")
	ErrConflict               = New(http.StatusConflict, "already exists")
	ErrNotFound               = New(http.StatusNotFound, "not found")
	ErrOperationFailed        = New(http.StatusInternalServerError, "operation failed")
	ErrUnauthorized           = New(http.StatusUnauthorized, "unauthorized")
)

// JsonError struct represents "application/problem+json"
type JsonError struct {
	Type     string                 `json:"type,omitempty"`
	Title    string                 `json:"title"`
	Detail   string                 `json:"detail"`
	Status   int                    `json:"status"`
	Instance string                 `json:"instance,omitempty"`
	Errors   []*jsonValidationError `json:"exception,omitempty"`
}

// Error implements error interface
func (p *JsonError) Error() string {
	return strconv.Itoa(p.Status) + " - " + p.Title + " - " + p.Detail
}

type jsonValidationError struct {
	Title    string `json:"title"`
	Type     string `json:"type"`
	Property string `json:"property,omitempty"`
	Param    string `json:"param,omitempty"`
}

func AsHTTPError(ctx *gin.Context, err error) {
	if jsonErr, ok := err.(*JsonError); ok {
		ctx.AbortWithStatusJSON(jsonErr.Status, jsonErr)
	} else {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, New(http.StatusInternalServerError, err.Error()))
	}
}

func New(statusCode int, message string) *JsonError {
	return &JsonError{
		Title:  http.StatusText(statusCode),
		Status: statusCode,
		Detail: message,
	}
}

func (r *JsonError) WithDetail(detail string) *JsonError {
	return &JsonError{
		Title:  r.Title,
		Status: r.Status,
		Detail: r.Detail + ": " + detail,
	}
}
