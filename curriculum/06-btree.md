# Module 06: B+tree

## Goal
Build a page-based B+tree behind the same `engine.Store` interface, then benchmark it head to head against your LSM tree.

## Why it matters
"Explain the latency differences between B-tree and LSM storage engines" is a literal line item in storage engine job postings. After this module you answer it with your own numbers.

## Before you code
Answer in `notes/06-btree.md`:
1. Why do B+trees keep values only in leaves and link the leaves together?
2. A page split happens and the process crashes halfway. What state is the tree in? How do real databases prevent this?
3. Predict, before measuring: which engine wins random writes, random reads, range scans, and space efficiency? Write your predictions down now.

## Build
- `engine/btree/`: fixed-size pages (4 KB) in a single file, a page cache, insert with split, delete (merge or lazy deletion, your call, documented), range scan across leaves.
- Crash safety via your WAL from module 01 (redo logging of page changes).

## Constraints
- Standard library only.
- Same conformance and crash suites as the LSM tree, unchanged.

## Graded deliverable
The head-to-head benchmark, same machine, same data sizes (1M and 10M keys), both engines:
- random writes/sec, sequential writes/sec
- random read p50/p99
- 100-key range scan p50/p99
- disk bytes used and write amplification

Compare against your predictions from question 3. Explain every place you were wrong.

## Done when
- Suites and crash tests pass.
- Benchmark table, predictions, and explanations are in your notes. This is the most important artifact in the repo.
- `git tag module-06`

## Reading
- *Database Internals*, Chapters 2 through 6
- Search: "Modern B-Tree Techniques" (Graefe)
- CMU 15-445 lectures on B+trees and buffer pools
