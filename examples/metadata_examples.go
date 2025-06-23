package main

import (
	"fmt"
	"path/filepath"

	"github.com/utsav-56/ufs" // Make sure your module path is correct
)

func main() {
	// Example file and directory paths (adjust as needed)
	filePath := filepath.Join("testdata", "example.txt")
	dirPath := filepath.Join("testdata")

	// Example: GetFileSize
	fileSize := ufs.GetFileSize(filePath)
	fmt.Printf("Size of file %s: %d bytes\n", filePath, fileSize)

	// Example: GetFolderSize
	folderSize := ufs.GetFolderSize(dirPath)
	fmt.Printf("Total size of folder %s: %d bytes\n", dirPath, folderSize)

	// Example: GetFolderList
	folders := ufs.GetFolderList(dirPath)
	fmt.Printf("Folders in %s: %v\n", dirPath, folders)

	// Example: GetFileList
	files := ufs.GetFileList(dirPath)
	fmt.Printf("Files in %s: %v\n", dirPath, files)
}
