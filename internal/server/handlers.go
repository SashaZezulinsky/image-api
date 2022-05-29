package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	imageHttp "image-api/internal/image/delivery/http"
	imageRepo "image-api/internal/image/repository/mongodb"
	imageUsecase "image-api/internal/image/usecase"
	"image-api/pkg/utils"
)

func (s *Server) MapHandlers(e *echo.Echo) error {
	iRepo, err := imageRepo.NewMongoDBImageRepository(s.mongoDB, s.cfg.MongoDB.Database)
	if err != nil {
		return err
	}
	imageUC := imageUsecase.NewImageUsecase(iRepo)

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
