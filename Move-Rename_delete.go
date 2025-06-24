package ufs

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
Move-Rename_delete.go

This file contains functions related to moving, renaming, and deleting files and directories.
It includes functions to move or rename files, delete files, and delete directories.

These functions are designed to be used with the UFS package, providing a unified interface for file system operations.

The functions handle errors gracefully and return boolean values indicating success or failure.
They also ensure that the operations are performed on absolute paths, resolving any relative paths as necessary.

// Functions:

- MoveFile: Moves or renames a file from one path to another
- DeleteFile: Deletes a file at the specified path
- DeleteDirectory: Deletes a directory at the specified path, including all its contents
- MoveDirectory: Moves or renames a directory from one path to another

Advance checked functions:
- MoveFileIfExists: Moves a file only if it exists at the source path
- MoveDirectoryIfExists: Moves a directory only if it exists at the source path
- DeleteFileIfExists: Deletes a file only if it exists at the specified path
- DeleteDirectoryIfExists: Deletes a directory only if it exists at the specified path

More advanced functions:
- MoveDirectoryIfEmpty: Moves a directory only if it is empty
- MoveFileIfEmpty: Moves a file only if it is empty
- DeleteFileIfEmpty: Deletes a file only if it is empty
- DeleteDirectoryIfEmpty: Deletes a directory only if it is empty
*/

// MoveFile moves or renames a file from one path to another.
// If the destination already exists, it will be overwritten.
// This function will create any parent directories for the destination if they don't exist.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source file
//   - destPath: The absolute or relative path where the file should be moved to
//
// Returns:
//   - bool: true if the file was moved successfully, false otherwise
//
// Example:
//
//	success := ufs.MoveFile("/path/to/source.txt", "/path/to/destination.txt")
//	if !success {
//	    fmt.Println("Failed to move file")
//	}
func (ufs *UFS) MoveFile(srcPath, destPath string) bool {
	// Verify source is a file
	if !ufs.IsFile(srcPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveFile: Source is not a file: %s", srcPath))
		return false
	}

	// Ensure destination directory exists
	destDir := filepath.Dir(destPath)
	if !ufs.IsDirectory(destDir) {
		if !ufs.CreateDirectory(destDir) {
			return false
		}
	}

	// If destination exists and is a file, remove it
	if ufs.IsFile(destPath) {
		if !ufs.RemoveFile(destPath) {
			return false
		}
	}

	// Move the file
	err := os.Rename(srcPath, destPath)
	if err != nil {
		// Try copy and delete if rename fails (e.g., across different filesystems)
		if !ufs.copyThenDelete(srcPath, destPath) {
			ufs.handleError(err, "MoveFile")
			return false
		}
	}

	return true
}

// DeleteFile deletes a file at the specified path.
// This is a wrapper around RemoveFile for consistency with naming.
//
// Parameters:
//   - path: The absolute or relative path to the file to delete
//
// Returns:
//   - bool: true if the file was deleted successfully, false otherwise
//
// Example:
//
//	success := ufs.DeleteFile("/path/to/file.txt")
//	if !success {
//	    fmt.Println("Failed to delete file")
//	}
func (ufs *UFS) DeleteFile(path string) bool {
	return ufs.RemoveFile(path)
}

// DeleteDirectory deletes a directory at the specified path, including all its contents.
// This is a wrapper around RemoveDirectoryRecursive for consistency with naming.
//
// Parameters:
//   - path: The absolute or relative path to the directory to delete
//
// Returns:
//   - bool: true if the directory was deleted successfully, false otherwise
//
// Example:
//
//	success := ufs.DeleteDirectory("/path/to/directory")
//	if !success {
//	    fmt.Println("Failed to delete directory")
//	}
func (ufs *UFS) DeleteDirectory(path string) bool {
	return ufs.RemoveDirectoryRecursive(path)
}

