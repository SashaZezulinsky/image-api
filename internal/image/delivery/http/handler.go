package http

import (
	"context"
	"github.com/image-api/internal/domain"
	"github.com/image-api/pkg/errors"
	"github.com/image-api/pkg/logger"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

type imageHandler struct {
	ImageUsecase    domain.ImageUseCase
	MetadataUsecase domain.MetadataUseCase
	logger          logger.Logger
}

func NewImageHandler(iUsecase domain.ImageUseCase, mUsecase domain.MetadataUseCase, logger logger.Logger) domain.ImageHandlers {
	return &imageHandler{
		ImageUsecase:    iUsecase,
		MetadataUsecase: mUsecase,
		logger:          logger,
	}
}

func (h *imageHandler) ListMetadata() echo.HandlerFunc {
	return func(c echo.Context) error {
		list, err := h.MetadataUsecase.List(context.Background())
		if err != nil {
			//todo test err handler
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, list)
	}
}

func (h *imageHandler) GetMetadata() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		metadata, err := h.MetadataUsecase.Get(context.Background(), id)
		switch err {
		case errors.ErrNotFound:
			return c.JSON(http.StatusOK, errors.ErrNotFound)
		case nil:
			return c.JSON(http.StatusOK, metadata)
		default:
			return err
		}
	}
}

func (h *imageHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		metadata, err := h.ImageUsecase.Get(context.Background(), id)
		switch err {
		case errors.ErrNotFound:
			return c.JSON(http.StatusOK, errors.ErrNotFound)
		case nil:
			return c.JSON(http.StatusOK, metadata)
		default:
			return err
		}
	}
}

func (h *imageHandler) Add() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		image, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		id, err := h.ImageUsecase.Add(ctx, image)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"id": id})

	}
}

func (h *imageHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		id := c.Param("id")

		image, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		err = h.ImageUsecase.Update(ctx, id, image)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{"success": "true"})

	}
}
