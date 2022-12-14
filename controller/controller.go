package controller

import (
	"fmt"
	"net/http"

	"C-SquaredService/service"
	"C-SquaredService/service/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	*gin.Engine
}

func NewController(e *gin.Engine) *Controller {
	return &Controller{e}
}

// 設定Router
func (r *Controller) Router() {
	r.Use(middlewares.CorsMiddleware())
	r.Use(cors.Default())
	api := r.Group("/api")
	{
		api.Use(middlewares.DBMiddleware())
		api.Use(middlewares.CorsMiddleware())
		api.GET("/version", version)
		v1 := api.Group("/v1")
		{
			v1.Use(middlewares.CorsMiddleware())
			v1.GET("/param/:first", Param)
			v1.GET("/query", service.Test)
			v1.POST("/user", service.CreateUser)
			v1.GET("/classify", service.GetProductClassify)
			v1.GET("/products/:id", service.GetProducts)
			v1.GET("/product/:id", service.GetProduct)
		}
		back := api.Group("/back")
		{
			back.Use(middlewares.CorsMiddleware())
			back.GET("/classify", service.GetProductClassify)
			back.POST("/product", service.InsertProduct)
		}
	}
}
func version(c *gin.Context) {
	c.String(http.StatusOK, fmt.Sprintf("%v V:%v-%v,Build:%v", "demo", "1.0.1", "Local", "2022/01/01 15:32:10"))
}

func Param(c *gin.Context) {
	firstParameter := c.Param("first")
	c.String(http.StatusOK, "Param = %v\n", firstParameter)
}
