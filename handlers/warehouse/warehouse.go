package warehouse

import (
	// "fmt"
	"net/http"
	// "time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an warehouse
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	warehouse := models.Warehouse{}
	err := c.Bind(&warehouse)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionWarehouse).Insert(warehouse)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all warehouse
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	warehouse := []models.Warehouse{}
	err := db.C(models.CollectionWarehouse).Find(nil).Sort("-_id").All(&warehouse)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, warehouse)
}

// Update an warehouse
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	warehouse := models.Warehouse{}
	err := c.Bind(&warehouse)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": warehouse.ID}
	doc := bson.M{
		"code":          warehouse.Code,
		"warehouse":     warehouse.Warehouse,
		"address":       warehouse.Address,
		"contactNo":     warehouse.ContactNo,
		"contactPerson": warehouse.ContactPerson,
	}
	err = db.C(models.CollectionWarehouse).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an warehouse
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	warehouse := models.Warehouse{}
	err := c.Bind(&warehouse)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": warehouse.ID}
	err = db.C(models.CollectionWarehouse).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
