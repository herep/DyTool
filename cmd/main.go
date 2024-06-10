package main

import (
	"gitlab.xfq.com/tech-lab/dionysus"
	"gitlab.xfq.com/tech-lab/dionysus/cmd/gincmd"
	"gitlab.xfq.com/tech-lab/dionysus/pkg/conf"
	"gitlab.xfq.com/tech-lab/dionysus/pkg/logger"
	dredis "gitlab.xfq.com/tech-lab/dionysus/pkg/redis"
	"go.uber.org/dig"
	"log"
	helloWorldHandler "tbTool/api/handler/helloWorld"
	itemsHandler "tbTool/api/handler/items"
	lvHandler "tbTool/api/handler/lvccx"
	"tbTool/api/routers"
	"tbTool/api/service/helloWorld"
	"tbTool/api/service/items"
	"tbTool/api/service/lvccx"
)

func initContainer() *dig.Container {
	c := dig.New()

	itemSrvErr := c.Provide(items.NewItemServiceImpl)
	itemHandErr := c.Provide(itemsHandler.NewItemsOnSaleGetHandler)
	logger.Fatalf("initContainer start items result:%v,%v", itemSrvErr, itemHandErr)

	helloSrvErr := c.Provide(helloWorld.NewHelloWorldServiceImpl)
	helloHandErr := c.Provide(helloWorldHandler.NewHelloWorldHandle)
	logger.Fatalf("initContainer start hello result:%v,%v", helloSrvErr, helloHandErr)

	//吕橙服务
	lvSrvErr := c.Provide(lvccx.NewLvOrderServiceImpl)
	lvHandErr := c.Provide(lvHandler.NewOrderOnGetHandler)
	logger.Fatalf("initContainer start hello result:%v,%v", lvSrvErr, lvHandErr)

	return c
}

func main() {
	g := gincmd.New()

	err := g.RegPreRunFunc("watch.redis", 1, func() error {
		return conf.RegisterEtcdWatch(&dredis.RedisEvent{Prefix: "watch.redis"})
	})
	if err != nil {
		log.Println("Reg pre run func err:", err)
	}

	err = g.RegPreRunFunc("business", 2, func() error {
		return conf.StartWatchConfig("business")
	})
	if err != nil {
		log.Println("Reg pre run func err:", err)
	}

	err = g.RegPreRunFunc("watch.mysql", 3, func() error {
		return conf.StartWatchConfig("watch.mysql")
	})
	if err != nil {
		log.Println("Reg pre run func err:", err)
	}

	_ = g.RegPreRunFunc("initContainer", 5, func() error {
		//依赖注入
		log.Println("initContainer start")
		c := initContainer()

		//路由注入
		log.Println("RegisterRouter start step")
		routers.RegisterRouter(c, g.Engine)
		return nil
	})

	dionysus.Start("gapi", g)
}
