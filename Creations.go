package ufs

import (
	"io/fs"
	"os"
	"path/filepath"
)

/*
Creations.go provides functions to create files and directories in the UFS (Ultimate File System).

It includes methods to create files, directories, and symbolic links, as well as to create a directory tree.

This file is part of the UFS library, which provides a unified interface for file system operations.

Some especial methods includes:
create file with content, create directory with permissions,
create symbolic link, and create a directory tree with specified permissions.
also provides option to symlink whole directory tree.
*/

// CreateFile creates a new empty file at the specified path.
// If the file already exists, it will be truncated to zero length.
//
// Parameters:
//   - path: The absolute or relative path to the file to create
//
// Returns:
//   - bool: true if the file was created successfully, false otherwise
//
// Example:
//
//	ok := ufs.CreateFile("/path/to/new_file.txt")
//	if !ok {
//	    fmt.Printf("Error creating file\n")
//	}
func (ufs *UFS) CreateFile(path string) bool {
	file, err := os.Create(path)
	if err != nil {
		ufs.handleError(err, "CreateFile")
		return false
	}
	defer file.Close()
	return true
}

// CreateFileWithContent creates a new file at the specified path with the given content.
// If the file already exists, it will be overwritten.
//
// Parameters:
//   - path: The absolute or relative path to the file to create
//   - content: The content to write to the file as a string
//
// Returns:
//   - bool: true if the file was created and written successfully, false otherwise
//
// Example:
//
//	ok := ufs.CreateFileWithContent("/path/to/new_file.txt", "Hello, World!")
//	if !ok {
//	    fmt.Printf("Error creating file with content\n")
//	}
func (ufs *UFS) CreateFileWithContent(path string, content string) bool {
	file, err := os.Create(path)
	if err != nil {
		ufs.handleError(err, "CreateFileWithContent")
		return false
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		ufs.handleError(err, "CreateFileWithContent")
		return false
	}

	return true
}

// CreateFileWithContentAndPermissions creates a new file at the specified path with the given content and permissions.
// If the file already exists, it will be overwritten.
//
// Parameters:
//   - path: The absolute or relative path to the file to create
//   - content: The content to write to the file as a string
//   - perm: The file permissions (e.g., 0644 for read/write by owner, read-only for others)
//
// Returns:
//   - bool: true if the file was created and written successfully, false otherwise
//
// Example:
//
//	ok := ufs.CreateFileWithContentAndPermissions("/path/to/new_file.txt", "Hello, World!", 0644)
//	if !ok {
//	    fmt.Printf("Error creating file with content and permissions\n")
//	}
func (ufs *UFS) CreateFileWithContentAndPermissions(path string, content string, perm fs.FileMode) bool {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		ufs.handleError(err, "CreateFileWithContentAndPermissions")
		return false
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		ufs.handleError(err, "CreateFileWithContentAndPermissions")
		return false
	}

	return true
}

// CreateFileWithPermissions creates a new empty file at the specified path with the given permissions.
// If the file already exists, it will be truncated to zero length.
//
// Parameters:
//   - path: The absolute or relative path to the file to create
//   - perm: The file permissions (e.g., 0644 for read/write by owner, read-only for others)
//
// Returns:
//   - bool: true if the file was created successfully, false otherwise
//
// Example:
//
//	ok := ufs.CreateFileWithPermissions("/path/to/new_file.txt", 0644)
//	if !ok {
//	    fmt.Printf("Error creating file with permissions\n")
//	}
func (ufs *UFS) CreateFileWithPermissions(path string, perm fs.FileMode) bool {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, perm)
	if err != nil {
		ufs.handleError(err, "CreateFileWithPermissions")
		return false
	}
	defer file.Close()
	return true
}

