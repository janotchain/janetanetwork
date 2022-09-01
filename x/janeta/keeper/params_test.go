package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "janeta/testutil/keeper"
	"janeta/x/janeta/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.JanetaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
