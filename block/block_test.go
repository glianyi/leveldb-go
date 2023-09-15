package block

import (
	"encoding/binary"
	"testing"
)

func TestBinary(t *testing.T) {
	buf := make([]byte, 10)
	// println(binary.PutUvarint(buf, 1))
	// println(binary.PutUvarint(buf, 10))
	// println(binary.PutUvarint(buf, 100))
	// println(binary.PutUvarint(buf, 1000))
	// println(binary.PutUvarint(buf, 10000))
	// println(binary.PutUvarint(buf, 100000))

	binary.LittleEndian.PutUint16(buf, 1)
}
