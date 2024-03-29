package roleType

import (
	// "fmt"
	"net/http"
	// "time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an roles
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	roles := models.Roles{}
	err := c.Bind(&roles)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionRoles).Insert(roles)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all roles
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roles := []models.Roles{}
	err := db.C(models.CollectionRoles).Find(nil).Sort("-_id").All(&roles)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, roles)
}

// Update an roles
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	roles := models.Roles{}
	err := c.Bind(&roles)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": roles.ID}
	doc := bson.M{
		"code":     roles.Code,
		"roleName": roles.RoleName,
	}
	err = db.C(models.CollectionRoles).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an roles
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	roles := models.Roles{}
	err := c.Bind(&roles)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": roles.ID}
	err = db.C(models.CollectionRoles).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
