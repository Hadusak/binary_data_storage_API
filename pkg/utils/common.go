package utils

import (
	"crypto/md5"
	"encoding/json"
	"github.com/Hadusak/binary_data_storage_API/pkg/models"
	"github.com/Hadusak/binary_data_storage_API/pkg/proto"
	"net/http"
	"time"
)

func GetEnv(env, fallback string) string {
	curr :=  (env)
	if curr != "" {
		return curr
	}
	return fallback
}

func JSONResponse(w http.ResponseWriter, code int, output interface{}) {
	// Convert our interface to JSON
	response, _ := json.Marshal(output)
	// Set the content type to json for browsers
	w.Header().Set("Content-Type", "application/json")
	// Our response code
	w.WriteHeader(code)

	w.Write(response)
}

func ProtoDataToInternalData(data *proto.Data) (string, *models.Data) {
	return data.Key, &models.Data{
		Value:     data.Data,
		Timestamp: time.Unix(data.Timestamp, 0),
		Md5Sum:    md5.Sum(data.Data),
	}
}