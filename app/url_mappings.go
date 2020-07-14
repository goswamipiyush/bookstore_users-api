//Use this file to map incoming URLs to their actual handler functions.

package app

import (
	hello "github.com/goswamipiyush/bookstore_users-api/controllers/hello"
	ping "github.com/goswamipiyush/bookstore_users-api/controllers/ping"
	users "github.com/goswamipiyush/bookstore_users-api/controllers/users"
)

func mapUrls() {

	router.GET("/ping", ping.Ping)
	router.GET("/user/:name", hello.Hello)

	router.GET("/users/:userid", users.GetUser)

	router.POST("/users", users.CreateUser)

}
