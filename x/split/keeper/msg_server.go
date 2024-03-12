package keeper

import (
	"github.com/99adarsh/orbit/x/split/types"
)

type msgServer struct {
	keeper *Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewSplitterMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{keeper: &keeper}
}

var _ types.MsgServer = msgServer{}
