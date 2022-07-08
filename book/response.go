package book

type BookResponse struct {
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Decription	string	`json:"description"`
	Price		int		`json:"price"`
	Rating		int		`json:"rating"`
}