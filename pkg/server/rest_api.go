package server

import (
	"encoding/json"
	"github.com/Hadusak/binary_data_storage_API/pkg/models"
	"github.com/Hadusak/binary_data_storage_API/pkg/storage"
	"github.com/Hadusak/binary_data_storage_API/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type RestApi interface {
	Run(addr string)
}

type RestApiImpl struct {
	Router *mux.Router
	Storage storage.Storage
}

func (rai *RestApiImpl) Run(addr string) {
	http.ListenAndServe(addr, rai.Router)
}

func (rai *RestApiImpl) GetDataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.FormValue("key")

		response := rai.Storage.Load(key)
		dataJson, err := json.Marshal(models.GetDataResponse{
			Key:  key,
			Data: *response,
		})
		if err != nil {
			//todo some err handling
		}
		utils.JSONResponse(w, 200, dataJson)
	})
}

func (rai *RestApiImpl) SaveDataHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var model models.SaveDataRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&model); err != nil {
			utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"error": "Invalid request"})
			return
		}
		defer r.Body.Close()

		rai.Storage.Save(model.Key, &models.Data{
			model.Data, time.Unix(model.Timestamp,0), nil, //todo md5sum
		})
		dataJson, err := json.Marshal(models.SaveDataResponse{
			Key:  model.Key,
			Ok: true,
		})
		if err != nil {
			utils.JSONResponse(w, 500, nil)
		}
		utils.JSONResponse(w, 200, dataJson)
	})
}

func NewRestApi(storage storage.Storage) RestApi {
	router := mux.NewRouter()
	rApi := &RestApiImpl{
		Router: router,
		Storage: storage,
	}
	router.Handle("/getData/", rApi.GetDataHandler() ).Methods("GET")
	router.Handle("/saveData/", rApi.SaveDataHandler()).Methods("POST")
	return rApi
}