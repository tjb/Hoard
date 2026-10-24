# Module 05: Compaction

## Goal
Merge SSTables in the background so reads stay fast and deleted data actually goes away. Implement size-tiered first, then leveled, and measure both.

## Why it matters
Compaction is where LSM engines win or lose in production. Write amplification, space amplification, and latency stalls all live here. The Airbnb kernel role explicitly lists tuning compaction, caching, and I/O.

## Before you code
Answer in `notes/05-compaction.md`:
1. Define write amplification, read amplification, and space amplification. Why can't you minimize all three?
2. When is it safe to drop a tombstone during compaction?
3. Compaction runs while reads and writes continue. What has to be true for a reader that started before compaction finished?

## Build
- `engine/compaction/`: a size-tiered strategy and a leveled strategy behind a common interface, selectable at `Open`.
- Background compaction goroutine with a manifest update that atomically swaps input SSTables for outputs.
- Tombstone garbage collection.

## Constraints
- Standard library only.
- Compaction must not block foreground reads.

## Graded deliverable
Run the same workloads under both strategies: write-heavy (90% Put), read-heavy (90% Get), and delete-heavy (50% Delete). For each, measure:
- write amplification (bytes written to disk / bytes written by user)
- space amplification (disk bytes / live data bytes)
- p99 read latency
- p99 write latency during compaction

Then write a paragraph: which strategy wins which workload, and why.

## Done when
- Suites and crash tests pass under both strategies, including crashes mid-compaction.
- The measurement table and the paragraph are in your notes.
- `git tag module-05`

## Reading
- RocksDB wiki: "Leveled Compaction" and "Universal Compaction"
- O'Neil et al., "The Log-Structured Merge-Tree" (1996)
- Search: "RUM conjecture" (Athanassoulis et al.)
