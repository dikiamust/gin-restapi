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
		// Auth routes
		authService := &services.AuthService{DB: db}
		authHandler := &handlers.AuthHandler{AuthService: authService}
		authGroup := api.Group("/auth")
		{
			authGroup.POST("signup", authHandler.RegisterHandler)
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
