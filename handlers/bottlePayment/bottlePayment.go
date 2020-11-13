package bottlePayment

import (
	//"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/marcidblue-sales-api/models"
)

// Create an bottlePayment
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottlePayment := models.BottlePayment{}
	err := c.Bind(&bottlePayment)
	if err != nil {
		c.Error(err)
		return
	}
	bottlePayment.DateUpdated = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionBottlePayment).Insert(bottlePayment)
	if err != nil {
		c.Error(err)
		return
	}
}

// List all bottlePaymentList
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottlePaymentList := []models.BottlePayment{}
	err := db.C(models.CollectionBottlePayment).Find(nil).Sort("-_id").All(&bottlePaymentList)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, bottlePaymentList)
}

// Update an bottlePayment
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottlePayment := models.BottlePayment{}
	err := c.Bind(&bottlePayment)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bottlePayment.ID}
	doc := bson.M{
		"paymentType":             bottlePayment.PaymentType,
		"arNumber":                bottlePayment.ArNumber,
		"customerName":            bottlePayment.CustomerName,
		"customerCode":            bottlePayment.CustomerCode,
		"customerType":            bottlePayment.CustomerType,
		"customerParent":          bottlePayment.CustomerParent,
		"dateUpdated":             time.Now().UnixNano() / int64(time.Millisecond),
		"paymentDate":             bottlePayment.PaymentDate,
		"cashPaymentAmount":       bottlePayment.CashPaymentAmount,
		"totalCheckPaymentAmount": bottlePayment.TotalCheckPaymentAmount,
		"totalPaymentAmount":      bottlePayment.TotalPaymentAmount,
		"preparedBy":              bottlePayment.PreparedBy,
		"checkPayment":            bottlePayment.CheckPayment,
		"otherCurrency":           bottlePayment.OtherCurrency,
		"countedBy":               bottlePayment.CountedBy,
	}
	err = db.C(models.CollectionBottlePayment).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an bottlePayment
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottlePayment := models.BottlePayment{}
	err := c.Bind(&bottlePayment)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": bottlePayment.ID}
	err = db.C(models.CollectionBottlePayment).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
