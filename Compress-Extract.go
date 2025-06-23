package ufs

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

/*
Compress-Extract.go contains functions for compressing and extracting files and directories.
These functions allow you to create compressed archives (like ZIP files) and extract their contents.

This will use system commands to perform compression and extraction.
Like the `tar` command on Unix-like systems
As windows 10 and later have built-in support for ZIP files and tar.exe we will use same command,

but for older versions of Windows, We are not currently supporting compression and extraction.

Note [For contributors]: both unix and windows systems may have same name of tar but the commands are way too different.
So, please be careful while writing code for compression and extraction.

This file is part of the UFS (Universal File System) library, which provides a unified interface
for file and directory operations across different platforms.

Basic Functions:
- CompressDirectory: Compresses a directory into a ZIP file.
- ExtractArchive: Extracts the contents of a ZIP file to a specified directory.
- CompressFile: Compresses a single file into a ZIP file.

Some utilities uses basic functions internally:
- CompressHere: Compresses the  directory into a ZIP file and outputs in cwd.
- ExtractHere: Extracts the contents of a ZIP file in the current working directory.
- CompressFileHere: Compresses a single file into a ZIP file and outputs in cwd.

Other utilities (Just for demonstration, not recommended for production use) all are mad codes:
- CompressAndRemove: [Dangerous] Compresses a directory and removes the original directory.
- ExtractAndRemove: [Dangerous] Extracts a ZIP file and removes the original ZIP file.
- CompressAndExtract: [Dangerous] Compresses a directory and extracts it to a specified location.
- ExtractAndCompress: [Dangerous] Extracts a ZIP file and compresses it to a specified location.
*/

// CompressDirectory compresses a directory into a ZIP file.
// This function will create a ZIP archive containing all files and subdirectories.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the directory to compress
//   - destPath: The absolute or relative path where the ZIP file will be created
//
// Returns:
//   - error: An error if the compression failed, nil otherwise
//
// Example:
//
//	err := ufs.CompressDirectory("/path/to/source_dir", "/path/to/archive.zip")
//	if err != nil {
//	    fmt.Printf("Error compressing directory: %v\n", err)
//	    return
//	}
//	fmt.Println("Directory compressed successfully")
func (ufs *UFS) CompressDirectory(sourcePath, destPath string) error {
	// Verify source is a directory
	if !ufs.IsDirectory(sourcePath) {
		return fmt.Errorf("source path is not a directory: %s", sourcePath)
	}

	// Get absolute paths to ensure consistent behavior
	sourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "CompressDirectory")
	}

	destPath, err = filepath.Abs(destPath)
	if err != nil {
		return ufs.wrapError(err, "CompressDirectory")
	}

	// Ensure destination directory exists
	destDir := filepath.Dir(destPath)
	if !ufs.IsDirectory(destDir) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "CompressDirectory")
		}
	}

	// Create zip file
	zipFile, err := os.Create(destPath)
	if err != nil {
		return ufs.wrapError(err, "CompressDirectory")
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Walk the directory and add files to the zip
	err = filepath.Walk(sourcePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the root directory itself
		if path == sourcePath {
			return nil
		}

		// Prevent compressing the destination zip itself
		if path == destPath {
			return nil
		}

		// Create a zip header
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}

		// Update the name to preserve directory structure
		relPath, err := filepath.Rel(sourcePath, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Set compression method
		header.Method = zip.Deflate

		// Create writer for the file header
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// If it's a directory, we're done
		if info.IsDir() {
			return nil
		}

		// Open the file for reading
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy file contents to the zip
		_, err = io.Copy(writer, file)
		return err
	})

	if err != nil {
		return ufs.wrapError(err, "CompressDirectory")
	}

	return nil
}

