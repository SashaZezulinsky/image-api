package usecase

import (
	"context"
	"github.com/image-api/internal/domain"
	"github.com/image-api/pkg/logger"
	"github.com/image-api/pkg/utils"
	"time"
)

type imageUsecase struct {
	imageRepo    domain.ImageRepository
	metadataRepo domain.MetadataRepository
	logger       logger.Logger
}

func NewImageUsecase(iRepo domain.ImageRepository, mRepo domain.MetadataRepository, logger logger.Logger) domain.ImageUseCase {
	return &imageUsecase{
		imageRepo:    iRepo,
		metadataRepo: mRepo,
		logger:       logger,
	}
}

func (uc *imageUsecase) Get(ctx context.Context, id string) (domain.Image, error) {
	image, err := uc.imageRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return image, nil
}

func (uc *imageUsecase) Add(ctx context.Context, i domain.Image) (id string, err error) {
	fileName, err := utils.GenerateRandomHex(24)
	if err != nil {
		return "", err
	}
	id, filesize, err := uc.imageRepo.Add(ctx, fileName, i)
	if err != nil {
		return "", err
	}

	metadata := &domain.Metadata{
		ImageName: fileName,
		ImageID:   id,
		Filesize:  filesize,
		ImageType: "", //TODO add image type and height/width
		Timestamp: time.Now().Unix(),
		ExtraData: "",
		Height:    0,
		Width:     0,
	}

	err = uc.metadataRepo.Add(ctx, metadata)
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

	fileName, err := utils.GenerateRandomHex(24)
	if err != nil {
		return err
	}

	id, filesize, err := uc.imageRepo.Add(ctx, fileName, i)
	if err != nil {
		return err
	}

	metadata, err := uc.metadataRepo.Get(ctx, id)
	if err != nil {
		return err
	}

	err = uc.metadataRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	metadata.ImageID = id
	metadata.Filesize = filesize
	//metadata.Height =
	//metadata.Width =
	metadata.ImageName = fileName
	//metadata.ExtraData =
	//metadata.ImageType =

	return uc.metadataRepo.Add(ctx, metadata)
}