// MoveDirectory moves or renames a directory from one path to another.
// If the destination already exists as a directory, it will attempt to merge the contents.
// This function will create any parent directories for the destination if they don't exist.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source directory
//   - destPath: The absolute or relative path where the directory should be moved to
//
// Returns:
//   - bool: true if the directory was moved successfully, false otherwise
//
// Example:
//
//	success := ufs.MoveDirectory("/path/to/source_dir", "/path/to/destination_dir")
//	if !success {
//	    fmt.Println("Failed to move directory")
//	}
func (ufs *UFS) MoveDirectory(srcPath, destPath string) bool {
	// Verify source is a directory
	if !ufs.IsDirectory(srcPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveDirectory: Source is not a directory: %s", srcPath))
		return false
	}

	// Ensure destination parent directory exists
	destParent := filepath.Dir(destPath)
	if !ufs.IsDirectory(destParent) {
		if !ufs.CreateDirectory(destParent) {
			return false
		}
	}

	// If destination doesn't exist, try simple rename
	if !ufs.PathExists(destPath) {
		err := os.Rename(srcPath, destPath)
		if err == nil {
			return true
		}
		// Continue to the more complex move if rename fails
	}

	// If destination exists and is a directory, merge contents
	if ufs.IsDirectory(destPath) {
		return ufs.mergeDirectories(srcPath, destPath)
	}

	// If destination exists and is a file, fail
	if ufs.IsFile(destPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveDirectory: Destination exists and is a file: %s", destPath))
		return false
	}

	// Fallback case: try to create destination directory and copy contents
	if !ufs.CreateDirectory(destPath) {
		return false
	}

	return ufs.mergeDirectories(srcPath, destPath)
}

// MoveFileIfExists moves a file only if it exists at the source path.
// If the source file doesn't exist, the function returns true without doing anything.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source file
//   - destPath: The absolute or relative path where the file should be moved to
//
// Returns:
//   - bool: true if the file was moved successfully or doesn't exist, false otherwise
//
// Example:
//
//	success := ufs.MoveFileIfExists("/path/to/maybe_existing.txt", "/path/to/destination.txt")
//	if !success {
//	    fmt.Println("Failed to move file (if it existed)")
//	}
func (ufs *UFS) MoveFileIfExists(srcPath, destPath string) bool {
	if !ufs.IsFile(srcPath) {
		return true // Success: nothing to move
	}
	return ufs.MoveFile(srcPath, destPath)
}

// MoveDirectoryIfExists moves a directory only if it exists at the source path.
// If the source directory doesn't exist, the function returns true without doing anything.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source directory
//   - destPath: The absolute or relative path where the directory should be moved to
//
// Returns:
//   - bool: true if the directory was moved successfully or doesn't exist, false otherwise
//
// Example:
//
//	success := ufs.MoveDirectoryIfExists("/path/to/maybe_existing_dir", "/path/to/destination_dir")
//	if !success {
//	    fmt.Println("Failed to move directory (if it existed)")
//	}
func (ufs *UFS) MoveDirectoryIfExists(srcPath, destPath string) bool {
	if !ufs.IsDirectory(srcPath) {
		return true // Success: nothing to move
	}
	return ufs.MoveDirectory(srcPath, destPath)
}

// DeleteFileIfExists deletes a file only if it exists at the specified path.
// If the file doesn't exist, the function returns true without doing anything.
//
// Parameters:
//   - path: The absolute or relative path to the file to delete
//
// Returns:
//   - bool: true if the file was deleted successfully or doesn't exist, false otherwise
//
// Example:
//
//	success := ufs.DeleteFileIfExists("/path/to/maybe_existing.txt")
//	if !success {
//	    fmt.Println("Failed to delete file (if it existed)")
//	}
func (ufs *UFS) DeleteFileIfExists(path string) bool {
	if !ufs.IsFile(path) {
		return true // Success: nothing to delete
	}
	return ufs.DeleteFile(path)
}

// DeleteDirectoryIfExists deletes a directory only if it exists at the specified path.
// If the directory doesn't exist, the function returns true without doing anything.
//
// Parameters:
//   - path: The absolute or relative path to the directory to delete
//
// Returns:
//   - bool: true if the directory was deleted successfully or doesn't exist, false otherwise
//
// Example:
//
//	success := ufs.DeleteDirectoryIfExists("/path/to/maybe_existing_dir")
//	if !success {
//	    fmt.Println("Failed to delete directory (if it existed)")
//	}
func (ufs *UFS) DeleteDirectoryIfExists(path string) bool {
	if !ufs.IsDirectory(path) {
		return true // Success: nothing to delete
	}
	return ufs.DeleteDirectory(path)
}

