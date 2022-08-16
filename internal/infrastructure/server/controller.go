package server

type ErrorResponse struct {
	Message string `json:"message"`
}

type CreateBookBody struct {
	Title    string  `json:"title" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	ISBN     string  `json:"isbn" binding:"required"`
	Stock    int     `json:"stock" binding:"required"`
	AuthorID uint    `json:"author_id" binding:"required"`
}
