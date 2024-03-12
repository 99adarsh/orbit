package keeper

import (
	"context"
	"fmt"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/99adarsh/orbit/x/split/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
)

// It updates the Pool's ica-accounnt which was created durign pool creation
func (server msgServer) UpdateSplitterPoolIcaAccount(ctx context.Context, msg *types.MsgUpdateplitterPoolAccount) (*types.MsgUpdateplitterPoolAccountResponse, error) {

	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Get ConnectionEnd (for counterparty connection)
	// connectionEnd, found := server.keeper.IBCKeeper.ConnectionKeeper.GetConnection(goCtx, msg.ConnectionId)
	// if !found {
	// 	return nil, errorsmod.Wrapf(connectiontypes.ErrConnectionNotFound, "connection %s not found", msg.ConnectionId)
	// }

	// only allow restoring an account if it already exists
	portID, err := icatypes.NewControllerPortID(msg.AccountOwner)
	if err != nil {
		return nil, err
	}
	address, found := server.keeper.icaControllerKeeper.GetInterchainAccountAddress(sdkCtx, msg.ConnectionId, portID)

	if !found {
		return nil, errorsmod.Wrapf(types.ErrFailedToUpdatePoolInterChainAccountAddress, "Could not find address for the interhcain account for this owner")
	}

	pool, found := server.keeper.GetPool(ctx, msg.PoolId)
	if !found {
		errMsg := fmt.Sprintf("Could not find Splitter pool for the PoolId %s", strconv.Itoa(int(msg.PoolId)))
		return nil, errorsmod.Wrapf(types.ErrFailedToFindSplitterPool, errMsg)
	}

	pool.IcaHostAddress = address

	server.keeper.SetPool(ctx, pool)

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeUpdatePoolIcaAddress,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.Itoa(int(msg.PoolId))),
			sdk.NewAttribute(types.AttributeKeyPoolIcaAddress, address),
		),
	)

	return &types.MsgUpdateplitterPoolAccountResponse{}, nil
}
