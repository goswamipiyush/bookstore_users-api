//Application package that is entry point and called from main.

package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartApplication() {

	//	mapUrls()
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	router.Run(":8080")
}
