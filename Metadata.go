package ufs

import (
	"os"
	"path/filepath"
)

// GetFileSize returns the size of the given file in bytes.
// This function checks if the path exists and is a file before retrieving its size.
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - int64: The size of the file in bytes
//   - Returns 0 if the file doesn't exist, is a directory, or if an error occurs
//
// Example:
//
//	size := ufs.GetFileSize("/path/to/file.txt")
//	fmt.Printf("File size: %d bytes\n", size)
func (ufs *UFS) GetFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "GetFileSize")
		return 0
	}
	if info.IsDir() {
		ufs.handleMistakeWarning("GetFileSize called on a directory, returning 0")
		return 0
	}
	return info.Size()
}

// GetFileMetadata retrieves basic metadata for a file at the specified path.
// This function collects essential information about a file including name, size,
// permissions, modification time, and whether it's a directory.
//
// Parameters:
//   - path: The absolute or relative path to the file
//
// Returns:
//   - map[string]interface{}: A map containing the file's metadata with the following keys:
//   - Name: The base name of the file
//   - Size: The size of the file in bytes
//   - Mode: The file's permission bits as a string
//   - ModTime: The last modification time
//   - IsDir: Boolean indicating if the path is a directory
//   - Returns nil if the file doesn't exist or if an error occurs
//
// Example:
//
//	metadata := ufs.GetFileMetadata("/path/to/file.txt")
//	fmt.Printf("File name: %s\n", metadata["Name"])
//	fmt.Printf("Last modified: %s\n", metadata["ModTime"])
func (ufs *UFS) GetFileMetadata(path string) map[string]interface{} {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "GetFileMetadata")
		return nil
	}
	metadata := map[string]interface{}{
		"Name":    info.Name(),
		"Size":    info.Size(),
		"Mode":    info.Mode().String(),
		"ModTime": info.ModTime(),
		"IsDir":   info.IsDir(),
	}

	return metadata
}

// GetFileList returns a list of file names under the given path (non-recursive).
// This function lists only files (not directories) in the specified directory.
//
// Parameters:
//   - path: The absolute or relative path to the directory to list files from
//
// Returns:
//   - []string: A slice containing the names of all files in the directory
//   - Returns an empty slice if the directory doesn't exist or if an error occurs
//
// Example:
//
//	files := ufs.GetFileList("/path/to/directory")
//	for _, file := range files {
//	    fmt.Printf("Found file: %s\n", file)
//	}
func (ufs *UFS) GetFileList(path string) []string {
	var files []string
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetFileList")
		return []string{}
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files
}

// GetFolderList returns a list of folder names under the given path.
// This function lists only directories (not files) in the specified directory.
//
// Parameters:
//   - path: The absolute or relative path to the directory to list folders from
//
// Returns:
//   - []string: A slice containing the names of all subdirectories in the directory
//   - Returns an empty slice if the directory doesn't exist or if an error occurs
//
// Example:
//
//	folders := ufs.GetFolderList("/path/to/directory")
//	for _, folder := range folders {
//	    fmt.Printf("Found subdirectory: %s\n", folder)
//	}
func (ufs *UFS) GetFolderList(path string) []string {
	var folders []string
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetFolderList")
		return []string{}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}
	return folders
}

// GetFolderFileCount returns the number of files (not directories) in the specified directory.
// This function counts only files at the top level and doesn't recurse into subdirectories.
//
// Parameters:
//   - path: The absolute or relative path to the directory to count files in
//
// Returns:
//   - int: The number of files in the directory
//   - Returns 0 if the directory doesn't exist or if an error occurs
//
// Example:
//
//	count := ufs.GetFolderFileCount("/path/to/directory")
//	fmt.Printf("Directory contains %d files\n", count)
func (ufs *UFS) GetFolderFileCount(path string) int {
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetFolderFileCount")
		return 0
	}
	count := 0
	for _, entry := range entries {
		if !entry.IsDir() {
			count++
		}
	}
	return count
}

