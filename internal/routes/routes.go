package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"go-restapi-gin/internal/handlers"
	"go-restapi-gin/internal/services"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	api := router.Group("/api")
	{
		// // Dummy route for users
		// api.GET("/users", func(c *gin.Context) {
		// 	c.JSON(200, gin.H{
		// 		"message": "Get all users (dummy route)",
		// 	})
		// })

		// User routes
		userService := &services.UserService{DB: db}
		userHandler := &handlers.UserHandler{UserService: userService}
		userGroup := api.Group("/users")
		{
			userGroup.POST("", userHandler.RegisterHandler)
		}

		// Role routes
		roleService := &services.RoleService{DB: db}
		roleHandler := &handlers.RoleHandler{RoleService: roleService}
		roleGroup := api.Group("/roles")
		{
			roleGroup.POST("", roleHandler.CreateRoleHandler)
			roleGroup.GET("", roleHandler.GetRolesHandler)
			roleGroup.PUT(":id", roleHandler.UpdateRoleHandler)
			roleGroup.DELETE(":id", roleHandler.DeleteRoleHandler)
		}
	}
}
