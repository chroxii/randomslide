package main

import (
	"errors"
	"fmt"
	"io/fs"
	"math/rand"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

var errPathNotFound = errors.New("path does not exist")
var errPathIsNotDir = errors.New("path is not a directory")

func main() {
	fmt.Println("RandomSlide starting...")

	if len(os.Args) != 2 {
		fmt.Println("Usage: randomslide <directory>")
		os.Exit(1)
	}

	pathFromUser := os.Args[1]

	err := validatePath(pathFromUser)

	if errors.Is(err, errPathNotFound) {
		fmt.Println("Error: path does not exist")
		os.Exit(1)
	}
	if errors.Is(err, errPathIsNotDir) {
		fmt.Println("Error: path is not a directory")
		os.Exit(1)
	}
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("Media directory:", pathFromUser)

	mediaFiles, err := findMedia(pathFromUser)

	if err != nil {
		fmt.Println("Error while walking directory:", err)
		os.Exit(1)
	}

	fmt.Println("Found", len(mediaFiles), "media files")

	if len(mediaFiles) == 0 {
		fmt.Println("No media files found")
		os.Exit(1)
	}
	fmt.Println("Random media:", mediaFiles[rand.Intn(len(mediaFiles))])
}

func findMedia(rootPath string) ([]string, error) {

	supportedExtensions := []string{".jpg", ".jpeg", ".png", ".webp", ".mp4", ".webm", ".mkv", ".mov", ".gif"}

	mediaFiles := []string{}

	walkErr := filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			currentExt := strings.ToLower(filepath.Ext(path))
			if slices.Contains(supportedExtensions, currentExt) {
				mediaFiles = append(mediaFiles, path)
			}
		}
		return nil
	},
	)

	if walkErr != nil {
		return nil, walkErr
	}
	return mediaFiles, nil
}

func validatePath(path string) error {

	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errPathNotFound
		}
		return fmt.Errorf("an error occurred: %w", err)
	}
	if !info.IsDir() {
		return errPathIsNotDir
	}
	return nil
}
