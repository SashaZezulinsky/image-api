package usecase

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"bytes"
	"context"
	"image"
	"net/http"
	"strconv"

	"image-api/internal/domain"
	errs "image-api/pkg/errors"
	"image-api/pkg/utils"
)

type imageUsecase struct {
	imageRepo domain.ImageRepository
}

func NewImageUsecase(iRepo domain.ImageRepository) domain.ImageUseCase {
	return &imageUsecase{
		imageRepo: iRepo,
	}
}

func (uc *imageUsecase) Get(ctx context.Context, id string) (domain.Image, error) {
	imageFile, err := uc.imageRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return imageFile, nil
}

func (uc *imageUsecase) Add(ctx context.Context, i domain.Image) (id string, err error) {
	fileName, metadata, err := prepareImage(i)
	if err != nil {
		return "", err
	}

	id, err = uc.imageRepo.Add(ctx, "", fileName, i, metadata)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (uc *imageUsecase) Update(ctx context.Context, id string, i domain.Image) error {
	err := uc.imageRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	fileName, metadata, err := prepareImage(i)
	if err != nil {
		return err
	}

	_, err = uc.imageRepo.Add(ctx, id, fileName, i, metadata)
	if err != nil {
		return err
	}
	return nil
}

func (uc *imageUsecase) ListAllMetadata(ctx context.Context) ([]map[string]interface{}, error) {
	return uc.imageRepo.ListAllMetadata(ctx)
}

func (uc *imageUsecase) GetMetadata(ctx context.Context, id string) (map[string]interface{}, error) {
	return uc.imageRepo.GetMetadata(ctx, id)
}

func prepareImage(i []byte) (fileName string, metadata map[string]interface{}, err error) {
	fileName, err = utils.GenerateRandomHex(24)
	if err != nil {
		return "", nil, err
	}

	im, _, err := image.DecodeConfig(bytes.NewReader(i))
	if err != nil {
		if err == image.ErrFormat {
			return "", nil, errs.ErrFormat
		}
		return "", nil, err
	}
	contentType := http.DetectContentType(i)

	metadata = map[string]interface{}{
		"image_type": contentType,
		"height":     strconv.Itoa(im.Height),
		"width":      strconv.Itoa(im.Width),
	}
	return fileName, metadata, nil
}
