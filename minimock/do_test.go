package minimock_test

import (
	mm "github.com/gojuno/minimock"
	"github.com/jamesjoshuahill/test-doubles-golang/minimock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Do", func() {
	It("logs a message", func() {
		mc := mm.NewController(GinkgoT())
		mockLogger := minimock.NewloggerMock(mc)
		mockLogger.InfoMock.Expect("running do function").Return()

		minimock.Do(mockLogger)

		Expect(mockLogger.InfoCounter).To(Equal(uint64(1)))
		mc.Finish()
	})
})
