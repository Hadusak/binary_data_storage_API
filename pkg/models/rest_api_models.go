package models

type GetDataResponse struct {
	Key string `json:"key"`
	Data Data `json:"data"`
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