package ufs

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

/*
File Reader/Writer functions
These functions are used to read and write files in the UFS (Universal File System) package.

This is only file in ufs package that will send error too where other files will not send error only boolean or other types.
We knowingly decided to send error because file operations are critical and often require error handling.
This file provides basic file operations such as reading, writing, appending, copying, moving, and deleting files.

Provided functions include:
- ReadFile: Reads the content of a file and returns it as a byte slice.
- WriteFile: Writes data to a file, creating it if it doesn't exist or overwriting it if it does.
- AppendToFile: Appends data to a file, creating it if it doesn't exist.
- CopyFile: Copies the content of one file to another.
- MoveFile: Moves a file from one location to another.
- DeleteFile: Deletes a specified file.

// These functions handle file operations with internal error handling and logging via handleError function already implemented.

Other utilities:
- ReadFileAsString: Reads the content of a file and returns it as a string.
- WriteStringToFile: Writes a string to a file, creating it if it doesn't exist or overwriting it if it does.
- AppendStringToFile: Appends a string to a file, creating it if it doesn't exist.

// - CopyFileWithPermissions: Copies a file to a new location, preserving its permissions.
- MoveFileWithPermissions: Moves a file to a new location, preserving its permissions.
// - DeleteFileWithPermissions: Deletes a file, preserving its permissions.

Advanced utilities includes:
- AssembleFiles : Combines multiple files into a single file in order of slice.,
- SplitFile : Splits a file into multiple files based on a specified size limit.
- CleanUpFiles : Cleans up files by removing empty files given in a slice.
- ReadFileWithLines : Reads a file and returns its content as a slice of strings, each representing a line in the file.
- AppendToLastLine : Appends a string to the last line of a file, creating the file if it doesn't exist. if file has 14 lines, it will append to 15th line. wont append to 14th line (same line).
- AppendToFirstLine : Appends a string to the first line of a file, creating the file if it doesn't exist. it will gracefully shift current first line to second line and append to first line.
*/

// ReadFile reads the content of a file and returns it as a byte slice.
// This function will read the entire content of the file into memory.
//
// Parameters:
//   - path: The absolute or relative path to the file to read
//
// Returns:
//   - []byte: The content of the file as a byte slice
//   - error: An error if the file couldn't be read or doesn't exist
//
// Example:
//
//	data, err := ufs.ReadFile("/path/to/file.txt")
//	if err != nil {
//	    fmt.Printf("Error reading file: %v\n", err)
//	    return
//	}
//	fmt.Printf("File content: %s\n", data)
func (ufs *UFS) ReadFile(path string) ([]byte, error) {
	if !ufs.IsFile(path) {
		return nil, fmt.Errorf("path is not a file: %s", path)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, ufs.wrapError(err, "ReadFile")
	}
	return data, nil
}

// ReadFileAsString reads the content of a file and returns it as a string.
// This function will read the entire content of the file into memory.
//
// Parameters:
//   - path: The absolute or relative path to the file to read
//
// Returns:
//   - string: The content of the file as a string
//   - error: An error if the file couldn't be read or doesn't exist
//
// Example:
//
//	content, err := ufs.ReadFileAsString("/path/to/file.txt")
//	if err != nil {
//	    fmt.Printf("Error reading file: %v\n", err)
//	    return
//	}
//	fmt.Printf("File content: %s\n", content)
func (ufs *UFS) ReadFileAsString(path string) (string, error) {
	data, err := ufs.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// WriteFile writes data to a file, creating it if it doesn't exist or overwriting it if it does.
// This function will create any parent directories if they don't exist.
//
// Parameters:
//   - path: The absolute or relative path to the file to write
//   - data: The data to write to the file as a byte slice
//
// Returns:
//   - error: An error if the file couldn't be written to
//
// Example:
//
//	err := ufs.WriteFile("/path/to/file.txt", []byte("Hello, World!"))
//	if err != nil {
//	    fmt.Printf("Error writing to file: %v\n", err)
//	    return
//	}
//	fmt.Println("File written successfully")
func (ufs *UFS) WriteFile(path string, data []byte) error {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if !ufs.IsDirectory(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return ufs.wrapError(err, "WriteFile")
		}
	}

	err := os.WriteFile(path, data, 0644)
	if err != nil {
		return ufs.wrapError(err, "WriteFile")
	}
	return nil
}

