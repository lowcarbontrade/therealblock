package types

import sdk "github.com/cosmos/cosmos-sdk/types"

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1
const devAddr string = "realblock18njy4mmunlt2dkw283l0x754d4yv6tgdv6nsxc"

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # genesis/types/default
		Params:        DefaultParams(),
		AdminAccounts: []Account{{Address: devAddr}},
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	if err := validateAcc(gs); err != nil {
		return err
	}
	return gs.Params.Validate()
}

func validateAcc(gs GenesisState) error {
	for _, account := range gs.AdminAccounts {
		var accMap = make(map[string]bool, len(gs.AdminAccounts))
		acc, err := sdk.AccAddressFromBech32(account.Address)
		if err != nil {
			return err
		}
		if accMap[acc.String()] {
			return ErrAccountExists
		} else {
			accMap[acc.String()] = true
		}
	}
	return gs.Params.Validate()
}
