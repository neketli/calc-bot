package events

type Fetcher interface {
	Fetch(limit int) ([]Event, error)
}

type Processor interface {
	Process([]Event)
}

type Type int

const (
	Unknown Type = iota
	Message
)

type Event struct {
	Type Type
	Text string
}