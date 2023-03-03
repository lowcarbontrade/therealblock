package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	"github.com/cosmos/cosmos-sdk/types/query"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/realblocknetwork/therealblock/x/therealblock/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListAdmins(goCtx context.Context, req *types.QueryListAdminsRequest) (*types.QueryListAdminsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(goCtx)
	var admins []types.Account
	adminStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GenAccountKey))
	pageRes, err := query.Paginate(adminStore, req.Pagination, func(key []byte, value []byte) error {
		var admin types.Account
		if err := k.cdc.Unmarshal(value, &admin); err != nil {
			return err
		}
		admins = append(admins, admin)
		return nil
	})
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &types.QueryListAdminsResponse{
		Accounts:   admins,
		Pagination: pageRes,
	}, nil
}
