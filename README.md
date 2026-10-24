<p align="center">
  <img src="assets/hoard.png" alt="Hoard" width="160">
</p>

<h1 align="center">Hoard</h1>

<p align="center"><em>Everything a wyrm keeps, nothing it loses.</em><br>A learn-by-building storage engine in Go.</p>

---

Hoard is a storage engine built from scratch to learn how databases work underneath: write-ahead logging, LSM trees, SSTables, Bloom filters, compaction, a B+tree for comparison, and Raft replication.

It is built one module at a time. Each module ends with a graded deliverable: a crash test, a benchmark, or a measurement compared against theory. Nothing counts as done until the numbers say so.

## Modules

| # | Module | Graded by |
|---|---|---|
| 00 | [Harness](curriculum/00-harness.md) | Catches a planted 1-in-10,000 bug |
| 01 | [Durable log](curriculum/01-durable-log.md) | 1,000 SIGKILLs, zero corrupt records |
| 02 | [Memtable + WAL](curriculum/02-memtable-wal.md) | Zero acknowledged writes lost across 500 crashes |
| 03 | [SSTables](curriculum/03-sstables.md) | Read amplification measured and explained |
| 04 | [Bloom filters](curriculum/04-bloom-filters.md) | Measured false positive rate matches theory |
| 05 | [Compaction](curriculum/05-compaction.md) | Size-tiered vs leveled, write and space amplification |
| 06 | [B+tree](curriculum/06-btree.md) | Head-to-head benchmark against the LSM tree |
| 07 | [Raft](curriculum/07-raft.md) | 10-minute chaos test, zero acknowledged losses |
| 08 | [Capstone](curriculum/08-capstone.md) | OpenHearth migrates a live bucket with zero divergence |

Progress and grades: [notes/progress.md](notes/progress.md)

## Layout

```
engine/          the storage engine
  store.go       the Store interface every backend implements
  oracle/        a simple in-memory reference used to verify everything else
bench/           benchmark runner
curriculum/      one spec per module: goal, questions, build, graded deliverable
notes/           answers, measurements, and what each module taught
```

## Rules

- Go standard library only inside `engine/`.
- Every backend passes the same conformance, differential, and crash suites, unchanged.
- Crash tests use real process kills, not simulated errors.
- Each passing module gets a tag: `git tag module-NN`.

## Running

```
go test ./...
```

## Reading list

- *Designing Data-Intensive Applications*, Martin Kleppmann
- *Database Internals*, Alex Petrov
- CMU 15-445/645 Database Systems (lectures are free online)
- MIT 6.5840 Distributed Systems (Raft labs)
