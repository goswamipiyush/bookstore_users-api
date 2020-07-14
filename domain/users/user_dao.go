//Only point in the application where we will interact with the database
package users

import "github.com/goswamipiyush/bookstore_users-api/utils/errors"

var (
	usersDB = make(map[int64]*User) //temporary database till we actually have a DB to persist data
)

func (user *User) Save() *errors.RestErr {

	if usersDB[user.Id] != nil {
		return errors.NewBadRequestError("User already exists")
	}
	usersDB[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError("User id not found")
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil

}
