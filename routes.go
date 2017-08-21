package main

import (
	"github.com/rof20004/go-echo-rest-template/modules"
)

// InitRoutes - start routes
func InitRoutes() {
	// Root path
	main := server.Group("/go-echo-rest-template")

	// User resources
	userResources := modules.NewUserController(DB)
	userRoutes := main.Group("/v1/users")
	// userRoutes.Use(RequestLog)
	// userRoutes.Use(middleware.BodyDump(ResponseLog))
	userRoutes.GET("", userResources.List)
	userRoutes.POST("", userResources.Create)
	userRoutes.GET("/:id", userResources.Read)
	userRoutes.PUT("/:id", userResources.Update)
	userRoutes.DELETE("/:id", userResources.Delete)
}
