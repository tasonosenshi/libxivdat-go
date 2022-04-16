package dat_test

import (
	"os"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"libxivdat-go/dat"
)

const defaultDatRelPath = "../resources/ACQ.DAT"

var _ = Describe("Acq", func() {
	var data []byte
	var acq dat.ACQ

	BeforeEach(func() {
		Ω(defaultDatRelPath).Should(BeARegularFile())

		var err error
		data, err = os.ReadFile(defaultDatRelPath)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(data).ShouldNot(BeEmpty())

		acq = dat.ACQ{}
		err = acq.UnmarshalBinary(data)
		Ω(err).ShouldNot(HaveOccurred())

		GinkgoWriter.Print(acq.String())
	})

	It("Should have the correct header info", func() {
		Ω(acq.FileType).Should(BeNumerically("==", 0x00640006))
		Ω(acq.MaxFileSize).Should(BeNumerically("==", len(data)-32))
		Ω(acq.ContentSize).ShouldNot(BeZero())
		Ω(acq.HeaderEnd).Should(BeNumerically("==", 0xFF))
	})
})
