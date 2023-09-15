package block

import "encoding/binary"

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

type Block struct {
	data    []byte
	offsets []int
}

func (b *Block) encode() []byte {
	buf := b.data[0:]
	for _, v := range b.offsets {
		binary.PutUvarint(buf, uint64(v))
	}
	binary.PutUvarint(buf, uint64(len(b.offsets)))
	return buf
}

func (b *Block) decode() *Block {
	return &Block{}
}
