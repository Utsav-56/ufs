package main

import (
	"fmt"
	"path/filepath"

	"github.com/utsav-56/ufs" // Make sure your module path is correct
	"github.com/utsav-56/ufs/ulog"
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

	// Example: GetFolderMetadata
	metadata := ufs.Get(dirPath)
	ulog.PrintMapWithIndent(metadata, " ")

	// Example: GetFileMetadata
	fileMetadata := ufs.GetFileMetadata(filePath)
	ulog.PrintMapWithIndent(fileMetadata, " ")

	// Example: GetFolderChildCount
	childCount := ufs.GetFolderChildCount(dirPath)
	fmt.Printf("Number of children in folder %s: %d\n", dirPath, childCount)

}
