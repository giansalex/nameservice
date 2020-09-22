package types

import "strings"

// QueryResResolve Queries Result Payload for a resolve query
type QueryResResolve struct {
	Value string `json:"value"`
}

// implement fmt.Stringer
func (r QueryResResolve) String() string {
	return r.Value
}

// QueryResNames Queries Result Payload for a names query
type QueryResNames []string

// implement fmt.Stringer
func (n QueryResNames) String() string {
	return strings.Join(n[:], "\n")
}

// QueryResParams Queries Result Payload for parameters query
type QueryResParams struct {
	Params Params `json:"params"`
}

// implement fmt.Stringer
func (r QueryResParams) String() string {
	return r.Params.String()
}
