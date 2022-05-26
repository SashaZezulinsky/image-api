package mongodb

import (
	"bytes"
	"context"
	"errors"
	"github.com/image-api/internal/domain"
	"github.com/image-api/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
)

type mongoDBRepo struct {
	client   *mongo.Client
	database string
	logger   logger.Logger
}

func NewMongoDBImageRepository(client *mongo.Client, database string, logger logger.Logger) domain.ImageRepository {
	return &mongoDBRepo{
		client:   client,
		database: database,
		logger:   logger,
	}
}

func (m *mongoDBRepo) Get(ctx context.Context, id string) (domain.Image, error) {
	bucket, _ := gridfs.NewBucket(
		m.client.Database(m.database),
	)
	var buf bytes.Buffer
	_, err := bucket.DownloadToStream(id, &buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}

func (m *mongoDBRepo) Add(ctx context.Context, fileName string, i domain.Image) (string, int, error) {
	//todo get image file name of generate id
	bucket, err := gridfs.NewBucket(
		m.client.Database(m.database),
	)
	if err != nil {
		return "", 0, err
	}
	uploadStream, err := bucket.OpenUploadStream(
		fileName,
	)
	if err != nil {
		return "", 0, err
	}
	defer uploadStream.Close()

	fileSize, err := uploadStream.Write(i)
	if err != nil {
		return "", 0, err
	}
	id, ok := uploadStream.FileID.(string)
	if !ok {
		return "", 0, errors.New("uploaded file id is not string")
	}
	return id, fileSize, nil
}

func (m *mongoDBRepo) Delete(ctx context.Context, id string) error {
	bucket, _ := gridfs.NewBucket(
		m.client.Database(m.database),
	)
	err := bucket.Delete(id)
	if err != nil {
		return err
	}
	return err
}
