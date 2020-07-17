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
	queryInsertUser       = "INSERT INTO users(first_name, last_name, email, date_created, status) VALUES(?,?,?,?,?);"
	queryDeleteUser       = "DELETE FROM users WHERE id =?;"
	queryFindUserByStatus = "SELECT id, first_name, last_name, email, date_created, status FROM USERS where status = ?;"
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

	insertResult, err := insertStmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status)
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

	queryGetUser := "SELECT id, first_name, last_name, email, date_created, status FROM users where id = ?"
	stringId := strconv.FormatInt(id, 10)
	rows, err := database.SqlDB.Query(queryGetUser, stringId)
	if err != nil {
		return nil, errors.NewInternalServerError("Could not fetch the record")
	}
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
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

func (user *User) Search(status string) ([]User, *errors.RestErr) {

	findStmt, err := database.SqlDB.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError("Could not prepare the query for a find")
	}
	defer findStmt.Close()

	rows, err := findStmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError("Could not query the database for the find by status")
	}
	defer rows.Close()
	if err != nil {
		return nil, errors.NewInternalServerError("Could not close rows")
	}

	results := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			return nil, errors.NewInternalServerError("Could not fetch the record(s)")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("User with status '%s' does not exist", status))
	}
	return results, nil
}
