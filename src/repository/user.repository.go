package repository

import (
	"context"
	"fmt"

	"github.com/sneicast/golang-crud-users-mongodb/src/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/sneicast/golang-crud-users-mongodb/src/models"
)

var collection = database.GetCollection("users")
var ctx = context.Background()

//FilterUsers obtener todos los usuarios
func FilterUsers(filter interface{}) ([]*models.User, error) {
	var users []*models.User

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return users, err
	}

	for cur.Next(ctx) {
		var u models.User
		err := cur.Decode(&u)
		if err != nil {
			return users, err
		}

		users = append(users, &u)
	}

	if err := cur.Err(); err != nil {
		return users, err
	}

	cur.Close(ctx)

	if len(users) == 0 {
		return users, mongo.ErrNoDocuments
	}

	return users, nil

}

//GetDetailUser detalle usuario
func GetDetailUser(objID primitive.ObjectID) (models.User, error) {
	var result models.User

	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	err := collection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, err
}

//CreateUser insertar registro en mongoDb
func CreateUser(user models.User) error {

	//_, err = collection.InsertOne(ctx, user)
	insertResult, err := collection.InsertOne(ctx, user)

	if err != nil {
		return err
	}
	fmt.Println(insertResult)
	return nil
}

//UpdateUser actualizar usuario
func UpdateUser(objID primitive.ObjectID, userUpdate models.User) error {
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}

	//update := bson.D{primitive.E{Key: "name", Value: userUpdate.Name}, primitive.E{Key: "email", Value: userUpdate.Email}, primitive.E{Key: "updated_at", Value: userUpdate.UpdatedAt}}

	update := bson.M{
		"$set": userUpdate,
	}

	updateResult, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(updateResult)
	return nil
}

//DeleteUser eliminar usuario
func DeleteUser(objID primitive.ObjectID) error {
	filter := bson.D{primitive.E{Key: "_id", Value: objID}}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
