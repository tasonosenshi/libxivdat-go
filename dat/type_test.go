package dat_test

import (
	"encoding/binary"
	"os"
	fp "path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"libxivdat-go/dat"
)

const defaultDatRelPath = "../resources/default_dats"

var _ = Describe("Type", func() {
	var FileTypeTable []TableEntry = []TableEntry{
		Entry(nil, dat.ACQ, fp.Join(defaultDatRelPath, "ACQ.DAT")),
		Entry(nil, dat.GEARSET, fp.Join(defaultDatRelPath, "GEARSET.DAT")),
		Entry(nil, dat.GS, fp.Join(defaultDatRelPath, "GS.DAT")),
		Entry(nil, dat.ITEMFDR, fp.Join(defaultDatRelPath, "ITEMFDR.DAT")),
		Entry(nil, dat.ITEMODR, fp.Join(defaultDatRelPath, "ITEMODR.DAT")),
		Entry(nil, dat.KEYBIND, fp.Join(defaultDatRelPath, "KEYBIND.DAT")),
		Entry(nil, dat.LOGFLTR, fp.Join(defaultDatRelPath, "LOGFLTR.DAT")),
		Entry(nil, dat.MACRO, fp.Join(defaultDatRelPath, "MACRO.DAT")),
		Entry(nil, dat.UISAVE, fp.Join(defaultDatRelPath, "UISAVE.DAT")),
	}

	DescribeTable("Getting type from header bytes", func(dtype dat.Type, fp string) {
		Ω(fp).Should(BeARegularFile())

		file, err := os.Open(fp)
		Ω(err).ShouldNot(HaveOccurred())

		header := make([]byte, 4)
		n, err := file.Read(header)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(n).Should(BeNumerically("==", 4))

		id_bytes := binary.LittleEndian.Uint32(header)
		Ω(id_bytes).Should(BeEquivalentTo(dtype))
	}, FileTypeTable)

	DescribeTable("Getting default end byte for type", func(dtype dat.Type, fp string) {
		_, err := dat.GetDefaultEndByteForType(dtype)
		Ω(err).ShouldNot(HaveOccurred())
	}, FileTypeTable)

	DescribeTable("Getting default max size for type", func(dtype dat.Type, fp string) {
		_, err := dat.GetDefaultEndByteForType(dtype)
		Ω(err).ShouldNot(HaveOccurred())
	}, FileTypeTable)

	DescribeTable("Getting bitmask for type", func(dtype dat.Type, fp string) {
		_, err := dat.GetDefaultEndByteForType(dtype)
		Ω(err).ShouldNot(HaveOccurred())
	}, FileTypeTable)
})
