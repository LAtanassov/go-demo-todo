package todo

// contains todo service

// New return new todo.Service
func New(i Inserter) *Service {
	return &Service{
		i: i
	}
}

// Service contains all methods for the todo functionality
type Service struct {
	i Inserter
}

// Item ...
type Item struct {
	Description string
}

//
type Inserter interface {
	Insert(i Item) error
}

// Create ...
func (s *Service) Create(t Item) error {
	return s.i.Insert(i)
}

// Read ...
func (s *Service) Read() ([]Item, error) {
	return nil, nil
}

// Update ...
func (s *Service) Update(t Item) error {
	return nil
}

// Delete ...
func (s *Service) Delete(t Item) error {
	return nil
}
