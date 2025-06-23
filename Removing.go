package ufs

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
Removing.go provides functions to remove files, directories, and symbolic links.

It includes methods for removing files, directories, and symbolic links, with options for recursive deletion and error handling.

This package is part of the ufs library, which provides a unified file system interface for Go applications.
*/

// RemoveFile removes a file at the specified path.
// This function will not remove directories; use RemoveDirectory for that purpose.
//
// Parameters:
//   - path: The absolute or relative path to the file to remove
//
// Returns:
//   - bool: true if the file was removed successfully, false otherwise
//
// Example:
//
//	ok := ufs.RemoveFile("/path/to/file.txt")
//	if !ok {
//	    fmt.Println("Error removing file")
//	}
func (ufs *UFS) RemoveFile(path string) bool {
	// Verify the path is a file
	if !ufs.IsFile(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveFile: Path is not a file: %s", path))
		return false
	}

	err := os.Remove(path)
	if err != nil {
		ufs.handleError(err, "RemoveFile")
		return false
	}
	return true
}

// RemoveDirectory removes an empty directory at the specified path.
// This function will fail if the directory is not empty.
// Use RemoveDirectoryRecursive to remove a directory and all its contents.
//
// Parameters:
//   - path: The absolute or relative path to the directory to remove
//
// Returns:
//   - bool: true if the directory was removed successfully, false otherwise
//
// Example:
//
//	ok := ufs.RemoveDirectory("/path/to/empty_directory")
//	if !ok {
//	    fmt.Println("Error removing directory")
//	}
func (ufs *UFS) RemoveDirectory(path string) bool {
	// Verify the path is a directory
	if !ufs.IsDirectory(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveDirectory: Path is not a directory: %s", path))
		return false
	}

	// Verify the directory is empty
	if !ufs.IsDirectoryEmpty(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveDirectory: Directory is not empty: %s", path))
		return false
	}

	err := os.Remove(path)
	if err != nil {
		ufs.handleError(err, "RemoveDirectory")
		return false
	}
	return true
}

// RemoveDirectoryRecursive removes a directory and all its contents recursively.
// This function will remove all files and subdirectories within the specified directory.
// Use with caution as this operation cannot be undone.
//
// Parameters:
//   - path: The absolute or relative path to the directory to remove
//
// Returns:
//   - bool: true if the directory and all its contents were removed successfully, false otherwise
//
// Example:
//
//	ok := ufs.RemoveDirectoryRecursive("/path/to/directory")
//	if !ok {
//	    fmt.Println("Error removing directory recursively")
//	}
func (ufs *UFS) RemoveDirectoryRecursive(path string) bool {
	// Verify the path is a directory
	if !ufs.IsDirectory(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveDirectoryRecursive: Path is not a directory: %s", path))
		return false
	}

	err := os.RemoveAll(path)
	if err != nil {
		ufs.handleError(err, "RemoveDirectoryRecursive")
		return false
	}
	return true
}

// RemoveSymlink removes a symbolic link at the specified path.
// This function only removes the symlink itself, not the target it points to.
//
// Parameters:
//   - path: The absolute or relative path to the symlink to remove
//
// Returns:
//   - bool: true if the symlink was removed successfully, false otherwise
//
// Example:
//
//	ok := ufs.RemoveSymlink("/path/to/symlink")
//	if !ok {
//	    fmt.Println("Error removing symlink")
//	}
func (ufs *UFS) RemoveSymlink(path string) bool {
	// Check if path is a symlink
	info, err := os.Lstat(path)
	if err != nil {
		ufs.handleError(err, "RemoveSymlink")
		return false
	}

	if info.Mode()&os.ModeSymlink == 0 {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveSymlink: Path is not a symlink: %s", path))
		return false
	}

	err = os.Remove(path)
	if err != nil {
		ufs.handleError(err, "RemoveSymlink")
		return false
	}
	return true
}

