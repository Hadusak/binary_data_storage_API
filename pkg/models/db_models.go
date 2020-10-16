package models

import "github.com/jinzhu/gorm"

type DBKey struct {
	gorm.Model
	Key string
	Data DBData
}

type DBData struct {
	gorm.Model
	Data []byte
	ValidTo int64
	Md5sum [16]byte
}