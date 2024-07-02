package server

type Store interface {
	Fetch() string
}

type SpyStore struct {
	response string
}

func (s *SpyStore) Fetch() string {
	return s.response
}
