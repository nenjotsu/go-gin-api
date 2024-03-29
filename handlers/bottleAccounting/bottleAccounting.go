package bottleAccounting

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an bottleAccounting
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottleAccounting := models.BottleAccounting{}
	err := c.Bind(&bottleAccounting)
	if err != nil {
		c.Error(err)
		return
	}
	bottleAccounting.DateUpdated = time.Now().UnixNano() / int64(time.Millisecond)
	err = db.C(models.CollectionBottleAccounting).Insert(bottleAccounting)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, bottleAccounting)
}

// List all bottleAccountList
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottleAccountList := []models.BottleAccounting{}
	err := db.C(models.CollectionBottleAccounting).Find(nil).Sort("-_id").All(&bottleAccountList)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, bottleAccountList)
}

// Update an bottleAccounting
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottleAccounting := models.BottleAccounting{}
	err := c.Bind(&bottleAccounting)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bottleAccounting.ID}
	doc := bson.M{
		"customerCode":             bottleAccounting.CustomerCode,
		"customerType":             bottleAccounting.CustomerType,
		"customerName":             bottleAccounting.CustomerName,
		"customerParent":           bottleAccounting.CustomerParent,
		"productCode":              bottleAccounting.ProductCode,
		"productName":              bottleAccounting.ProductName,
		"accumulatedOrderAmount":   bottleAccounting.AccumulatedOrderAmount,
		"accumulatedPaymentAmount": bottleAccounting.AccumulatedPaymentAmount,
		"totalBalance":             bottleAccounting.TotalBalance,
		"lastPaymentAmount":        bottleAccounting.LastPaymentAmount,
		"lastPaymentDate":          bottleAccounting.LastPaymentDate,
		"dateUpdated":              time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionBottleAccounting).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an bottleAccounting
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottleAccounting := models.BottleAccounting{}
	err := c.Bind(&bottleAccounting)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": bottleAccounting.ID}
	err = db.C(models.CollectionBottleAccounting).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