// ExtractArchive extracts the contents of a ZIP file to a specified directory.
// This function will create the destination directory if it doesn't exist.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the ZIP file
//   - destPath: The absolute or relative path where the contents will be extracted
//
// Returns:
//   - error: An error if the extraction failed, nil otherwise
//
// Example:
//
//	err := ufs.ExtractArchive("/path/to/archive.zip", "/path/to/extract_dir")
//	if err != nil {
//	    fmt.Printf("Error extracting archive: %v\n", err)
//	    return
//	}
//	fmt.Println("Archive extracted successfully")
func (ufs *UFS) ExtractArchive(sourcePath, destPath string) error {
	// Verify source is a file
	if !ufs.IsFile(sourcePath) {
		return fmt.Errorf("source path is not a file: %s", sourcePath)
	}

	// Get absolute paths to ensure consistent behavior
	sourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "ExtractArchive")
	}

	destPath, err = filepath.Abs(destPath)
	if err != nil {
		return ufs.wrapError(err, "ExtractArchive")
	}

	// Ensure destination directory exists
	if !ufs.IsDirectory(destPath) {
		err = os.MkdirAll(destPath, 0755)
		if err != nil {
			return ufs.wrapError(err, "ExtractArchive")
		}
	}

	// Open the zip file
	reader, err := zip.OpenReader(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "ExtractArchive")
	}
	defer reader.Close()

	// Extract each file
	for _, file := range reader.File {
		err := ufs.extractZipFile(file, destPath)
		if err != nil {
			return ufs.wrapError(err, "ExtractArchive")
		}
	}

	return nil
}

// extractZipFile is a helper function to extract a single file from a zip archive
func (ufs *UFS) extractZipFile(file *zip.File, destPath string) error {
	// Form the full path to the file
	filePath := filepath.Join(destPath, file.Name)

	// Check for zip slip vulnerability
	if !strings.HasPrefix(filePath, filepath.Clean(destPath)+string(os.PathSeparator)) {
		return fmt.Errorf("illegal file path: %s", filePath)
	}

	// If it's a directory, create it
	if file.FileInfo().IsDir() {
		err := os.MkdirAll(filePath, file.Mode())
		if err != nil {
			return err
		}
		return nil
	}

	// Ensure the parent directory exists
	err := os.MkdirAll(filepath.Dir(filePath), 0755)
	if err != nil {
		return err
	}

	// Create the file
	destFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Open the file from the zip
	zipFile, err := file.Open()
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// Copy the contents
	_, err = io.Copy(destFile, zipFile)
	return err
}

// CompressFile compresses a single file into a ZIP file.
// This function will create a ZIP archive containing just the specified file.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the file to compress
//   - destPath: The absolute or relative path where the ZIP file will be created
//
// Returns:
//   - error: An error if the compression failed, nil otherwise
//
// Example:
//
//	err := ufs.CompressFile("/path/to/myfile.txt", "/path/to/myfile.zip")
//	if err != nil {
//	    fmt.Printf("Error compressing file: %v\n", err)
//	    return
//	}
//	fmt.Println("File compressed successfully")
func (ufs *UFS) CompressFile(sourcePath, destPath string) error {
	// Verify source is a file
	if !ufs.IsFile(sourcePath) {
		return fmt.Errorf("source path is not a file: %s", sourcePath)
	}

	// Get absolute paths to ensure consistent behavior
	sourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}

	destPath, err = filepath.Abs(destPath)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}

	// Ensure destination directory exists
	destDir := filepath.Dir(destPath)
	if !ufs.IsDirectory(destDir) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "CompressFile")
		}
	}

	// Create zip file
	zipFile, err := os.Create(destPath)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// Get file info
	info, err := os.Stat(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}

	// Create a zip header
	header, err := zip.FileInfoHeader(info)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}

	// Use the base file name as the name in the archive
	header.Name = filepath.Base(sourcePath)
	header.Method = zip.Deflate

	// Create writer for the file header
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}

	// Open the file for reading
	file, err := os.Open(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}
	defer file.Close()

	// Copy file contents to the zip
	_, err = io.Copy(writer, file)
	if err != nil {
		return ufs.wrapError(err, "CompressFile")
	}

	return nil
}

// CompressHere compresses a directory into a ZIP file and outputs in the current working directory.
// This function is a convenience wrapper around CompressDirectory.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the directory to compress
//
// Returns:
//   - string: The path to the created ZIP file
//   - error: An error if the compression failed, nil otherwise
//
// Example:
//
//	zipPath, err := ufs.CompressHere("/path/to/source_dir")
//	if err != nil {
//	    fmt.Printf("Error compressing directory: %v\n", err)
//	    return
//	}
//	fmt.Printf("Directory compressed to: %s\n", zipPath)
func (ufs *UFS) CompressHere(sourcePath string) (string, error) {
	// Verify source is a directory
	if !ufs.IsDirectory(sourcePath) {
		return "", fmt.Errorf("source path is not a directory: %s", sourcePath)
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", ufs.wrapError(err, "CompressHere")
	}

	// Use the base name of the source directory as the ZIP file name
	zipName := filepath.Base(sourcePath) + ".zip"
	zipPath := filepath.Join(cwd, zipName)

	// Compress the directory
	err = ufs.CompressDirectory(sourcePath, zipPath)
	if err != nil {
		return "", err
	}

	return zipPath, nil
}

