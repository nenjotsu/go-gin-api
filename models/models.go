package models

import "gopkg.in/mgo.v2/bson"

// Models
const (
	CollectionUser             = "users"
	CollectionBottleSalesOrder = "bottle_sales_orders"
	CollectionTransactionType  = "transaction_type"
	CollectionBottleCustomer   = "bottle_customers"
	CollectionBottleAccounting = "bottle_accounting"
	CollectionBottlePayment    = "bottle_payment"
	CollectionRoles            = "roles"
	CollectionRemarks          = "remarks"
	CollectionWarehouse        = "warehouse"
	CollectionInventory        = "inventory"
	CollectionInventoryHistory = "inventory_history"
)

// User model
type User struct {
	ID        bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string        `json:"username" form:"username" binding:"required" bson:"username"`
	Password  string        `json:"password" form:"password" binding:"required" bson:"password"`
	Email     string        `json:"email" form:"email" bson:"email"`
	Role      string        `json:"role" form:"role" bson:"role"`
	CreatedOn int64         `json:"createdOn" bson:"createdOn"`
	UpdatedOn int64         `json:"updatedOn" bson:"updatedOn"`
}

// BottleSalesOrder model
type BottleSalesOrder struct {
	ID                       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerName             string        `json:"customerName" form:"customerName" bson:"customerName"`
	CustomerCode             string        `json:"customerCode" form:"customerCode" bson:"customerCode"`
	CustomerParent           string        `json:"customerParent" form:"customerParent" bson:"customerParent"`
	CustomerType             string        `json:"customerType" form:"customerType" bson:"customerType"`
	OrderType                string        `json:"orderType" form:"orderType" bson:"orderType"`
	ShipTo                   string        `json:"shipTo" form:"shipTo" bson:"shipTo"`
	PoNo                     string        `json:"poNo" form:"poNo" bson:"poNo"`
	WarehouseFrom            string        `json:"warehouseFrom" form:"warehouseFrom" bson:"warehouseFrom"`
	WarehouseTo              string        `json:"warehouseTo" form:"warehouseTo" bson:"warehouseTo"`
	Remarks                  string        `json:"remarks" form:"remarks" bson:"remarks"`
	IsPaid                   bool          `json:"isPaid" form:"isPaid" bson:"isPaid"`
	DateEncoded              int64         `json:"dateEncoded" bson:"dateEncoded"`
	OrderedDate              int64         `json:"orderedDate" bson:"orderedDate"`
	ScheduledDeliveryDate    int64         `json:"scheduledDeliveryDate" bson:"scheduledDeliveryDate"`
	ActualDeliveryDate       int64         `json:"actualDeliveryDate" bson:"actualDeliveryDate"`
	ReturnDate               int64         `json:"returnDate" bson:"returnDate"`
	ReceivedDateFromSupplier int64         `json:"receivedDateFromSupplier" bson:"receivedDateFromSupplier"`
	TotalNetPrice            float32       `json:"totalNetPrice" form:"totalNetPrice" bson:"totalNetPrice"`
	TotalGrossPrice          float32       `json:"totalGrossPrice" form:"totalGrossPrice" bson:"totalGrossPrice"`
	TransactionTypeCode      string        `json:"transactionTypeCode" form:"transactionTypeCode" bson:"transactionTypeCode"`
	TransactionTypeName      string        `json:"transactionTypeName" form:"transactionTypeName" bson:"transactionTypeName"`
	Status                   string        `json:"status" form:"status" bson:"status"`
	Details                  []struct {
		ProductName        string  `json:"productName" form:"productName" bson:"productName"`
		ProductDescription string  `json:"productDescription" form:"productDescription" bson:"productDescription"`
		ProductCode        string  `json:"productCode" form:"productCode" bson:"productCode"`
		ProductType        string  `json:"productType" form:"productType" bson:"productType"`
		PromoCode          string  `json:"promoCode" form:"promoCode" bson:"promoCode"`
		Qty                int     `json:"qty" form:"qty" bson:"qty"`
		ComplimentaryQty   int     `json:"complimentaryQty" form:"complimentaryQty" bson:"complimentaryQty"`
		UnitPrice          float32 `json:"unitPrice" form:"unitPrice" bson:"unitPrice"`
		NetPrice           float32 `json:"netPrice" form:"netPrice" bson:"netPrice"`
		GrossPrice         float32 `json:"grossPrice" form:"grossPrice" bson:"grossPrice"`
		TotalNetPrice      float32 `json:"totalNetPrice" form:"totalNetPrice" bson:"totalNetPrice"`
		TotalGrossPrice    float32 `json:"totalGrossPrice" form:"totalGrossPrice" bson:"totalGrossPrice"`
		Uom                string  `json:"uom" form:"uom" bson:"uom"`
	} `json:"details" form:"details" bson:"details"`
	Expenses []struct {
		ID           string  `json:"_id" form:"_id" bson:"_id"`
		Particulars  string  `json:"particulars" form:"particulars" bson:"particulars"`
		ReceivedBy   string  `json:"receivedBy" form:"receivedBy" bson:"receivedBy"`
		Amount       float32 `json:"amount" form:"amount" bson:"amount"`
		Remarks      string  `json:"remarks" form:"remarks" bson:"remarks"`
		ReleasedDate int64   `json:"releasedDate" bson:"releasedDate"`
		Operation    string  `json:"operation" bson:"operation"`
	} `json:"expenses" form:"expenses" bson:"expenses"`
}

