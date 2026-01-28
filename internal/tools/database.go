package tools

import (
	"sync"
)

type Database struct {
	mutex	sync.RWMutex		// Read/Write mutex for thread safety'
	data	map[string]string	// The core storage
}

func (d *Database) NewDatabase() *Database {
	return &Database{
		data: make(map[string]string),
	}
}

func (d *Database) SetItem(key string, value string) {
	d.mutex.Lock() // Lock the mutex for writing
	defer d.mutex.Unlock() // Unlock when function returns

	s.data[key] = value
}

func (d *Database) GetItem(key string) (string, bool) {
	d.mutex.RLock() // Lock the mutex for reading
	defer d.mutex.RUnlock()

	item, exists := d.data[key]
	if (!exists) {
		return ("", false)
	}

	return (item, true)
}
