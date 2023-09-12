package app

import (
	"github.com/flosch/pongo2"
	"github.com/gin-gonic/gin"
	"gitlab.com/go-box/pongo2gin/v6"
	"net/http"
)

type App interface {
	Run(addr ...string) error
	Use(middleware ...gin.HandlerFunc)
	GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes
}

type app struct {
	engine *gin.Engine
}

func (a app) GET(relativePath string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return a.engine.GET(relativePath, handlers...)
}

func (a app) Use(middleware ...gin.HandlerFunc) {
	a.engine.Use(middleware...)
}

func (a app) Run(addr ...string) error {
	return a.engine.Run(addr...)
}

func NewDefaultApp() App {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.HTMLRender = pongo2gin.New(pongo2gin.RenderOptions{
		TemplateDir: "templates",
		TemplateSet: nil,
	})
	engine.GET("/docs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "docs.html", pongo2.Context{})
	})
	return &app{
		engine: engine,
	}
}