// GetFolderChildCount returns the total number of entries (both files and directories)
// in the specified directory.
// This function counts all items at the top level without recursing into subdirectories.
//
// Parameters:
//   - path: The absolute or relative path to the directory to count children in
//
// Returns:
//   - int: The total number of files and directories in the directory
//   - Returns 0 if the directory doesn't exist or if an error occurs
//
// Example:
//
//	count := ufs.GetFolderChildCount("/path/to/directory")
//	fmt.Printf("Directory contains %d total items\n", count)
func (ufs *UFS) GetFolderChildCount(path string) int {
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetFolderChildCount")
		return 0
	}
	count := 0
	for range entries {
		count++
	}
	return count
}

// GetChildCount returns separate counts for the number of files and directories
// in the specified directory.
// This function counts items at the top level without recursing into subdirectories.
//
// Parameters:
//   - path: The absolute or relative path to the directory to count children in
//
// Returns:
//   - int: The number of directories (first return value)
//   - int: The number of files (second return value)
//   - Returns 0, 0 if the directory doesn't exist or if an error occurs
//
// Example:
//
//	folderCount, fileCount := ufs.GetChildCount("/path/to/directory")
//	fmt.Printf("Directory contains %d folders and %d files\n", folderCount, fileCount)
func (ufs *UFS) GetChildCount(path string) (int, int) {
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetChildCount")
		return 0, 0
	}
	folderCount := 0
	fileCount := 0
	for _, entry := range entries {
		if entry.IsDir() {
			folderCount++
		} else {
			fileCount++
		}
	}
	return folderCount, fileCount
}

// GetFolderMetadata retrieves detailed metadata for a folder at the specified path.
// This function collects information about a directory including its name, permissions,
// modification time, and indicates that it's a directory.
//
// Parameters:
//   - path: The absolute or relative path to the directory
//
// Returns:
//   - map[string]interface{}: A map containing the directory's metadata with the following keys:
//   - Name: The base name of the directory
//   - Size: The size of the directory entry itself (not its contents)
//   - Mode: The directory's permission bits as a string
//   - ModTime: The last modification time
//   - IsDir: Boolean indicating if the path is a directory (always true for this function)
//   - ChildNum: Placeholder for child count (defaults to 0)
//   - Returns nil if the directory doesn't exist or if an error occurs
//
// Example:
//
//	metadata := ufs.GetFolderMetadata("/path/to/directory")
//	fmt.Printf("Folder name: %s\n", metadata["Name"])
//	fmt.Printf("Last modified: %s\n", metadata["ModTime"])
func (ufs *UFS) GetFolderMetadata(path string) map[string]interface{} {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "GetFolderMetadata")
		return nil
	}
	metadata := map[string]interface{}{
		"Name":     info.Name(),
		"Size":     info.Size(),
		"Mode":     info.Mode().String(),
		"ModTime":  info.ModTime(),
		"IsDir":    info.IsDir(),
		"ChildNum": ufs.GetFolderChildCount(path), // Placeholder for child count
	}

	return metadata
}

// GetFolderSize recursively calculates the total size of a folder and all its contents.
// This function walks through the directory tree and sums the sizes of all files.
//
// Parameters:
//   - path: The absolute or relative path to the directory to calculate size for
//
// Returns:
//   - int64: The total size of all files in the directory tree in bytes
//   - Returns 0 if the directory doesn't exist or if an error occurs
//
// Example:
//
//	size := ufs.GetFolderSize("/path/to/directory")
//	fmt.Printf("Total folder size: %d bytes\n", size)
func (ufs *UFS) GetFolderSize(path string) int64 {
	var size int64
	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			ufs.handleError(err, "GetFolderSize")
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				ufs.handleError(err, "GetFolderSize")
				return nil
			}
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		ufs.handleError(err, "GetFolderSize")
		return 0
	}
	return size
}
