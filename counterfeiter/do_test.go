package counterfeiter_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/test-doubles-golang/counterfeiter"
	"github.com/jamesjoshuahill/test-doubles-golang/counterfeiter/counterfeiterfakes"
)

var _ = Describe("Do", func() {
	It("logs a message", func() {
		fakeLogger := new(counterfeiterfakes.FakeLogger)

		counterfeiter.Do(fakeLogger)

		Expect(fakeLogger.InfoCallCount()).To(Equal(1))
		Expect(fakeLogger.InfoArgsForCall(0)).To(Equal("running do function"))
	})
})
