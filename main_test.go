package main

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
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

func TestFindMedia(t *testing.T) {

	allFiles := []string{
		"picture.jpg",
		"otherpicture.JpG",
		"executable.exe",
		"movie.mp4",
		"OTHERMOVIE.MP4",
	}

	expectedFiles := []string{
		"picture.jpg",
		"otherpicture.JpG",
		"movie.mp4",
		"OTHERMOVIE.MP4",
	}

	tempDir := t.TempDir()

	tempSubDir := filepath.Join(tempDir, `unterordner`)

	err := os.Mkdir(tempSubDir, 0755)

	if err != nil {
		t.Fatalf("before testing the necessary temporary directory could not created")
	}

	for _, file := range allFiles {
		err = os.WriteFile(filepath.Join(tempDir, file), nil, 0644)
		if err != nil {
			t.Fatalf("could not create test file %s: %v", file, err)
		}

	}

	for _, file := range allFiles {
		err = os.WriteFile(filepath.Join(tempSubDir, file), nil, 0644)
		if err != nil {
			t.Fatalf("could not create test file %s: %v", file, err)
		}

	}

	foundMedia, err := findMedia(tempDir)
	if err != nil {
		t.Fatalf("could not use findMedia: %v", err)
	}

	if len(foundMedia) != 8 {
		t.Errorf("the expected number of files found is incorrect")
	}

	for _, expectedFile := range expectedFiles {
		expectedPath := filepath.Join(tempDir, expectedFile)
		if !slices.Contains(foundMedia, expectedPath) {
			t.Errorf("expected file not found: %v", expectedPath)
		}
	}
	for _, expectedFile := range expectedFiles {
		expectedPath := filepath.Join(tempSubDir, expectedFile)

		if !slices.Contains(foundMedia, expectedPath) {
			t.Errorf("expected file not found: %v", expectedPath)
		}
	}

}
