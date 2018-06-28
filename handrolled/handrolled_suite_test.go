package handrolled_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHandrolled(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Handrolled Suite")
}
