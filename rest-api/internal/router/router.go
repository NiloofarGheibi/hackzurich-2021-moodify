package router

import (
	"github.com/gin-gonic/gin"
	"rest-api/internal/controllers"
)

// Router type to define different routing groups and gin engine
type Router struct {
	Router     *gin.Engine
	UserGroup  *gin.RouterGroup
	RadioGroup *gin.RouterGroup
}

// Create functions returns gin Engine with specified handler functions
func Create() *Router {
	r := Router{Router: gin.New()}
	r.Router.Use(gin.Recovery())
	// SetUserGroup set user group in the router
	r.UserGroup = r.Router.Group("/users")
	r.RadioGroup = r.Router.Group("/radios")
	return &r
}

// SetHealthEndPoint get health handler as an input and sets health endpoint
func (r *Router) SetHealthEndPoint(health gin.HandlerFunc) {
	r.Router.GET("/health", health)
}

// InitUserRoutes initializes the routes in the user group in the Router
func (r *Router) InitUserRoutes(userController *controllers.User) *Router {
	r.UserGroup.GET(
		"",
		userController.GetAll,
	)

	r.UserGroup.GET(
		"/:userid",
		userController.GetByID,
	)

	r.UserGroup.POST(
		"",
		userController.RegisterUser,
	)

	r.UserGroup.PUT(
		"/:userid",
		userController.UpdateUser,
	)

	r.UserGroup.DELETE(
		"/:userid",
		userController.DeleteUser,
	)

	return r
}

// InitRadioRoutes initializes the routes in the radio group in the Router
func (r *Router) InitRadioRoutes(radioController *controllers.Radio) *Router {
	r.RadioGroup.POST(
		"",
		radioController.RegisterTrack,
	)

	return r
}
