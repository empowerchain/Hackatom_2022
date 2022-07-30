package plasticcredits_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "plasticcreditledger/testutil/keeper"
	"plasticcreditledger/testutil/nullify"
	"plasticcreditledger/x/plasticcredits"
	"plasticcreditledger/x/plasticcredits/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		IssuerList: []types.Issuer{
			{
				Addr: "0",
			},
			{
				Addr: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PlasticcreditsKeeper(t)
	plasticcredits.InitGenesis(ctx, *k, genesisState)
	got := plasticcredits.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.IssuerList, got.IssuerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
