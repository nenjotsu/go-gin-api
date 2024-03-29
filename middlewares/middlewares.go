// Package middlewares contains gin middlewares
// Usage: router.Use(middlewares.Connect)
package middlewares

import (
	"fmt"
	//"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/db"
)

// Connect middleware clones the database session for each request and
// makes the `db` object available for each handler
func Connect(c *gin.Context) {
	s := db.Session.Clone()

	defer s.Close()

	c.Set("db", s.DB(db.Mongo.Database))
	c.Next()
}

// ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	c.Next()

	fmt.Println(c.Errors)
	r := gin.Default()

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//if len(c.Errors) > 0 {
	//	fmt.Println(c.Errors)
	//c.Redirect(http.StatusMovedPermanently, "/public/404")
	//}
}
