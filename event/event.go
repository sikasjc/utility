package event

type Event interface {
	Type() uint
	Params() interface{}
}
