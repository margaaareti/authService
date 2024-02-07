package homerepo

import (
	"Test_derictory/models"
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"strings"
)

type PictureStorage struct {
	client   *minio.Client
	bucket   string
	endpoint string
}

func NewPictureStorage(client *minio.Client, bucket, endpoint string) *PictureStorage {
	return &PictureStorage{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (ps *PictureStorage) SaveImage(ctx context.Context, input models.UploadInput) (string, error) {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	_, err := ps.client.PutObject(ctx, ps.bucket, input.Name, input.File, input.Size, opts)
	if err != nil {
		logrus.Errorf("error occured while uploading file to bucket: %s", err.Error())
		return "", errors.Errorf("error occured while uploading file to bucket: %s", err.Error())
	}

	return ps.generateFileURL(input.Name), nil

}

func (ps *PictureStorage) generateFileURL(filename string) string {
	endpoint := strings.Replace(ps.endpoint, "localstack", "localhost", -1)
	return fmt.Sprintf("http://%s/%s/%s", endpoint, ps.bucket, filename)
}
