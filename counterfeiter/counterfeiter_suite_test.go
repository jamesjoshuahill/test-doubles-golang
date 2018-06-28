package counterfeiter_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCounterfeiter(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Counterfeiter Suite")
}
