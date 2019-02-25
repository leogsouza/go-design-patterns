package singleton

// Singleton provide AddOne counter method
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance get a Singleton instance
func GetInstance() Singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

// AddOne increments the counter
func (s *singleton) AddOne() int {
	s.count++
	return s.count
}
