package helper

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MetaTpl struct {
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalData int `json:"total_data"`
}

type BasePayload struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
	Meta    *MetaTpl    `json:"meta,omitempty"`
}

func WriteSuccess(c *gin.Context, message string, data interface{}, meta *MetaTpl) {
	c.JSON(http.StatusOK, BasePayload{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func WriteError(c *gin.Context, err error) {
	httpStatusCode := http.StatusInternalServerError
	payload := BasePayload{
		Error:   ErrorCodeGeneralError,
		Success: false,
		Message: fmt.Sprintf("fatal error: %s", err.Error()),
	}

	if err, ok := err.(*Err); ok {
		payload.Message = err.Error()
		payload.Error = err.GetErrorCode()
		httpStatusCode = err.GetHttpStatusCode()
	}

	c.JSON(httpStatusCode, payload)
}
