package upload

import (
	"context"
	"errors"

	"github.com/RicardoSandoval11/MediaPipeline/goservice/entities"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type UploadFileRequest struct {
	UserId    uuid.UUID `json:"userId"`
	Base64Img string    `json:"base64Img"`
}

type UploadFileResponse struct {
	Success      bool   `json:"success"`
	LimitReached bool   `json:"limitReached"`
	PublicFileId string `json:"publicFileId"`
}

func MakeUploadFileEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UploadFileRequest)
		uploadResult, err := svc.UploadFile(ctx, req.UserId, req.Base64Img)
		if err != nil {

			var limitReachedErr *LimitReachedError

			if errors.As(err, &limitReachedErr) {
				return UploadFileResponse{
					Success:      true,
					LimitReached: true,
				}, nil
			}

			return UploadFileResponse{
				Success: false,
			}, err
		}

		return UploadFileResponse{
			Success:      true,
			LimitReached: false,
			PublicFileId: uploadResult.PublicId,
		}, nil
	}
}

type GetFileRequest struct {
	PublicId uuid.UUID `json:"publicId"`
}

type GetFileResponse struct {
	Found bool              `json:"found"`
	File  *entities.AppFile `json:"file"`
}

func MakeGetFileEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFileRequest)
		result, err := svc.GetFile(ctx, req.PublicId)

		if err != nil {
			return GetFileResponse{
				Found: false,
			}, nil
		}

		return GetFileResponse{
			Found: true,
			File:  result,
		}, nil
	}
}
