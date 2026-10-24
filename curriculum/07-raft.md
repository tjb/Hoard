# Module 07: Raft replication

## Goal
Replicate Hoard across three nodes with Raft, so the store survives losing a machine and never loses an acknowledged write.

## Why it matters
Consensus and shard management are part of every distributed database role on the target list. Reddit, Airbnb, and Netflix all run replicated storage where this is the failure model.

## How this module works
This one is split across two repos:
1. Do the MIT 6.5840 Raft labs (3A leader election, 3B log, 3C persistence, 3D snapshots) in their own repo, exactly as given. Their test harness is the grader.
2. Port your Raft implementation into `engine/raft/` and put the LSM store behind it as the state machine.

## Before you code
Answer in `notes/07-raft.md`:
1. Why does Raft require a majority, and what exactly goes wrong with a two-node cluster?
2. A leader is partitioned but does not know it yet. A client writes to it. What happens to that write?
3. How do Raft snapshots relate to your LSM's SSTables and WAL? Can you avoid storing the data twice?

## Build
- `engine/raft/`: your 6.5840 implementation, adapted to real networking (gRPC or net/rpc) and real disk persistence (reuse `engine/wal`).
- A replicated `Store` where Put and Delete go through the Raft log; reads go to the leader (linearizable) with an option for follower reads (stale).

## Constraints
- Raft code is yours. No external Raft libraries.

## Graded deliverable
A chaos test: three nodes, continuous writes, and a fault injector that randomly kills nodes, partitions the leader, and restarts nodes, for 10 minutes. Afterward:
- every acknowledged write is present on all live nodes
- no node ever returned a value that was never written
- report: leader elections, writes acked, writes rejected, max unavailability window

## Done when
- 6.5840 lab tests pass in their repo.
- Hoard chaos test passes with zero acknowledged losses.
- `git tag module-07`

## Reading
- Ongaro and Ousterhout, "In Search of an Understandable Consensus Algorithm" (the Raft paper)
- MIT 6.5840 course page and lab handouts
- Jepsen analyses of etcd and CockroachDB (search "jepsen etcd")
