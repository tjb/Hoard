# Module 03: SSTables

## Goal
When the memtable gets big, flush it to an immutable, sorted file on disk. Reads now check the memtable first, then SSTables from newest to oldest.

## Why it matters
Immutable sorted files are why LSM trees write fast and why compaction exists. Understanding block layout is how you answer "how many disk reads does this lookup take?" in a storage interview.

## Before you code
Answer in `notes/03-sstables.md`:
1. Why make SSTables immutable? What does that buy you for concurrency and crash safety?
2. Why split the file into blocks with an index, instead of one big sorted run you binary search?
3. After a flush, when is it safe to delete the old WAL? What happens if you crash between writing the SSTable and deleting the WAL?

## Build
- `engine/sstable/`: a writer that takes sorted entries and produces data blocks, an index block, and a footer; a reader that does point lookups and range iteration.
- A manifest file that records which SSTables are live, updated atomically (write temp file, fsync, rename, fsync directory).
- `engine/lsm/` flushes the memtable when it passes a size threshold, then rotates the WAL.
- Merge iteration for Scan across memtable + all SSTables, newest version wins, tombstones hide older values.

## Constraints
- Standard library only.
- Block size configurable; default 4 KB.

## Graded deliverable
1. Suites pass with a tiny memtable threshold (forces many flushes).
2. Crash test including crashes during flush and manifest update. Zero acknowledged losses.
3. Measure: for a 1M-key store, average and p99 SSTables touched per point read, and block reads per lookup. Explain the numbers.

## Done when
- Suites and crash test pass.
- Notes include the read-amplification measurement and your explanation.
- `git tag module-03`

## Reading
- *Database Internals*, Chapter 7
- LevelDB `doc/table_format.md` (search "leveldb table format")
