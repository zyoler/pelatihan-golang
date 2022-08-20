package productController

import (
	"encoding/json"
	"log"
	"net/http"
	"randi_firmansyah/helper/helper"
	"randi_firmansyah/helper/helperRedis"
	res "randi_firmansyah/helper/response"
	"randi_firmansyah/models/productModel"
	"randi_firmansyah/repository/productRepository"
	"strconv"

	"github.com/go-chi/chi"
)

var (
	key_redis       = "list_product_randi"
	redisGetSuccess = "Berhasil get product data from redis"
	redisSetFailed  = "Gagal set product data ke redis"
	redisSetSuccess = "Berhasil set product data ke redis"
	dbGettingData   = "Sedang mengambil product data ke database"
	controllerName  = "Product"
)

func GetSemuaProduct(w http.ResponseWriter, r *http.Request) {
	// check redis
	if redisData, err := helperRedis.GetRedis(key_redis); err == nil {
		log.Println(redisGetSuccess)

		// unmarshall from redis
		var data []productModel.Product
		if err := helper.UnMarshall(redisData, &data); err != nil {
			log.Println(err.Error())
			return
		}

		res.Response(w, http.StatusOK, res.MsgGetAll(controllerName), data)
		return
	}

	// select ke db
	log.Println(dbGettingData)
	listProduct, err := productRepository.FindAll()
	if err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	// nge jadiin json
	result, err := json.Marshal(listProduct)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// ngeset ke redis dan ngecek nya
	if err := helperRedis.SetRedis(key_redis, result); err != nil {
		log.Println(redisSetFailed, err)
	}

	log.Println(redisSetSuccess)
	res.Response(w, http.StatusOK, res.MsgGetAll(controllerName), listProduct)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, "id")

	if id == "" {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// check redis
	if redisData, err := helperRedis.GetRedis(key_redis); err == nil {
		log.Println(redisGetSuccess)

		// unmarshall from redis
		var data []productModel.Product
		if err := helper.UnMarshall(redisData, &data); err != nil {
			log.Println(err.Error())
			return
		}

		var oneData productModel.Product
		for i := 0; i < len(data); i++ {
			if oneData.Id == newId {
				oneData = data[i]
				res.Response(w, http.StatusOK, res.MsgGetAll(controllerName), data)
				return
			}
		}
	}

	// select ke db
	log.Println(dbGettingData)
	cari, err := productRepository.FindByID(newId)
	if err != nil {
		log.Println(err)
		res.Response(w, http.StatusBadRequest, res.MsgNotFound(controllerName), nil)
		return
	}

	res.Response(w, http.StatusOK, res.MsgGetDetail(controllerName), cari)
}

func PostProduct(w http.ResponseWriter, r *http.Request) {
	// decode from json
	decoder := json.NewDecoder(r.Body)

	var datarequest productModel.Product

	// cek
	if err := decoder.Decode(&datarequest); err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// insert
	create, err := productRepository.Create(datarequest)
	if err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	helperRedis.ClearRedis(key_redis)

	res.Response(w, http.StatusOK, res.MsgTambah(controllerName), create)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	// ambil parameter
	id := chi.URLParam(r, "id")

	newId, err := strconv.Atoi(id)
	if err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	if id == "" {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// cari
	search, err := productRepository.FindByID(newId)
	if err != nil {
		log.Println(err)
		res.Response(w, http.StatusBadRequest, res.MsgNotFound(controllerName), nil)
		return
	}

	// set id
	var datarequest productModel.Product
	datarequest.Id = newId

	// delete
	if _, err := productRepository.Delete(datarequest); err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	helperRedis.ClearRedis(key_redis)

	res.Response(w, http.StatusOK, res.MsgHapus(controllerName), search)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	newId, errInt := strconv.Atoi(id)
	if errInt != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// cari
	if _, err := productRepository.FindByID(newId); err != nil {
		log.Println(err)
		res.Response(w, http.StatusBadRequest, res.MsgNotFound(controllerName), nil)
		return
	}

	// decode
	decoder := json.NewDecoder(r.Body)
	var datarequest productModel.Product
	if err := decoder.Decode(&datarequest); err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// update
	datarequest.Id = newId
	updated, err := productRepository.Update(newId, datarequest)
	if err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	helperRedis.ClearRedis(key_redis)

	res.Response(w, http.StatusOK, res.MsgUpdate(controllerName), updated)
}
