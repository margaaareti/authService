package mainservice

import (
	"Test_derictory/mainpage"
	"Test_derictory/models"
	"context"
	"io"
	"math/rand"
	"time"
)

const (
	letterBytes    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
	fileNameLength = 16
)

type HomeUseCase struct {
	CrudPage    mainpage.HomeRepo
	fileStorage mainpage.FilesStorage
}

func NewHomeUseCase(crudpage mainpage.HomeRepo, filestorage mainpage.FilesStorage) *HomeUseCase {
	return &HomeUseCase{
		CrudPage:    crudpage,
		fileStorage: filestorage}
}

func (h *HomeUseCase) AddStudent(ctx context.Context, userId uint64, student models.Student) (uint64, error) {
	return h.CrudPage.CreateStudent(ctx, userId, student)
}

func (h *HomeUseCase) GetAllNotice(ctx context.Context) ([]models.Student, error) {
	return h.CrudPage.PullAllNotice(ctx)
}

func (h *HomeUseCase) GetById(ctx context.Context, Id uint64) (models.Student, error) {
	return h.CrudPage.PullById(ctx, Id)
}

func (h *HomeUseCase) DeleteNoticeByID(ctx context.Context, Id int) error {
	return h.CrudPage.DeleteNotice(ctx, Id)
}

func (h *HomeUseCase) UpdateEntryUseCase(ctx context.Context, userId, studentId uint64, input models.UpdateStudentInput) error {
	if err := input.Validate(); err != nil {
		return err
	}
	return h.CrudPage.UpdateEntry(ctx, userId, studentId, input)
}

func (h *HomeUseCase) UploadImage(ctx context.Context, file io.Reader, size int64, contentType string) (string, error) {
	filename, err := h.fileStorage.SaveImage(ctx, models.UploadInput{
		File:        file,
		Name:        generateFileName(),
		Size:        size,
		ContentType: contentType,
	})
	if err != nil {
		return "", err
	}

	return filename, nil
}

func generateFileName() string {
	rand.Seed(time.Now().Unix())
	b := make([]byte, fileNameLength)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)

}
