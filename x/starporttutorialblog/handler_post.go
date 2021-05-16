package stellartutorialblog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/darthlukan/stellar-tutorial-blog/x/stellartutorialblog/keeper"
	"github.com/darthlukan/stellar-tutorial-blog/x/stellartutorialblog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
