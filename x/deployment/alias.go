package deployment

import (
	types "github.com/spheronFdn/akash-api-fork/go/node/deployment/v1beta3"

	"github.com/spheronFdn/akash-node/x/deployment/keeper"
)

const (
	// StoreKey represents storekey of deployment module
	StoreKey = types.StoreKey
	// ModuleName represents current module name
	ModuleName = types.ModuleName
)

type (
	// Keeper defines keeper of deployment module
	Keeper = keeper.Keeper
)

var (
	// NewKeeper creates new keeper instance of deployment module
	NewKeeper = keeper.NewKeeper
)
