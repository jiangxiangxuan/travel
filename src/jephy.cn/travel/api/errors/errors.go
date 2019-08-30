package errors

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorInfo struct {
	status int
	err    gin.H
}

func errorf(status int, msg string) *ErrorInfo {
	return &ErrorInfo{
		status: status,
		err:    gin.H{"msg": msg},
	}
}

func (self *ErrorInfo) Json(c *gin.Context) {
	c.JSON(self.status, self.err)
}

func Result(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

var (
	OK   = errorf(http.StatusOK, "success")
	Fail = errorf(http.StatusForbidden, "fail")
)
