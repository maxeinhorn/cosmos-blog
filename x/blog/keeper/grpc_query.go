package keeper

import (
	"github.com/maxeinhorn/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
