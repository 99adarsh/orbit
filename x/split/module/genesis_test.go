package split_test

import (
	"testing"

	keepertest "github.com/99adarsh/orbit/testutil/keeper"
	"github.com/99adarsh/orbit/testutil/nullify"
	split "github.com/99adarsh/orbit/x/split/module"
	"github.com/99adarsh/orbit/x/split/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SplitKeeper(t)
	split.InitGenesis(ctx, k, genesisState)
	got := split.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
