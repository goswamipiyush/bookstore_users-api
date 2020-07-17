//Users controller to take care of users endpoint requests

package users

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goswamipiyush/bookstore_users-api/domain/users"
	"github.com/goswamipiyush/bookstore_users-api/services"
	"github.com/goswamipiyush/bookstore_users-api/utils/errors"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)

	//One way of populating JSON structure from reqeust body
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO - handle the error
	// 	return
	// }
	//err = json.Unmarshal(bytes, &user)

	//Second way of populating JSON structure from reauest body
	err := c.ShouldBindJSON(&user)

	if err != nil {
		restErr := errors.NewBadRequestError("Invalid JSON body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	var user users.User
	val := c.Params[0].Value

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		//TODO - handle get error
		return
	}
	result, getErr := services.GetUser(i, user)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

func DeleteUser(c *gin.Context) {
	var user users.User
	val := c.Params[0].Value

	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		//TODO - handle get error
		return
	}
	result, deleteErr := services.DeleteUser(i, user)
	if deleteErr != nil {
		c.JSON(deleteErr.Status, deleteErr)
		return
	}
	//str := `{"message:resource deleted successfully"}`
	c.JSON(http.StatusOK, result)
}

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Please implement me!")
}
