// Package ufs provides a comprehensive Go library for unified file system operations.
//
// # UFS - Ultimate File System Utils
//
// # Overview
//
// UFS is a comprehensive Go library that provides a unified interface for file system operations across different platforms.
// It abstracts away the complexities of common file operations, allowing developers to write cleaner, more maintainable code
// with minimal platform-specific considerations.
//
// Key Features
//
//   - Platform Independence: Works consistently across Windows, macOS, and Linux
//   - Comprehensive API: Covers everything from basic file operations to advanced directory manipulation
//   - Error Handling: Built-in error management with customizable verbosity
//   - Safety Features: Includes validation and backup options for potentially destructive operations
//   - Flexible Usage: Can be used via static functions or through UFS instances
//   - Rich Documentation: Every function includes detailed explanations and examples
//
// # Core Functionality
//
// # Path Properties and Validation
//
// UFS provides a rich set of functions for checking path properties:
//   - Existence checks for files and directories
//   - Type validation (file vs directory)
//   - Empty status verification
//   - Location validation (system, user, or current directory)
//   - Permission and attribute checks (hidden, executable, readable, writable)
//
// # File and Directory Creation
//
// UFS simplifies the creation of file system objects:
//   - Create files with or without content
//   - Create files and directories with custom permissions
//   - Create symbolic and hard links
//   - Build complex directory trees from hierarchical maps
//   - Create symlinked directory structures
//
// # File Reading and Writing
//
// UFS handles all aspects of file content manipulation:
//   - Read files as bytes or strings
//   - Write or append data to files
//   - Read files line by line
//   - Manipulate specific lines in files
//   - Combine multiple files or split large files
//
// # File and Directory Management
//
// UFS provides robust management functions:
//   - Move and rename files and directories
//   - Delete files and directories with safety checks
//   - Conditional operations (if exists, if empty)
//   - Operations with automatic backups
//   - Pattern-based file removal
//
// # Compression and Extraction
//
// UFS includes comprehensive archive handling:
//   - Compress directories and files into ZIP archives
//   - Extract contents from ZIP archives
//   - Perform compression operations with system tools
//   - Combine compression with other operations (move, delete)
//
// # Why UFS is Helpful
//
// 1. Simplifies Common Tasks
//
// UFS dramatically reduces the amount of code needed for file operations by providing high-level functions that handle edge cases automatically.
//
// Without UFS:
//
//	// Check if directory exists
//	info, err := os.Stat(dirPath)
//	if os.IsNotExist(err) || !info.IsDir() {
//	    // Handle non-existent directory
//	}
//
//	// Create parent directory if it doesn't exist
//	parentDir := filepath.Dir(filePath)
//	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
//	    if err := os.MkdirAll(parentDir, 0755); err != nil {
//	        return err
//	    }
//	}
//
//	// Write to file
//	file, err := os.Create(filePath)
//	if err != nil {
//	    return err
//	}
//	defer file.Close()
//	_, err = file.Write(data)
//	return err
//
// With UFS:
//
//	fs := ufs.New()
//	fs.CreateFileWithContent(filePath, string(data))
//
// 2. Improves Code Readability
//
// UFS's consistent API design and meaningful function names make code more readable and self-documenting.
//
// 3. Reduces Platform-Specific Code
//
// UFS handles platform differences internally, eliminating the need for conditional code based on operating systems.
//
// 4. Prevents Common Errors
//
// Built-in validation prevents common mistakes like:
//   - Attempting to create files in non-existent directories
//   - Accidentally overwriting important files
//   - Path traversal vulnerabilities
//   - Resource leaks from unclosed files
//
// 5. Enhances Safety
//
// Functions like MoveWithBackup, DeleteWithBackup, and SafeRemoveFile protect against accidental data loss.
//
// 6. Streamlines Complex Operations
//
// UFS makes complex operations straightforward:
//   - Creating nested directory structures
//   - Compressing and extracting archives
//   - Moving directories while preserving permissions
//   - Recursively processing file trees
//
// # Usage Examples
//
// Basic Usage with Static Functions
//
//	import "github.com/utsav-56/ufs"
//
//	func main() {
//	    // Check if a file exists
//	    if ufs.PathExists("config.json") {
//	        // Read file content
//	        content, err := ufs.ReadFileAsString("config.json")
//	        if err == nil {
//	            // Process content
//	        }
//	    } else {
//	        // Create default config
//	        ufs.WriteStringToFile("config.json", "{\"setting\": \"default\"}")
//	    }
//
//	    // Compress a directory
//	    ufs.CompressDirectory("./logs", "./logs_archive.zip")
//	}
//
// Advanced Usage with UFS Instance
//
//	import "github.com/utsav-56/ufs"
//
//	func main() {
//	    // Create UFS instance with custom options
//	    fs := ufs.NewUfs(&ufs.Options{
//	        ShowError: true,
//	        ReturnReadable: true,
//	    })
//
//	    // Create a complex directory structure
//	    structure := map[string]interface{}{
//	        "src": map[string]interface{}{
//	            "api": nil,
//	            "models": nil,
//	            "utils": nil,
//	        },
//	        "docs": map[string]interface{}{
//	            "api": nil,
//	        },
//	        "tests": nil,
//	    }
//
//	    fs.CreateDirectoryTree("./project", structure)
//
//	    // Add some files
//	    fs.CreateFileWithContent("./project/src/api/server.go", "package api\n\nfunc StartServer() {\n\t// TODO\n}")
//
//	    // Back up the entire project
//	    fs.CompressDirectory("./project", "./backup/project_backup.zip")
//	}
//
// # Design Philosophy
//
// UFS was designed with several key principles in mind:
//
//  1. Simplicity: Complex operations should be accessible through simple function calls
//  2. Safety: Potentially destructive operations should have safeguards
//  3. Consistency: Error handling and return values should follow predictable patterns
//  4. Flexibility: The library should work with different coding styles and preferences
//  5. Comprehensiveness: Cover the full range of common file system operations
//
// # Conclusion
//
// UFS fills a gap in Go's standard library by providing a higher-level abstraction for file system operations.
// It saves development time, reduces errors, and makes code more maintainable by encapsulating complex operations into simple,
// well-documented functions. Whether for simple scripts or complex applications, UFS provides a consistent and reliable interface
// for all file system needs.
package ufs