// WriteStringToFile writes a string to a file, creating it if it doesn't exist or overwriting it if it does.
// This function will create any parent directories if they don't exist.
//
// Parameters:
//   - path: The absolute or relative path to the file to write
//   - content: The string content to write to the file
//
// Returns:
//   - error: An error if the file couldn't be written to
//
// Example:
//
//	err := ufs.WriteStringToFile("/path/to/file.txt", "Hello, World!")
//	if err != nil {
//	    fmt.Printf("Error writing to file: %v\n", err)
//	    return
//	}
//	fmt.Println("File written successfully")
func (ufs *UFS) WriteStringToFile(path string, content string) error {
	return ufs.WriteFile(path, []byte(content))
}

// AppendToFile appends data to a file, creating it if it doesn't exist.
// This function will create any parent directories if they don't exist.
//
// Parameters:
//   - path: The absolute or relative path to the file to append to
//   - data: The data to append to the file as a byte slice
//
// Returns:
//   - error: An error if the file couldn't be appended to
//
// Example:
//
//	err := ufs.AppendToFile("/path/to/file.txt", []byte("More data to append"))
//	if err != nil {
//	    fmt.Printf("Error appending to file: %v\n", err)
//	    return
//	}
//	fmt.Println("Data appended to file successfully")
func (ufs *UFS) AppendToFile(path string, data []byte) error {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if !ufs.IsDirectory(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return ufs.wrapError(err, "AppendToFile")
		}
	}

	// Open file in append mode
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return ufs.wrapError(err, "AppendToFile")
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return ufs.wrapError(err, "AppendToFile")
	}
	return nil
}

// AppendStringToFile appends a string to a file, creating it if it doesn't exist.
// This function will create any parent directories if they don't exist.
//
// Parameters:
//   - path: The absolute or relative path to the file to append to
//   - content: The string content to append to the file
//
// Returns:
//   - error: An error if the file couldn't be appended to
//
// Example:
//
//	err := ufs.AppendStringToFile("/path/to/file.txt", "More text to append")
//	if err != nil {
//	    fmt.Printf("Error appending to file: %v\n", err)
//	    return
//	}
//	fmt.Println("Text appended to file successfully")
func (ufs *UFS) AppendStringToFile(path string, content string) error {
	return ufs.AppendToFile(path, []byte(content))
}

// CopyFile copies the content of one file to another.
// If the destination file already exists, it will be overwritten.
// This function will create any parent directories for the destination if they don't exist.
//
// Parameters:
//   - src: The absolute or relative path to the source file
//   - dst: The absolute or relative path to the destination file
//
// Returns:
//   - error: An error if the file couldn't be copied
//
// Example:
//
//	err := ufs.CopyFile("/path/to/source.txt", "/path/to/destination.txt")
//	if err != nil {
//	    fmt.Printf("Error copying file: %v\n", err)
//	    return
//	}
//	fmt.Println("File copied successfully")
func (ufs *UFS) CopyFile(src, dst string) error {
	// Verify source is a file
	if !ufs.IsFile(src) {
		return fmt.Errorf("source is not a file: %s", src)
	}

	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if !ufs.IsDirectory(dstDir) {
		err := os.MkdirAll(dstDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "CopyFile")
		}
	}

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return ufs.wrapError(err, "CopyFile")
	}
	defer srcFile.Close()

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return ufs.wrapError(err, "CopyFile")
	}
	defer dstFile.Close()

	// Copy the contents
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return ufs.wrapError(err, "CopyFile")
	}

	return nil
}

// MoveFile moves a file from one location to another.
// If the destination file already exists, it will be overwritten.
// This function will create any parent directories for the destination if they don't exist.
//
// Parameters:
//   - src: The absolute or relative path to the source file
//   - dst: The absolute or relative path to the destination file
//
// Returns:
//   - error: An error if the file couldn't be moved
//
// Example:
//
//	err := ufs.MoveFile("/path/to/source.txt", "/path/to/destination.txt")
//	if err != nil {
//	    fmt.Printf("Error moving file: %v\n", err)
//	    return
//	}
//	fmt.Println("File moved successfully")
func (ufs *UFS) MoveFile(src, dst string) error {
	// Verify source is a file
	if !ufs.IsFile(src) {
		return fmt.Errorf("source is not a file: %s", src)
	}

	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if !ufs.IsDirectory(dstDir) {
		err := os.MkdirAll(dstDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "MoveFile")
		}
	}

	// Try to rename the file (only works on same file system)
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}

	// If rename fails, try copy and delete
	err = ufs.CopyFile(src, dst)
	if err != nil {
		return err
	}

	// Delete the source file
	err = os.Remove(src)
	if err != nil {
		return ufs.wrapError(err, "MoveFile")
	}

	return nil
}

