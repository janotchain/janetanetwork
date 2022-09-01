package janeta_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "janeta/testutil/keeper"
	"janeta/testutil/nullify"
	"janeta/x/janeta"
	"janeta/x/janeta/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.JanetaKeeper(t)
	janeta.InitGenesis(ctx, *k, genesisState)
	got := janeta.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
