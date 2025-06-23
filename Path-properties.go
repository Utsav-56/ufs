package ufs

// PathDetails holds the details of a file or directory path.
/*
	Contains basic functions to know about the file or directory.
	contains pathExists,

	Includes methods like isFile, isDirectory, isDirectoryEmpty, isFileEmpty

	also some especial methods like
	isInSystemPath, isInUserPath, isInCurrentPath,

	for files especial it has methods like
	isFileHidden, isFileExecutable, isFileReadable, isFileWritable

	For directories it has methods like
	isDirectoryHidden, isDirectoryReadable, isDirectoryWritable

	These methods can be used to check the properties of a file or directory.

*/

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

// PathExists checks if a file or directory exists at the specified path.
//
// Parameters:
//   - path: The absolute or relative path to check
//
// Returns:
//   - bool: True if the path exists, false otherwise
//
// Example:
//
//	if ufs.PathExists("/path/to/check") {
//	    fmt.Println("Path exists!")
//	}
func (ufs *UFS) PathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
		ufs.handleError(err, "PathExists")
		return false
	}
	return true
}

// IsFile checks if the specified path points to a regular file.
//
// Parameters:
//   - path: The absolute or relative path to check
//
// Returns:
//   - bool: True if the path exists and is a regular file, false otherwise
//
// Example:
//
//	if ufs.IsFile("/path/to/check") {
//	    fmt.Println("This is a file!")
//	}
func (ufs *UFS) IsFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "IsFile")
		return false
	}
	return !info.IsDir()
}

// IsDirectory checks if the specified path points to a directory.
//
// Parameters:
//   - path: The absolute or relative path to check
//
// Returns:
//   - bool: True if the path exists and is a directory, false otherwise
//
// Example:
//
//	if ufs.IsDirectory("/path/to/check") {
//	    fmt.Println("This is a directory!")
//	}
func (ufs *UFS) IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "IsDirectory")
		return false
	}
	return info.IsDir()
}

// IsDirectoryEmpty checks if the specified directory is empty.
//
// Parameters:
//   - path: The absolute or relative path to the directory
//
// Returns:
//   - bool: True if the directory exists and is empty, false otherwise
//
// Example:
//
//	if ufs.IsDirectoryEmpty("/path/to/directory") {
//	    fmt.Println("The directory is empty!")
//	}
func (ufs *UFS) IsDirectoryEmpty(path string) bool {
	if !ufs.IsDirectory(path) {
		return false
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "IsDirectoryEmpty")
		return false
	}

	return len(entries) == 0
}

// IsFileEmpty checks if the specified file is empty (zero bytes).
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - bool: True if the file exists and is empty, false otherwise
//
// Example:
//
//	if ufs.IsFileEmpty("/path/to/file.txt") {
//	    fmt.Println("The file is empty!")
//	}
func (ufs *UFS) IsFileEmpty(path string) bool {
	if !ufs.IsFile(path) {
		return false
	}

	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "IsFileEmpty")
		return false
	}

	return info.Size() == 0
}

// IsInSystemPath checks if the specified path is in one of the system directories.
//
// Parameters:
//   - path: The absolute or relative path to check
//
// Returns:
//   - bool: True if the path is in a system directory, false otherwise
//
// Example:
//
//	if ufs.IsInSystemPath("/usr/bin/python") {
//	    fmt.Println("This is a system path!")
//	}
func (ufs *UFS) IsInSystemPath(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		ufs.handleError(err, "IsInSystemPath")
		return false
	}

	systemPaths := []string{}

	// Define system paths based on OS
	if runtime.GOOS == "windows" {
		systemPaths = []string{
			os.Getenv("WINDIR"),
			os.Getenv("SYSTEMROOT"),
			filepath.Join(os.Getenv("SYSTEMROOT"), "System32"),
			filepath.Join(os.Getenv("SYSTEMROOT"), "SysWOW64"),
			os.Getenv("PROGRAMFILES"),
			os.Getenv("PROGRAMFILES(X86)"),
		}
	} else {
		// Unix-like systems
		systemPaths = []string{
			"/bin", "/sbin", "/usr/bin", "/usr/sbin",
			"/usr/local/bin", "/usr/local/sbin",
			"/etc", "/var", "/lib", "/opt",
		}
	}

	// Check if the path starts with any system path
	for _, sysPath := range systemPaths {
		if sysPath != "" && strings.HasPrefix(absPath, sysPath) {
			return true
		}
	}

	return false
}

// IsInUserPath checks if the specified path is in the user's home directory.
//
// Parameters:
//   - path: The absolute or relative path to check
//
// Returns:
//   - bool: True if the path is in the user's home directory, false otherwise
//
// Example:
//
//	if ufs.IsInUserPath("~/Documents/file.txt") {
//	    fmt.Println("This is in the user's home directory!")
//	}
func (ufs *UFS) IsInUserPath(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		ufs.handleError(err, "IsInUserPath")
		return false
	}

	currentUser, err := user.Current()
	if err != nil {
		ufs.handleError(err, "IsInUserPath")
		return false
	}

	return strings.HasPrefix(absPath, currentUser.HomeDir)
}

