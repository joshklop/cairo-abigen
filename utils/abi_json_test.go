package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joshklop/cairo-abigen/utils"
	"github.com/stretchr/testify/require"
)

func TestABIFromJSON(t *testing.T) {
	path := filepath.Join("testdata", "0x03db66c8da0e47340f65030aa4c076c5c4864923b54dcc9ee13d87559c1c0b96.json")
	file, err := os.ReadFile(path)
	require.NoError(t, err)
	got, err := utils.ABIFromJSON(file)
	require.NoError(t, err)
	want := []*utils.ABIEntry{
		{
			Type:          utils.TypeImpl,
			Name:          "FrancoImpl",
			InterfaceName: "hunt::IFranco",
		},
		{
			Type: utils.TypeStruct,
			Name: "core::integer::u256",
			Members: []*utils.Member{
				{
					Name: "low",
					Type: "core::integer::u128",
				},
				{
					Name: "high",
					Type: "core::integer::u128",
				},
			},
		},
		{
			Type: utils.TypeInterface,
			Name: "hunt::IFranco",
			Items: []*utils.ABIEntry{
				{
					Type: utils.TypeFunction,
					Name: "add_count",
					Inputs: []*utils.Input{
						{
							Name: "amount",
							Type: "core::integer::u256",
						},
					},
					Outputs:         make([]*utils.Output, 0),
					StateMutability: utils.External,
				},
				{
					Type:   utils.TypeFunction,
					Name:   "get_count",
					Inputs: make([]*utils.Input, 0),
					Outputs: []*utils.Output{
						{
							Type: "core::integer::u256",
						},
					},
					StateMutability: utils.View,
				},
			},
		},
		{
			Type:     utils.TypeEvent,
			Name:     "hunt::Franco::Event",
			Kind:     utils.KindEnum,
			Variants: make([]*utils.Member, 0),
		},
	}
	require.ElementsMatch(t, want, got)
}
