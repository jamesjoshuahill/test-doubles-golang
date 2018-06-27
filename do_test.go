package doubles_test

import (
	"github.com/gojuno/minimock"
	"github.com/jamesjoshuahill/test-doubles-golang"
	"github.com/jamesjoshuahill/test-doubles-golang/charlatan"
	"github.com/jamesjoshuahill/test-doubles-golang/counterfeiter"
	"github.com/jamesjoshuahill/test-doubles-golang/handrolled"
	"github.com/jamesjoshuahill/test-doubles-golang/minimocks"
	"github.com/jamesjoshuahill/test-doubles-golang/pegomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Do", func() {
	Context("with a hand-rolled fake logger", func() {
		It("logs a message", func() {
			fakeLogger := new(handrolled.FakeLogger)

			doubles.Do(fakeLogger)

			Expect(fakeLogger.InfoCallCount).To(Equal(1))
			Expect(fakeLogger.InfoReceived).To(Equal("running do function"))
		})
	})

	Context("with a counterfeiter fake logger", func() {
		It("logs a message", func() {
			fakeLogger := new(counterfeiter.FakeLogger)

			doubles.Do(fakeLogger)

			Expect(fakeLogger.InfoCallCount()).To(Equal(1))
			Expect(fakeLogger.InfoArgsForCall(0)).To(Equal("running do function"))
		})
	})

	Context("with a charlatan fake logger", func() {
		It("logs a message", func() {
			fakeLogger := new(charlatan.Fakelogger)
			fakeLogger.InfoHook = func(msg string) {}

			doubles.Do(fakeLogger)

			Expect(fakeLogger.InfoCalledOnceWith("running do function")).To(BeTrue())
		})
	})

	Context("with a pegomock mock logger", func() {
		It("logs a message", func() {
			mockLogger := pegomock.NewMocklogger()

			doubles.Do(mockLogger)

			mockLogger.VerifyWasCalledOnce().Info("running do function")
		})
	})

	Context("with a minimock mock logger", func() {
		It("logs a message", func() {
			mc := minimock.NewController(GinkgoT())
			mockLogger := minimocks.NewloggerMock(mc)
			mockLogger.InfoMock.Expect("running do function").Return()

			doubles.Do(mockLogger)

			Expect(mockLogger.InfoCounter).To(Equal(uint64(1)))
			mc.Finish()
		})
	})
})
