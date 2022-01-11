package webbase

import "github.com/gin-gonic/gin"

type CommonResp struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ServeResponse(c *gin.Context, errMsg *CommonResp, args ...interface{}) {
	if nil == errMsg {
		return
	}

	data := errMsg.Data
	if len(args) > 0 {
		if nil != args[0] {
			data = args[0]
		}
	}

	result := &CommonResp{
		Code:    errMsg.Code,
		Message: errMsg.Message,
		Data:    data,
	}
	c.JSON(200, result)
}
