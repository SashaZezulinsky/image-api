package http

import (
	"github.com/image-api/internal/domain"
	"github.com/labstack/echo/v4"
)

// Map image routes
func MapImageRoutes(imageGroup *echo.Group, h domain.ImageHandlers) {
	imageGroup.GET("/", h.ListMetadata())
	imageGroup.GET("/:id", h.GetMetadata())

	imageGroup.GET("/:id/data", h.Get())
	imageGroup.POST("/", h.Add())
	imageGroup.PUT("/:id", h.Update())
}
