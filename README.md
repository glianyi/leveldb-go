### lsm
https://github.com/skyzh/mini-lsm/blob/main/mini-lsm-book/src/00-get-started.md
https://izualzhy.cn/leveldb-block
https://bean-li.github.io/leveldb-sstable/

### binary
https://nakabonne.dev/posts/binary-encoding-go/
https://protobuf.dev/programming-guides/encoding/#varints

### plan
Day 1: Block encoding. SSTs are composed of multiple data blocks. We will implement the block encoding.
Day 2: SST encoding.
Day 3: MemTable and Merge Iterators.
Day 4: Block cache and Engine. To reduce disk I/O and maximize performance, we will use moka-rs to build a block cache for the LSM tree. In this day we will get a functional (but not persistent) key-value engine with get, put, scan, delete API.
Day 5: Compaction. Now it's time to maintain a leveled structure for SSTs.
Day 6: Recovery. We will implement WAL and manifest so that the engine can recover after restart.
Day 7: Bloom filter and key compression. They are widely-used optimizations in LSM tree structures.

