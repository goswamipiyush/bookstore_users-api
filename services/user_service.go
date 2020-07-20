//Services layer that will actually have ALL business logic. Whether you use to write to an RDBMS here or NOSQL OR anything else; you can change whenever you want
//Controller should never ever change and be aware of service layer

package services

import (
	"database/sql"

	"github.com/goswamipiyush/bookstore_users-api/domain/users"
	"github.com/goswamipiyush/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	//We have a clean "savable" user at this point; so save to database possible
	saveErr := user.Save()
	if saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}

func GetUser(id int64, user users.User) (*users.User, *errors.RestErr) {
	result, err := user.Get(id)
	if err != nil {
		return nil, err
	}
	//Seems we have got a 'good' user back, return it
	return result, nil
}

func DeleteUser(id int64, user users.User) (*sql.Result, *errors.RestErr) {
	result, err := user.Delete(id)
	if err != nil {
		return nil, err
	}
	//Seems like the user is deleted now
	return result, nil
}

func Search(status string) ([]users.User, *errors.RestErr) {
	dao := &users.User{}
	users, err := dao.Search(status)
	if err != nil {
		return nil, err
	}
	//Seems we have got a 'good' user back, return it
	return users, nil
}

func UpdateUser(id int64, user users.User) (*users.User, *errors.RestErr) {
	err := user.Validate()
	if err != nil {
		return nil, err
	}

	//We have a clean "savable" user at this point; so send to 'update database'
	saveErr := user.Update(id)
	if saveErr != nil {
		return nil, saveErr
	}
	return &user, nil
}