// CreateDirectory creates a new directory at the specified path.
// If the directory already exists, no error is returned.
//
// Parameters:
//   - path: The absolute or relative path to the directory to create
//
// Returns:
//   - bool: true if the directory was created successfully or already exists, false otherwise
//
// Example:
//
//	ok := ufs.CreateDirectory("/path/to/new_directory")
//	if !ok {
//	    fmt.Printf("Error creating directory\n")
//	}
func (ufs *UFS) CreateDirectory(path string) bool {
	err := os.MkdirAll(path, 0755) // Default permissions: rwxr-xr-x
	if err != nil {
		ufs.handleError(err, "CreateDirectory")
		return false
	}
	return true
}

// CreateDirectoryWithPermissions creates a new directory at the specified path with the given permissions.
// If the directory already exists, no error is returned but permissions won't be changed.
//
// Parameters:
//   - path: The absolute or relative path to the directory to create
//   - perm: The directory permissions (e.g., 0755 for read/write/execute by owner, read/execute for others)
//
// Returns:
//   - bool:  if the directory could not be created, false otherwise
//
// Example:
//
//	err := ufs.CreateDirectoryWithPermissions("/path/to/new_directory", 0755)
//	if err != nil {
//	    fmt.Printf("Error creating directory with permissions: %v\n", err)
//	}
func (ufs *UFS) CreateDirectoryWithPermissions(path string, perm fs.FileMode) bool {
	err := os.MkdirAll(path, perm)
	if err != nil {
		ufs.handleError(err, "CreateDirectoryWithPermissions")
		return false
	}
	return true
}

// CreateSymlink creates a symbolic link at the specified path pointing to the target.
//
// Parameters:
//   - target: The file or directory that the symlink will point to
//   - symlink: The path where the symlink will be created
//
// Returns:
//   - bool: true if the symlink was created successfully, false otherwise
//
// Example:
//
//	ok := ufs.CreateSymlink("/path/to/target", "/path/to/symlink")
//	if !ok {
//	    fmt.Printf("Error creating symlink\n")
//	}
func (ufs *UFS) CreateSymlink(target string, symlink string) bool {
	err := os.Symlink(target, symlink)
	if err != nil {
		ufs.handleError(err, "CreateSymlink")
		return false
	}
	return true
}

// CreateHardLink creates a hard link at the specified path pointing to the target.
// Both the target and link paths must be on the same file system.
//
// Parameters:
//   - target: The file that the hard link will refer to
//   - link: The path where the hard link will be created
//
// Returns:
//   - bool: true if the hard link was created successfully, false otherwise
//
// Example:
//
//	ok := ufs.CreateHardLink("/path/to/target", "/path/to/link")
//	if !ok {
//	    fmt.Printf("Error creating hard link\n")
//	}
func (ufs *UFS) CreateHardLink(target string, link string) bool {
	err := os.Link(target, link)
	if err != nil {
		ufs.handleError(err, "CreateHardLink")
		return false
	}
	return true
}

// CreateDirectoryTree creates a directory tree based on the provided structure.
// The structure is a map where keys are directory names and values are either
// nil (for empty directories) or nested maps (for subdirectories).
//
// Parameters:
//   - basePath: The base directory path where the tree will be created
//   - structure: A map representing the directory structure to create
//
// Returns:
//   - error: An error if any directory in the tree could not be created, nil otherwise
//
// Example:
//
//	structure := map[string]interface{}{
//	    "dir1": nil,
//	    "dir2": map[string]interface{}{
//	        "subdir1": nil,
//	        "subdir2": map[string]interface{}{
//	            "subsubdir": nil,
//	        },
//	    },
//	}
//	err := ufs.CreateDirectoryTree("/path/to/base", structure)
//	if err != nil {
//	    fmt.Printf("Error creating directory tree: %v\n", err)
//	}
func (ufs *UFS) CreateDirectoryTree(basePath string, structure map[string]interface{}) bool {
	// Create the base directory if it doesn't exist
	if !ufs.CreateDirectory(basePath) {
		return false
	}

	// Iterate through the structure and create subdirectories
	for dirName, subStructure := range structure {
		dirPath := filepath.Join(basePath, dirName)
		err := ufs.CreateDirectory(dirPath)
		if err {
			return false
		}

		// If subStructure is not nil, it's a nested directory structure
		if subStructure != nil {
			if subMap, ok := subStructure.(map[string]interface{}); ok {
				if !ufs.CreateDirectoryTree(dirPath, subMap) {
					return false
				}
			}
		}
	}

	return true
}

