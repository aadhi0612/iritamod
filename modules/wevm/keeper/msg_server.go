package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ethereum/go-ethereum/common"

	"github.com/bianjieai/iritamod/modules/wevm/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) AddToContractDenyList(goCtx context.Context, list *types.MsgAddToContractDenyList) (*types.MsgAddToContractDenyListResponse, error) {
	if !common.IsHexAddress(list.ContractAddress) {
		return &types.MsgAddToContractDenyListResponse{},
			errors.Wrapf(types.ErrInvalidContractAddress, "contract Address %s is invalid", list.ContractAddress)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.AddToContractDenyList(ctx, list.ContractAddress)
	if err != nil {
		return &types.MsgAddToContractDenyListResponse{}, err
	}

	return &types.MsgAddToContractDenyListResponse{}, nil
}

func (k msgServer) RemoveFromContractDenyList(goCtx context.Context, list *types.MsgRemoveFromContractDenyList) (*types.MsgRemoveFromContractDenyListResponse, error) {
	if !common.IsHexAddress(list.ContractAddress) {
		return &types.MsgRemoveFromContractDenyListResponse{},
			errors.Wrapf(types.ErrInvalidContractAddress, "contract Address %s is invalid", list.ContractAddress)
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.RemoveFromContractDenyList(ctx, list.ContractAddress)
	if err != nil {
		return &types.MsgRemoveFromContractDenyListResponse{}, err
	}
	return &types.MsgRemoveFromContractDenyListResponse{}, nil
}
