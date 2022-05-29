package http

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"

	"image-api/internal/domain"
	"image-api/pkg/errors"
)

type imageHandler struct {
	ImageUsecase domain.ImageUseCase
}

func NewImageHandler(iUsecase domain.ImageUseCase) domain.ImageHandlers {
	return &imageHandler{
		ImageUsecase: iUsecase,
	}
}

func (h *imageHandler) ListMetadata() echo.HandlerFunc {
	return func(c echo.Context) error {
		list, err := h.ImageUsecase.ListAllMetadata(context.Background())
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, list)
	}
}

func (h *imageHandler) GetMetadata() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := getIDParam(c)
		if err != nil {
			return err
		}

		metadata, err := h.ImageUsecase.GetMetadata(context.Background(), id)
		switch err {
		case errors.ErrNotFound, errors.ErrBadID:
			return c.JSON(http.StatusOK, map[string]interface{}{"error": err.Error()})
		case nil:
			return c.JSON(http.StatusOK, metadata)
		default:
			return err
		}
	}
}

func (h *imageHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := getIDParam(c)
		if err != nil {
			return err
		}

		imageBytes, err := h.ImageUsecase.Get(context.Background(), id)
		contentType := http.DetectContentType(imageBytes)
		switch err {
		case errors.ErrNotFound, errors.ErrBadID:
			return c.JSON(http.StatusOK, map[string]interface{}{"error": err.Error()})
		case nil:
			return c.Stream(http.StatusOK, contentType, bytes.NewReader(imageBytes))
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
		switch err {
		case errors.ErrFormat:
			return c.JSON(http.StatusOK, map[string]interface{}{"error": errors.ErrFormat.Error()})
		case nil:
			return c.JSON(http.StatusCreated, map[string]interface{}{"success": "true", "id": id})
		default:
			return err
		}
	}
}

func (h *imageHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.Background()
		id, err := getIDParam(c)
		if err != nil {
			return err
		}

		image, err := ioutil.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		err = h.ImageUsecase.Update(ctx, id, image)
		switch err {
		case errors.ErrNotFound, errors.ErrFormat, errors.ErrBadID:
			return c.JSON(http.StatusOK, map[string]interface{}{"error": err.Error()})
		case nil:
			return c.JSON(http.StatusOK, map[string]interface{}{"success": "true"})
		default:
			return err
		}
	}
}

func getIDParam(c echo.Context) (string, error) {
	id := c.Param("id")
	if id == "" {
		params, err := c.FormParams()
		if err != nil {
			return "", err
		}
		id = params.Get("id")
	}
	return id, nil
}