// DeleteFile deletes a specified file.
// This is a wrapper around os.Remove that adds error handling.
//
// Parameters:
//   - path: The absolute or relative path to the file to delete
//
// Returns:
//   - error: An error if the file couldn't be deleted
//
// Example:
//
//	err := ufs.DeleteFile("/path/to/file.txt")
//	if err != nil {
//	    fmt.Printf("Error deleting file: %v\n", err)
//	    return
//	}
//	fmt.Println("File deleted successfully")
func (ufs *UFS) DeleteFile(path string) error {
	// Verify it's a file
	if !ufs.IsFile(path) {
		return fmt.Errorf("path is not a file: %s", path)
	}

	err := os.Remove(path)
	if err != nil {
		return ufs.wrapError(err, "DeleteFile")
	}
	return nil
}

// CopyFileWithPermissions copies a file to a new location, preserving its permissions.
// If the destination file already exists, it will be overwritten.
// This function will create any parent directories for the destination if they don't exist.
//
// Parameters:
//   - src: The absolute or relative path to the source file
//   - dst: The absolute or relative path to the destination file
//
// Returns:
//   - error: An error if the file couldn't be copied or permissions couldn't be preserved
//
// Example:
//
//	err := ufs.CopyFileWithPermissions("/path/to/source.txt", "/path/to/destination.txt")
//	if err != nil {
//	    fmt.Printf("Error copying file with permissions: %v\n", err)
//	    return
//	}
//	fmt.Println("File copied with permissions successfully")
func (ufs *UFS) CopyFileWithPermissions(src, dst string) error {
	// Verify source is a file
	if !ufs.IsFile(src) {
		return fmt.Errorf("source is not a file: %s", src)
	}

	// Get source file info for permissions
	srcInfo, err := os.Stat(src)
	if err != nil {
		return ufs.wrapError(err, "CopyFileWithPermissions")
	}

	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if !ufs.IsDirectory(dstDir) {
		err := os.MkdirAll(dstDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "CopyFileWithPermissions")
		}
	}

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return ufs.wrapError(err, "CopyFileWithPermissions")
	}
	defer srcFile.Close()

	// Create destination file with same permissions
	dstFile, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, srcInfo.Mode())
	if err != nil {
		return ufs.wrapError(err, "CopyFileWithPermissions")
	}
	defer dstFile.Close()

	// Copy the contents
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return ufs.wrapError(err, "CopyFileWithPermissions")
	}

	return nil
}

// MoveFileWithPermissions moves a file to a new location, preserving its permissions.
// If the destination file already exists, it will be overwritten.
// This function will create any parent directories for the destination if they don't exist.
//
// Parameters:
//   - src: The absolute or relative path to the source file
//   - dst: The absolute or relative path to the destination file
//
// Returns:
//   - error: An error if the file couldn't be moved or permissions couldn't be preserved
//
// Example:
//
//	err := ufs.MoveFileWithPermissions("/path/to/source.txt", "/path/to/destination.txt")
//	if err != nil {
//	    fmt.Printf("Error moving file with permissions: %v\n", err)
//	    return
//	}
//	fmt.Println("File moved with permissions successfully")
func (ufs *UFS) MoveFileWithPermissions(src, dst string) error {
	// Verify source is a file
	if !ufs.IsFile(src) {
		return fmt.Errorf("source is not a file: %s", src)
	}

	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if !ufs.IsDirectory(dstDir) {
		err := os.MkdirAll(dstDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "MoveFileWithPermissions")
		}
	}

	// Try to rename the file (only works on same file system)
	err := os.Rename(src, dst)
	if err == nil {
		return nil
	}

	// If rename fails, try copy with permissions and delete
	err = ufs.CopyFileWithPermissions(src, dst)
	if err != nil {
		return err
	}

	// Delete the source file
	err = os.Remove(src)
	if err != nil {
		return ufs.wrapError(err, "MoveFileWithPermissions")
	}

	return nil
}

