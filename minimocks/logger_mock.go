package minimocks

/*
DO NOT EDIT!
This code was generated automatically using github.com/gojuno/minimock v1.9
The original interface "logger" can be found in github.com/jamesjoshuahill/test-doubles-golang
*/
import (
	"sync/atomic"
	"time"

	"github.com/gojuno/minimock"
	testify_assert "github.com/stretchr/testify/assert"
)

//loggerMock implements github.com/jamesjoshuahill/test-doubles-golang.logger
type loggerMock struct {
	t minimock.Tester

	InfoFunc       func(p string)
	InfoCounter    uint64
	InfoPreCounter uint64
	InfoMock       mloggerMockInfo
}

//NewloggerMock returns a mock for github.com/jamesjoshuahill/test-doubles-golang.logger
func NewloggerMock(t minimock.Tester) *loggerMock {
	m := &loggerMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.InfoMock = mloggerMockInfo{mock: m}

	return m
}

type mloggerMockInfo struct {
	mock             *loggerMock
	mockExpectations *loggerMockInfoParams
}

//loggerMockInfoParams represents input parameters of the logger.Info
type loggerMockInfoParams struct {
	p string
}

//Expect sets up expected params for the logger.Info
func (m *mloggerMockInfo) Expect(p string) *mloggerMockInfo {
	m.mockExpectations = &loggerMockInfoParams{p}
	return m
}

//Return sets up a mock for logger.Info to return Return's arguments
func (m *mloggerMockInfo) Return() *loggerMock {
	m.mock.InfoFunc = func(p string) {
		return
	}
	return m.mock
}

//Set uses given function f as a mock of logger.Info method
func (m *mloggerMockInfo) Set(f func(p string)) *loggerMock {
	m.mock.InfoFunc = f
	return m.mock
}

//Info implements github.com/jamesjoshuahill/test-doubles-golang.logger interface
func (m *loggerMock) Info(p string) {
	atomic.AddUint64(&m.InfoPreCounter, 1)
	defer atomic.AddUint64(&m.InfoCounter, 1)

	if m.InfoMock.mockExpectations != nil {
		testify_assert.Equal(m.t, *m.InfoMock.mockExpectations, loggerMockInfoParams{p},
			"logger.Info got unexpected parameters")

		if m.InfoFunc == nil {

			m.t.Fatal("No results are set for the loggerMock.Info")

			return
		}
	}

	if m.InfoFunc == nil {
		m.t.Fatal("Unexpected call to loggerMock.Info")
		return
	}

	m.InfoFunc(p)
}

//InfoMinimockCounter returns a count of loggerMock.InfoFunc invocations
func (m *loggerMock) InfoMinimockCounter() uint64 {
	return atomic.LoadUint64(&m.InfoCounter)
}

//InfoMinimockPreCounter returns the value of loggerMock.Info invocations
func (m *loggerMock) InfoMinimockPreCounter() uint64 {
	return atomic.LoadUint64(&m.InfoPreCounter)
}

//ValidateCallCounters checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *loggerMock) ValidateCallCounters() {

	if m.InfoFunc != nil && atomic.LoadUint64(&m.InfoCounter) == 0 {
		m.t.Fatal("Expected call to loggerMock.Info")
	}

}

//CheckMocksCalled checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish method or use Finish method of minimock.Controller
func (m *loggerMock) CheckMocksCalled() {
	m.Finish()
}

//Finish checks that all mocked methods of the interface have been called at least once
//Deprecated: please use MinimockFinish or use Finish method of minimock.Controller
func (m *loggerMock) Finish() {
	m.MinimockFinish()
}

//MinimockFinish checks that all mocked methods of the interface have been called at least once
func (m *loggerMock) MinimockFinish() {

	if m.InfoFunc != nil && atomic.LoadUint64(&m.InfoCounter) == 0 {
		m.t.Fatal("Expected call to loggerMock.Info")
	}

}

//Wait waits for all mocked methods to be called at least once
//Deprecated: please use MinimockWait or use Wait method of minimock.Controller
func (m *loggerMock) Wait(timeout time.Duration) {
	m.MinimockWait(timeout)
}

//MinimockWait waits for all mocked methods to be called at least once
//this method is called by minimock.Controller
func (m *loggerMock) MinimockWait(timeout time.Duration) {
	timeoutCh := time.After(timeout)
	for {
		ok := true
		ok = ok && (m.InfoFunc == nil || atomic.LoadUint64(&m.InfoCounter) > 0)

		if ok {
			return
		}

		select {
		case <-timeoutCh:

			if m.InfoFunc != nil && atomic.LoadUint64(&m.InfoCounter) == 0 {
				m.t.Error("Expected call to loggerMock.Info")
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
func (m *loggerMock) AllMocksCalled() bool {

	if m.InfoFunc != nil && atomic.LoadUint64(&m.InfoCounter) == 0 {
		return false
	}

	return true
}
