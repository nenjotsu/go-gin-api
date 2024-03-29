package transactionType

import (
	//"fmt"
	"net/http"
	// "time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an transactionType
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	transactionType := models.TransactionType{}
	err := c.Bind(&transactionType)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionTransactionType).Insert(transactionType)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all transactionTypes
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	transactionTypes := []models.TransactionType{}
	err := db.C(models.CollectionTransactionType).Find(nil).Sort("-_id").All(&transactionTypes)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, transactionTypes)
}

// Update an transactionType
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	transactionType := models.TransactionType{}
	err := c.Bind(&transactionType)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": transactionType.ID}
	doc := bson.M{
		"_id":             transactionType.ID,
		"code":            transactionType.Code,
		"transactionType": transactionType.TransactionType,
	}
	err = db.C(models.CollectionTransactionType).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an transactionType
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	transactionType := models.TransactionType{}
	err := c.Bind(&transactionType)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": transactionType.ID}
	err = db.C(models.CollectionTransactionType).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
