package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
	Update(book BookRequest, ID int) (Book, error)
	Delete(ID int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service  {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error)  {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindByID(ID int) (Book, error)  {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error)  {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()
	book := Book {
		Title: bookRequest.Title,
		Price: int(price),
		Rating: int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Update(bookRequest BookRequest, ID int) (Book, error)  {
	book, err := s.repository.FindByID(ID)

	if err != nil {
		return book, err
	}
	
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Decription = bookRequest.Description
	book.Price = int(price)
	book.Rating = int(rating)

	newBook, err := s.repository.Update(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error)  {
	book, err := s.repository.FindByID(ID)

	if err != nil {
		return book, err
	}

	newBook, err := s.repository.Delete(book)
	return newBook, err
}