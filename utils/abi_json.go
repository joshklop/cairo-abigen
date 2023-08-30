package utils

import (
	"encoding/json"
	"fmt"
)

func ABIFromJSON(data []byte) ([]*ABIEntry, error) {
	abi := make([]*ABIEntry, 0)
	if err := json.Unmarshal(data, &abi); err != nil {
		return nil, err
	}
	return abi, nil
}

type ABIEntry struct {
	Type            Type            `json:"type"`
	Name            string          `json:"name"`
	InterfaceName   string          `json:"interface_name"`
	Inputs          []*Input        `json:"inputs"`
	Outputs         []*Output       `json:"outputs"`
	StateMutability StateMutability `json:"state_mutability"`
	Kind            EventKind       `json:"kind"`
	Members         []*Member       `json:"members"`
	Variants        []*Member       `json:"variants"`
	Items           []*ABIEntry     `json:"items"` // These will all be TypeFunctions.
}

type Member struct {
	Name string
	Type string
	Kind EventFieldKind `json:"kind,omitempty"`
}

type EventFieldKind uint8

const (
	Key EventFieldKind = iota + 1
	Data
	Nested
)

type Input struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type Output struct {
	Type string `json:"type"`
}

type EventKind uint8

const (
	KindStruct EventKind = iota + 1
	KindEnum
)

func (ek *EventKind) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"enum"`:
		*ek = KindEnum
	case `"struct"`:
		*ek = KindStruct
	default:
		return fmt.Errorf("unknown event kind: %s", string(data))
	}
	return nil
}

type Type uint8

const (
	TypeFunction Type = iota + 1
	TypeConstructor
	TypeL1Handler
	TypeEvent
	TypeStruct
	TypeEnum
	TypeInterface
	TypeImpl
)

func (et *Type) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"function"`:
		*et = TypeFunction
	case `"constructor"`:
		*et = TypeConstructor
	case `"l1_handler"`:
		*et = TypeL1Handler
	case `"event"`:
		*et = TypeEvent
	case `"struct"`:
		*et = TypeStruct
	case `"enum"`:
		*et = TypeEnum
	case `"interface"`:
		*et = TypeInterface
	case `"impl"`:
		*et = TypeImpl
	default:
		return fmt.Errorf("unknown entry type: %s", string(data))
	}
	return nil
}

type StateMutability uint8

const (
	View StateMutability = iota + 1
	External
)

func (sm *StateMutability) UnmarshalJSON(data []byte) error {
	switch string(data) {
	case `"view"`:
		*sm = View
	case `"external"`:
		*sm = External
	default:
		return fmt.Errorf("unknown state mutability: %s", string(data))
	}
	return nil
}
