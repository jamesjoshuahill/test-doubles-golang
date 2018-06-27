package handrolled

import "github.com/jamesjoshuahill/test-doubles-golang"

type FakeRepository struct {
	QueryMock struct {
		CallCount int
		Received  struct {
			Category string
		}
		Returns struct {
			Records []doubles.Record
			Error   error
		}
	}
}

func (f *FakeRepository) Query(category string) ([]doubles.Record, error) {
	f.QueryMock.CallCount++
	f.QueryMock.Received.Category = category

	return f.QueryMock.Returns.Records, f.QueryMock.Returns.Error
}
