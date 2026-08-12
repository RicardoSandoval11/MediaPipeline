package upload

import (
	"context"
	"fmt"

	"github.com/RicardoSandoval11/MediaPipeline/goservice/entities"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type Service interface {
	GetFile(ctx context.Context, publicId uuid.UUID) (*entities.AppFile, error)
	UploadFile(ctx context.Context, userId uuid.UUID, base64Img string) (UploadFileResult, error)
}

type uploadService struct {
	repository    Repository
	fileCounterEp endpoint.Endpoint
}

func NewUploadService(fileCounterEp endpoint.Endpoint, repository Repository) Service {
	return &uploadService{
		repository:    repository,
		fileCounterEp: fileCounterEp,
	}
}

func (s *uploadService) GetFile(ctx context.Context, publicId uuid.UUID) (*entities.AppFile, error) {
	result, err := s.repository.GetFileByPublicId(ctx, publicId)

	if err != nil {
		return nil, err
	}

	return &result, err
}

type UploadFileResult struct {
	Success  bool
	PublicId string
}

type LimitReachedError struct {
	UserId string
}

func (e *LimitReachedError) Error() string {
	return fmt.Sprintf("User %s has reached the maximum uploaded files per day", e.UserId)
}

func (s *uploadService) UploadFile(ctx context.Context, userId uuid.UUID, base64Img string) (UploadFileResult, error) {
	resp, err := s.fileCounterEp(ctx, UpsertFileCounterRequest{
		UserId: userId,
	})

	if err != nil {
		return UploadFileResult{}, err
	}

	response := resp.(UpsertFileCounterResponse)

	if response.LimitReached {
		return UploadFileResult{}, &LimitReachedError{UserId: userId.String()}
	}

	publicId, err := s.repository.SaveFile(ctx, base64Img)

	if err != nil {
		return UploadFileResult{}, err
	}

	return UploadFileResult{
		Success:  true,
		PublicId: publicId,
	}, nil
}
