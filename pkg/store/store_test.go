package store

import (
	"reflect"
	"testing"
)

func TestReadSaveState(t *testing.T) {
	// Save some data
	err := SaveKeyVal("my key", 0xDEADBEEF)
	err = SaveKeyVal("my other key", 0xBEEFCAFE)
	if err != nil {
		t.Error(err)
	}

	// Wipe the local state
	localState = saveState{}

	// Read the old state from the disk
	got, err := ReadSaveState()
	if err != nil {
		t.Error(err)
	}

	expected := saveState{"my key": 0xDEADBEEF, "my other key": 0xBEEFCAFE}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("\nExpected:<%v>\nGot:<%v>", expected, got)
	}
}
