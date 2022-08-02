package order

type RepositoryI interface {
	Querier
	Exercer
}

type Querier interface {
}

type Exercer interface {
}
type RepositoryMemmory struct {
	//TODO: put database here
}

func NewRepository() *RepositoryMemmory {
	return &RepositoryMemmory{}
}
