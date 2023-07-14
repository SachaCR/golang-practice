package todo

type TodoRepository interface {
	Save(todoTask TodoTask) error
	GetById(id string) (TodoTask, error)
	GetAll() ([]TodoTask, error)
}
