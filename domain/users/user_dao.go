//Only point in the application where we will interact with the database
package users

import (
	"fmt"

	database "github.com/goswamipiyush/bookstore_users-api/datasources/mysql/users_db"
	utils "github.com/goswamipiyush/bookstore_users-api/utils/datetime"
	"github.com/goswamipiyush/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
)

var (
	usersDB = make(map[int64]*User) //temporary database till we actually have a DB to persist data
)

func (user *User) Save() *errors.RestErr {
	insertStmt, err := database.SqlDB.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError("Could not prepare the query for insertion")
	}
	defer insertStmt.Close()
	//Format the date before saving to the database
	timeNow := utils.FormatDate()
	user.DateCreated = timeNow

	insertResult, err := insertStmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return errors.NewInternalServerError("Could not insert the record")
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError("Error while trying to insert id")
	}
	user.Id = userId

	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError("User already exists")
	}

	usersDB[user.Id] = user
	return nil
}

func (user *User) Get(id int64) (*User, *errors.RestErr) {
	//Check for DB heartbeat
	// err := users_db.sqldb.Ping()
	// if err != nil {
	// 	panic(err)
	// }

	result := usersDB[id]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User id %d not found", id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return user, nil

}
