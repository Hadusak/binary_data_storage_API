package models

import "time"

type GetDataResponse struct {
	Key string `json:"key"`
	Data DataResponse `json:"data"`
}

type SaveDataResponse struct {
	Key string `json:"key"`
	Ok bool `json:"ok"`
}

type SaveDataRequest struct {
	Key string `json:"key"`
	Data []byte `json:"data"`
	ValidTo int64 `json:"validTo"`
}

type DataResponse struct {
	Value string `json:"value"`
	Timestamp time.Time `json:"validTo"`
}