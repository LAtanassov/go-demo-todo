package todo

// contains todo service

// NewService return new todo.Service
func NewService() *Service {
	return &Service{}
}

// Service contains all methods for the todo functionality
type Service struct {
}

// Item ...
type Item struct {
	Description string
}

// Inserter ...
type Inserter interface {
	Insert(i Item) error
}

// Create ...
func (s *Service) Create(t Item) error {
	return nil
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
