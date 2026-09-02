package main

import (
	"errors"
	"testing"
)

func TestValidatePath(t *testing.T) {

	err := validatePath(`C:\gibtsnicht`)

	if !errors.Is(err, errPathNotFound) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", err, errPathNotFound)
	}

	err = validatePath(`C:\coding\randomslide\main.go`)

	if !errors.Is(err, errPathIsNotDir) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", err, errPathIsNotDir)
	}

	err = validatePath(`C:\coding`)

	if err != nil {

		t.Errorf("Result was incorrect, got: %v, want: nil.", err)
	}
}
