package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"

	"gitlab.bianjie.ai/irita-pro/iritamod/modules/slashing/types"
)

type Keeper struct {
	slashingkeeper.Keeper
	validatorKeeper types.ValidatorKeeper
}

// NewKeeper creates a slashing keeper
func NewKeeper(slashingKeeper slashingkeeper.Keeper, validatorKeeper types.ValidatorKeeper) Keeper {
	return Keeper{
		slashingKeeper,
		validatorKeeper,
	}
}

// HandleValidatorSignature handles a validator signature, must be called once per validator per block.
// Block all subsequent logic if this validator has been removed.
func (k Keeper) HandleValidatorSignature(ctx sdk.Context, addr crypto.Address, power int64, signed bool) {
	logger := k.Logger(ctx)

	// fetch the validator public key
	consAddr := sdk.ConsAddress(addr)
	if _, err := k.GetPubkey(ctx, addr); err != nil {
		logger.Info(fmt.Sprintf("Validator consensus-address %s not found", consAddr))
		return
	}

	// fetch signing info
	if _, found := k.GetValidatorSigningInfo(ctx, consAddr); !found {
		logger.Info(fmt.Sprintf("Expected signing info for validator %s but not found", consAddr))
		return
	}

	k.Keeper.HandleValidatorSignature(ctx, addr, power, signed)
}

// HandleUnjail handles ths unjail msg
func (k Keeper) HandleUnjail(ctx sdk.Context, msg types.MsgUnjailValidator) error {
	validator := k.validatorKeeper.ValidatorByID(ctx, msg.Id)
	if validator == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unknown validator: %s", msg.Id)
	}
	return k.Unjail(ctx, validator.GetOperator())
}
