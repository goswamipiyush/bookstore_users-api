//Use this file to map incoming URLs to their actual handler functions.

package app

import (
	"github.com/goswamipiyush/bookstore_users-api/controllers"
)

func mapUrls() {

	router.GET("/ping", controllers.Ping)
	router.GET("/user/:name", controllers.Hello)

}
