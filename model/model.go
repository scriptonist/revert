package model

// Person --
type Person struct {
	ID          int64
	First, Last string
}

type db interface {
	SelectPeople() ([]*Person, error)
}

// Model --
type Model struct {
	db
}

// New --
func New(db db) *Model {
	return &Model{
		db: db,
	}
}

// People --
func (m *Model) People() ([]*Person, error) {
	return m.SelectPeople()
}
