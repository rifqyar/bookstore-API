package dto

type CreateBook struct {
	Title       string  `json:"title" form:"title" binding:"required"`
	Author      string  `json:"author" form:"author" binding:"required"`
	Price       float64 `json:"price" form:"price" binding:"required"`
	Stock       int     `json:"stock" form:"stock" binding:"required"`
	Year        int     `json:"year" form:"year" binding:"required"`
	CategoryID  uint    `json:"category_id" form:"category_id" binding:"required"`
	ImageBase64 string  `json:"image_base64" form:"image_base64" binding:"required"`
}

type UpdateBook struct {
	Title       *string  `json:"title" form:"title"`
	Author      *string  `json:"author" form:"author"`
	Price       *float64 `json:"price" form:"price"`
	Stock       *int     `json:"stock" form:"stock"`
	Year        *int     `json:"year" form:"year"`
	CategoryID  *uint    `json:"category_id" form:"category_id"`
	ImageBase64 *string  `json:"image_base64" form:"image_base64"`
}
