package todoservice

type TodoNotFoundError struct{}

func (m *TodoNotFoundError) Error() string {
	return "Todo not found"
}

func TodoNotFound() *TodoNotFoundError {
	return &TodoNotFoundError{}
}
