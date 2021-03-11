package services

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sneicast/golang-crud-users-mongodb/src/models"
	"github.com/sneicast/golang-crud-users-mongodb/src/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//CreateService servicio crear usuario
func CreateService(user models.User) error {

	err := repository.CreateUser(user)

	if err != nil {
		return err
	}

	return nil
}

//GetUsers obtener todos los usuarios
func GetUsers(c echo.Context) error {

	filter := bson.D{{}}
	users, err := repository.FilterUsers(filter)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}

//GetDetailUser Obtener detalle usuario
func GetDetailUser(c echo.Context) error {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}

	user, err := repository.GetDetailUser(objID)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

//CreateUser crear usuario
func CreateUser(c echo.Context) error {
	u := new(models.User)

	if err := c.Bind(u); err != nil {
		return err
	}
	u.ID = primitive.NewObjectID()
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	err := repository.CreateUser(*u)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}

//UpdatelUser Actualizar usuario
func UpdatelUser(c echo.Context) error {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}

	//	userUpdate := new(models.User)

	var userData models.User //interface{}
	errBody := json.NewDecoder(c.Request().Body).Decode(&userData)
	if errBody != nil {
		return errBody
	}

	userUpdate, errDetailUser := repository.GetDetailUser(objID)
	if errDetailUser != nil {
		return errDetailUser
	}

	userUpdate.UpdatedAt = time.Now()
	userUpdate.Name = userData.Name
	userUpdate.Email = userData.Email

	errorUpdate := repository.UpdateUser(objID, userUpdate)
	if errorUpdate != nil {
		return errorUpdate
	}

	return c.JSON(http.StatusOK, userUpdate)
}

//DeleteUser Eliminar usuario usuario
func DeleteUser(c echo.Context) error {
	objID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return err
	}
	errDeleteUser := repository.DeleteUser(objID)
	if errDeleteUser != nil {
		return errDeleteUser
	}
	return c.String(http.StatusOK, "Usuario Eliminado")
}
