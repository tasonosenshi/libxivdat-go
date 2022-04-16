package dat

import (
	"encoding/binary"
	"fmt"
	"strings"
)

const acqType uint32 = 0x00640006
const acqDefaultMaxSize uint32 = 2048
const xorMask byte = 0x73

type ACQ struct {
	FileType    uint32
	MaxFileSize uint32
	ContentSize uint32
	HeaderEnd   byte
	Sections    []Section // encoded with xor mask
}

func (f *ACQ) UnmarshalBinary(data []byte) error {
	f.FileType = binary.LittleEndian.Uint32(data[:0x4])
	// Max file size doesn't account for 17 byte header and 15 byte null-padded footer
	f.MaxFileSize = binary.LittleEndian.Uint32(data[0x4:0x8])
	f.ContentSize = binary.LittleEndian.Uint32(data[0x8:0xc])
	f.HeaderEnd = data[0x10]

	// Remove the header and null-padded footer
	data = data[0x11 : len(data)-0xe]

	// xor rest of the file
	for i, b := range data {
		data[i] = b ^ xorMask
	}

	var pos uint32 = 0
	for pos < f.ContentSize-1 {
		s := Section{}
		s.UnmarshalText(data[pos:])
		f.Sections = append(f.Sections, s)

		pos += uint32(3 + s.ContentSize)
	}

	return nil
}

func (f *ACQ) String() string {
	var b strings.Builder

	b.WriteString("ACQ File (Recent Tells)\n")

	b.WriteString(fmt.Sprintf("File Max Size: %d bytes\n", f.MaxFileSize))

	b.WriteString(fmt.Sprintf("File Content Size: %d bytes\n", f.ContentSize))

	for _, s := range f.Sections {
		b.WriteString(s.String())
		b.WriteRune('\n')
	}

	return b.String()
}
