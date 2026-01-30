package tools

import (
	"sync"
)

type Database struct {
	setMutex	sync.RWMutex	// Read/Write mutex for thread safety'
	hsetMutex	sync.RWMutex
	sets	map[string]string	// The core storage for SET
	hsets	map[string]map[string]string // The core storage for HSET
}

func NewDatabase() *Database {
	return &Database{
		sets: map[string]string{},
		hsets: map[string]map[string]string{},
	}
}

func (d *Database) SetItem(key string, value string) {
	d.setMutex.Lock() // Lock the mutex for writing
	defer d.setMutex.Unlock() // Unlock when function returns

	d.sets[key] = value
}

func (d *Database) GetItem(key string) (string, bool) {
	d.setMutex.RLock() // Lock the mutex for reading
	defer d.setMutex.RUnlock()

	item, exists := d.sets[key]
	if (!exists) {
		return "", false
	}

	return item, true
}

func (d *Database) SetHash(hash string, key string, value string) {
	d.hsetMutex.Lock()
	defer d.hsetMutex.Unlock()

	if _, ok := d.hsets[hash]; !ok {
		d.hsets[hash] = map[string]string{}
	}

	d.hsets[hash][key] = value
}

func (d *Database) GetHash(hash string, key string) (string, bool) {
	d.hsetMutex.RLock()
	defer d.hsetMutex.RUnlock()

	item, exists := d.hsets[hash][key]

	if !exists {
		return "", false
	}

	return item, true
}

func (d *Database) GetAllHash(hash string) (map[string]string, bool) {
	d.hsetMutex.RLock()
	defer d.hsetMutex.RUnlock()

	items, exists := d.hsets[hash]

	if !exists {
		return map[string]string{}, false
	}

	return items, true
}
