package cli

import (
	"encoding/json"
	"errors"
	"time"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	deploymentTypes "github.com/spheronFdn/akash-api-fork/go/node/deployment/v1beta3"
	types "github.com/spheronFdn/akash-api-fork/go/node/escrow/v1beta3"

	marketTypes "github.com/spheronFdn/akash-api-fork/go/node/market/v1beta4"

	aclient "github.com/spheronFdn/akash-node/client"
	netutil "github.com/spheronFdn/akash-node/util/network"
	"github.com/spheronFdn/akash-node/x/deployment/client/cli"
	"github.com/spheronFdn/akash-node/x/escrow/client/util"
)

func GetQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Escrow query commands",
		SuggestionsMinimumDistance: 2,
		RunE:                       sdkclient.ValidateCmd,
	}

	cmd.AddCommand(
		cmdBlocksRemaining(),
	)

	return cmd
}

var errNoLeaseMatches = errors.New("leases for deployment do not exist")

func cmdBlocksRemaining() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "blocks-remaining",
		Short: "Compute the number of blocks remaining for an ecrow account",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()

			cctx, err := sdkclient.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			qq, err := aclient.DiscoverQueryClient(ctx, cctx)
			if err != nil {
				return err
			}

			id, err := cli.DeploymentIDFromFlags(cmd.Flags())
			if err != nil {
				return err
			}

			// Fetch leases matching owner & dseq
			leaseRequest := marketTypes.QueryLeasesRequest{
				Filters: marketTypes.LeaseFilters{
					Owner:    id.Owner,
					DSeq:     id.DSeq,
					GSeq:     0,
					OSeq:     0,
					Provider: "",
					State:    "active",
				},
				Pagination: nil,
			}

			leasesResponse, err := qq.Leases(cmd.Context(), &leaseRequest)
			if err != nil {
				return err
			}

			if len(leasesResponse.Leases) == 0 {
				return errNoLeaseMatches
			}

			// Fetch the balance of the escrow account
			totalLeaseAmount := leasesResponse.TotalPriceAmount()
			blockchainHeight, err := cli.CurrentBlockHeight(qq.ClientContext())
			if err != nil {
				return err
			}

			res, err := qq.Deployment(cmd.Context(), &deploymentTypes.QueryDeploymentRequest{
				ID: deploymentTypes.DeploymentID{Owner: id.Owner, DSeq: id.DSeq},
			})
			if err != nil {
				return err
			}

			balanceRemain := util.LeaseCalcBalanceRemain(res.EscrowAccount.TotalBalance().Amount,
				int64(blockchainHeight),
				res.EscrowAccount.SettledAt,
				totalLeaseAmount)

			blocksRemain := util.LeaseCalcBlocksRemain(balanceRemain, totalLeaseAmount)

			output := struct {
				BalanceRemain       float64       `json:"balance_remaining" yaml:"balance_remaining"`
				BlocksRemain        int64         `json:"blocks_remaining" yaml:"blocks_remaining"`
				EstimatedTimeRemain time.Duration `json:"estimated_time_remaining" yaml:"estimated_time_remaining"`
			}{
				BalanceRemain:       balanceRemain,
				BlocksRemain:        blocksRemain,
				EstimatedTimeRemain: netutil.AverageBlockTime * time.Duration(blocksRemain),
			}

			outputType, err := cmd.Flags().GetString("output")
			if err != nil {
				return err
			}

			var data []byte
			if outputType == "json" {
				data, err = json.MarshalIndent(output, " ", "\t")
			} else {
				data, err = yaml.Marshal(output)
			}

			if err != nil {
				return err
			}

			return qq.ClientContext().PrintBytes(data)

		},
	}

	flags.AddQueryFlagsToCmd(cmd)
	cli.AddDeploymentIDFlags(cmd.Flags())
	cli.MarkReqDeploymentIDFlags(cmd)
	return cmd
}
