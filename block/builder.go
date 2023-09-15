package block

import (
	"encoding/binary"
)

// Day 1: Block encoding. SSTs are composed of multiple data blocks. We will implement the block encoding.

/*
---------------------------------------------------------------------
|          data         |          offsets          |      meta     |
|-----------------------|---------------------------|---------------|
|entry|entry|entry|entry|offset|offset|offset|offset|num_of_elements|
---------------------------------------------------------------------

+---------------+---------------+---------------+---------------+
| entry 1 | entry 2 |      ...   |entry n-1 | entry n | trailer |
+---------------+---------------+---------------+---------------+

-----------------------------------------------------------------------
|                           Entry #1                            | ... |
-----------------------------------------------------------------------
| key_len (2B) | key (keylen) | value_len (2B) | value (varlen) | ... |
-----------------------------------------------------------------------


-------------------------------------------------------------------------
|                                 trailer
-------------------------------------------------------------------------
| offset 1 | offset 2 | offset 2 | ... | offset last | number_of_elements|
--------------------------------------------------------------------------
*/

const (
	DEFAULT_BLOCK_SIZE = 4 * 1024 // 4KB
)

type BlockBuilder struct {
	data      []byte
	offsets   []int
	blockSize int
	totalSize int
}

func New(blockSize int) *BlockBuilder {
	return &BlockBuilder{
		data:      make([]byte, blockSize),
		offsets:   make([]int, 0),
		blockSize: blockSize,
	}
}

func (bb *BlockBuilder) add(key, value []byte) bool {
	l := binary.PutUvarint([]byte{}, uint64(bb.totalSize))
	if !bb.isEmpty() && bb.totalSize+len(key)+len(value)+l > bb.blockSize {
		return false
	}
	n := binary.PutUvarint([]byte{}, uint64(bb.totalSize))
	bb.totalSize += n
	bb.offsets = append(bb.offsets, len(bb.data))
	n = binary.PutUvarint(bb.data, uint64(len(key)))
	bb.totalSize += n
	bb.data = append(bb.data, key...)
	n = binary.PutUvarint(bb.data, uint64(len(value)))
	bb.totalSize += n
	bb.data = append(bb.data, value...)

	return true
}

func (bb *BlockBuilder) isEmpty() bool {
	return len(bb.offsets) == 0
}

func (bb *BlockBuilder) build() *Block {
	return &Block{
		data:    bb.data,
		offsets: bb.offsets,
	}
}
