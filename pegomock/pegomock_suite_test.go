package pegomock_test

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
)

var (
	Describe = ginkgo.Describe
	Context  = ginkgo.Context
	It       = ginkgo.It
)

func TestPegomock(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	pegomock.RegisterMockFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "Pegomock Suite")
}
