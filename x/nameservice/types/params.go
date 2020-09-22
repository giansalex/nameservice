package types

import (
	"errors"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params"
)

// Default parameter namespace
const (
	DefaultParamspace        = ModuleName
	DefaultMinPrice   uint64 = 100000
)

var (
	// KeyMinPrice Min Price Key
	KeyMinPrice = []byte("MinPrice")
	// KeyBondDenom Coin Denom key
	KeyBondDenom = []byte("BondDenom")
)

// ParamKeyTable for nameservice module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// Params - used for initializing default parameter for nameservice at genesis
type Params struct {
	MinPrice  uint64 `json:"min_price"`
	BondDenom string `json:"bond_denom"`
}

// NewParams creates a new Params object
func NewParams(minPrice uint64, bondDenom string) Params {
	return Params{
		MinPrice:  minPrice,
		BondDenom: bondDenom,
	}
}

// ParamSetPairs - Implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		// TODO: Pair your key with the param
		paramtypes.NewParamSetPair(KeyMinPrice, &p.MinPrice, validateMinPrice),
		paramtypes.NewParamSetPair(KeyBondDenom, &p.BondDenom, validateBondDenom),
	}
}

// DefaultParams defines the parameters for this module
func DefaultParams() Params {
	return NewParams(
		DefaultMinPrice,
		sdk.DefaultBondDenom,
	)
}

// Validate a set of params
func (p Params) Validate() error {
	if err := validateMinPrice(p.MinPrice); err != nil {
		return err
	}

	if err := validateBondDenom(p.BondDenom); err != nil {
		return err
	}

	return nil
}

func validateMinPrice(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v <= 0 {
		return fmt.Errorf("min price must be positive: %d", v)
	}

	return nil
}

func validateBondDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if strings.TrimSpace(v) == "" {
		return errors.New("bond denom cannot be blank")
	}

	if err := sdk.ValidateDenom(v); err != nil {
		return err
	}

	return nil
}
