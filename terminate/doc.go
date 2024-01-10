// Package terminate implements routines for previewing the costs of terminating
// the sectors on the miners owned by an agent
//
// It is possible to preview the cost of terminating a single sector, or all
// the sectors on a miner.“
//
// # Calculation methods
//
// There are two methods that can be used to preview the cost of terminating sector(s):
//
//   - A slower "on-chain" method can be used that utilizes the `StateCompute`
//     JSON-RPC API to calculate the cost by simulating the execution of the
//     termination messages on a Filecoin node using the system actors.
//
//   - A faster "off-chain" method can be used that retrieves the sector data
//     and feeds it into calculations that match what the system actors would do.
//
// # Sampling for approximate results
//
// It is possible to preview the cost of terminating all the sectors by
// examining every sector. However, this can be slow, particularly for very
// large miners. For situations where an speed vs. accuracy trade-off is
// acceptable, it is also possible to sample a subset of sectors and extrapolate
// the results for an approximate result.
//
// For example:
//
//	f096133 On-chain: 52385/52385 0m12.557s Liquidation Value: 7822.794 @3450134 Recovery Pct: 76.848%
//	f096133 Offchain: 840/52385 0m1.501s Liquidation Value: 7820.798 @3450134 Recovery Pct: 76.829%
//
// Using the "off-chain" method, only 840 sectors were sampled, but the query
// completed in 1.5 seconds instead of 12.5 seconds, but the approximated
// result was very close to the "on-chain" method which simulated the termination
// of every sector using the system actors.
//
// # Previewing terminating a single sector on a miner
//
// The [PreviewTerminateSector] function can be used to preview the termination
// costs for a single sector on a miner. This function is flexible, and can do
// either "on-chain" or "off-chain" calculations.
//
// # Previewing terminating all sectors on a miner
//
// The [PreviewTerminateSectors] function can be used preview the termination
// costs for all the sectors on a miner. This function is flexible, and can
// do either "on-chain" or "off-chain" calculations. It can terminate all
// sectors or sample a subset of sectors for an approximate result.
//
// The [PreviewTerminateSectorsQuick] function is preferred for quick
// results using opinionated settings, sampling, "off-chain" calcuation, and
// a simpler API (no channels for streaming progress)
//
// # Previewing terminating all sectors on all miners for an agent
//
// The [PreviewAgentTermination] function will preview terminating all the
// sectors on all the miners for an agent (using sampling and "off-chain"
// calculation) and will return aggregated stats ([AgentCollateralStats]).
package terminate