// TransactionType model
type TransactionType struct {
	ID              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code            string        `json:"code" form:"code" bson:"code"`
	TransactionType string        `json:"transactionType" form:"transactionType" bson:"transactionType"`
}

// Inventory model
type Inventory struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductCode   string        `json:"productCode" form:"productCode" bson:"productCode"`
	ProductName   string        `json:"productName" form:"productName" bson:"productName"`
	Uom           string        `json:"uom" form:"uom" bson:"uom"`
	WarehouseCode string        `json:"warehouseCode" form:"warehouseCode" bson:"warehouseCode"`
	Warehouse     string        `json:"warehouse" form:"warehouse" bson:"warehouse"`
	StockCount    int32         `json:"stockCount" form:"stockCount" bson:"stockCount"`
	DateUpdated   int64         `json:"dateUpdated" form:"dateUpdated" bson:"dateUpdated"`
}

// InventoryHistory model
type InventoryHistory struct {
	ID             bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerName   string        `json:"customerName" form:"customerName" bson:"customerName"`
	CustomerCode   string        `json:"customerCode" form:"customerCode" bson:"customerCode"`
	CustomerParent string        `json:"customerParent" form:"customerParent" bson:"customerParent"`
	CustomerType   string        `json:"customerType" form:"customerType" bson:"customerType"`
	PoNo           string        `json:"poNo" form:"poNo" bson:"poNo"`
	ProductCode    string        `json:"productCode" form:"productCode" bson:"productCode"`
	ProductName    string        `json:"productName" form:"productName" bson:"productName"`
	Uom            string        `json:"uom" form:"uom" bson:"uom"`
	WarehouseCode  string        `json:"warehouseCode" form:"warehouseCode" bson:"warehouseCode"`
	Warehouse      string        `json:"warehouse" form:"warehouse" bson:"warehouse"`
	WarehouseFrom  string        `json:"warehouseFrom" form:"warehouseFrom" bson:"warehouseFrom"`
	WarehouseTo    string        `json:"warehouseTo" form:"warehouseTo" bson:"warehouseTo"`
	StockBeginning int32         `json:"stockBeginning" form:"stockBeginning" bson:"stockBeginning"`
	StockIn        int32         `json:"stockIn" form:"stockIn" bson:"stockIn"`
	StockOut       int32         `json:"stockOut" form:"stockOut" bson:"stockOut"`
	StockEnding    int32         `json:"stockEnding" form:"stockEnding" bson:"stockEnding"`
	DateUpdated    int64         `json:"dateUpdated" form:"dateUpdated" bson:"dateUpdated"`
}

// BottleCustomer model
type BottleCustomer struct {
	ID           bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code         int32         `json:"code" form:"code" bson:"code"`
	CustomerType string        `json:"customerType" form:"customerType" bson:"customerType"`
	Customer     string        `json:"customer" form:"customer" bson:"customer"`
	Parent       string        `json:"parent" form:"parent" bson:"parent"`
	Status       bool          `json:"status" form:"status" bson:"status"`
}

