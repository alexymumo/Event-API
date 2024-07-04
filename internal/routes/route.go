package routes

import (
	"events/internal/controllers"
	"events/internal/repository"
	"events/internal/service"

	"github.com/gin-gonic/gin"
)

var (
	repo          repository.EventRepository  = repository.EventRepsotoryImpl()
	eventservices service.EventService        = service.EventServiceImpl(repo)
	controller    controllers.EventController = controllers.EventControllerImpl(eventservices)
)

func UserRoutes(route *gin.Engine) {
	route.POST("v1/auth/register")
	route.POST("v1/auth/login")
}

func EventRoutes(route *gin.Engine) {
	route.POST("v1/event")
	route.GET("v1/events")
	route.GET("v1/event:id")

}
