package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Default nameservice params
var (
	DefaultMinPriceTokens = sdk.NewInt(10)
)

var (
	KeyMinPrice = []byte("MinPrice")
)

// ParamKeyTable - Key declaration for parameters
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(minPrice sdk.Coins) Params {
	return Params{
		MinPrice: minPrice,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMinPrice, &p.MinPrice, validateMinPrice),
	}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return NewParams(
		sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, DefaultMinPriceTokens)),
	)
}

// String returns a human readable string representation of the parameters.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// unmarshal the current staking params value from store key or panic
func MustUnmarshalParams(cdc *codec.LegacyAmino, value []byte) Params {
	params, err := UnmarshalParams(cdc, value)
	if err != nil {
		panic(err)
	}

	return params
}

// unmarshal the current staking params value from store key
func UnmarshalParams(cdc *codec.LegacyAmino, value []byte) (params Params, err error) {
	err = cdc.UnmarshalBinaryBare(value, &params)
	if err != nil {
		return
	}

	return
}

// validate a set of params
func (p Params) Validate() error {
	if err := validateMinPrice(p.MinPrice); err != nil {
		return err
	}

	return nil
}

func validateMinPrice(i interface{}) error {
	v, ok := i.(sdk.Coins)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if !v.IsAllPositive() {
		return fmt.Errorf("min price must be positive: %d", v)
	}

	return nil
}
