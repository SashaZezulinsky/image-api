package domain

import "context"

//Metadata is representation of image's metadata
type Metadata struct {
	ImageID   string `json:"_id" bson:"_id"`
	ImageName string `json:"image_name" bson:"image_name"`

	Filesize  int    `json:"filesize" bson:"filesize"`
	ImageType string `json:"image_type" bson:"image_type"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
	ExtraData string `json:"extra_data" bson:"extra_data"`

	Height uint64 `json:"height" bson:"height"`
	Width  uint64 `json:"width" bson:"width"`
}

//MetadataUseCase represent the metadata's usecases
type MetadataUseCase interface {
	List(ctx context.Context) ([]*Metadata, error)
	Get(ct context.Context, id string) (*Metadata, error)
}

//MetadataRepository represent the metadata's repository contract
type MetadataRepository interface {
	Add(ctx context.Context, metadata *Metadata) error
	GetAll(ctx context.Context) ([]*Metadata, error)
	Get(ctx context.Context, id string) (*Metadata, error)
	Delete(ctx context.Context, id string) error
}
