package routes

import (
	"book/api/v1"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	Db *gorm.DB
}

func (routes *Routes) Init() {
	engine := gin.Default()
	routeV1 := route_v1.RouteV1{Db: routes.Db, RouterGroup: engine.Group("/v1")}
	routeV1.Init()
	engine.Run("localhost:3000")
}
