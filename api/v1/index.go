package route_v1

import (
	controller_v1 "book/controllers/v1"
	"book/repositories/repo_impl"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RouteV1 struct {
	Db          *gorm.DB
	RouterGroup *gin.RouterGroup
}

func (r *RouteV1) Init() {
	r.userRoute()
	r.bookRoute()
}

func (r *RouteV1) bookRoute(){
	bookRouteV1 := BookRouteV1{
		Controller: controller_v1.BookController{
			BookRepo: repo_impl.NewInstanceBookRepo(r.Db),
		},
		RouterGroup: r.RouterGroup,
	}
	bookRouteV1.InitRoutes()
}

func (r *RouteV1) userRoute() {
	userRouteV1 := UserRouteV1{
		Controller: controller_v1.UserController{
			UserRepo: repo_impl.NewInstanceUserRepo(r.Db),
		},
		RouterGroup: r.RouterGroup,
	}
	userRouteV1.InitRoutes()
}
