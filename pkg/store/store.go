package store

import (
	"encoding/gob"
	"os"
	"sync"
)

const filename = ".store.bruh"

type saveState map[string]any

var localState = make(saveState)
var mu = new(sync.Mutex)

// ReadSaveState updates the local save state variable by reading data from the disk.
func ReadSaveState() (saveState, error) {
	file, err := os.Open(filename)
	if err != nil {
		// File doesn't exist; return fresh local save state
		return localState, nil
	}
	defer file.Close()

	err = gob.NewDecoder(file).Decode(&localState)
	if err != nil {
		return localState, err
	}
	return localState, nil
}

// SaveKeyVal saves a key value pair to the store.
func SaveKeyVal(key string, val any) error {
	mu.Lock()
	defer mu.Unlock()

	// Create save file if it doesn't exist
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Update last saveState variable and write it to the disk
	localState[key] = val
	err = gob.NewEncoder(file).Encode(localState)

	if err != nil {
		return err
	}
	return nil
}
