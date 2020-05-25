package config

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"goadmin/module/ai/controller"
)

func InitRouter() {
	urlPath := g.Config().GetString("url-path")
	s := g.Server()
	s.Group(urlPath+"/animals", func(g *ghttp.RouterGroup) {
		// 首页展示
		birdController := new(controller.AnimalsController)
		g.ALL("/", birdController.Index)
		g.ALL("/page", birdController.Page)
		g.ALL("/get/{id}", birdController.Get)
		g.ALL("/save", birdController.Save)
		g.ALL("/update", birdController.Update)
		g.ALL("/delete", birdController.Delete)

	})
}