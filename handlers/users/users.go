package users

import (
	"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword ...
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash ...
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Create an user
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}

	i := bson.NewObjectId()
	user.ID = i
	password := user.Password
	hash, _ := HashPassword(password)
	user.Password = hash

	user.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	user.UpdatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	query := bson.M{
		"username": user.Username,
	}
	//existing := true
	err = db.C(models.CollectionUser).Find(query).Sort("-_id").One(&user)
	if err == nil {
		fmt.Println("error", err)
		fmt.Println("existing", user)
		c.Error(err)
		return
	}

	//if existing == true {
	err = db.C(models.CollectionUser).Insert(user)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, user)

}

// Options Method
func Options(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	users := []models.User{}
	err := db.C(models.CollectionUser).Find(nil).Sort("-_id").All(&users)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// List all users
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	users := []models.User{}
	err := db.C(models.CollectionUser).Find(nil).Sort("-_id").All(&users)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, users)
}

// Update an user
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}

	password := user.Password
	hash, _ := HashPassword(password)
	user.Password = hash

	query := bson.M{"_id": user.ID}
	doc := bson.M{
		"username":  user.Username,
		"password":  user.Password,
		"email":     user.Email,
		"role":      user.Role,
		"updatedOn": time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionUser).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an user
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	user := models.User{}
	err := c.Bind(&user)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": user.ID}
	err = db.C(models.CollectionUser).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
