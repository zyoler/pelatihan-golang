package userController

import (
	"encoding/json"
	"log"
	"net/http"
	"randi_firmansyah/helper/helper"
	"randi_firmansyah/helper/helperRedis"
	res "randi_firmansyah/helper/response"
	"randi_firmansyah/models/userModel"
	"randi_firmansyah/repository/userRepository"
	"strconv"

	"github.com/go-chi/chi"
)

var (
	key_redis       = "list_user_randi"
	redisGetSuccess = "Berhasil get user data from redis"
	redisSetFailed  = "Gagal set user data ke redis"
	redisSetSuccess = "Berhasil set user data ke redis"
	dbGettingData   = "Sedang mengambil user data ke database"
	controllerName  = "User"
)

func GetSemuaUser(w http.ResponseWriter, r *http.Request) {
	// check redis
	if redisData, err := helperRedis.GetRedis(key_redis); err == nil {
		log.Println(redisGetSuccess)

		// unmarshall from redis
		var data []userModel.User
		if err := helper.UnMarshall(redisData, &data); err != nil {
			log.Println(err.Error())
			return
		}

		res.Response(w, http.StatusOK, res.MsgGetAll(controllerName), data)
		return
	}

	// select ke db
	log.Println(dbGettingData)
	listUser, err := userRepository.FindAll()
	if err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	// nge jadiin json
	result, err := json.Marshal(listUser)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// ngeset ke redis dan ngecek nya
	if err := helperRedis.SetRedis(key_redis, result); err != nil {
		log.Println(redisSetFailed, err)
	}

	log.Println(redisSetSuccess)
	res.Response(w, http.StatusOK, res.MsgGetAll(controllerName), listUser)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
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
		var data []userModel.User
		if err := helper.UnMarshall(redisData, &data); err != nil {
			log.Println(err.Error())
			return
		}

		var oneData userModel.User
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
	cari, err := userRepository.FindByID(newId)
	if err != nil {
		log.Println(err)
		res.Response(w, http.StatusBadRequest, res.MsgNotFound(controllerName), nil)
		return
	}

	res.Response(w, http.StatusOK, res.MsgGetDetail(controllerName), cari)
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	// decode from json
	decoder := json.NewDecoder(r.Body)

	var datarequest userModel.User

	// cek
	if err := decoder.Decode(&datarequest); err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// hash password
	newPassword, err := helper.HashPassword(datarequest.Password)
	if err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	datarequest.Password = newPassword
	create, err := userRepository.Create(datarequest)
	if err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	helperRedis.ClearRedis(key_redis)

	res.Response(w, http.StatusOK, res.MsgTambah(controllerName), create)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
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
	search, err := userRepository.FindByID(newId)
	if err != nil {
		log.Println(err)
		res.Response(w, http.StatusBadRequest, res.MsgNotFound(controllerName), nil)
		return
	}

	// set id
	var datarequest userModel.User
	datarequest.Id = newId

	// delete
	if _, err := userRepository.Delete(datarequest); err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	helperRedis.ClearRedis(key_redis)

	res.Response(w, http.StatusOK, res.MsgHapus(controllerName), search)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	newId, errInt := strconv.Atoi(id)
	if errInt != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// cari
	if _, err := userRepository.FindByID(newId); err != nil {
		log.Println(err)
		res.Response(w, http.StatusBadRequest, res.MsgNotFound(controllerName), nil)
		return
	}

	// decode
	decoder := json.NewDecoder(r.Body)
	var datarequest userModel.User
	if err := decoder.Decode(&datarequest); err != nil {
		res.Response(w, http.StatusBadRequest, res.MsgInvalidReq(), nil)
		return
	}

	// update
	datarequest.Id = newId
	updated, err := userRepository.Update(newId, datarequest)
	if err != nil {
		res.Response(w, http.StatusInternalServerError, res.MsgServiceErr(), nil)
		return
	}

	helperRedis.ClearRedis(key_redis)

	res.Response(w, http.StatusOK, res.MsgUpdate(controllerName), updated)
}
