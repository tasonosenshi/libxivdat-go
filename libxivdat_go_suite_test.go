package libxivdatgo_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestLibxivdatGo(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LibxivdatGo Suite")
}
