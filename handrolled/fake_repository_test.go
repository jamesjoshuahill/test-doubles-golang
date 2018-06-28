package handrolled_test

import "github.com/jamesjoshuahill/test-doubles-golang/handrolled"

type FakeRepository struct {
	QueryMock struct {
		CallCount int
		Received  struct {
			Name     string
			Category string
		}
		Returns struct {
			Records []handrolled.Record
			Error   error
		}
	}
}

func (f *FakeRepository) Query(name, category string) ([]handrolled.Record, error) {
	f.QueryMock.CallCount++
	f.QueryMock.Received.Name = name
	f.QueryMock.Received.Category = category

	return f.QueryMock.Returns.Records, f.QueryMock.Returns.Error
}
