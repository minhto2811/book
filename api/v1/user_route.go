package route_v1

import (
	controller_v1 "book/controllers/v1"

	"github.com/gin-gonic/gin"
)

type UserRouteV1 struct {
	Controller  controller_v1.UserController
	RouterGroup *gin.RouterGroup
}

func (route *UserRouteV1) InitRoutes() {
	items := route.RouterGroup.Group("/users")
	{
		items.POST("/sign-in", route.Controller.SignIn)
		items.POST("/sign-up", route.Controller.SignUp)
		items.GET("/refresh-token", route.Controller.RefreshToken)
	}
}