// RemoveFileWithBackup removes a file at the specified path after creating a backup.
// The backup file will have the same name with ".bak" appended to it.
//
// Parameters:
//   - path: The absolute or relative path to the file to remove
//
// Returns:
//   - bool: true if the file was backed up and removed successfully, false otherwise
//   - string: The path to the backup file, or an empty string if the operation failed
//
// Example:
//
//	ok, backupPath := ufs.RemoveFileWithBackup("/path/to/file.txt")
//	if !ok {
//	    fmt.Println("Error removing file with backup")
//	} else {
//	    fmt.Printf("File backed up to: %s\n", backupPath)
//	}
func (ufs *UFS) RemoveFileWithBackup(path string) (bool, string) {
	// Verify the path is a file
	if !ufs.IsFile(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveFileWithBackup: Path is not a file: %s", path))
		return false, ""
	}

	// Create backup path
	backupPath := path + ".bak"

	// Read the original file
	content, err := os.ReadFile(path)
	if err != nil {
		ufs.handleError(err, "RemoveFileWithBackup")
		return false, ""
	}

	// Write to backup file
	err = os.WriteFile(backupPath, content, 0644)
	if err != nil {
		ufs.handleError(err, "RemoveFileWithBackup")
		return false, ""
	}

	// Remove the original file
	err = os.Remove(path)
	if err != nil {
		ufs.handleError(err, "RemoveFileWithBackup")
		return false, backupPath
	}

	return true, backupPath
}

// RemoveEmptyFiles removes all empty files in the specified directory.
// This function does not recurse into subdirectories.
//
// Parameters:
//   - dirPath: The absolute or relative path to the directory to clean
//
// Returns:
//   - bool: true if all empty files were removed successfully, false if any removal failed
//   - int: The number of files removed
//
// Example:
//
//	ok, count := ufs.RemoveEmptyFiles("/path/to/directory")
//	if !ok {
//	    fmt.Println("Error removing some empty files")
//	} else {
//	    fmt.Printf("Removed %d empty files\n", count)
//	}
func (ufs *UFS) RemoveEmptyFiles(dirPath string) (bool, int) {
	// Verify the path is a directory
	if !ufs.IsDirectory(dirPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveEmptyFiles: Path is not a directory: %s", dirPath))
		return false, 0
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		ufs.handleError(err, "RemoveEmptyFiles")
		return false, 0
	}

	success := true
	count := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filePath := filepath.Join(dirPath, entry.Name())
		info, err := entry.Info()
		if err != nil {
			ufs.handleError(err, "RemoveEmptyFiles")
			success = false
			continue
		}

		// Check if file is empty
		if info.Size() == 0 {
			if ufs.RemoveFile(filePath) {
				count++
			} else {
				success = false
			}
		}
	}

	return success, count
}

// RemoveEmptyDirectories removes all empty directories in the specified directory.
// This function does not recurse into subdirectories.
//
// Parameters:
//   - dirPath: The absolute or relative path to the directory to clean
//
// Returns:
//   - bool: true if all empty directories were removed successfully, false if any removal failed
//   - int: The number of directories removed
//
// Example:
//
//	ok, count := ufs.RemoveEmptyDirectories("/path/to/directory")
//	if !ok {
//	    fmt.Println("Error removing some empty directories")
//	} else {
//	    fmt.Printf("Removed %d empty directories\n", count)
//	}
func (ufs *UFS) RemoveEmptyDirectories(dirPath string) (bool, int) {
	// Verify the path is a directory
	if !ufs.IsDirectory(dirPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveEmptyDirectories: Path is not a directory: %s", dirPath))
		return false, 0
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		ufs.handleError(err, "RemoveEmptyDirectories")
		return false, 0
	}

	success := true
	count := 0

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		subDirPath := filepath.Join(dirPath, entry.Name())

		// Check if directory is empty
		if ufs.IsDirectoryEmpty(subDirPath) {
			if ufs.RemoveDirectory(subDirPath) {
				count++
			} else {
				success = false
			}
		}
	}

	return success, count
}

