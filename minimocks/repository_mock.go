package minimocks

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "repository" can be found in github.com/jamesjoshuahill/test-doubles-golang
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	doubles "github.com/jamesjoshuahill/test-doubles-golang"

	testify_assert "github.com/stretchr/testify/assert"
)

//repositoryMock implements github.com/jamesjoshuahill/test-doubles-golang.repository
type repositoryMock struct {
	t minimock.Tester

	QueryFunc       func(p string) (r []doubles.Record, r1 error)
	QueryCounter    uint64
	QueryPreCounter uint64
	QueryMock       mrepositoryMockQuery
}

//NewrepositoryMock returns a mock for github.com/jamesjoshuahill/test-doubles-golang.repository
func NewrepositoryMock(t minimock.Tester) *repositoryMock {
	m := &repositoryMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.QueryMock = mrepositoryMockQuery{mock: m}

	return m
}

type mrepositoryMockQuery struct {
	mock             *repositoryMock
	mockExpectations *repositoryMockQueryParams
}

//repositoryMockQueryParams represents input parameters of the repository.Query
type repositoryMockQueryParams struct {
	p string
}

//Expect sets up expected params for the repository.Query
func (m *mrepositoryMockQuery) Expect(p string) *mrepositoryMockQuery {
	m.mockExpectations = &repositoryMockQueryParams{p}
	return m
}

//Return sets up a mock for repository.Query to return Return's arguments
func (m *mrepositoryMockQuery) Return(r []doubles.Record, r1 error) *repositoryMock {
	m.mock.QueryFunc = func(p string) ([]doubles.Record, error) {
		return r, r1
	}
	return m.mock
}

//Set uses given function f as a mock of repository.Query method
func (m *mrepositoryMockQuery) Set(f func(p string) (r []doubles.Record, r1 error)) *repositoryMock {
	m.mock.QueryFunc = f
	return m.mock
}

//Query implements github.com/jamesjoshuahill/test-doubles-golang.repository interface
func (m *repositoryMock) Query(p string) (r []doubles.Record, r1 error) {
	atomic.AddUint64(&m.QueryPreCounter, 1)
	defer atomic.AddUint64(&m.QueryCounter, 1)

	if m.QueryMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.QueryMock.mockExpectations, repositoryMockQueryParams{p},
			"repository.Query got unexpected parameters")

		if m.QueryFunc == nil {

			m.t.Fatal("No results are set for the repositoryMock.Query")

			return
		}
	}

	if m.QueryFunc == nil {
		m.t.Fatal("Unexpected call to repositoryMock.Query")
		return
	}

	return m.QueryFunc(p)
}

//QueryMinimockCounter returns a count of repositoryMock.QueryFunc invocations
func (m *repositoryMock) QueryMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.QueryCounter)
}

//QueryMinimockPreCounter returns the value of repositoryMock.Query invocations
func (m *repositoryMock) QueryMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.QueryPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *repositoryMock) ValidateCallCounters() {

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to repositoryMock.Query")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *repositoryMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *repositoryMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *repositoryMock) MinimockFinish() {

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		m.t.Fatal("Expected call to repositoryMock.Query")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *repositoryMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *repositoryMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.QueryFunc == nil || atomic.LoadUint64(&m.QueryCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
				m.t.Error("Expected call to repositoryMock.Query")
			}

			m.t.Fatalf("Some mocks were not called on time: %s", timeout)
			return
		default:
			time.Sleep(time.Millisecond)
		}
	}
}

//AllMocksCalled returns true if all mocked methods were called before the execution of AllMocksCalled,
//it can be used with assert/require, i.e. assert.True(mock.AllMocksCalled())
func (m *repositoryMock) AllMocksCalled() bool {

	if m.QueryFunc != nil && atomic.LoadUint64(&m.QueryCounter) == 0 {
		return false
	}

	return true
}
