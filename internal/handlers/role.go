package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "go-restapi-gin/internal/models"
	"go-restapi-gin/internal/services"
)

type RoleHandler struct {
	RoleService *services.RoleService
}

// CreateRoleHandler handles the creation of a new role
func (h *RoleHandler) CreateRoleHandler(c *gin.Context) {
	var input services.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.RoleService.CreateRole(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, role)
}

// GetRolesHandler retrieves all roles
func (h *RoleHandler) GetRolesHandler(c *gin.Context) {
	roles, err := h.RoleService.GetRoles()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, roles)
}

// UpdateRoleHandler handles updating an existing role
func (h *RoleHandler) UpdateRoleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	var input services.RoleInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, err := h.RoleService.UpdateRole(uint(id), input)
	if err != nil {
		if err.Error() == "role not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, role)
}

// DeleteRoleHandler handles deleting a role
func (h *RoleHandler) DeleteRoleHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role ID"})
		return
	}

	if err := h.RoleService.DeleteRole(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Role deleted successfully"})
}

// RegisterRoleRoutes registers the role-related routes
func (h *RoleHandler) RegisterRoleRoutes(router *gin.Engine) {
	roleGroup := router.Group("/roles")
	{
		roleGroup.POST("", h.CreateRoleHandler)
		roleGroup.GET("", h.GetRolesHandler)
		roleGroup.PUT(":id", h.UpdateRoleHandler)
		roleGroup.DELETE(":id", h.DeleteRoleHandler)
	}
}
