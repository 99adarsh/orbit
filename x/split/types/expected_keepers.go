package types

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	icaControllerKeeper "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/controller/types"
)

// AccountKeeper defines the expected interface for the Account module.
type AccountKeeper interface {
	GetAccount(context.Context, sdk.AccAddress) sdk.AccountI // only used for simulation
	// Methods imported from account should be defined here
	NewAccount(context.Context, sdk.AccountI) sdk.AccountI
}

// BankKeeper defines the expected interface for the Bank module.
type BankKeeper interface {
	SpendableCoins(context.Context, sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}

// ParamSubspace defines the expected Subspace interface for parameters.
type ParamSubspace interface {
	Get(context.Context, []byte, interface{})
	Set(context.Context, []byte, interface{})
}

// ICAControllerKeeper defines the expected ICA controller keeper
type ICAControllerKeeper interface {
	RegisterInterchainAccount(ctx sdk.Context, connectionID, owner, version string) error
	GetInterchainAccountAddress(ctx sdk.Context, connectionID, portID string) (string, bool)
}

type IBCKeeper interface {
	RegisterInterchainAccount(ctx context.Context, msg *icaControllerKeeper.MsgRegisterInterchainAccount) (*icaControllerKeeper.MsgRegisterInterchainAccountResponse, error)
}
