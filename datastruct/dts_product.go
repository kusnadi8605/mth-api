package datastruct

//Product ..
type Product struct {
	ProductID   int64
	SKU         string
	ProductName string
	ProductDesc string
	Price       float64
	Quantity    int32
	CreatedDate string
	CreatedBy   string
	UpdatedDate string
	UpdatedBy   string
}

//ProductReq ..
type ProductReq struct {
	SKU         string  `json:"sku" validate:"required"`
	ProductName string  `json:"productName" validate:"required"`
	ProductDesc string  `json:"productDesc" validate:"required"`
	Price       float64 `json:"price" validate:"gte=0"`
	Quantity    int32   `json:"quantity" validate:"gte=0"`
	UserID      string  `json:"userId" validate:"required"`
}

//ProductAct for detail & delete
type ProductAct struct {
	SKU string `json:"sku" validate:"required"`
}

//ProductResponse data
type ProductResponse struct {
	ResponseCode string    `json:"responseCode"`
	ResponseDesc string    `json:"responseDesc"`
	Payload      []Product `json:"payload,omitempty"`
}
