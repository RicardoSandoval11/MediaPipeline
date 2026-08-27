package upload

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/google/uuid"
)

func DecodeUploadFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	authHeader := r.Header.Get("Authorization")

	if len(authHeader) == 0 || !strings.HasPrefix(authHeader, "Bearer ") {
		return nil, errors.New("auth header is required")
	}

	userId := r.FormValue("userId")
	file, header, err := r.FormFile("file")

	if err != nil {
		return nil, errors.New("Failed parsing file")
	}

	defer file.Close()

	fileBytes, err := io.ReadAll(file)

	if err != nil {
		return nil, errors.New("Error reading file")
	}

	base64String := base64.StdEncoding.EncodeToString(fileBytes)

	mimeType := header.Header.Get("Content-Type")

	if mimeType == "" {
		mimeType = http.DetectContentType(fileBytes)
	}

	parsedUId, err := uuid.Parse(userId)

	if err != nil {
		return nil, err
	}

	return UploadFileRequest{
		Base64Img: base64String,
		UserId:    parsedUId,
	}, nil
}

func EncodeUploadFileResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, ok := response.(UploadFileResponse); !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	return json.NewEncoder(w).Encode(response)
}

func DecodeGetFileRequest(_ context.Context, r *http.Request) (interface{}, error) {
	publicId := r.URL.Query().Get("publicId")

	if publicId == "" {
		return nil, errors.New("publicId is required")
	}

	parsed, err := uuid.Parse(publicId)

	if err != nil {
		return nil, errors.New("invalid publicId")
	}

	return GetFileRequest{
		PublicId: parsed,
	}, nil
}

func EncodeGetFileResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, ok := response.(GetFileResponse); !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	return json.NewEncoder(w).Encode(response)
}