// IsInCurrentPath checks if the specified path is in the current working directory.
//
// Parameters:
//   - path: The absolute or relative path to check
//
// Returns:
//   - bool: True if the path is in the current working directory, false otherwise
//
// Example:
//
//	if ufs.IsInCurrentPath("./data/file.txt") {
//	    fmt.Println("This is in the current working directory!")
//	}
func (ufs *UFS) IsInCurrentPath(path string) bool {
	absPath, err := filepath.Abs(path)
	if err != nil {
		ufs.handleError(err, "IsInCurrentPath")
		return false
	}

	cwd, err := os.Getwd()
	if err != nil {
		ufs.handleError(err, "IsInCurrentPath")
		return false
	}

	return strings.HasPrefix(absPath, cwd)
}

// IsFileHidden checks if a file is hidden according to the OS conventions.
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - bool: True if the file exists and is hidden, false otherwise
//
// Example:
//
//	if ufs.IsFileHidden("/path/to/.hidden_file") {
//	    fmt.Println("This is a hidden file!")
//	}
func (ufs *UFS) IsFileHidden(path string) bool {
	if !ufs.IsFile(path) {
		return false
	}

	// Get the base name of the file
	baseName := filepath.Base(path)

	// On Unix-like systems, files starting with a dot are hidden
	if runtime.GOOS != "windows" {
		return strings.HasPrefix(baseName, ".")
	}

	// On Windows, use file attributes
	fileInfo, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "IsFileHidden")
		return false
	}

	// Check if the file has the hidden attribute (Windows only)
	// The hidden attribute is represented by the constant 0x2 (FILE_ATTRIBUTE_HIDDEN)
	if runtime.GOOS == "windows" {
		attributes := fileInfo.Sys().(*syscall.Win32FileAttributeData).FileAttributes
		return attributes&0x2 != 0
	}

	return false
}

// IsFileExecutable checks if a file is executable by the current user.
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - bool: True if the file exists and is executable, false otherwise
//
// Example:
//
//	if ufs.IsFileExecutable("/path/to/script.sh") {
//	    fmt.Println("This file is executable!")
//	}
func (ufs *UFS) IsFileExecutable(path string) bool {
	if !ufs.IsFile(path) {
		return false
	}

	// On Windows, executable status is determined by file extension
	if runtime.GOOS == "windows" {
		ext := strings.ToLower(filepath.Ext(path))
		execExts := []string{".exe", ".bat", ".cmd", ".com", ".ps1"}
		for _, execExt := range execExts {
			if ext == execExt {
				return true
			}
		}
		return false
	}

	// On Unix-like systems, check execution permission
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "IsFileExecutable")
		return false
	}

	return info.Mode()&0111 != 0
}

// IsFileReadable checks if a file is readable by the current user.
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - bool: True if the file exists and is readable, false otherwise
//
// Example:
//
//	if ufs.IsFileReadable("/path/to/file.txt") {
//	    fmt.Println("This file is readable!")
//	}
func (ufs *UFS) IsFileReadable(path string) bool {
	if !ufs.IsFile(path) {
		return false
	}

	// Try to open the file for reading
	file, err := os.Open(path)
	if err != nil {
		ufs.handleError(err, "IsFileReadable")
		return false
	}
	defer file.Close()

	return true
}

// IsFileWritable checks if a file is writable by the current user.
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - bool: True if the file exists and is writable, false otherwise
//
// Example:
//
//	if ufs.IsFileWritable("/path/to/file.txt") {
//	    fmt.Println("This file is writable!")
//	}
func (ufs *UFS) IsFileWritable(path string) bool {
	if !ufs.IsFile(path) {
		return false
	}

	// Try to open the file for writing (append mode to avoid destroying content)
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		if errors.Is(err, os.ErrPermission) {
			return false
		}
		ufs.handleError(err, "IsFileWritable")
		return false
	}
	defer file.Close()

	return true
}

// IsDirectoryHidden checks if a directory is hidden according to the OS conventions.
//
// Parameters:
//   - path: The absolute or relative path to the directory
//
// Returns:
//   - bool: True if the directory exists and is hidden, false otherwise
//
// Example:
//
//	if ufs.IsDirectoryHidden("/path/to/.hidden_dir") {
//	    fmt.Println("This is a hidden directory!")
//	}
func (ufs *UFS) IsDirectoryHidden(path string) bool {
	if !ufs.IsDirectory(path) {
		return false
	}

	// Get the base name of the directory
	baseName := filepath.Base(path)

	// On Unix-like systems, directories starting with a dot are hidden
	if runtime.GOOS != "windows" {
		return strings.HasPrefix(baseName, ".")
	}

	// On Windows, use file attributes
	fileInfo, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "IsDirectoryHidden")
		return false
	}

	// Check if the directory has the hidden attribute (Windows only)
	// The hidden attribute is represented by the constant 0x2 (FILE_ATTRIBUTE_HIDDEN)
	if runtime.GOOS == "windows" {
		attributes := fileInfo.Sys().(*syscall.Win32FileAttributeData).FileAttributes
		return attributes&0x2 != 0
	}

	return false
}

// IsDirectoryReadable checks if a directory is readable by the current user.
//
// Parameters:
//   - path: The absolute or relative path to the directory
//
// Returns:
//   - bool: True if the directory exists and is readable, false otherwise
//
// Example:
//
//	if ufs.IsDirectoryReadable("/path/to/directory") {
//	    fmt.Println("This directory is readable!")
//	}
func (ufs *UFS) IsDirectoryReadable(path string) bool {
	if !ufs.IsDirectory(path) {
		return false
	}

	// Try to read the directory entries
	_, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "IsDirectoryReadable")
		return false
	}

	return true
}
