package utils_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joshklop/cairo-abigen"
	"github.com/joshklop/cairo-abigen/utils"
	"github.com/stretchr/testify/require"
)

func TestABIFromJSON(t *testing.T) {
	path := filepath.Join("testdata", "0x03db66c8da0e47340f65030aa4c076c5c4864923b54dcc9ee13d87559c1c0b96.json")
	file, err := os.ReadFile(path)
	require.NoError(t, err)
	got, err := utils.ABIFromJSON(file)
	require.NoError(t, err)

	want := &abi.ABI{
		Functions:    []*abi.Function{},
		Constructors: []*abi.Constructor{},
		L1Handlers:   []*abi.L1Handler{},
		Events: []*abi.Event{
			{
				Name:     "hunt::Franco::Event",
				Variants: make([]*abi.Variant, 0),
			},
		},
		Structs: []*abi.Struct{
			{
				Name: "core::integer::u256",
				Members: []*abi.Member{
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
		},
		Enums: []*abi.Enum{},
		Impls: []*abi.Impl{
			{
				Name:          "FrancoImpl",
				InterfaceName: "hunt::IFranco",
			},
		},

		Interfaces: []*abi.Interface{
			{
				Name: "hunt::IFranco",
				Items: []*abi.Function{
					{
						Name: "add_count",
						Inputs: []*abi.Input{
							{
								Name: "amount",
								Type: "core::integer::u256",
							},
						},
						Outputs:         make([]*abi.Output, 0),
						StateMutability: abi.External,
					},
					{
						Name:   "get_count",
						Inputs: make([]*abi.Input, 0),
						Outputs: []*abi.Output{
							{
								Type: "core::integer::u256",
							},
						},
						StateMutability: abi.View,
					},
				},
			},
		},
	}
	require.Equal(t, want, got)
}
