package doubles_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/petergtz/pegomock"
)

func TestTestDoublesGolang(t *testing.T) {
	RegisterFailHandler(Fail)
	pegomock.RegisterMockFailHandler(Fail)
	RunSpecs(t, "Doubles Suite")
}
