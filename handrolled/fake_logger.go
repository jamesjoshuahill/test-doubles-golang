package handrolled

type FakeLogger struct {
	InfoCallCount int
	InfoReceived  string
}

func (f *FakeLogger) Info(msg string) {
	f.InfoCallCount++
	f.InfoReceived = msg
}
