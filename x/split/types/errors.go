package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/split module sentinel errors
var (
	ErrInvalidSigner                              = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample                                     = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrFailedToCreateSplitPool                    = sdkerrors.Register(ModuleName, 1102, "Could not create splitter pool")
	ErrClientStateNotTendermint                   = sdkerrors.Register(ModuleName, 1103, "Liquid staking Provider chain's client is not a tendermint client")
	ErrFailedToUpdatePoolInterChainAccountAddress = sdkerrors.Register(ModuleName, 1104, "Could not find address for the interchain account register for the pool")
	ErrFailedToFindSplitterPool                   = sdkerrors.Register(ModuleName, 1105, "Could not find splitter pool")
)
