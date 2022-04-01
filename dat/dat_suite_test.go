package dat_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDat(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DAT Suite")
}
