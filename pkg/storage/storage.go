package storage

import (
	"github.com/Hadusak/binary_data_storage_API/pkg/models"
	"github.com/jinzhu/gorm"
	"time"
)
type StorageImpl struct {
	storageMap map[string]*models.Data
	dbConn *gorm.DB
}

type Storage interface {
	Save(key string, data *models.Data)
	Load(key string) *models.Data
}

func (s *StorageImpl) Save(key string, data *models.Data) {
	s.storageMap[key] = data
}

func (s *StorageImpl) Load(key string) *models.Data {
	return s.storageMap[key]
}

func (s *StorageImpl) DeleteNonValid() {
	for {
		for key, value := range s.storageMap {
			if value.Timestamp.Before(time.Now()) {
				delete(s.storageMap, key)
			}
		}
		time.Sleep(time.Second)
	}
}

func NewStorage(db *gorm.DB) Storage{
	return &StorageImpl{
		storageMap: make(map[string]*models.Data),
		dbConn:     db,
	}
}