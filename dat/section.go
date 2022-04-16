package dat

import (
	"encoding/binary"
	"fmt"
)

type Section struct {
	Type        rune
	ContentSize uint16
	Content     string // remove null character first
}

func (s *Section) UnmarshalText(text []byte) error {
	s.Type = rune(text[0])
	s.ContentSize = binary.LittleEndian.Uint16(text[1:3])
	if s.ContentSize != 0 {
		s.Content = string(text[3 : 3+s.ContentSize-1]) // drop null terminator

	}

	return nil
}

func (s *Section) String() string {
	return fmt.Sprintf("Section:\n\tType: %c\n\tContent Length: %d bytes\n\tContent: %q\n", s.Type, s.ContentSize-1, s.Content)
}
