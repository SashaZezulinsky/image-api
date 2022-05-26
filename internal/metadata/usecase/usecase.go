package metadata

import (
	"context"
	"github.com/image-api/internal/domain"
	"github.com/image-api/pkg/logger"
)

type metadataUsecase struct {
	metadataRepo domain.MetadataRepository
	logger       logger.Logger
}

func NewMetadataUsecase(mRepo domain.MetadataRepository, logger logger.Logger) domain.MetadataUseCase {
	return &metadataUsecase{
		metadataRepo: mRepo,
		logger:       logger,
	}
}

func (uc *metadataUsecase) List(ctx context.Context) ([]*domain.Metadata, error) {
	//TODO implement me
	panic("implement me")
}

func (uc *metadataUsecase) Get(ctx context.Context, id string) (*domain.Metadata, error) {
	//TODO implement me
	panic("implement me")
}
