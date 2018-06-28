package pegomock_test

import "github.com/jamesjoshuahill/test-doubles-golang/pegomock"

var _ = Describe("Do", func() {
	It("logs a message", func() {
		mockLogger := NewMocklogger()

		pegomock.Do(mockLogger)

		mockLogger.VerifyWasCalledOnce().Info("running do function")
	})
})