// AssembleFiles combines multiple files into a single file in the order of the provided slice.
// This function will create the destination file if it doesn't exist or overwrite it if it does.
//
// Parameters:
//   - srcFiles: A slice of file paths to be combined
//   - dst: The path to the destination file
//
// Returns:
//   - error: An error if the files couldn't be combined
//
// Example:
//
//	files := []string{"/path/to/part1.txt", "/path/to/part2.txt", "/path/to/part3.txt"}
//	err := ufs.AssembleFiles(files, "/path/to/combined.txt")
//	if err != nil {
//	    fmt.Printf("Error combining files: %v\n", err)
//	    return
//	}
//	fmt.Println("Files combined successfully")
func (ufs *UFS) AssembleFiles(srcFiles []string, dst string) error {
	// Ensure all source files exist
	for _, src := range srcFiles {
		if !ufs.IsFile(src) {
			return fmt.Errorf("source file does not exist: %s", src)
		}
	}

	// Ensure the destination directory exists
	dstDir := filepath.Dir(dst)
	if !ufs.IsDirectory(dstDir) {
		err := os.MkdirAll(dstDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "AssembleFiles")
		}
	}

	// Create destination file
	dstFile, err := os.Create(dst)
	if err != nil {
		return ufs.wrapError(err, "AssembleFiles")
	}
	defer dstFile.Close()

	// Combine files
	for _, src := range srcFiles {
		srcFile, err := os.Open(src)
		if err != nil {
			return ufs.wrapError(err, "AssembleFiles")
		}

		_, err = io.Copy(dstFile, srcFile)
		srcFile.Close()
		if err != nil {
			return ufs.wrapError(err, "AssembleFiles")
		}
	}

	return nil
}

// SplitFile splits a file into multiple files based on a specified size limit.
// This function will create the split files in the same directory as the original with suffixes _1, _2, etc.
//
// Parameters:
//   - src: The path to the source file to split
//   - chunkSize: The maximum size in bytes of each split file
//
// Returns:
//   - []string: A slice of paths to the created split files
//   - error: An error if the file couldn't be split
//
// Example:
//
//	splitFiles, err := ufs.SplitFile("/path/to/large_file.dat", 1024*1024) // Split into 1MB chunks
//	if err != nil {
//	    fmt.Printf("Error splitting file: %v\n", err)
//	    return
//	}
//	fmt.Printf("File split into %d parts\n", len(splitFiles))
//	for i, file := range splitFiles {
//	    fmt.Printf("Part %d: %s\n", i+1, file)
//	}
func (ufs *UFS) SplitFile(src string, chunkSize int64) ([]string, error) {
	// Verify source is a file
	if !ufs.IsFile(src) {
		return nil, fmt.Errorf("source is not a file: %s", src)
	}

	// Open source file
	srcFile, err := os.Open(src)
	if err != nil {
		return nil, ufs.wrapError(err, "SplitFile")
	}
	defer srcFile.Close()

	// Get file info
	srcInfo, err := srcFile.Stat()
	if err != nil {
		return nil, ufs.wrapError(err, "SplitFile")
	}

	// Calculate number of parts
	fileSize := srcInfo.Size()
	numParts := (fileSize + chunkSize - 1) / chunkSize // Round up division

	if numParts == 0 {
		return nil, fmt.Errorf("file is empty, nothing to split: %s", src)
	}

	// Generate split file paths
	baseDir := filepath.Dir(src)
	baseExt := filepath.Ext(src)
	baseName := filepath.Base(src)
	baseName = strings.TrimSuffix(baseName, baseExt)

	splitFiles := make([]string, numParts)
	for i := range splitFiles {
		splitFiles[i] = filepath.Join(baseDir, fmt.Sprintf("%s_%d%s", baseName, i+1, baseExt))
	}

	// Split the file
	buffer := make([]byte, 4096) // 4KB read buffer
	for i := int64(0); i < numParts; i++ {
		// Create part file
		partFile, err := os.Create(splitFiles[i])
		if err != nil {
			return splitFiles[:i], ufs.wrapError(err, "SplitFile")
		}

		// Write chunk
		bytesWritten := int64(0)
		for bytesWritten < chunkSize {
			bytesToRead := chunkSize - bytesWritten
			if bytesToRead > int64(len(buffer)) {
				bytesToRead = int64(len(buffer))
			}

			n, err := srcFile.Read(buffer[:bytesToRead])
			if err != nil && err != io.EOF {
				partFile.Close()
				return splitFiles[:i+1], ufs.wrapError(err, "SplitFile")
			}

			if n == 0 {
				break // End of file
			}

			_, err = partFile.Write(buffer[:n])
			if err != nil {
				partFile.Close()
				return splitFiles[:i+1], ufs.wrapError(err, "SplitFile")
			}

			bytesWritten += int64(n)
			if err == io.EOF {
				break
			}
		}

		partFile.Close()
	}

	return splitFiles, nil
}

