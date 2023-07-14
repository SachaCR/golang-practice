package repositories

type RepositoryEngine int64

const (
	Postgres RepositoryEngine = iota
	InMemory
)

func (s RepositoryEngine) String() string {
	switch s {
	case Postgres:
		return "Postgres"

	case InMemory:
		return "InMemory"
	}

	return "InMemory"
}

func RepositoryEngineFromString(statusString string) RepositoryEngine {

	switch statusString {
	case "Postgres":
		return Postgres

	case "InMemory":
		return InMemory
	}

	return InMemory
}
