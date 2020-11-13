package inventoryHistory

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/marcidblue-sales-api/models"
)

// Create an inventoryHistory
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	inventoryHistory := models.InventoryHistory{}
	err := c.Bind(&inventoryHistory)
	if err != nil {
		c.Error(err)
		return
	}

	i := bson.NewObjectId()
	inventoryHistory.ID = i
	inventoryHistory.DateUpdated = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionInventoryHistory).Insert(inventoryHistory)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, inventoryHistory)
}

// List all inventoryHistory
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	inventoryHistory := []models.InventoryHistory{}
	err := db.C(models.CollectionInventoryHistory).Find(nil).Sort("-_id").All(&inventoryHistory)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, inventoryHistory)
}

// Update an inventoryHistory
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	inventoryHistory := models.InventoryHistory{}
	err := c.Bind(&inventoryHistory)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": inventoryHistory.ID}
	doc := bson.M{
		"productCode": inventoryHistory.ProductCode,
		"productName": inventoryHistory.ProductName,
		// "transactionTypeCode": inventoryHistory.TransactionTypeCode,
		// "transactionTypeName": inventoryHistory.TransactionTypeName,
		// "warehouseFromCode":   inventoryHistory.WarehouseFromCode,
		// "warehouseFrom":       inventoryHistory.WarehouseFrom,
		// "warehouseToCode":     inventoryHistory.WarehouseToCode,
		// "warehouseTo":         inventoryHistory.WarehouseTo,
		"warehouseCode":  inventoryHistory.WarehouseCode,
		"warehouse":      inventoryHistory.Warehouse,
		"stockBeginning": inventoryHistory.StockBeginning,
		"stockIn":        inventoryHistory.StockIn,
		"stockOut":       inventoryHistory.StockOut,
		"stockEnding":    inventoryHistory.StockEnding,
		"dateUpdated":    time.Now().UnixNano() / int64(time.Millisecond),
	}

	err = db.C(models.CollectionInventoryHistory).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an inventoryHistory
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	inventoryHistory := models.InventoryHistory{}
	err := c.Bind(&inventoryHistory)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": inventoryHistory.ID}
	err = db.C(models.CollectionInventoryHistory).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
