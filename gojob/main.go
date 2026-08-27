package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"os"
	"strings"

	imgcrop "github.com/btfriar/go_imgcrop"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const RETRIEVED_ITEMS_PER_PAGE = 4

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

func initDb() (*gorm.DB, error) {
	connString := os.Getenv("DATABASE_CONN_STRING")

	if connString == "" {
		return nil, errors.New("No connection string found")
	}

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func base64ToReader(imageBase64 string) io.Reader {
	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageBase64))
}

func imageToBase64(img image.Image) (string, error) {
	var buffer bytes.Buffer

	if err := png.Encode(&buffer, img); err != nil {
		return "", fmt.Errorf("error encoding image: %w", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return base64Str, nil
}

func main() {
	db, err := initDb()

	if err != nil {
		fmt.Printf("Could not initialize db connection: %s\n", err.Error())
		os.Exit(1)
	}

	var pendingFiles []AppFile = []AppFile{}
	totalSuccess := 0
	totalFailed := 0
	totalItemsProcessed := 0

	for {
		result := db.Limit(RETRIEVED_ITEMS_PER_PAGE).Where("status = ?", NewFile).Find(&pendingFiles)

		if result.Error != nil {
			fmt.Printf("Database query error: %v\n", result.Error)
			break
		}

		if len(pendingFiles) == 0 {
			break
		}

		for _, file := range pendingFiles {
			totalItemsProcessed++
			reader := base64ToReader(file.OriginalImg)
			result, err := imgcrop.CropToAspectRatio(reader, 16, 9)

			if err != nil {
				fmt.Printf("failed to crop file %s: %s\n", file.PublicId, err.Error())
				db.Model(&file).Update("status", FailedFile)
				totalFailed++
				continue
			}

			croppedImg, err := imageToBase64(result.Image)

			if err != nil {
				fmt.Printf("failed processing file %s\n", err.Error())
				db.Model(&file).Update("status", FailedFile)
				totalFailed++
				continue
			}

			file.CroppedImg = croppedImg
			file.Status = ProcessedFile

			if err := db.Save(&file).Error; err != nil {
				fmt.Printf("failed updating database for file %s: %s\n", file.PublicId, err.Error())
				totalFailed++
				continue
			}

			totalSuccess++
		}
	}

	fmt.Printf("process completed - items processed: %d - success: %d - failed: %d\n", totalItemsProcessed, totalSuccess, totalFailed)
	os.Exit(0)
}
