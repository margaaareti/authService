package repository

import "github.com/minio/minio-go/v7"

func NewPictureStorage(endpoint string) *minio.Client {

	//Initialize minio client object
	fileStorage, err := minio.New(endpoint, &minio.Options{
		Secure: false,
	})
	if err != nil {

	}

	return fileStorage
}
