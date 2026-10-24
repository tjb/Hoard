# Module 00: Harness

## Goal
Build the tools that will grade every later module: a randomized differential tester and a benchmark runner. Nothing in this repo gets trusted until it survives these.

## Why it matters
Storage bugs hide in rare interleavings, not happy paths. Every storage team runs some version of "replay random operations against a reference and diff the results." You are building your own grader.

## Before you code
Answer these in `notes/00-harness.md`:
1. Why compare against an oracle instead of writing hand-picked test cases?
2. When a random test fails on operation 48,213 of 100,000, how will you find the smallest sequence that still fails?
3. What makes a benchmark lie to you? Name at least three ways.

## Build
- `engine/storetest/`: a reusable conformance suite. Given a constructor `func(t *testing.T) engine.Store`, it runs:
  - basic Get/Put/Delete/Scan semantics
  - a differential test: N random operations (weighted Put, Delete, Get, Scan) applied to the store under test and to `engine/oracle`, failing on the first divergence
  - a fixed seed by default, with a flag or env var to randomize and print the seed on failure
- `bench/`: a benchmark runner that reports ops/sec and p50, p99, p999 latency for: sequential writes, random writes, random reads, range scans of 100 keys. Output a Markdown table you can paste into notes.

## Constraints
- Standard library only.
- The conformance suite must pass against `engine/oracle` with zero changes.

## Graded deliverable
Plant a bug: write a throwaway wrapper around the oracle that silently drops 1 in 10,000 deletes. Your differential test must catch it, and print the seed and the failing operation index.

## Done when
- `go test ./...` passes against the oracle.
- The planted bug is caught in under 10 seconds.
- `notes/00-harness.md` has your answers and the benchmark table for the oracle (your baseline).

## Reading
- *Database Internals* (Petrov), Chapter 1
- Search: "differential testing" and "property-based testing shrinking"
