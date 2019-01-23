package todo

// compiled 2017, does not contains any interface

// NewRepository ...
func NewRepository() *Repository {
	return &Repository{}
}

type Repository struct {
}

// Insert ...
func (r *Repository) Insert(t Item) error {
	return nil
}
