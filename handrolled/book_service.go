package handrolled

import (
	"fmt"
)

type repository interface {
	Query(name, category string) ([]Record, error)
}

type Record struct {
	Name     string
	Category string
}

type BookService struct {
	repository repository
}

func NewBookService(r repository) *BookService {
	return &BookService{
		repository: r,
	}
}

func (s BookService) First(name string) (Record, error) {
	records, err := s.repository.Query(name, "book")
	if err != nil {
		return Record{}, fmt.Errorf("finding books with name %q failed: %s", name, err)
	}

	if len(records) == 0 {
		return Record{}, fmt.Errorf("no books found with name %q", name)
	}

	return records[0], nil
}
