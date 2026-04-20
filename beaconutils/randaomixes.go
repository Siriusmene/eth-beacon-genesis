package beaconutils

import (
	"github.com/ethpandaops/eth-beacon-genesis/beaconconfig"
	"github.com/ethpandaops/go-eth2-client/spec/phase0"
)

func SeedRandomMixes(genesisBlockHash phase0.Hash32, cfg *beaconconfig.Config) []phase0.Root {
	epochsPerHistoricalVector := cfg.GetUintDefault("EPOCHS_PER_HISTORICAL_VECTOR", 65536)
	randomMixes := make([]phase0.Root, epochsPerHistoricalVector)

	for i := range randomMixes {
		randomMixes[i] = phase0.Root(genesisBlockHash)
	}

	return randomMixes
}
