package handler

import (
	"fortune/pkg/errno"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListModel struct {
	Code     int         `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
	Total    int         `json:"total"`
	TotalNum int         `json:"total_num"`
}

type EditorModel struct {
	Code    int      `json:"code"`
	Message string   `json:"message"`
	Data    []string `json:"data"`
	ErrNo   int      `json:"errno"`
}

type ErrData struct {
	ErrMessage string `json:"err_message"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendList(c *gin.Context, err error, data interface{}, total, totalNum int) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, ListModel{
		Code:     code,
		Message:  message,
		Data:     data,
		Total:    total,
		TotalNum: totalNum,
	})
}

// wangEditor富文本编辑器文件上传返回专用
func EditorResponse(c *gin.Context, err error, data []string, errNo int) {
	code, message := errno.DecodeErr(err)

	c.JSON(http.StatusOK, EditorModel{
		Code:    code,
		Message: message,
		Data:    data,
		ErrNo:   errNo,
	})
}
