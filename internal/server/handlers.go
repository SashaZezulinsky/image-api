package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	imageHttp "github.com/image-api/internal/image/delivery/http"
	imageRepo "github.com/image-api/internal/image/repository/mongodb"
	imageUsecase "github.com/image-api/internal/image/usecase"
	"github.com/image-api/pkg/utils"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	iRepo, err := imageRepo.NewMongoDBImageRepository(s.mongoDB, s.cfg.MongoDB.Database)
	if err != nil {
		return err
	}
	// Init useCases
	imageUC := imageUsecase.NewImageUsecase(iRepo)

	// Init handlers
	imagesHandlers := imageHttp.NewImageHandler(imageUC)

	v1 := e.Group("/v1")

	health := v1.Group("/health")
	imageGroup := v1.Group("/images")

	imageHttp.MapImageRoutes(imageGroup, imagesHandlers)

	health.GET("", func(c echo.Context) error {
		log.Printf("Health check RequestID: %s\n", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
