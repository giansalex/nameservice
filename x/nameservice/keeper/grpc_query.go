package keeper

import (
	"github.com/giansalex/nameservice/x/nameservice/types"
)

var _ types.QueryServer = Keeper{}
