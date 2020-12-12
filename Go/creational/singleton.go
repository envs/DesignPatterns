package singleton

// Singleton interface
type Singleton interface {
	AddOne() int
}

type singleton struct {
	count int
}

var instance *singleton

// GetInstance Function
func GetInstance() Singleton {
	return nil
}

func (s *singleton) AddOne() int {
	return 0
}
