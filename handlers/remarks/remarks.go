package remarks

import (
	// "fmt"
	"net/http"
	// "time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an remarks
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	remarks := models.Remarks{}
	err := c.Bind(&remarks)
	if err != nil {
		c.Error(err)
		return
	}

	err = db.C(models.CollectionRemarks).Insert(remarks)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all remarks
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	remarks := []models.Remarks{}
	err := db.C(models.CollectionRemarks).Find(nil).Sort("-_id").All(&remarks)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, remarks)
}

// Update an remarks
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	remarks := models.Remarks{}
	err := c.Bind(&remarks)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": remarks.ID}
	doc := bson.M{
		"code":    remarks.Code,
		"remarks": remarks.Remarks,
	}
	err = db.C(models.CollectionRemarks).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an remarks
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	remarks := models.Remarks{}
	err := c.Bind(&remarks)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": remarks.ID}
	err = db.C(models.CollectionRemarks).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
