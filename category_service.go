package doubles

import (
	"fmt"
)

//go:generate charlatan -output charlatan/fake_repository.go -package charlatan repository
//go:generate counterfeiter -o counterfeiter/fake_repository.go . repository
//go:generate minimock -i github.com/jamesjoshuahill/test-doubles-golang.repository -o minimocks -s _mock.go
//go:generate pegomock generate --use-experimental-model-gen -o pegomock/mock_repository.go --package pegomock repository
type repository interface {
	Query(name, kind string) ([]Record, error)
}

type Record struct {
	Category string
}

type CategoryService struct {
	repository repository
}

func NewCategoryService(r repository) *CategoryService {
	return &CategoryService{
		repository: r,
	}
}

func (s CategoryService) First(name string) (Record, error) {
	records, err := s.repository.Query(name, "category")
	if err != nil {
		return Record{}, fmt.Errorf("finding categories with name %q failed: %s", name, err)
	}

	if len(records) == 0 {
		return Record{}, fmt.Errorf("no categories found with name %q", name)
	}

	return records[0], nil
}
