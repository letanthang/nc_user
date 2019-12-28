package db

import (
	"context"
	"fmt"
	"time"

	"github.com/letanthang/nc_user/model"
	"github.com/letanthang/nc_user/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func AddUser(user *model.User) (interface{}, error) {
	user.Password = utils.MD5(user.Password)
	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	res, err := collection.InsertOne(ctx, user)
	return res, err
}

func UpdateUser(req *model.UserUpdateReq) (interface{}, error) {
	user, _ := FindUserByID(req.ID)

	if req.FirstName != "" {
		user.FirstName = req.FirstName
	}
	if req.LastName != "" {
		user.LastName = req.LastName
	}
	if req.Password != "" {
		user.Password = utils.MD5(req.Password)
	}

	fmt.Printf("user :%+v", user)

	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
