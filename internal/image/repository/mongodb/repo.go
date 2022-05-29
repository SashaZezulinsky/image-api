package mongodb

import (
	"bytes"
	"context"
	"errors"

	"github.com/fatih/structs"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"

	"image-api/internal/domain"
	errs "image-api/pkg/errors"
)

type mongoDBRepo struct {
	client *mongo.Client
	bucket *gridfs.Bucket
}

func NewMongoDBImageRepository(client *mongo.Client, database string) (domain.ImageRepository, error) {
	bucket, err := gridfs.NewBucket(client.Database(database))
	if err != nil {
		return nil, err
	}

	return &mongoDBRepo{
		client: client,
		bucket: bucket,
	}, nil
}

func (m *mongoDBRepo) Get(ctx context.Context, id string) (domain.Image, error) {
	var buf bytes.Buffer
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errs.ErrBadID
	}

	_, err = m.bucket.DownloadToStream(ObjID, &buf)
	if err != nil {
		if err == gridfs.ErrFileNotFound {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return buf.Bytes(), err
}

func (m *mongoDBRepo) Add(ctx context.Context, fileId, fileName string, i domain.Image, metadata map[string]interface{}) (string, error) {
	var (
		uploadStream *gridfs.UploadStream
		err          error
	)

	if fileId == "" {
		uploadStream, err = m.bucket.OpenUploadStream(
			fileName,
			&options.UploadOptions{Metadata: metadata},
		)
	} else {
		ObjID, err := primitive.ObjectIDFromHex(fileId)
		if err != nil {
			return "", errs.ErrBadID
		}
		uploadStream, err = m.bucket.OpenUploadStreamWithID(
			ObjID,
			fileName,
			&options.UploadOptions{Metadata: metadata},
		)
	}
	if err != nil {
		return "", err
	}
	defer uploadStream.Close()

	_, err = uploadStream.Write(i)
	if err != nil {
		return "", err
	}

	objID, ok := uploadStream.FileID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("uploaded file id is not primitive.ObjectID")
	}
	return objID.Hex(), nil
}

func (m *mongoDBRepo) Delete(ctx context.Context, id string) error {
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errs.ErrBadID
	}
	err = m.bucket.Delete(ObjID)
	if err != nil {
		if err == gridfs.ErrFileNotFound {
			return errs.ErrNotFound
		}
		return err
	}
	return err
}

func (m *mongoDBRepo) ListAllMetadata(ctx context.Context) ([]map[string]interface{}, error) {
	var results []map[string]interface{}
	cur, err := m.bucket.GetFilesCollection().Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var elem *gridfs.File
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		resMap := structs.Map(elem)
		resMap["Metadata"] = elem.Metadata.String()
		results = append(results, resMap)
	}
	return results, nil
}

func (m *mongoDBRepo) GetMetadata(ctx context.Context, id string) (map[string]interface{}, error) {
	var result map[string]interface{}
	ObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errs.ErrBadID
	}

	err = m.bucket.GetFilesCollection().FindOne(ctx, bson.M{"_id": ObjID}).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errs.ErrNotFound
		}
		return nil, err
	}
	return result, nil

}
