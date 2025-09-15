package dto

type OrderItemRequest struct {
	BookID   uint `json:"book_id" example:"1"`
	Quantity int  `json:"quantity" example:"2"`
}

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}
