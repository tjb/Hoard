# Module 01: Durable log

## Goal
Build an append-only log file that survives crashes: every record is either fully there and verified, or detectably absent. This is the foundation every later module writes through.

## Why it matters
"Once Put returns nil, the data survives a crash" is the hardest sentence in storage. Write-ahead logs, Kafka partitions, Raft logs, and OpenHearth's replication log are all this module.

## Before you code
Answer in `notes/01-durable-log.md`:
1. What does `write()` returning successfully actually guarantee? What does `fsync()` add? What about the directory entry for a newly created file?
2. A crash happens halfway through writing a 4 KB record. What can the file look like afterward? (There is more than one answer.)
3. How will a reader tell a torn record from a valid one? Why is a length prefix alone not enough?

## Build
- `engine/wal/`: `Open(path)`, `Append(record []byte) (offset, error)`, `Sync() error`, and a reader that iterates valid records from the start.
- Record format of your design. Document it in a comment at the top of the file, byte by byte.
- On open, recovery scans the log and truncates any torn tail.

## Constraints
- Standard library only.
- Checksums: use `hash/crc32` (Castagnoli table).

## Graded deliverable
A crash test in `engine/wal/crash_test.go`: a child process appends records in a loop and is killed with SIGKILL at a random moment, 1,000 times. After each kill, reopen and verify:
- every record acknowledged after a `Sync` is present and intact
- zero corrupt records are ever returned
- the log accepts new appends after recovery

Report in your notes: records acknowledged vs recovered, and any losses.

## Done when
- 1,000 kills, zero corrupt records returned, zero synced records lost.
- You can explain your record format and recovery logic without looking.
- `git tag module-01`

## Reading
- *Designing Data-Intensive Applications*, Chapter 3 (the log section) and Chapter 7 (durability)
- Search: "files are hard" by Dan Luu
- Search: "PostgreSQL fsyncgate"
