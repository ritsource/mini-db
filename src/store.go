package src

import (
	"errors"
	"fmt"
	"os/user"
	"path"
	"time"
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
	Delay time.Duration
	// Map stores all the key value pairs in-memory
	Map map[string]interface{}
}

// Init initializes store, reads data from output file,
// and concurrently runs a function that saves data to output in specific time delay
func (s *Store) Init(persist bool, delay int, output string) {
	// If persist == true, If data needs to be saved in fs
	if persist {
		// Setup output = default output if not provided
		if output == "" {
			usr, err := user.Current() // Current user in the OS
			if err != nil {
				panic(err)
			}
			output = path.Join(usr.HomeDir, ".minidb", "data.out") // Output path, default
		}

		// Initializing store instance
		*s = Store{Persist: persist}  // defining store instance
		rmap, err := ReadFile(output) // Reading data from output file to popultate previously saved data
		if err != nil {
			fmt.Println("Unable to read data from file:", err)
		}
		s.Map = rmap // Populating store.Map

		// Backing up data in FS in provided time delay
		go (func(st *Store) {
			for {
				time.Sleep(time.Second * time.Duration(delay))

				err := WriteFile(output, s.Map) // Writing data to FS
				if err != nil {
					panic(err)
				}
			}
		})(s)
	} else {
		// If persist == false,
		*s = Store{Persist: false}
		s.Map = make(map[string]interface{})
	}
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

// MGet reads multiple entries from the Store-Map
func (s *Store) MGet(keys []string) map[string]interface{} {
	cmap := make(map[string]interface{})

	for _, key := range keys {
		cmap[key] = s.Map[key]
	}

	return cmap
}

// MSet adds multiple kkey value pair on the Store-Map
func (s *Store) MSet(pairs map[string]interface{}) error {
	var err error

	for key, val := range pairs {
		s.Map[key] = val
	}

	// If data persistence is true and snapshot false
	// Anotherwords, have to write data to the disk
	if s.Persist && !s.Snapshot {
		// Write data to Disk
	}

	return err
}

// Flush deletes all key value pairs
func (s *Store) Flush() error {
	var err error
	s.Map = make(map[string]interface{})

	// Again, if data persistence is true and snapshot false
	if s.Persist && !s.Snapshot {
		// Write data to Disk
	}

	return err
}
