package singleton

// Singleton interface
// type Singleton interface {
// 	AddOne() int
// }

type singleton struct {
	count int
}

var instance *singleton

// GetInstance Function
func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func (s *singleton) GetCount() int {
	return s.count
}
