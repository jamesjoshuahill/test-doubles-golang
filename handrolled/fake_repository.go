package handrolled

import "github.com/jamesjoshuahill/test-doubles-golang"

type FakeRepository struct {
	QueryMock struct {
		CallCount int
		Received  struct {
			Name string
			Kind string
		}
		Returns struct {
			Records []doubles.Record
			Error   error
		}
	}
}

func (f *FakeRepository) Query(name, kind string) ([]doubles.Record, error) {
	f.QueryMock.CallCount++
	f.QueryMock.Received.Name = name
	f.QueryMock.Received.Kind = kind

	return f.QueryMock.Returns.Records, f.QueryMock.Returns.Error
}
