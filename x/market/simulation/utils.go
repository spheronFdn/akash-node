package simulation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	types "github.com/spheronFdn/akash-api-fork/go/node/market/v1beta4"

	ptypes "github.com/spheronFdn/akash-api-fork/go/node/provider/v1beta3"

	keepers "github.com/spheronFdn/akash-node/x/market/handler"
)

func getOrdersWithState(ctx sdk.Context, ks keepers.Keepers, state types.Order_State) []types.Order {
	var orders []types.Order

	ks.Market.WithOrders(ctx, func(order types.Order) bool {
		if order.State == state {
			orders = append(orders, order)
		}

		return false
	})

	return orders
}

func getProviders(ctx sdk.Context, ks keepers.Keepers) []ptypes.Provider {
	var providers []ptypes.Provider

	ks.Provider.WithProviders(ctx, func(provider ptypes.Provider) bool {
		providers = append(providers, provider)

		return false
	})

	return providers
}
