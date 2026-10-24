# Module 02: Memtable + WAL

## Goal
Your first real `engine.Store`: writes go to the WAL from module 01, then into an in-memory sorted structure. On restart, replay the WAL to rebuild memory.

## Why it matters
This is the write path of RocksDB, Cassandra, Pebble, and LevelDB. It is also the moment the durability contract in `engine/store.go` becomes something you can prove.

## Before you code
Answer in `notes/02-memtable-wal.md`:
1. Why a skiplist instead of a hash map? Instead of a balanced tree?
2. How do you represent a delete in the WAL and the memtable? Why can't you just remove the key?
3. If you ack a Put before fsync, what breaks? If you fsync every single write, what does it cost? What is group commit?

## Build
- `engine/memtable/`: a skiplist keyed by bytes, supporting point lookup and ordered iteration, with tombstones for deletes.
- `engine/lsm/`: a `Store` that combines WAL + memtable. `Open(dir)` replays the WAL.
- Wire `engine/storetest` against it.

## Constraints
- Standard library only. Write the skiplist yourself.
- Memory size is unbounded for now. Flushing to disk is module 03.

## Graded deliverable
1. Conformance and differential suites pass.
2. Crash test: kill the process mid-workload 500 times; after recovery, the store matches the oracle for every acknowledged operation.
3. The experiment: add a flag that acks before fsync. Run the crash test again and record how many acknowledged writes are lost. Then implement group commit and benchmark writes/sec for: fsync every write, group commit, no fsync.

## Done when
- Both suites pass, the crash test shows zero acknowledged losses.
- `notes/02-memtable-wal.md` has the lost-write count for ack-before-fsync and the three-way write benchmark.
- `git tag module-02`

## Reading
- *Database Internals*, Chapter 7 (Log-Structured Storage)
- Pugh, "Skip Lists: A Probabilistic Alternative to Balanced Trees" (1990)
