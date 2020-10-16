package storage

import (
	"bytes"
	"github.com/Hadusak/binary_data_storage_API/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type StorageImpl struct {
	dbConn *gorm.DB
}

type Storage interface {
	Save(key string, data *models.Data)
	Load(key string) *models.Data
}

func (s *StorageImpl) Save(key string, data *models.Data) {
	if s.Compare(data.Md5Sum) {
		dbKey := models.DBKey{
			Model: gorm.Model{},
			Key:   key,
			Data:  s.dbConn.Where("md5sum = ?", data.Md5Sum).First(models.DBData{}).Value.(models.DBData),
		}
		s.dbConn.Save(&dbKey)
		return
	}
	dbKey := models.DBKey{
		Model: gorm.Model{},
		Key:   key,
		Data:  models.DBData{
			Model:   gorm.Model{},
			Data:    data.Value,
			ValidTo: data.Timestamp.Unix(),
			Md5sum:  data.Md5Sum,
		},
	}
	s.dbConn.Save(&dbKey)
}

func (s *StorageImpl) Load(key string) *models.Data {
	dbKey := s.dbConn.Where("key=?", key).First(models.DBKey{}).Value.(models.DBKey)
	return &models.Data{
		Value:     dbKey.Data.Data,
		Timestamp: time.Unix(dbKey.Data.ValidTo, 0),
		Md5Sum:    dbKey.Data.Md5sum,
	}
}

func (s *StorageImpl) DeleteNonValid() {
	for {
		for _, value := range s.dbConn.Find(models.DBKey{}).Value.([]models.DBKey) {
			if time.Unix(value.Data.ValidTo,0).Before(time.Now()) {
				s.dbConn.Delete(&value)
				s.dbConn.Delete(&value.Data)
			}
		}
		time.Sleep(time.Second)
	}
}

func (s *StorageImpl) Compare(hash [16]byte) bool {
	var data []models.DBData
	s.dbConn.Find(&data)

	for _, value := range data {
		if bytes.Equal(value.Md5sum[:], hash[:]) {
			return true
		}
	}
	return false
}

func NewStorage(db *gorm.DB) Storage{
	return &StorageImpl{
		dbConn:     db,
	}
}