// ExtractHere extracts the contents of a ZIP file in the current working directory.
// This function is a convenience wrapper around ExtractArchive.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the ZIP file
//
// Returns:
//   - string: The path to the directory where the archive was extracted
//   - error: An error if the extraction failed, nil otherwise
//
// Example:
//
//	extractPath, err := ufs.ExtractHere("/path/to/archive.zip")
//	if err != nil {
//	    fmt.Printf("Error extracting archive: %v\n", err)
//	    return
//	}
//	fmt.Printf("Archive extracted to: %s\n", extractPath)
func (ufs *UFS) ExtractHere(sourcePath string) (string, error) {
	// Verify source is a file
	if !ufs.IsFile(sourcePath) {
		return "", fmt.Errorf("source path is not a file: %s", sourcePath)
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", ufs.wrapError(err, "ExtractHere")
	}

	// Use the base name of the zip file (without extension) as the destination directory
	zipBase := filepath.Base(sourcePath)
	destName := strings.TrimSuffix(zipBase, filepath.Ext(zipBase))
	destPath := filepath.Join(cwd, destName)

	// Extract the archive
	err = ufs.ExtractArchive(sourcePath, destPath)
	if err != nil {
		return "", err
	}

	return destPath, nil
}

// CompressFileHere compresses a single file into a ZIP file and outputs in the current working directory.
// This function is a convenience wrapper around CompressFile.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the file to compress
//
// Returns:
//   - string: The path to the created ZIP file
//   - error: An error if the compression failed, nil otherwise
//
// Example:
//
//	zipPath, err := ufs.CompressFileHere("/path/to/myfile.txt")
//	if err != nil {
//	    fmt.Printf("Error compressing file: %v\n", err)
//	    return
//	}
//	fmt.Printf("File compressed to: %s\n", zipPath)
func (ufs *UFS) CompressFileHere(sourcePath string) (string, error) {
	// Verify source is a file
	if !ufs.IsFile(sourcePath) {
		return "", fmt.Errorf("source path is not a file: %s", sourcePath)
	}

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", ufs.wrapError(err, "CompressFileHere")
	}

	// Use the base name of the source file as the ZIP file name
	zipName := filepath.Base(sourcePath) + ".zip"
	zipPath := filepath.Join(cwd, zipName)

	// Compress the file
	err = ufs.CompressFile(sourcePath, zipPath)
	if err != nil {
		return "", err
	}

	return zipPath, nil
}

// CompressAndRemove compresses a directory into a ZIP file and removes the original directory.
// WARNING: This is a dangerous operation as it permanently removes the original directory.
// Use with extreme caution.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the directory to compress and remove
//   - destPath: The absolute or relative path where the ZIP file will be created
//
// Returns:
//   - error: An error if the operation failed, nil otherwise
//
// Example:
//
//	err := ufs.CompressAndRemove("/path/to/source_dir", "/path/to/archive.zip")
//	if err != nil {
//	    fmt.Printf("Error compressing and removing directory: %v\n", err)
//	    return
//	}
//	fmt.Println("Directory compressed and removed successfully")
func (ufs *UFS) CompressAndRemove(sourcePath, destPath string) error {
	// First compress the directory
	err := ufs.CompressDirectory(sourcePath, destPath)
	if err != nil {
		return err
	}

	// Then remove the directory
	err = os.RemoveAll(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "CompressAndRemove")
	}

	return nil
}

