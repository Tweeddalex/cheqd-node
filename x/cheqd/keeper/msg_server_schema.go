package keeper

import (
	"context"
	"github.com/cheqd/cheqd-node/x/cheqd/types"
)

func (k msgServer) CreateSchema(goCtx context.Context, msg *types.MsgWriteRequest) (*types.MsgCreateSchemaResponse, error) {
	return nil, types.ErrNotImplemented.Wrap("Schema")

	/*
		ctx := sdk.UnwrapSDKContext(goCtx)
		prefix := types.DidPrefix + ":" + types.DidMethod + ":" + ctx.ChainID() + ":"

		var schemaMsg types.MsgCreateSchema
		err := k.cdc.Unmarshal(msg.Data.Value, &schemaMsg)
		if err != nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized %s message type: %T", types.ModuleName, msg)
		}

		if err := schemaMsg.ValidateBasic(prefix); err != nil {
			return nil, err
		}

		if err := k.VerifySignature(&ctx, msg, schemaMsg.GetSigners()); err != nil {
			return nil, err
		}

		// Checks that the did doesn't exist
		if err := k.EnsureDidIsNotUsed(ctx, schemaMsg.GetDid()); err != nil {
			return nil, err
		}

		k.AppendSchema(
			ctx,
			schemaMsg.Id,
			schemaMsg.Type,
			schemaMsg.Name,
			schemaMsg.Version,
			schemaMsg.AttrNames,
			schemaMsg.Controller,
		)

		return &types.MsgCreateSchemaResponse{
			Id: schemaMsg.Id,
		}, nil*/
}
