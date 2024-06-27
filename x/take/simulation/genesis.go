package simulation

import (
	"github.com/cosmos/cosmos-sdk/types/module"
	types "github.com/spheronFdn/akash-api-fork/go/node/take/v1beta3"
)

// RandomizedGenState generates a random GenesisState for supply
func RandomizedGenState(simState *module.SimulationState) {
	takeGenesis := &types.GenesisState{
		Params: types.DefaultParams(),
	}

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(takeGenesis)
}