// ExtractAndRemove extracts a ZIP file and removes the original ZIP file.
// WARNING: This is a dangerous operation as it permanently removes the original ZIP file.
// Use with extreme caution.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the ZIP file to extract and remove
//   - destPath: The absolute or relative path where the contents will be extracted
//
// Returns:
//   - error: An error if the operation failed, nil otherwise
//
// Example:
//
//	err := ufs.ExtractAndRemove("/path/to/archive.zip", "/path/to/extract_dir")
//	if err != nil {
//	    fmt.Printf("Error extracting and removing archive: %v\n", err)
//	    return
//	}
//	fmt.Println("Archive extracted and removed successfully")
func (ufs *UFS) ExtractAndRemove(sourcePath, destPath string) error {
	// First extract the archive
	err := ufs.ExtractArchive(sourcePath, destPath)
	if err != nil {
		return err
	}

	// Then remove the ZIP file
	err = os.Remove(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "ExtractAndRemove")
	}

	return nil
}

// CompressAndExtract compresses a directory and extracts it to a specified location.
// WARNING: This operation compresses a directory and then immediately extracts it elsewhere.
// It's generally inefficient and should only be used for specific purposes.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the directory to compress
//   - tempPath: The absolute or relative path where the temporary ZIP file will be created
//   - finalPath: The absolute or relative path where the contents will be extracted
//
// Returns:
//   - error: An error if the operation failed, nil otherwise
//
// Example:
//
//	err := ufs.CompressAndExtract("/path/to/source_dir", "/path/to/temp.zip", "/path/to/extract_dir")
//	if err != nil {
//	    fmt.Printf("Error compressing and extracting directory: %v\n", err)
//	    return
//	}
//	fmt.Println("Directory compressed and extracted successfully")
func (ufs *UFS) CompressAndExtract(sourcePath, tempPath, finalPath string) error {
	// First compress the directory
	err := ufs.CompressDirectory(sourcePath, tempPath)
	if err != nil {
		return err
	}

	// Then extract it to the final path
	err = ufs.ExtractArchive(tempPath, finalPath)
	if err != nil {
		return err
	}

	// Remove the temporary ZIP file
	err = os.Remove(tempPath)
	if err != nil {
		return ufs.wrapError(err, "CompressAndExtract")
	}

	return nil
}

// ExtractAndCompress extracts a ZIP file and compresses it to a specified location.
// WARNING: This operation extracts an archive and then immediately recompresses it elsewhere.
// It's generally inefficient and should only be used for specific purposes.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the ZIP file to extract
//   - tempPath: The absolute or relative path where the contents will be temporarily extracted
//   - finalPath: The absolute or relative path where the new ZIP file will be created
//
// Returns:
//   - error: An error if the operation failed, nil otherwise
//
// Example:
//
//	err := ufs.ExtractAndCompress("/path/to/archive.zip", "/path/to/temp_dir", "/path/to/new_archive.zip")
//	if err != nil {
//	    fmt.Printf("Error extracting and compressing archive: %v\n", err)
//	    return
//	}
//	fmt.Println("Archive extracted and compressed successfully")
func (ufs *UFS) ExtractAndCompress(sourcePath, tempPath, finalPath string) error {
	// First extract the archive
	err := ufs.ExtractArchive(sourcePath, tempPath)
	if err != nil {
		return err
	}

	// Then compress it to the final path
	err = ufs.CompressDirectory(tempPath, finalPath)
	if err != nil {
		return err
	}

	// Remove the temporary directory
	err = os.RemoveAll(tempPath)
	if err != nil {
		return ufs.wrapError(err, "ExtractAndCompress")
	}

	return nil
}

