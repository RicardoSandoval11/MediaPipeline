package entities

import "github.com/google/uuid"

const (
	NewFile = iota
	FailedFile
	ProcessedFile
)

type AppFile struct {
	Id          uint64    `gorm:"column:id;primaryKey;autoIncremental"`
	OriginalImg string    `gorm:"column:original_img;not null"`
	CroppedImg  string    `gorm:"column:cropped_img"`
	Status      int       `gorm:"column:status;not null"`
	PublicId    uuid.UUID `gorm:"column:public_id;not null;unique"`
}
