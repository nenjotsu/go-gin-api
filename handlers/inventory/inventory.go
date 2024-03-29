package inventory

import (
	"net/http"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/models"
)

// Create an inventory
func Create(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	inventory := models.Inventory{}
	err := c.Bind(&inventory)
	if err != nil {
		c.Error(err)
		return
	}

	i := bson.NewObjectId()
	inventory.ID = i
	inventory.DateUpdated = time.Now().UnixNano() / int64(time.Millisecond)

	err = db.C(models.CollectionInventory).Insert(inventory)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// List all inventory
func List(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	inventory := []models.Inventory{}
	err := db.C(models.CollectionInventory).Find(nil).Sort("-_id").All(&inventory)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, inventory)
}

// Update an inventory
func Update(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)

	inventory := models.Inventory{}
	err := c.Bind(&inventory)
	if err != nil {
		c.Error(err)
		return
	}

	query := bson.M{"_id": inventory.ID}
	doc := bson.M{
		"productCode":   inventory.ProductCode,
		"productName":   inventory.ProductName,
		"uom":           inventory.Uom,
		"warehouseCode": inventory.WarehouseCode,
		"warehouse":     inventory.Warehouse,
		"stockCount":    inventory.StockCount,
		"dateUpdated":   time.Now().UnixNano() / int64(time.Millisecond),
	}
	err = db.C(models.CollectionInventory).Update(query, doc)
	if err != nil {
		c.Error(err)
		return
	}
}

// Delete an inventory
func Delete(c *gin.Context) {
	db := c.MustGet("db").(*mgo.Database)
	inventory := models.Inventory{}
	err := c.Bind(&inventory)
	if err != nil {
		c.Error(err)
		return
	}
	query := bson.M{"_id": inventory.ID}
	err = db.C(models.CollectionInventory).Remove(query)
	if err != nil {
		c.Error(err)
		return
	}
}
