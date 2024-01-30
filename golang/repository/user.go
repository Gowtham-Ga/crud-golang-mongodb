package repository

import (
	"context"
	"curdapi/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	MongoCollecion *mongo.Collection
}

func (r *UserRepo) InsertUser(usr *model.User) (interface{}, error) {
	result, err := r.MongoCollecion.InsertOne(context.Background(), usr)

	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *UserRepo) FindUserById(usrID string) (*model.User, error) {
	var usr model.User

	err := r.MongoCollecion.FindOne(context.Background(), bson.D{{Key: "user_id", Value: usrID}}).Decode(&usr)

	if err != nil {
		return nil, err
	}

	return &usr, nil
}

func (r *UserRepo) FindAllUser() ([]model.User, error) {
	results, err := r.MongoCollecion.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var usrs []model.User
	err = results.All(context.Background(), &usrs)
	if err != nil {
		return nil, fmt.Errorf("results decode error %s", err.Error())
	}
	return usrs, nil
}

func (r *UserRepo) UpdateUserID(usrID string, updateUsr *model.User) (int64, error) {
	result, err := r.MongoCollecion.UpdateOne(context.Background(),
		bson.D{{Key: "user_id", Value: usrID}},
		bson.D{{Key: "$set", Value: updateUsr}})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *UserRepo) DeleteUesrID(usrID string) (int64, error) {
	result, err := r.MongoCollecion.DeleteOne(context.Background(), bson.D{{Key: "user_id", Value: usrID}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *UserRepo) DeleteAllUser() (int64, error) {
	result, err := r.MongoCollecion.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