// RemoveDirectoryContents removes all contents of a directory without removing the directory itself.
// This function will remove all files and subdirectories within the specified directory.
//
// Parameters:
//   - dirPath: The absolute or relative path to the directory whose contents will be removed
//
// Returns:
//   - bool: true if all contents were removed successfully, false otherwise
//
// Example:
//
//	ok := ufs.RemoveDirectoryContents("/path/to/directory")
//	if !ok {
//	    fmt.Println("Error removing directory contents")
//	}
func (ufs *UFS) RemoveDirectoryContents(dirPath string) bool {
	// Verify the path is a directory
	if !ufs.IsDirectory(dirPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveDirectoryContents: Path is not a directory: %s", dirPath))
		return false
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		ufs.handleError(err, "RemoveDirectoryContents")
		return false
	}

	success := true

	for _, entry := range entries {
		path := filepath.Join(dirPath, entry.Name())
		if entry.IsDir() {
			if !ufs.RemoveDirectoryRecursive(path) {
				success = false
			}
		} else {
			if !ufs.RemoveFile(path) {
				success = false
			}
		}
	}

	return success
}

// RemoveDirectoryTree removes a directory tree structure matching the provided structure.
// The structure is a map where keys are directory names and values are either
// nil (for empty directories) or nested maps (for subdirectories).
// This function is the inverse of CreateDirectoryTree.
//
// Parameters:
//   - basePath: The base directory path where the tree structure exists
//   - structure: A map representing the directory structure to remove
//
// Returns:
//   - bool: true if the directory tree was removed successfully, false otherwise
//
// Example:
//
//	structure := map[string]interface{}{
//	    "dir1": nil,
//	    "dir2": map[string]interface{}{
//	        "subdir1": nil,
//	    },
//	}
//	ok := ufs.RemoveDirectoryTree("/path/to/base", structure)
//	if !ok {
//	    fmt.Println("Error removing directory tree")
//	}
func (ufs *UFS) RemoveDirectoryTree(basePath string, structure map[string]interface{}) bool {
	// Verify the path is a directory
	if !ufs.IsDirectory(basePath) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveDirectoryTree: Base path is not a directory: %s", basePath))
		return false
	}

	success := true

	// First, process subdirectories recursively
	for dirName, subStructure := range structure {
		dirPath := filepath.Join(basePath, dirName)

		if !ufs.IsDirectory(dirPath) {
			continue // Skip if directory doesn't exist
		}

		// If subStructure is not nil, it's a nested directory structure
		if subStructure != nil {
			if subMap, ok := subStructure.(map[string]interface{}); ok {
				// First remove the contents recursively
				if !ufs.RemoveDirectoryTree(dirPath, subMap) {
					success = false
				}
			}
		}

		// Then remove the directory itself if it's empty
		if ufs.IsDirectoryEmpty(dirPath) {
			if !ufs.RemoveDirectory(dirPath) {
				success = false
			}
		}
	}

	return success
}

// RemoveAllLinks removes all symbolic links in the specified directory.
// This function does not recurse into subdirectories.
//
// Parameters:
//   - dirPath: The absolute or relative path to the directory to clean
//
// Returns:
//   - bool: true if all links were removed successfully, false if any removal failed
//   - int: The number of links removed
//
// Example:
//
//	ok, count := ufs.RemoveAllLinks("/path/to/directory")
//	if !ok {
//	    fmt.Println("Error removing some links")
//	} else {
//	    fmt.Printf("Removed %d symbolic links\n", count)
//	}
func (ufs *UFS) RemoveAllLinks(dirPath string) (bool, int) {
	// Verify the path is a directory
	if !ufs.IsDirectory(dirPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveAllLinks: Path is not a directory: %s", dirPath))
		return false, 0
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		ufs.handleError(err, "RemoveAllLinks")
		return false, 0
	}

	success := true
	count := 0

	for _, entry := range entries {
		entryPath := filepath.Join(dirPath, entry.Name())

		// Check if it's a symlink
		info, err := os.Lstat(entryPath)
		if err != nil {
			ufs.handleError(err, "RemoveAllLinks")
			success = false
			continue
		}

		if info.Mode()&os.ModeSymlink != 0 {
			if ufs.RemoveSymlink(entryPath) {
				count++
			} else {
				success = false
			}
		}
	}

	return success, count
}

