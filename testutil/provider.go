package testutil

import (
	"testing"

	ptypes "github.com/spheronFdn/akash-api-fork/go/node/provider/v1beta3"
)

func Provider(t testing.TB) ptypes.Provider {
	t.Helper()

	return ptypes.Provider{
		Owner:      AccAddress(t).String(),
		HostURI:    Hostname(t),
		Attributes: Attributes(t),
		Info: ptypes.ProviderInfo{
			EMail:   "test@example.com",
			Website: ProviderHostname(t),
		},
	}
}
