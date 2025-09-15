package dto

type SalesReportResponse struct {
	Revenue   float64 `json:"revenue" example:"1500000"`
	BooksSold int64   `json:"books_sold" example:"25"`
}

type BestsellerReportResponse struct {
	BookID uint   `json:"book_id" example:"1"`
	Title  string `json:"title" example:"Laskar Pelangi"`
	Sold   int64  `json:"sold" example:"15"`
}

type PriceStatsReportResponse struct {
	Max float64 `json:"max" example:"120000"`
	Min float64 `json:"min" example:"25000"`
	Avg float64 `json:"avg" example:"75000"`
}
