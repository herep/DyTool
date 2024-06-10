package helloWorld

import (
	"github.com/gin-gonic/gin"
	"tbTool/api/service/helloWorld"
	"tbTool/api/tools/common"
	"tbTool/api/tools/lvccx"
	"tbTool/pkg"
)

type HelloWorldHandle struct {
	he helloWorld.HelloWorldService
}

func NewHelloWorldHandle(he helloWorld.HelloWorldService) *HelloWorldHandle {
	return &HelloWorldHandle{
		he: he,
	}
}

func (hh *HelloWorldHandle) GetInfo(c *gin.Context) pkg.Render {
	logToken, cookies := lvccx.GetLogTokenAndCookie()

	return common.Succ(logToken + "/n" + cookies)
}
