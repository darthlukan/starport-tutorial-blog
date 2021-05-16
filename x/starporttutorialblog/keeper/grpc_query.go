package keeper

import (
	"github.com/darthlukan/starport-tutorial-blog/x/starporttutorialblog/types"
)

var _ types.QueryServer = Keeper{}