// BottleAccounting model
type BottleAccounting struct {
	ID                       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerName             string        `json:"customerName" form:"customerName" bson:"customerName"`
	CustomerCode             string        `json:"customerCode" form:"customerCode" bson:"customerCode"`
	CustomerParent           string        `json:"customerParent" form:"customerParent" bson:"customerParent"`
	CustomerType             string        `json:"customerType" form:"customerType" bson:"customerType"`
	AccumulatedOrderAmount   float32       `json:"accumulatedOrderAmount" form:"accumulatedOrderAmount" bson:"accumulatedOrderAmount"`       // 220,000
	AccumulatedPaymentAmount float32       `json:"accumulatedPaymentAmount" form:"accumulatedPaymentAmount" bson:"accumulatedPaymentAmount"` // 50,000
	TotalBalance             float32       `json:"totalBalance" form:"totalBalance" bson:"totalBalance"`
	LastPaymentAmount        float32       `json:"lastPaymentAmount" form:"lastPaymentAmount" bson:"lastPaymentAmount"` // 50,000                                  // 200,000
	LastPaymentDate          int64         `json:"lastPaymentDate" form:"lastPaymentDate" bson:"lastPaymentDate"`
	DateUpdated              int64         `json:"dateUpdated" form:"dateUpdated" bson:"dateUpdated"`
	ProductName              string        `json:"productName" form:"productName" bson:"productName"`
	ProductCode              string        `json:"productCode" form:"productCode" bson:"productCode"`
}

// BottlePayment model
type BottlePayment struct {
	ID                      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	PaymentType             string        `json:"paymentType" form:"paymentType" bson:"paymentType"` // Previous Balance, Current Order, Advanced Payment
	ArNumber                string        `json:"arNumber" form:"arNumber" bson:"arNumber"`
	CustomerName            string        `json:"customerName" form:"customerName" bson:"customerName"`
	CustomerCode            string        `json:"customerCode" form:"customerCode" bson:"customerCode"`
	CustomerParent          string        `json:"customerParent" form:"customerParent" bson:"customerParent"`
	CustomerType            string        `json:"customerType" form:"customerType" bson:"customerType"`
	DateUpdated             int64         `json:"dateUpdated" form:"dateUpdated" bson:"dateUpdated"`
	PaymentDate             int64         `json:"paymentDate" form:"paymentDate" bson:"paymentDate"`
	CashPaymentAmount       float32       `json:"cashPaymentAmount" form:"cashPaymentAmount" bson:"cashPaymentAmount"`
	TotalCheckPaymentAmount float32       `json:"totalCheckPaymentAmount" form:"totalCheckPaymentAmount" bson:"totalCheckPaymentAmount"`
	TotalPaymentAmount      float32       `json:"totalPaymentAmount" form:"totalPaymentAmount" bson:"totalPaymentAmount"`
	PreparedBy              string        `json:"preparedBy" form:"preparedBy" bson:"preparedBy"`
	ProductName             string        `json:"productName" form:"productName" bson:"productName"`
	ProductCode             string        `json:"productCode" form:"productCode" bson:"productCode"`
	CheckPayment            []struct {
		ID                 string  `json:"_id" form:"_id" bson:"_id"`
		CheckPaymentAmount float32 `json:"checkPaymentAmount" form:"checkPaymentAmount" bson:"checkPaymentAmount"`
		CheckNo            int32   `json:"checkNo" form:"checkNo" bson:"checkNo"`
		Date               int64   `json:"date" form:"date" bson:"date"`
	} `json:"checkPayment" form:"checkPayment" bson:"checkPayment"`
	OtherCurrency []struct {
		ID       string  `json:"_id" form:"_id" bson:"_id"`
		Currency string  `json:"currency" form:"currency" bson:"currency"`
		Amount   float32 `json:"amount" form:"amount" bson:"amount"`
	} `json:"otherCurrency" form:"otherCurrency" bson:"otherCurrency"`
	CountedBy []struct {
		ID   string `json:"_id" form:"_id" bson:"_id"`
		Name string `json:"name" form:"name" bson:"name"`
	} `json:"countedBy" form:"countedBy" bson:"countedBy"`
}

// Remarks model
type Remarks struct {
	ID      bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code    string        `json:"code" form:"code" bson:"code"`
	Remarks string        `json:"remarks" form:"remarks" bson:"remarks"`
}

// Roles model
type Roles struct {
	ID       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" form:"code" bson:"code"`
	RoleName string        `json:"roleName" form:"roleName" bson:"roleName"`
}

// Warehouse model
type Warehouse struct {
	ID            bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code          string        `json:"code" form:"code" bson:"code"`
	Warehouse     string        `json:"warehouse" form:"warehouse" bson:"warehouse"`
	Address       string        `json:"address" form:"address" bson:"address"`
	ContactNo     string        `json:"contactNo" form:"contactNo" bson:"contactNo"`
	ContactPerson string        `json:"contactPerson" form:"contactPerson" bson:"contactPerson"`
}
