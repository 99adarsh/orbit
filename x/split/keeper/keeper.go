package keeper

import (
	"fmt"

	"cosmossdk.io/core/store"
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/log"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/99adarsh/orbit/x/split/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	ibckeeper "github.com/cosmos/ibc-go/v8/modules/core/keeper"
	ibctmtypes "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
)

type (
	Keeper struct {
		cdc          codec.BinaryCodec
		storeService store.KVStoreService
		logger       log.Logger

		// the address capable of executing a MsgUpdateParams message. Typically, this
		// should be the x/gov module account.
		authority string

		bankKeeper          types.BankKeeper
		accountKeeper       types.AccountKeeper
		icaControllerKeeper types.ICAControllerKeeper
		IBCKeeper           ibckeeper.Keeper
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeService store.KVStoreService,
	logger log.Logger,
	authority string,

	bankKeeper types.BankKeeper,
	IBCKeeper ibckeeper.Keeper,
	accountKeeper types.AccountKeeper,
	icaControllerKeeper types.ICAControllerKeeper,
) Keeper {
	if _, err := sdk.AccAddressFromBech32(authority); err != nil {
		panic(fmt.Sprintf("invalid authority address: %s", authority))
	}

	return Keeper{
		cdc:          cdc,
		storeService: storeService,
		authority:    authority,
		logger:       logger,

		bankKeeper:          bankKeeper,
		IBCKeeper:           IBCKeeper,
		accountKeeper:       accountKeeper,
		icaControllerKeeper: icaControllerKeeper,
	}
}

// GetAuthority returns the module's authority.
func (k Keeper) GetAuthority() string {
	return k.authority
}

// Logger returns a module-specific logger.
func (k Keeper) Logger() log.Logger {
	return k.logger.With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetChainIdFromConnectionId(ctx sdk.Context, connectionID string) (string, error) {
	connection, found := k.IBCKeeper.ConnectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		return "", errorsmod.Wrapf(connectiontypes.ErrConnectionNotFound, "connection %s not found", connectionID)
	}
	clientState, found := k.IBCKeeper.ClientKeeper.GetClientState(ctx, connection.ClientId)
	if !found {
		return "", errorsmod.Wrapf(clienttypes.ErrClientNotFound, "client %s not found", connection.ClientId)
	}
	client, ok := clientState.(*ibctmtypes.ClientState)
	if !ok {
		return "", types.ErrClientStateNotTendermint
	}

	return client.ChainId, nil
}