// CreateDirectoryTreeWithPermissions creates a directory tree with the specified permissions.
// The structure is a map where keys are directory names and values are either
// nil (for empty directories) or nested maps (for subdirectories).
//
// Parameters:
//   - basePath: The base directory path where the tree will be created
//   - structure: A map representing the directory structure to create
//   - perm: The permissions to apply to all directories in the tree
//
// Returns:
//   - error: An error if any directory in the tree could not be created, nil otherwise
//
// Example:
//
//	structure := map[string]interface{}{
//	    "dir1": nil,
//	    "dir2": map[string]interface{}{
//	        "subdir1": nil,
//	    },
//	}
//	err := ufs.CreateDirectoryTreeWithPermissions("/path/to/base", structure, 0755)
//	if err != nil {
//	    fmt.Printf("Error creating directory tree with permissions: %v\n", err)
//	}
func (ufs *UFS) CreateDirectoryTreeWithPermissions(basePath string, structure map[string]interface{}, perm fs.FileMode) bool {
	// Create the base directory if it doesn't exist
	ok := ufs.CreateDirectoryWithPermissions(basePath, perm)
	if !ok {
		return false
	}

	// Iterate through the structure and create subdirectories
	for dirName, subStructure := range structure {
		dirPath := filepath.Join(basePath, dirName)
		ok := ufs.CreateDirectoryWithPermissions(dirPath, perm)
		if !ok {
			return false
		}

		// If subStructure is not nil, it's a nested directory structure
		if subStructure != nil {
			if subMap, ok := subStructure.(map[string]interface{}); ok {
				if !ufs.CreateDirectoryTreeWithPermissions(dirPath, subMap, perm) {
					return false
				}
			}
		}
	}

	return true
}

// SymlinkDirectoryTree creates symbolic links for an entire directory tree.
// This function walks through the source directory tree and creates corresponding
// symbolic links in the destination directory.
//
// Parameters:
//   - sourceDir: The source directory tree to be symlinked
//   - destDir: The destination directory where symlinks will be created
//   - recursive: If true, symlinks subdirectories recursively; otherwise, only
//     symlinks files in the top-level directory
//
// Returns:
//   - error: An error if any symlink could not be created, nil otherwise
//
// Example:
//
//	err := ufs.SymlinkDirectoryTree("/path/to/source", "/path/to/dest", true)
//	if err != nil {
//	    fmt.Printf("Error symlinking directory tree: %v\n", err)
//	}
func (ufs *UFS) SymlinkDirectoryTree(sourceDir string, destDir string, recursive bool) bool {
	// Ensure the source directory exists
	if !ufs.IsDirectory(sourceDir) {
		return false
	}

	// Create the destination directory if it doesn't exist
	ok := ufs.CreateDirectory(destDir)
	if !ok {
		return false
	}

	// Read the source directory
	entries, err := os.ReadDir(sourceDir)
	if err != nil {
		ufs.handleError(err, "SymlinkDirectoryTree")
		return false
	}

	// Process each entry in the source directory
	for _, entry := range entries {
		sourcePath := filepath.Join(sourceDir, entry.Name())
		destPath := filepath.Join(destDir, entry.Name())

		if entry.IsDir() {
			if recursive {
				// If recursive is true, process subdirectories
				if !ufs.SymlinkDirectoryTree(sourcePath, destPath, true) {
					return false
				}
			}
		} else {
			// Create a symlink for the file
			ok := ufs.CreateSymlink(sourcePath, destPath)
			if !ok {
				return false
			}
		}
	}

	return true
}
