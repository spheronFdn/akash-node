package cert

import (
	types "github.com/spheronFdn/akash-api-fork/go/node/cert/v1beta3"

	"github.com/spheronFdn/akash-node/x/cert/keeper"
)

const (
	// StoreKey represents storekey of provider module
	StoreKey = types.StoreKey
	// ModuleName represents current module name
	ModuleName = types.ModuleName
)

type (
	// Keeper defines keeper of provider module
	Keeper = keeper.Keeper
)

var (
	// NewKeeper creates new keeper instance of provider module
	NewKeeper = keeper.NewKeeper
)
