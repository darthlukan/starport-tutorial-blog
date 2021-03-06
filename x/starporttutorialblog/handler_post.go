package starporttutorialblog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/darthlukan/starport-tutorial-blog/x/starporttutorialblog/keeper"
	"github.com/darthlukan/starport-tutorial-blog/x/starporttutorialblog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
