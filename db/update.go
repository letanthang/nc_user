package db

import (
	"context"
	"encoding/json"
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

	// user, _ := FindUserByID(req.ID)

	// if req.FirstName != "" {
	// 	user.FirstName = req.FirstName
	// }
	// if req.LastName != "" {
	// 	user.LastName = req.LastName
	// }
	// if req.Password != "" {
	// 	user.Password = utils.MD5(req.Password)
	// }
	var user map[string]interface{}

	bs, _ := json.Marshal(req)
	json.Unmarshal(bs, &user)
	if _, ok := user["password"]; ok {
		user["password"] = utils.MD5(user["password"].(string))
	}

	fmt.Printf("user :%+v", user)

	collection := Client.Database(DbName).Collection(ColName)
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	filter := bson.M{"id": req.ID}
	update := bson.M{"$set": user}
	res, err := collection.UpdateOne(ctx, filter, update)
	return res, err
}
