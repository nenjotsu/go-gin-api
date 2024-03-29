package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sales-api/db"
	"github.com/sales-api/handlers/bottleAccounting"
	"github.com/sales-api/handlers/bottleCustomer"
	"github.com/sales-api/handlers/bottlePayment"
	"github.com/sales-api/handlers/bottleSalesOrder"
	"github.com/sales-api/handlers/inventory"
	"github.com/sales-api/handlers/inventoryHistory"
	"github.com/sales-api/handlers/remarks"
	"github.com/sales-api/handlers/roleType"
	"github.com/sales-api/handlers/transactionType"
	"github.com/sales-api/handlers/users"
	"github.com/sales-api/handlers/warehouse"
	"github.com/sales-api/middlewares"
	"github.com/sales-api/models"

	//"github.com/fvbock/endless" //this is for production
	"github.com/itsjamie/gin-cors"
	"gopkg.in/appleboy/gin-jwt.v2"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Port at which the server starts listening
const Port = "7324"

func init() {
	db.Connect()
}

// OptionsUser : this is add methods in login user
func OptionsUser(c *gin.Context) {
	c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, DELETE, POST, PATCH, PUT, OPTIONS")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
	c.Next()
}

func main() {
	router := gin.Default()
	//gin.SetMode(gin.ReleaseMode) // this is for production

	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE, OPTIONS",
		RequestHeaders:  "Origin, Authorization, Content-Type, X-Requested-With",
		ExposedHeaders:  "",
		MaxAge:          time.Hour * 12,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	router.RedirectTrailingSlash = false
	router.RedirectFixedPath = false

	secretKey := "SOME_SECRETE_HERE"
	hashSecretKey, _ := users.HashPassword(secretKey)

	authMiddleware := &jwt.GinJWTMiddleware{
		Realm:      hashSecretKey,
		Key:        []byte(hashSecretKey),
		Timeout:    time.Hour * 12,
		MaxRefresh: time.Hour * 12,
		Authenticator: func(username string, password string, c *gin.Context) (string, bool) {
			db := c.MustGet("db").(*mgo.Database)
			query := bson.M{
				"username": username,
			}
			user := models.User{}
			err := db.C(models.CollectionUser).Find(query).Sort("-_id").One(&user)
			bpassword := password
			match := users.CheckPasswordHash(bpassword, user.Password)

			credentials := user.Password + user.Username + user.Email + user.Role
			if err != nil || match != true {
				fmt.Println(err)
				c.Error(err)
				return credentials, false
			}

			return credentials, true
		},
		Authorizator: func(username string, c *gin.Context) bool {
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":         code,
				"errorMessage": message,
				"success":      false,
			})

		},
		TokenLookup: "header:Authorization",
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Connect)
	router.Use(middlewares.ErrorHandler)

	//Routes
	router.POST("/login", authMiddleware.LoginHandler)
	router.OPTIONS("/login", OptionsUser)

	auth := router.Group("/api")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)

		// Users
		auth.GET("/get/user", users.List)
		auth.POST("/create/user", users.Create)
		auth.POST("/update/user", users.Update)
		auth.POST("/delete/user", users.Delete)

		// SalesOrder
		auth.GET("/get/bottle_sales_order", bottleSalesOrder.List)
		auth.POST("/create/bottle_sales_order", bottleSalesOrder.Create)
		auth.POST("/update/bottle_sales_order", bottleSalesOrder.Update)
		auth.POST("/delete/bottle_sales_order", bottleSalesOrder.Delete)

		// TransactionType
		auth.GET("/get/transaction_type", transactionType.List)
		auth.POST("/create/transaction_type", transactionType.Create)
		auth.POST("/update/transaction_type", transactionType.Update)
		auth.POST("/delete/transaction_type", transactionType.Delete)

		// RoleType
		auth.GET("/get/role_type", roleType.List)
		auth.POST("/create/role_type", roleType.Create)
		auth.POST("/update/role_type", roleType.Update)
		auth.POST("/delete/role_type", roleType.Delete)

		// BottleCustomerList
		auth.GET("/get/bottle_customers", bottleCustomer.List)
		auth.POST("/create/bottle_customers", bottleCustomer.Create)
		auth.POST("/update/bottle_customers", bottleCustomer.Update)
		auth.POST("/delete/bottle_customers", bottleCustomer.Delete)

		// RemarksList
		auth.GET("/get/remarks", remarks.List)
		auth.POST("/create/remarks", remarks.Create)
		auth.POST("/update/remarks", remarks.Update)
		auth.POST("/delete/remarks", remarks.Delete)

		// WarehouseList
		auth.GET("/get/warehouse", warehouse.List)
		auth.POST("/create/warehouse", warehouse.Create)
		auth.POST("/update/warehouse", warehouse.Update)
		auth.POST("/delete/warehouse", warehouse.Delete)

		// InventoryList
		auth.GET("/get/inventory", inventory.List)
		auth.POST("/create/inventory", inventory.Create)
		auth.POST("/update/inventory", inventory.Update)
		auth.POST("/delete/inventory", inventory.Delete)

		// InventoryHisotryList
		auth.GET("/get/inventory_history", inventoryHistory.List)
		auth.POST("/create/inventory_history", inventoryHistory.Create)
		auth.POST("/update/inventory_history", inventoryHistory.Update)
		auth.POST("/delete/inventory_history", inventoryHistory.Delete)

		// BottleAccountingList
		auth.GET("/get/bottle_accounting", bottleAccounting.List)
		auth.POST("/create/bottle_accounting", bottleAccounting.Create)
		auth.POST("/update/bottle_accounting", bottleAccounting.Update)
		auth.POST("/delete/bottle_accounting", bottleAccounting.Delete)

		// BottlePayment
		auth.GET("/get/bottle_payment", bottlePayment.List)
		auth.POST("/create/bottle_payment", bottlePayment.Create)
		auth.POST("/update/bottle_payment", bottlePayment.Update)
		auth.POST("/delete/bottle_payment", bottlePayment.Delete)
	}

	// If the routes is not existing
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "ERROR"})
	})

	http.Handle("/api", router)

	// Start listening
	port := Port
	if len(os.Getenv("PORT")) > 0 {
		port = os.Getenv("PORT")
	}
	router.Run(":" + port)
	//endless.ListenAndServe(":"+port, router)
}
