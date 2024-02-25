package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "github.com/99adarsh/orbit/testutil/keeper"
	"github.com/99adarsh/orbit/x/split/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.SplitKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
