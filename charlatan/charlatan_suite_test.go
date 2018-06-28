package charlatan_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//go:generate charlatan repository logger

func TestCharlatan(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Charlatan Suite")
}
