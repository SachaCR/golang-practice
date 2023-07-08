package todo

type TodoStatus int64

const (
	Todo TodoStatus = iota
	InProgress
	Done
	Cancelled
)

func (s TodoStatus) String() string {
	switch s {
	case Todo:
		return "Todo"

	case InProgress:
		return "InProgress"

	case Done:
		return "Done"

	case Cancelled:
		return "Cancelled"
	}

	return "unknown"
}
