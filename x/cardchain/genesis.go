package cardchain

import (
	"fmt"

	"github.com/DecentralCardGame/Cardchain/x/cardchain/keeper"
	"github.com/DecentralCardGame/Cardchain/x/cardchain/types"
	"github.com/DecentralCardGame/cardobject/keywords"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	for id := range genState.Users {
		address, _ := sdk.AccAddressFromBech32(genState.Addresses[id])
		k.SetUser(ctx, address, *genState.Users[id])
	}
	for id, match := range genState.Matches {
		k.SetMatch(ctx, uint64(id), *match)
	}
	fmt.Println("reading cards with id:")
	for _, record := range genState.CardRecords {
		lastId := k.GetLastCardSchemeId(ctx)
		currId := lastId + 1

		_, err := keywords.Unmarshal(record.Content)
		if err != nil {
			fmt.Println(currId, ":")
			fmt.Println(err.Error())
			fmt.Println(string(record.Content))
			fmt.Println("-----")
		}

		k.SetLastCardSchemeId(ctx, currId)
		k.SetCard(ctx, currId, *record)
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	// this line is used by starport scaffolding # genesis/module/export
	params := k.GetParams(ctx)
	records := k.GetAllCards(ctx)
	matches := k.GetAllMatches(ctx)
	users, addresses := k.GetAllUsers(ctx)
	return &types.GenesisState{
		Params:           params,
		CardRecords:      records,
		Users:            users,
		Matches:					matches,
		Addresses:        addresses,
		LastCardSchemeId: k.GetLastCardSchemeId(ctx),
	}
}
