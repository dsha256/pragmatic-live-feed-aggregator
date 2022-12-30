package server

import (
	"github.com/dsha256/pragmatic-live-feed-aggregator/pkg/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleSuccessfulResp(c *gin.Context, msg string, data interface{}) {
	resp := newSuccessfulResp(msg, http.StatusOK, data)
	c.AbortWithStatusJSON(http.StatusOK, resp)
}

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
