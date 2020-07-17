//Only point in the application where we will interact with the database
package users

import (
	"database/sql"
	"fmt"
	"strconv"

	database "github.com/goswamipiyush/bookstore_users-api/datasources/mysql/users_db"
	utils "github.com/goswamipiyush/bookstore_users-api/utils/datetime"
	"github.com/goswamipiyush/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?,?,?,?);"
	queryDeleteUser = "DELETE FROM users WHERE id =?;"
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
	return nil
}

func (user *User) Get(id int64) (*User, *errors.RestErr) {

	queryGetUser := "SELECT id, first_name, last_name, email, date_created FROM users where id = ?"
	stringId := strconv.FormatInt(id, 10)
	rows, err := database.SqlDB.Query(queryGetUser, stringId)
	if err != nil {
		return nil, errors.NewInternalServerError("Could not fetch the record")
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
		if err != nil {
			return nil, errors.NewInternalServerError("Could not fetch the record(s)")
		}
	}
	//At this point, if no rows were returned, just return an error
	if user.Id == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with id %d does not exist", id))
	}
	return user, nil
}

func (user *User) Delete(id int64) (*sql.Result, *errors.RestErr) {
	deleteStmt, err := database.SqlDB.Prepare(queryDeleteUser)
	if err != nil {
		return nil, errors.NewInternalServerError("Could not prepare the query for deletion")
	}
	defer deleteStmt.Close()

	deleteResult, err := deleteStmt.Exec(id)
	if err != nil {
		return nil, errors.NewInternalServerError("Could not delete the record")
	}
	return &deleteResult, nil
}
