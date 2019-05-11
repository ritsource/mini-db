package src

import "errors"

// Store contains all the data
// var Datamap map[string]interface{}
type Store struct {
	Map map[string]interface{}
}

// Get reads data from the Store-Map
func (s *Store) Get(key string) (interface{}, error) {
	v := s.Map[key]
	if v == nil {
		return nil, errors.New("not found")
	}
	return v, nil
}

// Set adds data on the Store-Map and invokes a function
// that saves data in the persistent storage, but only if option provided
func (s *Store) Set(key string, val interface{}) error {
	s.Map[key] = val
	return nil
}

// Delete deletes data from Store-Map and also Cache-Map
func (s *Store) Delete(key string) error {
	delete(s.Map, key)
	return nil
}
