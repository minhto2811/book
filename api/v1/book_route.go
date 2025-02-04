package route_v1

import (
	"book/controllers/v1"
	"book/midleware"

	"github.com/gin-gonic/gin"
)

type BookRouteV1 struct {
	Controller  controller_v1.BookController
	RouterGroup *gin.RouterGroup
}

func (route *BookRouteV1) InitRoutes() {
	items := route.RouterGroup.Group("/books", midleware.Authentication())
	{
		items.POST("/create", route.Controller.Create)
		items.GET("/get-all",route.Controller.GetAll)
		items.GET("/:id",route.Controller.GetById)
		items.DELETE("/:id",route.Controller.DeleteById)
		items.PUT("/:id",route.Controller.UpdateAllField)
		items.PATCH("/:id",route.Controller.UpdateFields)
	}
	
}
