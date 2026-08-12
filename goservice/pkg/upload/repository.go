package upload

import (
	"context"

	"github.com/RicardoSandoval11/MediaPipeline/goservice/entities"
	"github.com/google/uuid"
)

type Repository interface {
	GetFileByPublicId(ctx context.Context, publicId uuid.UUID) (entities.AppFile, error)
	SaveFile(ctx context.Context, base64Img string) (string, error)
}
