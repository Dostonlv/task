package http

import (
	"github.com/Dostonlv/task/internal/blogs"
	"github.com/labstack/echo/v4"
)

// MapBlogsRoutes maps routes for blogs
func MapBlogsRoutes(blogGroup *echo.Group, h blogs.Handlers) {
	blogGroup.POST("", h.Create())
	blogGroup.PUT("/:id", h.Update())
	blogGroup.DELETE("/:id", h.Delete())
	blogGroup.GET("/:id", h.GetByID())
	blogGroup.GET("", h.GetAll())
}
