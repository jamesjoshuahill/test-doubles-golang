package handrolled_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/test-doubles-golang/handrolled"
)

var _ = Describe("Do", func() {
	It("logs a message", func() {
		fakeLogger := new(FakeLogger)

		handrolled.Do(fakeLogger)

		Expect(fakeLogger.InfoCallCount).To(Equal(1))
		Expect(fakeLogger.InfoReceived).To(Equal("running do function"))
	})
})