// RemoveByPattern removes all files matching a specified pattern in the given directory.
// This function uses filepath.Match for pattern matching.
//
// Parameters:
//   - dirPath: The absolute or relative path to the directory to clean
//   - pattern: The pattern to match files against (e.g., "*.tmp", "backup-*")
//
// Returns:
//   - bool: true if all matching files were removed successfully, false if any removal failed
//   - int: The number of files removed
//
// Example:
//
//	ok, count := ufs.RemoveByPattern("/path/to/directory", "*.tmp")
//	if !ok {
//	    fmt.Println("Error removing some temporary files")
//	} else {
//	    fmt.Printf("Removed %d temporary files\n", count)
//	}
func (ufs *UFS) RemoveByPattern(dirPath, pattern string) (bool, int) {
	// Verify the path is a directory
	if !ufs.IsDirectory(dirPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("RemoveByPattern: Path is not a directory: %s", dirPath))
		return false, 0
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		ufs.handleError(err, "RemoveByPattern")
		return false, 0
	}

	success := true
	count := 0

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		match, err := filepath.Match(pattern, entry.Name())
		if err != nil {
			ufs.handleError(err, "RemoveByPattern")
			success = false
			continue
		}

		if match {
			filePath := filepath.Join(dirPath, entry.Name())
			if ufs.RemoveFile(filePath) {
				count++
			} else {
				success = false
			}
		}
	}

	return success, count
}

// SafeRemoveFile removes a file only if it matches the expected size and/or modification time.
// This provides a safety check before deletion to prevent accidental removal of important files.
//
// Parameters:
//   - path: The absolute or relative path to the file to remove
//   - expectedSize: The expected size of the file in bytes, or -1 to skip this check
//   - expectedModTime: The expected modification time of the file, or nil to skip this check
//
// Returns:
//   - bool: true if the file was removed successfully, false otherwise
//
// Example:
//
//	modTime := time.Date(2023, 6, 15, 12, 0, 0, 0, time.Local)
//	ok := ufs.SafeRemoveFile("/path/to/file.txt", 1024, &modTime)
//	if !ok {
//	    fmt.Println("Error: File did not match expected criteria or couldn't be removed")
//	}
func (ufs *UFS) SafeRemoveFile(path string, expectedSize int64, expectedModTime *os.FileInfo) bool {
	// Verify the path is a file
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "SafeRemoveFile")
		return false
	}

	if info.IsDir() {
		ufs.handleMistakeWarning(fmt.Sprintf("SafeRemoveFile: Path is not a file: %s", path))
		return false
	}

	// Check file size if specified
	if expectedSize >= 0 && info.Size() != expectedSize {
		ufs.handleMistakeWarning(fmt.Sprintf("SafeRemoveFile: File size mismatch: expected %d, got %d",
			expectedSize, info.Size()))
		return false
	}

	// Check modification time if specified
	if expectedModTime != nil && (*expectedModTime).ModTime() != info.ModTime() {
		ufs.handleMistakeWarning(fmt.Sprintf("SafeRemoveFile: File modification time mismatch"))
		return false
	}

	// All checks passed, remove the file
	err = os.Remove(path)
	if err != nil {
		ufs.handleError(err, "SafeRemoveFile")
		return false
	}

	return true
}
