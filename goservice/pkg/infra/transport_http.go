package infra

import (
	"context"
	"encoding/json"
	"net/http"
)

func DecodeLivenessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeLivenessResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, ok := response.(LivenessResponse); !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	return json.NewEncoder(w).Encode(response)
}

func DecodeReadinessRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}

func EncodeReadinessResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, ok := response.(ReadinessResponse); !ok {
		w.WriteHeader(http.StatusBadRequest)
	}
	return json.NewEncoder(w).Encode(response)
}
