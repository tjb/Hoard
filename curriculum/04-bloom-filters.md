# Module 04: Bloom filters

## Goal
Add a Bloom filter to every SSTable so point lookups skip files that cannot contain the key.

## Why it matters
Read amplification is the price of an LSM tree. Bloom filters are the main discount. Tuning bits-per-key is a real production knob in RocksDB, Cassandra, and ScyllaDB.

## Before you code
Answer in `notes/04-bloom-filters.md`:
1. Why can a Bloom filter return false positives but never false negatives?
2. Derive the expected false positive rate given bits-per-key and number of hash functions. What is the optimal number of hash functions?
3. Why don't Bloom filters help range scans?

## Build
- `engine/bloom/`: build from a set of keys, serialize, deserialize, `MayContain(key)`.
- Store the filter in each SSTable (a new block referenced from the footer).
- Point lookups consult the filter before reading any data block.

## Constraints
- Standard library only. Derive k hash values from one or two hashes (double hashing).

## Graded deliverable
1. Measured false positive rate vs bits-per-key (4, 8, 10, 16) on 1M keys, plotted or tabled next to the theoretical curve. They should match within a few percent.
2. Re-run module 03's read-amplification measurement for lookups of keys that do not exist. Before and after numbers.

## Done when
- Measured FPR matches theory.
- Missing-key lookups touch close to zero data blocks.
- `git tag module-04`

## Reading
- Bloom, "Space/Time Trade-offs in Hash Coding with Allowable Errors" (1970)
- Kirsch and Mitzenmacher, "Less Hashing, Same Performance" (2006)
