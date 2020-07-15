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
	//bytes, err := ioutil.ReadAll(c.Request.Body)
	// if err != nil {
	// 	//TODO - handle the error
	// 	return
	// }
	//err = json.Unmarshal(bytes, &user)

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

	key := c.Params[0].Key
	fmt.Println(key)
	val := c.Params[0].Value
	fmt.Println(val)

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

func SearchUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "Please implement me!")
}
