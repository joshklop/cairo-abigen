package abi

type ABI struct {
	Functions    []*Function
	Constructors []*Constructor
	L1Handlers   []*L1Handler
	Events       []*Event
	Structs      []*Struct
	Enums        []*Enum
	Impls        []*Impl
	Interfaces   []*Interface
}

type Function struct {
	Name            string
	Inputs          []*Input
	Outputs         []*Output
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
	Members  []*Member
	Variants []*Variant
}

type Variant struct {
	Name string
	Type string
}

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
	Variants []*Variant
}

type Interface struct {
	Name  string
	Items []*Function
}

type Impl struct {
	Name          string
	InterfaceName string
}
