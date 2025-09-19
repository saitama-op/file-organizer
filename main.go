package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// CLI flags
	dirPath := flag.String("path", ".", "Path of directory to organize")
	dryRun := flag.Bool("dry-run", false, "Show what would be done without moving files")
	flag.Parse()

	files, err := os.ReadDir(*dirPath)
	if err != nil {
		log.Fatalf("Failed to read directory: %v", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := filepath.Ext(file.Name())
		if ext == "" {
			ext = "no_extension"
		} else {
			ext = ext[1:] // remove dot
		}

		destDir := filepath.Join(*dirPath, ext)
		srcPath := filepath.Join(*dirPath, file.Name())
		destPath := filepath.Join(destDir, file.Name())

		if *dryRun {
			fmt.Printf("[DRY-RUN] Move %s → %s\n", srcPath, destPath)
			continue
		}

		if _, err := os.Stat(destDir); os.IsNotExist(err) {
			if err := os.Mkdir(destDir, 0755); err != nil {
				log.Printf("Failed to create directory %s: %v", destDir, err)
				continue
			}
		}

		err := os.Rename(srcPath, destPath)
		if err != nil {
			log.Printf("Failed to move %s: %v", srcPath, err)
		} else {
			fmt.Printf("Moved %s → %s\n", srcPath, destPath)
		}
	}
}
