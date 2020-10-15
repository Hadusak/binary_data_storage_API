package models

import "time"

type Data struct {
	Value []byte `json:"value"`
	Timestamp time.Time `json:"validTo"`
	Md5Sum [16]byte `json:"md5sum"`
}