// MoveDirectoryIfEmpty moves a directory only if it is empty.
// If the source directory is not empty, the function returns false without doing anything.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source directory
//   - destPath: The absolute or relative path where the directory should be moved to
//
// Returns:
//   - bool: true if the directory was moved successfully or is not empty, false otherwise
//
// Example:
//
//	success := ufs.MoveDirectoryIfEmpty("/path/to/maybe_empty_dir", "/path/to/destination_dir")
//	if !success {
//	    fmt.Println("Failed to move directory (it might not be empty)")
//	}
func (ufs *UFS) MoveDirectoryIfEmpty(srcPath, destPath string) bool {
	// Verify source is a directory
	if !ufs.IsDirectory(srcPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveDirectoryIfEmpty: Source is not a directory: %s", srcPath))
		return false
	}

	// Check if directory is empty
	if !ufs.IsDirectoryEmpty(srcPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveDirectoryIfEmpty: Source directory is not empty: %s", srcPath))
		return false
	}

	return ufs.MoveDirectory(srcPath, destPath)
}

// MoveFileIfEmpty moves a file only if it is empty (zero bytes).
// If the source file is not empty, the function returns false without doing anything.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source file
//   - destPath: The absolute or relative path where the file should be moved to
//
// Returns:
//   - bool: true if the file was moved successfully or is not empty, false otherwise
//
// Example:
//
//	success := ufs.MoveFileIfEmpty("/path/to/maybe_empty.txt", "/path/to/destination.txt")
//	if !success {
//	    fmt.Println("Failed to move file (it might not be empty)")
//	}
func (ufs *UFS) MoveFileIfEmpty(srcPath, destPath string) bool {
	// Verify source is a file
	if !ufs.IsFile(srcPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveFileIfEmpty: Source is not a file: %s", srcPath))
		return false
	}

	// Check if file is empty
	if !ufs.IsFileEmpty(srcPath) {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveFileIfEmpty: Source file is not empty: %s", srcPath))
		return false
	}

	return ufs.MoveFile(srcPath, destPath)
}

// DeleteFileIfEmpty deletes a file only if it is empty (zero bytes).
// If the file is not empty, the function returns false without doing anything.
//
// Parameters:
//   - path: The absolute or relative path to the file to delete
//
// Returns:
//   - bool: true if the file was deleted successfully or is not empty, false otherwise
//
// Example:
//
//	success := ufs.DeleteFileIfEmpty("/path/to/maybe_empty.txt")
//	if !success {
//	    fmt.Println("Failed to delete file (it might not be empty)")
//	}
func (ufs *UFS) DeleteFileIfEmpty(path string) bool {
	// Verify path is a file
	if !ufs.IsFile(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("DeleteFileIfEmpty: Path is not a file: %s", path))
		return false
	}

	// Check if file is empty
	if !ufs.IsFileEmpty(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("DeleteFileIfEmpty: File is not empty: %s", path))
		return false
	}

	return ufs.DeleteFile(path)
}

// DeleteDirectoryIfEmpty deletes a directory only if it is empty.
// If the directory is not empty, the function returns false without doing anything.
//
// Parameters:
//   - path: The absolute or relative path to the directory to delete
//
// Returns:
//   - bool: true if the directory was deleted successfully or is not empty, false otherwise
//
// Example:
//
//	success := ufs.DeleteDirectoryIfEmpty("/path/to/maybe_empty_dir")
//	if !success {
//	    fmt.Println("Failed to delete directory (it might not be empty)")
//	}
func (ufs *UFS) DeleteDirectoryIfEmpty(path string) bool {
	// Verify path is a directory
	if !ufs.IsDirectory(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("DeleteDirectoryIfEmpty: Path is not a directory: %s", path))
		return false
	}

	// Check if directory is empty
	if !ufs.IsDirectoryEmpty(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("DeleteDirectoryIfEmpty: Directory is not empty: %s", path))
		return false
	}

	// Use RemoveDirectory (not RemoveDirectoryRecursive) since we know it's empty
	return ufs.RemoveDirectory(path)
}