// CompressWithSystemCommand compresses a directory using the system's compression tool.
// This function uses 'tar' on Unix-like systems and 'tar.exe' on Windows 10 and later.
// Not supported on older Windows versions.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the directory to compress
//   - destPath: The absolute or relative path where the archive will be created
//   - format: The compression format to use (e.g., "gzip", "bzip2", "xz")
//
// Returns:
//   - error: An error if the compression failed, nil otherwise
//
// Example:
//
//	err := ufs.CompressWithSystemCommand("/path/to/source_dir", "/path/to/archive.tar.gz", "gzip")
//	if err != nil {
//	    fmt.Printf("Error compressing directory: %v\n", err)
//	    return
//	}
//	fmt.Println("Directory compressed successfully using system command")
func (ufs *UFS) CompressWithSystemCommand(sourcePath, destPath, format string) error {
	// Verify source is a directory
	if !ufs.IsDirectory(sourcePath) {
		return fmt.Errorf("source path is not a directory: %s", sourcePath)
	}

	// Get absolute paths to ensure consistent behavior
	sourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "CompressWithSystemCommand")
	}

	destPath, err = filepath.Abs(destPath)
	if err != nil {
		return ufs.wrapError(err, "CompressWithSystemCommand")
	}

	// Ensure destination directory exists
	destDir := filepath.Dir(destPath)
	if !ufs.IsDirectory(destDir) {
		err = os.MkdirAll(destDir, 0755)
		if err != nil {
			return ufs.wrapError(err, "CompressWithSystemCommand")
		}
	}

	// Set compression flag based on format
	var compressFlag string
	switch format {
	case "gzip":
		compressFlag = "z"
		if !strings.HasSuffix(destPath, ".tar.gz") && !strings.HasSuffix(destPath, ".tgz") {
			destPath += ".tar.gz"
		}
	case "bzip2":
		compressFlag = "j"
		if !strings.HasSuffix(destPath, ".tar.bz2") && !strings.HasSuffix(destPath, ".tbz2") {
			destPath += ".tar.bz2"
		}
	case "xz":
		compressFlag = "J"
		if !strings.HasSuffix(destPath, ".tar.xz") && !strings.HasSuffix(destPath, ".txz") {
			destPath += ".tar.xz"
		}
	default:
		return fmt.Errorf("unsupported compression format: %s", format)
	}

	var cmd *exec.Cmd
	sourceDir := filepath.Base(sourcePath)
	parentDir := filepath.Dir(sourcePath)

	if runtime.GOOS == "windows" {
		// Check if tar.exe is available
		_, err := exec.LookPath("tar.exe")
		if err != nil {
			return fmt.Errorf("tar.exe not found, compression not supported on this Windows version")
		}
		cmd = exec.Command("tar.exe", "-c"+compressFlag+"f", destPath, "-C", parentDir, sourceDir)
	} else {
		// Unix-like systems
		cmd = exec.Command("tar", "-c"+compressFlag+"f", destPath, "-C", parentDir, sourceDir)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("compression failed: %v, output: %s", err, output)
	}

	return nil
}

// ExtractWithSystemCommand extracts an archive using the system's extraction tool.
// This function uses 'tar' on Unix-like systems and 'tar.exe' on Windows 10 and later.
// Not supported on older Windows versions.
//
// Parameters:
//   - sourcePath: The absolute or relative path to the archive to extract
//   - destPath: The absolute or relative path where the contents will be extracted
//
// Returns:
//   - error: An error if the extraction failed, nil otherwise
//
// Example:
//
//	err := ufs.ExtractWithSystemCommand("/path/to/archive.tar.gz", "/path/to/extract_dir")
//	if err != nil {
//	    fmt.Printf("Error extracting archive: %v\n", err)
//	    return
//	}
//	fmt.Println("Archive extracted successfully using system command")
func (ufs *UFS) ExtractWithSystemCommand(sourcePath, destPath string) error {
	// Verify source is a file
	if !ufs.IsFile(sourcePath) {
		return fmt.Errorf("source path is not a file: %s", sourcePath)
	}

	// Get absolute paths to ensure consistent behavior
	sourcePath, err := filepath.Abs(sourcePath)
	if err != nil {
		return ufs.wrapError(err, "ExtractWithSystemCommand")
	}

	destPath, err = filepath.Abs(destPath)
	if err != nil {
		return ufs.wrapError(err, "ExtractWithSystemCommand")
	}

	// Ensure destination directory exists
	if !ufs.IsDirectory(destPath) {
		err = os.MkdirAll(destPath, 0755)
		if err != nil {
			return ufs.wrapError(err, "ExtractWithSystemCommand")
		}
	}

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		// Check if tar.exe is available
		_, err := exec.LookPath("tar.exe")
		if err != nil {
			return fmt.Errorf("tar.exe not found, extraction not supported on this Windows version")
		}
		cmd = exec.Command("tar.exe", "-xf", sourcePath, "-C", destPath)
	} else {
		// Unix-like systems
		cmd = exec.Command("tar", "-xf", sourcePath, "-C", destPath)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("extraction failed: %v, output: %s", err, output)
	}

	return nil
}
