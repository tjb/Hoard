# Module 08: Capstone (OpenHearth)

## Goal
Use what you built to ship a real tool: OpenHearth, an S3-compatible proxy that migrates a live bucket between two backends with zero downtime and zero divergence.

## Why it matters
Hoard proves you understand storage engines. OpenHearth proves you can apply that understanding to a problem people actually have. Together they cover storage engine roles, storage infrastructure roles, and data movement platform roles.

## How this module works
OpenHearth lives in its own repo. Hoard contributes one package: `engine/wal`, imported as the proxy's durable replication log.

## Before you code
Answer in `notes/08-capstone.md`:
1. During dual-write, the source write succeeds and the destination write fails. Walk through exactly how the system converges.
2. Backfill copies object X just as a user deletes X through the proxy. How do tombstones prevent resurrection?
3. Which guarantees from your WAL crash tests does OpenHearth now depend on?

## Build (in the OpenHearth repo)
- S3 passthrough for PutObject, GetObject, HeadObject, DeleteObject, ListObjectsV2.
- Dual-write: source synchronous, replication entry appended to `hoard/engine/wal` and synced before acking, async replay to destination with backoff.
- Tombstones for deletes; backfill workers that list the source and copy missing or mismatched objects.
- A verifier that diffs both backends.
- Docker Compose with SeaweedFS and Garage.

## Graded deliverable
Chaos run: mixed PUT/DELETE load against at least 100k objects, kill the destination for 5 minutes mid-run, restart, drain. The verifier reports 0 missing, 0 extra, 0 mismatched.

## Done when
- The chaos run passes.
- OpenHearth's README leads with that result.
- `git tag module-08` in Hoard, `v0.1.0` in OpenHearth.
