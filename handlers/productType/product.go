package productType

import (
	//"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an product
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	product := models.ProductType{}
	err := c.Bind(&product)
	if err != nil {
		c.Error(err)
		return
	}

	product.CreatedOn = time.Now().UnixNano() / int64(time.Millisecond)
	product.UpdatedOn = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionProductType).Insert(product)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all products
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	products := []models.Product{}
	err := db.C(models.CollectionProductType).Find(nil).Sort("-_id").All(&products)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, products)
}

// Update an product
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	product := models.ProductType{}
	err := c.Bind(&product)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": product.ID}
	doc := bson.M{
		"title":      product.Title,
		"body":       product.Body,
		"created_on": product.CreatedOn,
		"updated_on": time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionProductType).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an product
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	product := models.Product{}
	err := c.Bind(&product)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": product.ID}
	err = db.C(models.CollectionProductType).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
