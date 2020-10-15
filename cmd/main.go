package main

import (
	"crypto/md5"
	"fmt"
	"github.com/Hadusak/binary_data_storage_API/pkg/models"
	"github.com/Hadusak/binary_data_storage_API/pkg/server"
	"github.com/Hadusak/binary_data_storage_API/pkg/storage"
	"github.com/Hadusak/binary_data_storage_API/pkg/utils"
	"github.com/jinzhu/gorm"
	"time"
)

func main() {

	db, err := initDB()
	if err != nil {
		// todo error handling
	}

	storage := storage.NewStorage(db)

	s := "Some words in bytes"

	storage.Save("mara", &models.Data{
		Value:     []byte(s),
		Timestamp: time.Now().Add(time.Minute),
		Md5Sum:    md5.Sum([]byte(s)),
	})

	server.NewRestApi(storage)
	//server.NewGRPCServer(storage)


}

func initDB() (*gorm.DB, error){
	usr := utils.GetEnv("POSTGRES_USER", "admin")
	pass := utils.GetEnv("POSTGRES_PASSWORD", "securepasswordwhichnobodyknows")
	dbName := utils.GetEnv("POSTGRES_DB", "bds_db")
	host := utils.GetEnv("DB_HOST", "0.0.0.0")
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable", host, usr, dbName, pass))
	if err != nil {
		return nil, err
	}
	return db, nil
}