// RenameFile renames a file without moving it to a different directory.
// This is a convenience wrapper around MoveFile for cases where only the name changes.
//
// Parameters:
//   - path: The absolute or relative path to the file to rename
//   - newName: The new name for the file (not a path, just the filename)
//
// Returns:
//   - bool: true if the file was renamed successfully, false otherwise
//
// Example:
//
//	success := ufs.RenameFile("/path/to/old_name.txt", "new_name.txt")
//	if !success {
//	    fmt.Println("Failed to rename file")
//	}
func (ufs *UFS) RenameFile(path string, newName string) bool {
	// Verify source is a file
	if !ufs.IsFile(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RenameFile: Source is not a file: %s", path))
		return false
	}

	// Ensure newName is just a filename, not a path
	if filepath.Base(newName) != newName {
		ufs.handleMistakeWarning(fmt.Sprintf("RenameFile: New name should not be a path: %s", newName))
		return false
	}

	// Compute new path
	dir := filepath.Dir(path)
	newPath := filepath.Join(dir, newName)

	return ufs.MoveFile(path, newPath)
}

// RenameDirectory renames a directory without moving it to a different location.
// This is a convenience wrapper around MoveDirectory for cases where only the name changes.
//
// Parameters:
//   - path: The absolute or relative path to the directory to rename
//   - newName: The new name for the directory (not a path, just the directory name)
//
// Returns:
//   - bool: true if the directory was renamed successfully, false otherwise
//
// Example:
//
//	success := ufs.RenameDirectory("/path/to/old_name_dir", "new_name_dir")
//	if !success {
//	    fmt.Println("Failed to rename directory")
//	}
func (ufs *UFS) RenameDirectory(path string, newName string) bool {
	// Verify source is a directory
	if !ufs.IsDirectory(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("RenameDirectory: Source is not a directory: %s", path))
		return false
	}

	// Ensure newName is just a directory name, not a path
	if filepath.Base(newName) != newName {
		ufs.handleMistakeWarning(fmt.Sprintf("RenameDirectory: New name should not be a path: %s", newName))
		return false
	}

	// Compute new path
	dir := filepath.Dir(path)
	newPath := filepath.Join(dir, newName)

	return ufs.MoveDirectory(path, newPath)
}

// MoveWithBackup moves a file or directory after creating a backup of the destination if it exists.
// The backup will have the same name with ".bak" appended.
//
// Parameters:
//   - srcPath: The absolute or relative path to the source file or directory
//   - destPath: The absolute or relative path where the file or directory should be moved to
//
// Returns:
//   - bool: true if the move and backup were successful, false otherwise
//   - string: The path to the backup if one was created, or an empty string
//
// Example:
//
//	success, backupPath := ufs.MoveWithBackup("/path/to/source.txt", "/path/to/destination.txt")
//	if !success {
//	    fmt.Println("Failed to move with backup")
//	} else if backupPath != "" {
//	    fmt.Printf("Destination was backed up to: %s\n", backupPath)
//	}
func (ufs *UFS) MoveWithBackup(srcPath, destPath string) (bool, string) {
	backupPath := ""

	// If destination exists, create a backup
	if ufs.PathExists(destPath) {
		backupPath = destPath + ".bak"

		// Remove any existing backup
		if ufs.PathExists(backupPath) {
			if ufs.IsFile(backupPath) {
				if !ufs.RemoveFile(backupPath) {
					return false, ""
				}
			} else if ufs.IsDirectory(backupPath) {
				if !ufs.RemoveDirectoryRecursive(backupPath) {
					return false, ""
				}
			}
		}

		// Create backup
		if ufs.IsFile(destPath) {
			if !ufs.MoveFile(destPath, backupPath) {
				return false, ""
			}
		} else if ufs.IsDirectory(destPath) {
			if !ufs.MoveDirectory(destPath, backupPath) {
				return false, ""
			}
		}
	}

	// Perform the move
	success := false
	if ufs.IsFile(srcPath) {
		success = ufs.MoveFile(srcPath, destPath)
	} else if ufs.IsDirectory(srcPath) {
		success = ufs.MoveDirectory(srcPath, destPath)
	} else {
		ufs.handleMistakeWarning(fmt.Sprintf("MoveWithBackup: Source is neither a file nor a directory: %s", srcPath))
		return false, backupPath
	}

	if !success && backupPath != "" {
		// Restore from backup if move failed
		if ufs.IsFile(backupPath) {
			ufs.MoveFile(backupPath, destPath)
		} else if ufs.IsDirectory(backupPath) {
			ufs.MoveDirectory(backupPath, destPath)
		}
		return false, ""
	}

	return success, backupPath
}

