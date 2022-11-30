package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	HTTP_Header_Key_Content_type        string = "Content-type"
	HTTP_Header_Value_Content_type_JSON string = "application/json"
	HTTP_Header_Key_Response_format     string = "Response-Type"
	HTTP_Header_Value_Response_format   string = "LINE"
	HTTP_Header_Key_UUid                string = "chat_user_id"
)

type ResponseFail struct {
	Status  string `json:"status" example:"400"`
	Message string `json:"message" example:"Bad request"`
}

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Datas   interface{} `json:"datas"`
}

func success(_msg string, _datas interface{}) ResponseSuccess {
	return ResponseSuccess{
		Status:  "success",
		Message: _msg,
		Datas:   _datas}
}

func fail(_msg string) ResponseFail {
	return ResponseFail{
		Status:  "fail",
		Message: _msg}
}

func HandleError(c *gin.Context, err error) {
	// logf.Logs.e(err.Error())
	// logf.Logs.
	switch e := err.(type) {
	// case AppError:
	// 	c.JSON(e.Code, fail(e.Message))
	case error:
		c.JSON(http.StatusInternalServerError, fail(e.Error()))
	}
}

func HandleSuccess(c *gin.Context, message string, responseData interface{}) {
	c.JSON(http.StatusOK, success(message, responseData))
}

func HandleFail(c *gin.Context, message string) {
	c.JSON(http.StatusOK, fail(message))
}

// type AppError struct {
// 	Code    int
// 	Message string
// }
