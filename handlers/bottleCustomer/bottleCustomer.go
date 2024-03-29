package bottleCustomer

import (
	// "fmt"
	"net/http"
	// "time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an bottleCustomer
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottleCustomer := models.BottleCustomer{}
	err := c.Bind(&bottleCustomer)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionBottleCustomer).Insert(bottleCustomer)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all bottleCustomers
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottleCustomers := []models.BottleCustomer{}
	err := db.C(models.CollectionBottleCustomer).Find(nil).Sort("-_id").All(&bottleCustomers)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, bottleCustomers)
}

// Update an bottleCustomer
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottleCustomer := models.BottleCustomer{}
	err := c.Bind(&bottleCustomer)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bottleCustomer.ID}
	doc := bson.M{
		"code":         bottleCustomer.Code,
		"customerType": bottleCustomer.CustomerType,
		"customer":     bottleCustomer.Customer,
		"parent":       bottleCustomer.Parent,
		"status":       bottleCustomer.Status,
	}
	err = db.C(models.CollectionBottleCustomer).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an bottleCustomer
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottleCustomer := models.BottleCustomer{}
	err := c.Bind(&bottleCustomer)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": bottleCustomer.ID}
	err = db.C(models.CollectionBottleCustomer).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
