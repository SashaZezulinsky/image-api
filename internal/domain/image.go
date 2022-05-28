//go:generate mockgen -source image.go -destination ../image/mock/mock.go -package mock

package domain

import (
	"context"

	"github.com/labstack/echo/v4"
)

//Image is representation of image in slice of bytes
type Image []byte

//ImageUseCase represent the image's usecases
type ImageUseCase interface {
	ListAllMetadata(ctx context.Context) ([]map[string]interface{}, error)
	GetMetadata(ctx context.Context, id string) (map[string]interface{}, error)
	Get(ctx context.Context, id string) (Image, error)
	Add(ctx context.Context, i Image) (id string, err error)
	Update(ctx context.Context, id string, i Image) error
}

//ImageRepository represent the image's repository contract
type ImageRepository interface {
	ListAllMetadata(ctx context.Context) ([]map[string]interface{}, error)
	GetMetadata(ctx context.Context, id string) (map[string]interface{}, error)
	Get(ctx context.Context, id string) (Image, error)
	Add(ctx context.Context, fileId, fileName string, i Image, metadata map[string]interface{}) (id string, err error)
	Delete(ctx context.Context, id string) error
}

//ImageHandlers represent the image's http delivery interface
type ImageHandlers interface {
	ListMetadata() echo.HandlerFunc
	GetMetadata() echo.HandlerFunc
	Get() echo.HandlerFunc
	Add() echo.HandlerFunc
	Update() echo.HandlerFunc
}
