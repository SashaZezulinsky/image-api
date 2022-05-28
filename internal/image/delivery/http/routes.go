package http

import (
	"github.com/labstack/echo/v4"

	"github.com/image-api/internal/domain"
)

//MapImageRoutes Map image routes
func MapImageRoutes(imageGroup *echo.Group, h domain.ImageHandlers) {
	imageGroup.GET("", h.ListMetadata())
	imageGroup.GET("/:id", h.GetMetadata())

	imageGroup.GET("/:id/data", h.Get())
	imageGroup.POST("", h.Add())
	imageGroup.PUT("/:id", h.Update())
}
