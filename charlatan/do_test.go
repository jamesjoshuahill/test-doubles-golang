package charlatan_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/jamesjoshuahill/test-doubles-golang/charlatan"
)

var _ = Describe("Do", func() {
	It("logs a message", func() {
		fakeLogger := new(charlatan.Fakelogger)
		fakeLogger.InfoHook = func(msg string) {}

		charlatan.Do(fakeLogger)

		Expect(fakeLogger.InfoCalledOnceWith("running do function")).To(BeTrue())
	})
})
