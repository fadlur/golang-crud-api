package book

import "time"
type Book struct {
	ID			int
	Title		string
	Decription	string
	Price		int
	Rating		int
	CreatedAt	time.Time
	UpdatedAt	time.Time
}