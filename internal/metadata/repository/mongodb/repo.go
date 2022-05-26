package mongodb

import (
	"context"
	errs "errors"
	"github.com/image-api/internal/domain"
	"github.com/image-api/pkg/errors"
	"github.com/image-api/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//todo use *mongo.Collection instead of collection and database
type mongoDBRepo struct {
	client     *mongo.Client
	database   string
	collection string
	logger     logger.Logger
}

func NewMongoDBMetadataRepository(client *mongo.Client, database, collection string, logger logger.Logger) domain.MetadataRepository {
	return &mongoDBRepo{
		client:     client,
		database:   database,
		collection: collection,
		logger:     logger,
	}
}

func (m *mongoDBRepo) GetAll(ctx context.Context) ([]*domain.Metadata, error) {
	coll := m.client.Database(m.database).Collection(m.collection)

	var metadataList []*domain.Metadata
	metadataCursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = metadataCursor.All(ctx, &metadataList); err != nil {
		return nil, err
	}
	return metadataList, nil
}

func (m *mongoDBRepo) Get(ctx context.Context, id string) (*domain.Metadata, error) {
	coll := m.client.Database(m.database).Collection(m.collection)

	var result *domain.Metadata
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func (m *mongoDBRepo) Delete(ctx context.Context, id string) error {
	coll := m.client.Database(m.database).Collection(m.collection)

	res, err := coll.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return errors.ErrNotFound
	}
	return nil
}

func (m *mongoDBRepo) Add(ctx context.Context, metadata *domain.Metadata) error {
	coll := m.client.Database(m.database).Collection(m.collection)

	res, err := coll.InsertOne(ctx, metadata)
	if err != nil {
		return err
	}
	if res.InsertedID == 0 {
		return errs.New("image was not inserted due to unknown error")
	}
	return nil
}
