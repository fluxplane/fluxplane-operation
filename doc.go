// Package operation defines the pure model for callable capability contracts.
//
// An Operation is the smallest executable domain primitive: it receives one
// input value and returns one result. This package describes operation identity,
// input/output contracts, semantic effect claims, result/error shapes, and inert
// event payloads.
//
// Operation intentionally does not define middleware, validators, persistence,
// logging, approval prompts, retries, timeouts, rendering, or live event
// emission. Those are runtime or orchestration concerns.
package operation
