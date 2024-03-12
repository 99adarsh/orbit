package keeper

import (
	"context"
	"fmt"
	"strconv"

	errorsmod "cosmossdk.io/errors"
	"github.com/99adarsh/orbit/x/split/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
)

func (server msgServer) CreateSplitterPool(ctx context.Context, msg *types.MsgCreateSplitterPool) (*types.MsgCreateSplitterPoolResponse, error) {
	// call keeper's CreatePool method which will return id of the pool or error
	poolId, err := server.keeper.CreatePool(ctx, msg)

	if err != nil {
		return &types.MsgCreateSplitterPoolResponse{}, err
	}
	return &types.MsgCreateSplitterPoolResponse{PoolId: poolId}, nil
}

func (k Keeper) CreatePool(ctx context.Context, msg *types.MsgCreateSplitterPool) (uint64, error) {

	// err := ValidateBasic(*msg)
	// if err != nil {
	// 	return 0, err
	// }
	sdkCtx := sdk.UnwrapSDKContext(ctx)
	nextPoolId := k.GetNextPoolId(ctx)
	// pool address where this pool's fund will be stored
	poolAddrKey := append([]byte(strconv.FormatUint(nextPoolId, 10)), []byte(strconv.FormatUint(msg.Maturity, 10))...)
	// create a pool address
	poolAddr := address.Module(types.ModuleName, poolAddrKey)
	// create account for this pool adderess
	_ = k.accountKeeper.NewAccount(
		ctx,
		authtypes.NewBaseAccountWithAddress(poolAddr),
	)
	chainId, connErr := k.GetChainIdFromConnectionId(sdkCtx, msg.ConnectionId)
	if connErr != nil {
		return 0, errorsmod.Wrapf(
			types.ErrFailedToCreateSplitPool,
			"Error while fetching chainId from connection",
		)
	}

	// Create Ica account to claim rewards from the ls chains, address will be updated in the acknowledgement
	// Interchain account registration  creates a new channel and create a new port identifier

	// Get ConnectionEnd (for counterparty connection)
	connectionEnd, found := k.IBCKeeper.ConnectionKeeper.GetConnection(sdkCtx, msg.ConnectionId)
	if !found {
		err := fmt.Sprintf("invalid connection id, %s not found", msg.ConnectionId)
		return 0, errorsmod.Wrapf(types.ErrFailedToCreateSplitPool, err)
	}

	counterpartyConnection := connectionEnd.Counterparty
	// delegation account is used to capture reward amount from the liquid staking provider chain
	delegateAccount := chainId + "." + "delegate"
	appVersion := string(icatypes.ModuleCdc.MustMarshalJSON(&icatypes.Metadata{
		Version:                icatypes.Version,
		ControllerConnectionId: msg.ConnectionId,
		HostConnectionId:       counterpartyConnection.ConnectionId,
		Encoding:               icatypes.EncodingProtobuf,
		TxType:                 icatypes.TxTypeSDKMultiMsg,
	}))

	err := k.icaControllerKeeper.RegisterInterchainAccount(sdkCtx, msg.ConnectionId, delegateAccount, appVersion)
	if err != nil {
		errMsg := fmt.Sprintf("unable to register delegation account, err: %s", err.Error())
		return 0, errorsmod.Wrapf(types.ErrFailedToCreateSplitPool, errMsg)
	}

	pool := types.Pool{
		Id:      nextPoolId,
		ChainId: chainId,
		// IcahostAddress: nil, // will update after acknowledgement
		Maturity:          msg.Maturity,
		IbcDenom:          msg.IbcDenom,
		HostDenom:         msg.HostDenom,
		PoolAddress:       string(poolAddr),
		ConnectionId:      msg.ConnectionId,
		TransferChannelId: msg.TransferChannelId,
	}

	k.SetPool(ctx, pool)

	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSplitterPoolCreated,
			sdk.NewAttribute(types.AttributeKeyPoolId, strconv.Itoa(int(pool.Id))),
			sdk.NewAttribute(types.AttributeKeyChainId, pool.ChainId),
			sdk.NewAttribute(types.AttributeKeyHostDenom, pool.HostDenom),
			sdk.NewAttribute(types.AttributeKeyIbcDenom, pool.IbcDenom),
		),
	)

	return pool.Id, nil
}
