package routers

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
	"log"
	"tbTool/api/handler/helloWorld"
	"tbTool/api/handler/items"
	"tbTool/api/handler/lvccx"
	. "tbTool/api/middleware"
)

func RegisterRouter(c *dig.Container, e *gin.Engine) {

	api := e.Group("/tbApi")
	api.Use(Sign())

	if err := c.Invoke(func(h *items.ItemOnSaleGetHandler) {

		api.POST("items/ItemsOnSaleGet", func(ctx *gin.Context) { ctx.Render(200, h.TaoBaoItemsOnSaleGet(ctx)) })

	}); err != nil {
		log.Fatalf("%s", err)
	}

	if err := c.Invoke(func(h *helloWorld.HelloWorldHandle) {

		api.POST("helloWorld/GetInfo", func(ctx *gin.Context) { ctx.Render(200, h.GetInfo(ctx)) })

	}); err != nil {
		log.Fatalf("%s", err)
	}

	//吕橙服务 路由
	lvApi := e.Group("/lvApi")
	if err := c.Invoke(func(l *lvccx.OrderOnGetHandler) {

		lvApi.POST("lvC/GetInfo", func(ctx *gin.Context) { ctx.Render(200, l.GetOrderList(ctx)) })

	}); err != nil {
		log.Fatalf("%s", err)
	}
}
