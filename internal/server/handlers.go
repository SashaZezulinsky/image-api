package server

import (
	"github.com/AleksK1NG/api-mc/pkg/utils"
	imageHttp "github.com/image-api/internal/image/delivery/http"
	imageRepo "github.com/image-api/internal/image/repository/mongodb"
	imageUsecase "github.com/image-api/internal/image/usecase"
	metadataRepo "github.com/image-api/internal/metadata/repository/mongodb"
	metadataUsecase "github.com/image-api/internal/metadata/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Map Server Handlers
func (s *Server) MapHandlers(e *echo.Echo) error {
	// Init repositories
	iRepo := imageRepo.NewMongoDBImageRepository(s.mongoDB, s.logger)
	mRepo := metadataRepo.NewMongoDBMetadataRepository(s.mongoDB, s.logger)

	// Init useCases
	imageUC := imageUsecase.NewImageUsecase(iRepo, mRepo, s.logger)
	metadataUC := metadataUsecase.NewMetadataUsecase(mRepo, s.logger)

	// Init handlers
	imagesHandlers := imageHttp.NewImageHandler(imageUC, metadataUC, s.logger)

	v1 := e.Group("/v1")

	health := v1.Group("/health")
	imageGroup := v1.Group("/images")

	imageHttp.MapImageRoutes(imageGroup, imagesHandlers)

	health.GET("", func(c echo.Context) error {
		s.logger.Infof("Health check RequestID: %s", utils.GetRequestID(c))
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	return nil
}
