package abi

type ABI struct {
	Functions    []*Function
	Constructors []*Constructor
	L1Handlers   []*L1Handler
	Events       []*Event
	Structs      []*Struct
	Enums        []*Enum
	Impls        []*Impl
}

type Function struct {
	Name            string
	Inputs          []*Input
	Output          []*Output
	StateMutability StateMutability
}

type StateMutability uint8

const (
	View StateMutability = iota + 1
	External
)

type Input struct {
	Name string
	Type string
}

type Output struct {
	Type string
}

type Constructor struct {
	Inputs []*Input
}

type L1Handler struct {
	Name            string
	Inputs          []*Input
	Outputs         []*Output
	StateMutability StateMutability
}

type Event struct {
	Name     string
	Members  []*EventField // STRUCT_EVENT
	Variants []*EventField // ENUM_EVENT
}

type EventField struct {
	Name string
	Type string
	Kind EventFieldKind
}

type EventFieldKind uint8

const (
	Key EventFieldKind = iota + 1
	Data
	Nested
)

type Struct struct {
	Name    string
	Members []*Member
}

type Member struct {
	Name string
	Type string
}

type Enum struct {
	Name     string
	Variants []*Member
}

type Interface struct {
	Name  string
	Items []*Function
}

type Impl struct {
	Name          string
	InterfaceName string
}
