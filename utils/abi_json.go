package utils

import (
	"encoding/json"
	"fmt"

	"github.com/joshklop/cairo-abigen"
)

func ABIFromJSON(data []byte) (*abi.ABI, error) {
	abiEntries := make([]*ABIEntry, 0)
	if err := json.Unmarshal(data, &abiEntries); err != nil {
		return nil, err
	}

	theABI := &abi.ABI{
		Functions:    make([]*abi.Function, 0),
		Constructors: make([]*abi.Constructor, 0),
		L1Handlers:   make([]*abi.L1Handler, 0),
		Events:       make([]*abi.Event, 0),
		Structs:      make([]*abi.Struct, 0),
		Enums:        make([]*abi.Enum, 0),
		Impls:        make([]*abi.Impl, 0),
	}
	for _, entry := range abiEntries {
		switch entry.Type {
		case TypeFunction:
			theABI.Functions = append(theABI.Functions, adaptFunction(entry))
		case TypeConstructor:
			theABI.Constructors = append(theABI.Constructors, adaptConstructor(entry))
		case TypeL1Handler:
			theABI.L1Handlers = append(theABI.L1Handlers, adaptL1Handler(entry))
		case TypeEvent:
			theABI.Events = append(theABI.Events, adaptEvent(entry))
		case TypeStruct:
			theABI.Structs = append(theABI.Structs, adaptStruct(entry))
		case TypeEnum:
			theABI.Enums = append(theABI.Enums, adaptEnum(entry))
		case TypeInterface:
			theABI.Interfaces = append(theABI.Interfaces, adaptInterface(entry))
		case TypeImpl:
			theABI.Impls = append(theABI.Impls, adaptImpl(entry))
		default:
			panic(fmt.Sprintf("unknown abi entry type: %d", entry.Type))
		}

	}

	return theABI, nil
}

func adaptImpl(impl *ABIEntry) *abi.Impl {
	return &abi.Impl{
		Name: impl.Name,
		InterfaceName: impl.InterfaceName,
	}	
}

func adaptInterface(iface *ABIEntry) *abi.Interface {
	items := make([]*abi.Function, 0)
	for _, item := range iface.Items {
		items = append(items, adaptFunction(item))
	}
	return &abi.Interface{
		Name: iface.Name,
		Items: items,
	}
}

func adaptEnum(enum *ABIEntry) *abi.Enum {
	variants := make([]*abi.Variant, 0)
	for _, variant := range enum.Variants {
		variants = append(variants, adaptVariant(variant))
	}
	return &abi.Enum{
		Name: enum.Name,
		Variants: variants,
	}
}

func adaptStruct(theStruct *ABIEntry) *abi.Struct {
	members := make([]*abi.Member, 0)
	for _, member := range theStruct.Members {
		members = append(members, adaptMember(member))
	}
	return &abi.Struct{
		Name: theStruct.Name,
		Members: members,
	}
}

func adaptEvent(entry *ABIEntry) *abi.Event {
	event := &abi.Event{
		Name: entry.Name,
	}
	switch entry.Kind {
	case KindStruct:
		event.Members = make([]*abi.Member, 0)
		for _, member := range entry.Members {
			event.Members = append(event.Members, adaptMember(member))
		}
	case KindEnum:
		event.Variants = make([]*abi.Variant, 0)
		for _, member := range entry.Variants {
			event.Variants = append(event.Variants, adaptVariant(member))
		}
	default:
		panic(fmt.Sprintf("unknown event kind: %d", entry.Kind))
	}
	return event
}

func adaptVariant(m *Member) *abi.Variant {
	return &abi.Variant{
		Name: m.Name,
		Type: m.Type,
	}
}

func adaptMember(m *Member) *abi.Member {
	return &abi.Member{
		Name: m.Name,
		Type: m.Type,
	}
}

func adaptL1Handler(entry *ABIEntry) *abi.L1Handler {
	return &abi.L1Handler{
		Name:            entry.Name,
		Inputs:          adaptInputs(entry.Inputs),
		Outputs:         adaptOutputs(entry.Outputs),
		StateMutability: adaptStateMutability(entry.StateMutability),
	}
}

func adaptConstructor(entry *ABIEntry) *abi.Constructor {
	return &abi.Constructor{
		Inputs: adaptInputs(entry.Inputs),
	}
}

func adaptFunction(entry *ABIEntry) *abi.Function {
	return &abi.Function{
		Name:            entry.Name,
		Inputs:          adaptInputs(entry.Inputs),
		Outputs:         adaptOutputs(entry.Outputs),
		StateMutability: adaptStateMutability(entry.StateMutability),
	}
}

func adaptInputs(inputs []*Input) []*abi.Input {
	newInputs := make([]*abi.Input, 0)
	for _, input := range inputs {
		newInputs = append(newInputs, &abi.Input{
			Name: input.Name,
			Type: input.Type,
		})
	}
	return newInputs
}

func adaptOutputs(outputs []*Output) []*abi.Output {
	newOutputs := make([]*abi.Output, 0)
	for _, output := range outputs {
		newOutputs = append(newOutputs, &abi.Output{
			Type: output.Type,
		})
	}
	return newOutputs
}

func adaptStateMutability(sm StateMutability) abi.StateMutability {
	switch sm {
	case View:
		return abi.View
	case External:
		return abi.External
	default:
		panic(fmt.Sprintf("unknown state mutability: %d", sm))
	}
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
