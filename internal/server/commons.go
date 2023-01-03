package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// newSuccessfulResp returns Response for successful response
func newSuccessfulResp(msg string, code int, data interface{}) dto.Response {
	return dto.Response{
		Data:    data,
		Error:   false,
		Message: msg,
		Status:  code,
		Code:    code,
	}
}

func newErrResp(msg string, code int, internalCode int) dto.Response {
	return dto.Response{
		Message: msg,
		Error:   true,
		Status:  code,
		Code:    internalCode,
	}
}

func handleSuccessfulResp(c *gin.Context, msg string, data interface{}) {
	resp := newSuccessfulResp(msg, http.StatusOK, data)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

func handleErrResp(c *gin.Context, msg string, code int) {
	resp := newErrResp(msg, code, code)
	c.AbortWithStatusJSON(code, resp)
}