// DeleteWithBackup deletes a file or directory after creating a backup.
// The backup will have the same name with ".bak" appended.
//
// Parameters:
//   - path: The absolute or relative path to the file or directory to delete
//
// Returns:
//   - bool: true if the deletion and backup were successful, false otherwise
//   - string: The path to the backup that was created
//
// Example:
//
//	success, backupPath := ufs.DeleteWithBackup("/path/to/file.txt")
//	if !success {
//	    fmt.Println("Failed to delete with backup")
//	} else {
//	    fmt.Printf("File was backed up to: %s before deletion\n", backupPath)
//	}
func (ufs *UFS) DeleteWithBackup(path string) (bool, string) {
	// Verify path exists
	if !ufs.PathExists(path) {
		ufs.handleMistakeWarning(fmt.Sprintf("DeleteWithBackup: Path does not exist: %s", path))
		return false, ""
	}

	backupPath := path + ".bak"

	// Remove any existing backup
	if ufs.PathExists(backupPath) {
		if ufs.IsFile(backupPath) {
			if !ufs.RemoveFile(backupPath) {
				return false, ""
			}
		} else if ufs.IsDirectory(backupPath) {
			if !ufs.RemoveDirectoryRecursive(backupPath) {
				return false, ""
			}
		}
	}

	// Create backup
	if ufs.IsFile(path) {
		if err := ufs.CopyFile(path, backupPath); err != nil {
			return false, ""
		}
		return ufs.DeleteFile(path), backupPath
	} else if ufs.IsDirectory(path) {
		// For directories, we need to copy the entire structure
		success := ufs.copyDirectoryRecursive(path, backupPath)
		if !success {
			return false, ""
		}
		return ufs.DeleteDirectory(path), backupPath
	}

	ufs.handleMistakeWarning(fmt.Sprintf("DeleteWithBackup: Path is neither a file nor a directory: %s", path))
	return false, ""
}

// copyThenDelete is a helper function that copies a file and then deletes the source
// Used when os.Rename fails (e.g., across filesystems)
func (ufs *UFS) copyThenDelete(srcPath, destPath string) bool {
	// Copy the file
	if err := ufs.CopyFile(srcPath, destPath); err != nil {
		return false
	}

	// Delete the source
	if !ufs.DeleteFile(srcPath) {
		// If delete fails, try to remove the destination to avoid duplicates
		ufs.DeleteFile(destPath)
		return false
	}

	return true
}

// mergeDirectories is a helper function that merges the contents of srcPath into destPath
// and then removes srcPath if successful
func (ufs *UFS) mergeDirectories(srcPath, destPath string) bool {
	// Get all entries in the source directory
	entries, err := os.ReadDir(srcPath)
	if err != nil {
		ufs.handleError(err, "mergeDirectories")
		return false
	}

	success := true

	// Move each entry to the destination
	for _, entry := range entries {
		srcItemPath := filepath.Join(srcPath, entry.Name())
		destItemPath := filepath.Join(destPath, entry.Name())

		if entry.IsDir() {
			// If it's a directory, recursively move it
			if !ufs.MoveDirectoryIfExists(srcItemPath, destItemPath) {
				success = false
			}
		} else {
			// If it's a file, move it
			if !ufs.MoveFileIfExists(srcItemPath, destItemPath) {
				success = false
			}
		}
	}

	// If all items were moved successfully, remove the source directory
	if success {
		err := os.Remove(srcPath) // Only removes if empty
		if err != nil {
			ufs.handleError(err, "mergeDirectories")
			success = false
		}
	}

	return success
}

// copyDirectoryRecursive is a helper function that copies a directory and all its contents
func (ufs *UFS) copyDirectoryRecursive(srcPath, destPath string) bool {
	// Create the destination directory
	if !ufs.CreateDirectory(destPath) {
		return false
	}

	// Get all entries in the source directory
	entries, err := os.ReadDir(srcPath)
	if err != nil {
		ufs.handleError(err, "copyDirectoryRecursive")
		return false
	}

	success := true

	// Copy each entry to the destination
	for _, entry := range entries {
		srcItemPath := filepath.Join(srcPath, entry.Name())
		destItemPath := filepath.Join(destPath, entry.Name())

		if entry.IsDir() {
			// If it's a directory, recursively copy it
			if !ufs.copyDirectoryRecursive(srcItemPath, destItemPath) {
				success = false
			}
		} else {
			// If it's a file, copy it
			if err := ufs.CopyFile(srcItemPath, destItemPath); err != nil {
				success = false
			}
		}
	}

	return success
}
