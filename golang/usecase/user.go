package usecase

import (
	"curdapi/model"
	"curdapi/repository"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	MongoCollecion *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func (svc *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	var usr model.User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		return
	}

	usr.UserID = uuid.NewString()

	repo := repository.UserRepo{MongoCollecion: svc.MongoCollecion}

	insertID, err := repo.InsertUser(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid error", err)
		res.Error = err.Error()
		return
	}

	res.Data = usr.UserID
	w.WriteHeader(http.StatusOK)

	log.Println("user inserted with id", insertID, usr)

}
func (svc *UserService) GetUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	usrID := mux.Vars(r)["id"]
	log.Println("user id", usrID)

	repo := repository.UserRepo{MongoCollecion: svc.MongoCollecion}

	usr, err := repo.FindUserById(usrID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = usr
	w.WriteHeader(http.StatusOK)
	log.Println("user get", usr)
}
func (svc *UserService) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.UserRepo{MongoCollecion: svc.MongoCollecion}

	usr, err := repo.FindAllUser()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = usr
	w.WriteHeader(http.StatusOK)
	log.Println("user get", usr)
}
func (svc *UserService) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	usrID := mux.Vars(r)["id"]
	log.Println("user id", usrID)

	if usrID == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid usr id")
		res.Error = "Invalid UserID"
		return
	}

	var usr model.User

	err := json.NewDecoder(r.Body).Decode(&usr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid body", err)
		res.Error = "Invalid Request Data"
		return
	}

	usr.UserID = usrID

	repo := repository.UserRepo{MongoCollecion: svc.MongoCollecion}

	count, err := repo.UpdateUserID(usrID, &usr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
	log.Println("user updated", count)

}
func (svc *UserService) DeleteUserID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	usrID := mux.Vars(r)["id"]
	log.Println("user id", usrID)

	repo := repository.UserRepo{MongoCollecion: svc.MongoCollecion}
	count, err := repo.DeleteUesrID(usrID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
	log.Println("user deleted", count)
}
func (svc *UserService) DeleteAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	res := &Response{}
	defer json.NewEncoder(w).Encode(res)

	repo := repository.UserRepo{MongoCollecion: svc.MongoCollecion}
	count, err := repo.DeleteAllUser()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("error:", err)
		res.Error = err.Error()
		return
	}

	res.Data = count
	w.WriteHeader(http.StatusOK)
	log.Println("all user deleted", count)
}
