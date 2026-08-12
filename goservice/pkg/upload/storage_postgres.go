package upload

import (
	"context"
	"errors"

	"github.com/RicardoSandoval11/MediaPipeline/goservice/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type postgresRepository struct {
	db *gorm.DB
}

func NewPostgresqlRepository(db *gorm.DB) Repository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) GetFileByPublicId(ctx context.Context, publicId uuid.UUID) (entities.AppFile, error) {
	var result entities.AppFile

	err := r.db.WithContext(ctx).Where("public_id = ?", publicId).First(&result).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.AppFile{}, errors.New("file not found")
		}

		return entities.AppFile{}, err
	}

	return result, nil
}

func (r *postgresRepository) SaveFile(ctx context.Context, base64Img string) (string, error) {
	newFile := entities.AppFile{
		OriginalImg: base64Img,
		Status:      entities.NewFile,
		PublicId:    uuid.New(),
	}

	result := r.db.WithContext(ctx).Create(&newFile)

	if result.Error != nil {
		return "", result.Error
	}

	return newFile.PublicId.String(), nil
}
