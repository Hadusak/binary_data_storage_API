package models

import "time"

type Data struct {
	Value []byte
	Timestamp time.Time
	Md5Sum [16]byte
}



