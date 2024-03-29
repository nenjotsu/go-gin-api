package bottleSalesOrder

import (
	//"fmt"
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an bottleSalesOrder
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottleSalesOrder := models.BottleSalesOrder{}
	err := c.Bind(&bottleSalesOrder)
	if err != nil {
		c.Error(err)
		return
	}

	i := bson.NewObjectId()
	bottleSalesOrder.ID = i
	bottleSalesOrder.DateEncoded = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionBottleSalesOrder).Insert(bottleSalesOrder)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, bottleSalesOrder)
}

// List all bottleSalesOrders
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottleSalesOrders := []models.BottleSalesOrder{}
	err := db.C(models.CollectionBottleSalesOrder).Find(nil).Sort("-_id").All(&bottleSalesOrders)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, bottleSalesOrders)
}

// Update an bottleSalesOrder
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	bottleSalesOrder := models.BottleSalesOrder{}
	err := c.Bind(&bottleSalesOrder)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": bottleSalesOrder.ID}
	doc := bson.M{
		"customerName":             bottleSalesOrder.CustomerName,
		"customerCode":             bottleSalesOrder.CustomerCode,
		"customerParent":           bottleSalesOrder.CustomerParent,
		"customerType":             bottleSalesOrder.CustomerType,
		"orderType":                bottleSalesOrder.OrderType,
		"shipTo":                   bottleSalesOrder.ShipTo,
		"poNo":                     bottleSalesOrder.PoNo,
		"warehouseFrom":            bottleSalesOrder.WarehouseFrom,
		"warehouseTo":              bottleSalesOrder.WarehouseTo,
		"remarks":                  bottleSalesOrder.Remarks,
		"isPaid":                   bottleSalesOrder.IsPaid,
		"dateEncoded":              bottleSalesOrder.DateEncoded,
		"orderedDate":              bottleSalesOrder.OrderedDate,
		"scheduledDeliveryDate":    bottleSalesOrder.ScheduledDeliveryDate,
		"actualDeliveryDate":       bottleSalesOrder.ActualDeliveryDate,
		"returnDate":               bottleSalesOrder.ReturnDate,
		"receivedDateFromSupplier": bottleSalesOrder.ReceivedDateFromSupplier,
		"totalNetPrice":            bottleSalesOrder.TotalNetPrice,
		"totalGrossPrice":          bottleSalesOrder.TotalGrossPrice,
		"transactionTypeCode":      bottleSalesOrder.TransactionTypeCode,
		"transactionTypeName":      bottleSalesOrder.TransactionTypeName,
		"status":                   bottleSalesOrder.Status,
		"details":                  bottleSalesOrder.Details,
		"expenses":                 bottleSalesOrder.Expenses,
	}
	err = db.C(models.CollectionBottleSalesOrder).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an bottleSalesOrder
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	bottleSalesOrder := models.BottleSalesOrder{}
	err := c.Bind(&bottleSalesOrder)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": bottleSalesOrder.ID}
	err = db.C(models.CollectionBottleSalesOrder).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
