package provider

import (
	types "github.com/spheronFdn/akash-api-fork/go/node/provider/v1beta3"

	"github.com/spheronFdn/akash-node/x/provider/keeper"
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