// CleanUpFiles removes empty files from the given slice of file paths.
// This function is useful for removing temporary or empty files after processing.
//
// Parameters:
//   - files: A slice of file paths to check and potentially remove
//
// Returns:
//   - []string: A slice of paths to files that were removed
//   - error: An error if any file couldn't be checked or removed
//
// Example:
//
//	files := []string{"/path/to/file1.txt", "/path/to/file2.txt", "/path/to/file3.txt"}
//	removedFiles, err := ufs.CleanUpFiles(files)
//	if err != nil {
//	    fmt.Printf("Error cleaning up files: %v\n", err)
//	    return
//	}
//	fmt.Printf("Removed %d empty files\n", len(removedFiles))
func (ufs *UFS) CleanUpFiles(files []string) ([]string, error) {
	var removedFiles []string
	var lastError error

	for _, file := range files {
		if !ufs.IsFile(file) {
			continue
		}

		// Check if file is empty
		if ufs.IsFileEmpty(file) {
			err := os.Remove(file)
			if err != nil {
				lastError = ufs.wrapError(err, "CleanUpFiles")
				continue
			}
			removedFiles = append(removedFiles, file)
		}
	}

	if lastError != nil {
		return removedFiles, lastError
	}
	return removedFiles, nil
}

// ReadFileWithLines reads a file and returns its content as a slice of strings,
// each representing a line in the file.
//
// Parameters:
//   - path: The path to the file to read
//
// Returns:
//   - []string: A slice of strings, each representing a line in the file
//   - error: An error if the file couldn't be read
//
// Example:
//
//	lines, err := ufs.ReadFileWithLines("/path/to/file.txt")
//	if err != nil {
//	    fmt.Printf("Error reading file: %v\n", err)
//	    return
//	}
//	for i, line := range lines {
//	    fmt.Printf("Line %d: %s\n", i+1, line)
//	}
func (ufs *UFS) ReadFileWithLines(path string) ([]string, error) {
	// Verify source is a file
	if !ufs.IsFile(path) {
		return nil, fmt.Errorf("path is not a file: %s", path)
	}

	// Open file
	file, err := os.Open(path)
	if err != nil {
		return nil, ufs.wrapError(err, "ReadFileWithLines")
	}
	defer file.Close()

	// Read lines
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return lines, ufs.wrapError(err, "ReadFileWithLines")
	}

	return lines, nil
}

// AppendToLastLine appends a string to the last line of a file,
// or creates a new line if the file is empty or doesn't exist.
// If the file has 14 lines, it will append to the 15th line (add a new line).
//
// Parameters:
//   - path: The path to the file
//   - content: The string to append
//
// Returns:
//   - error: An error if the file couldn't be read or written to
//
// Example:
//
//	err := ufs.AppendToLastLine("/path/to/file.txt", "New line content")
//	if err != nil {
//	    fmt.Printf("Error appending to last line: %v\n", err)
//	    return
//	}
//	fmt.Println("Content appended to last line successfully")
func (ufs *UFS) AppendToLastLine(path string, content string) error {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if !ufs.IsDirectory(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return ufs.wrapError(err, "AppendToLastLine")
		}
	}

	// If file doesn't exist, create it with the content
	if !ufs.IsFile(path) {
		return ufs.WriteStringToFile(path, content)
	}

	// Append a newline and the content
	return ufs.AppendStringToFile(path, "\n"+content)
}

// AppendToFirstLine appends a string to the first line of a file,
// gracefully shifting the current first line to the second line.
// If the file doesn't exist, it will be created with the content.
//
// Parameters:
//   - path: The path to the file
//   - content: The string to add as the new first line
//
// Returns:
//   - error: An error if the file couldn't be read or written to
//
// Example:
//
//	err := ufs.AppendToFirstLine("/path/to/file.txt", "New first line")
//	if err != nil {
//	    fmt.Printf("Error appending to first line: %v\n", err)
//	    return
//	}
//	fmt.Println("Content added as first line successfully")
func (ufs *UFS) AppendToFirstLine(path string, content string) error {
	// Ensure the directory exists
	dir := filepath.Dir(path)
	if !ufs.IsDirectory(dir) {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			return ufs.wrapError(err, "AppendToFirstLine")
		}
	}

	// If file doesn't exist, create it with the content
	if !ufs.IsFile(path) {
		return ufs.WriteStringToFile(path, content)
	}

	// Read existing content
	lines, err := ufs.ReadFileWithLines(path)
	if err != nil {
		return err
	}

	// Insert the new content at the beginning
	newLines := append([]string{content}, lines...)

	// Join lines with newlines and write back to file
	newContent := strings.Join(newLines, "\n")

	return ufs.WriteStringToFile(path, newContent)
}

