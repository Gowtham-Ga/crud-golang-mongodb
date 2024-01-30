package main

import (
	"context"
	"curdapi/usecase"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("error on loading .env", err)
	}

	log.Println("env loaded")

	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGO_URI")))

	if err != nil {
		log.Fatal("connection error", err)
	}

	err = mongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

	log.Println("Mongo Connected")
}

func main() {
	defer mongoClient.Disconnect(context.Background())

	coll := mongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	usrService := usecase.UserService{MongoCollecion: coll}

	r := mux.NewRouter()

	r.HandleFunc("/user", usrService.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/user/{id}", usrService.GetUserID).Methods(http.MethodGet)
	r.HandleFunc("/user", usrService.GetAllUser).Methods(http.MethodGet)
	r.HandleFunc("/user/{id}", usrService.UpdateUser).Methods(http.MethodPut)
	r.HandleFunc("/user/{id}", usrService.DeleteUserID).Methods(http.MethodDelete)
	r.HandleFunc("/user", usrService.DeleteAllUser).Methods(http.MethodDelete)

	log.Println("server is running on 4444")
	http.ListenAndServe(":4444", r)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running"))
}
