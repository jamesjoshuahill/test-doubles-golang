package pegomocks_test

import "github.com/jamesjoshuahill/test-doubles-golang/pegomocks"

var _ = Describe("Do", func() {
	It("logs a message", func() {
		mockLogger := NewMocklogger()

		pegomocks.Do(mockLogger)

		mockLogger.VerifyWasCalledOnce().Info("running do function")
	})
})
