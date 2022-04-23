package model

type Book struct {
	ID    int64  `json:"id" form:"id" db:"id"`
	Title string `json:"title" form:"title" db:"title"`
	Price int64  `json:"price" form:"price" db:"price"`
}

// NewBook 构造函数
func NewBook() *Book {
	return &Book{}
}
