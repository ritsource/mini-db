package src

import (
	"errors"
)

// Store contains all the data
type Store struct {
	// Persist indicates if data should be written in the disk or not
	Persist bool
	// Snapshot indicates how to Persist data
	// if Snapshot == false, then save data while writing (Set) value
	// Snapshot == false makes write a bit slow, but more reliable
	// else save data to teh disk at a time delay
	// Snapshot is only useul if Persist == true
	Snapshot bool
	// Delay is the time delay for each data snapshot
	Delay float64
	// Map stores all the key value pairs in-memory
	Map map[string]interface{}
}

// Get reads data from the Store-Map
func (s *Store) Get(key string) (interface{}, error) {
	val := s.Map[key]
	if val == nil {
		return nil, errors.New("not found")
	}
	return val, nil
}

// Set adds data on the Store-Map and invokes a function
// that saves data in the persistent storage, but only if option provided
func (s *Store) Set(key string, val interface{}) error {
	var err error
	s.Map[key] = val

	// If data persistence is true and snapshot false
	// Anotherwords, have to write data to the disk
	if s.Persist && !s.Snapshot {
		// Write data to Disk
	}

	return err
}

// Delete deletes data from Store-Map and also Cache-Map
func (s *Store) Delete(key string) error {
	var err error
	val := s.Map[key]
	if val == nil {
		return errors.New("not found")
	}

	delete(s.Map, key)

	// Again, if data persistence is true and snapshot false
	if s.Persist && !s.Snapshot {
		// Write data to Disk
	}

	return err
}
