# Removing.go

The `Removing.go` file provides a comprehensive set of functions for deleting files, directories, and symbolic links from the file system. It offers a unified and safe approach to removing content with built-in validation and error handling.

## Overview

This module includes functionality for:

-   Removing individual files and empty directories
-   Recursively deleting directory structures
-   Removing symbolic links
-   Safely removing files with backup options
-   Cleaning up empty files and directories
-   Pattern-based file removal

All functions handle errors internally and return boolean values to indicate success or failure, making them easy to use in conditional statements.

## Basic Removal Functions

### RemoveFile

Removes a file at the specified path. This function will not remove directories.

**Parameters:**

-   `path`: The absolute or relative path to the file to remove

**Returns:**

-   `bool`: true if the file was removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // First create a file to remove
    fs.CreateFileWithContent("./test.txt", "This is a test file")

    // Now remove the file
    success := fs.RemoveFile("./test.txt")
    if success {
        fmt.Println("File removed successfully")
    } else {
        fmt.Println("Failed to remove file")
    }

    // Verify the file no longer exists
    exists := fs.PathExists("./test.txt")
    fmt.Printf("File still exists: %t\n", exists)
}
```

**Expected Output:**

```
File removed successfully
File still exists: false
```

</details>

### RemoveDirectory

Removes an empty directory at the specified path. This function will fail if the directory is not empty.

**Parameters:**

-   `path`: The absolute or relative path to the directory to remove

**Returns:**

-   `bool`: true if the directory was removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create an empty directory
    fs.CreateDirectory("./empty_dir")

    // Remove the empty directory
    success := fs.RemoveDirectory("./empty_dir")
    if success {
        fmt.Println("Empty directory removed successfully")
    } else {
        fmt.Println("Failed to remove directory")
    }

    // Now create a directory with content
    fs.CreateDirectory("./non_empty_dir")
    fs.CreateFile("./non_empty_dir/file.txt")

    // Try to remove the non-empty directory
    success = fs.RemoveDirectory("./non_empty_dir")
    if success {
        fmt.Println("Non-empty directory removed successfully")
    } else {
        fmt.Println("Failed to remove non-empty directory (expected)")
    }
}
```

**Expected Output:**

```
Empty directory removed successfully
Failed to remove non-empty directory (expected)
```

</details>

### RemoveDirectoryRecursive

Removes a directory and all its contents recursively. This will delete all files and subdirectories within the specified directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory to remove

**Returns:**

-   `bool`: true if the directory and all its contents were removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a directory with nested content
    fs.CreateDirectory("./nested_dir")
    fs.CreateFileWithContent("./nested_dir/file1.txt", "File 1 content")
    fs.CreateDirectory("./nested_dir/subdir")
    fs.CreateFileWithContent("./nested_dir/subdir/file2.txt", "File 2 content")

    // Remove the directory and all its contents
    success := fs.RemoveDirectoryRecursive("./nested_dir")
    if success {
        fmt.Println("Directory and all contents removed successfully")
    } else {
        fmt.Println("Failed to remove directory recursively")
    }

    // Verify the directory no longer exists
    exists := fs.PathExists("./nested_dir")
    fmt.Printf("Directory still exists: %t\n", exists)
}
```

**Expected Output:**

```
Directory and all contents removed successfully
Directory still exists: false
```

</details>

### RemoveSymlink

Removes a symbolic link at the specified path. This function only removes the symlink itself, not the target it points to.

**Parameters:**

-   `path`: The absolute or relative path to the symlink to remove

**Returns:**

-   `bool`: true if the symlink was removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a target file and a symlink to it
    fs.CreateFileWithContent("./target.txt", "This is the target file")
    fs.CreateSymlink("./target.txt", "./link_to_target")

    // Remove just the symlink
    success := fs.RemoveSymlink("./link_to_target")
    if success {
        fmt.Println("Symlink removed successfully")
    } else {
        fmt.Println("Failed to remove symlink")
    }

    // Verify the symlink is gone but the target remains
    linkExists := fs.PathExists("./link_to_target")
    targetExists := fs.PathExists("./target.txt")
    fmt.Printf("Symlink still exists: %t\n", linkExists)
    fmt.Printf("Target file still exists: %t\n", targetExists)
}
```

**Expected Output:**

```
Symlink removed successfully
Symlink still exists: false
Target file still exists: true
```

</details>

## Advanced Removal Functions

### RemoveFileWithBackup

Removes a file at the specified path after creating a backup. The backup file will have the same name with ".bak" appended.

**Parameters:**

-   `path`: The absolute or relative path to the file to remove

**Returns:**

-   `bool`: true if the file was backed up and removed successfully, false otherwise
-   `string`: The path to the backup file, or an empty string if the operation failed

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a file to remove with backup
    fs.CreateFileWithContent("./important.txt", "This is important data that should be backed up")

    // Remove with backup
    success, backupPath := fs.RemoveFileWithBackup("./important.txt")
    if success {
        fmt.Println("File removed with backup successfully")
        fmt.Printf("Backup created at: %s\n", backupPath)
    } else {
        fmt.Println("Failed to remove file with backup")
    }

    // Verify original is gone and backup exists
    originalExists := fs.PathExists("./important.txt")
    backupExists := fs.PathExists(backupPath)
    fmt.Printf("Original file still exists: %t\n", originalExists)
    fmt.Printf("Backup file exists: %t\n", backupExists)

    // Check backup content
    if backupExists {
        content, _ := fs.ReadFileAsString(backupPath)
        fmt.Printf("Backup content: %s\n", content)
    }
}
```

**Expected Output:**

```
File removed with backup successfully
Backup created at: ./important.txt.bak
Original file still exists: false
Backup file exists: true
Backup content: This is important data that should be backed up
```

</details>

### RemoveEmptyFiles

Removes all empty files in the specified directory. This function does not recurse into subdirectories.

**Parameters:**

-   `dirPath`: The absolute or relative path to the directory to clean

**Returns:**

-   `bool`: true if all empty files were removed successfully, false if any removal failed
-   `int`: The number of files removed

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a directory with some empty and non-empty files
    fs.CreateDirectory("./cleanup_dir")
    fs.CreateFile("./cleanup_dir/empty1.txt")  // Empty file
    fs.CreateFile("./cleanup_dir/empty2.txt")  // Empty file
    fs.CreateFileWithContent("./cleanup_dir/nonempty.txt", "This file has content")

    // Remove all empty files
    success, count := fs.RemoveEmptyFiles("./cleanup_dir")
    if success {
        fmt.Printf("Successfully removed %d empty files\n", count)
    } else {
        fmt.Println("Some empty files could not be removed")
    }

    // List remaining files
    files := fs.GetFileList("./cleanup_dir")
    fmt.Println("Remaining files:")
    for _, file := range files {
        fmt.Printf("- %s\n", file)
    }
}
```

**Expected Output:**

```
Successfully removed 2 empty files
Remaining files:
- nonempty.txt
```

</details>

### RemoveEmptyDirectories

Removes all empty directories in the specified directory. This function does not recurse into subdirectories.

**Parameters:**

-   `dirPath`: The absolute or relative path to the directory to clean

**Returns:**

-   `bool`: true if all empty directories were removed successfully, false if any removal failed
-   `int`: The number of directories removed

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a directory with some empty and non-empty subdirectories
    fs.CreateDirectory("./parent_dir")
    fs.CreateDirectory("./parent_dir/empty1")  // Empty directory
    fs.CreateDirectory("./parent_dir/empty2")  // Empty directory
    fs.CreateDirectory("./parent_dir/nonempty")
    fs.CreateFile("./parent_dir/nonempty/file.txt")  // Makes this directory non-empty

    // Remove all empty directories
    success, count := fs.RemoveEmptyDirectories("./parent_dir")
    if success {
        fmt.Printf("Successfully removed %d empty directories\n", count)
    } else {
        fmt.Println("Some empty directories could not be removed")
    }

    // List remaining directories
    dirs := fs.GetFolderList("./parent_dir")
    fmt.Println("Remaining directories:")
    for _, dir := range dirs {
        fmt.Printf("- %s\n", dir)
    }
}
```

**Expected Output:**

```
Successfully removed 2 empty directories
Remaining directories:
- nonempty
```

</details>

### RemoveDirectoryContents

Removes all contents of a directory without removing the directory itself. This will remove all files and subdirectories within the specified directory.

**Parameters:**

-   `dirPath`: The absolute or relative path to the directory whose contents will be removed

**Returns:**

-   `bool`: true if all contents were removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a directory with contents
    fs.CreateDirectory("./container_dir")
    fs.CreateFileWithContent("./container_dir/file1.txt", "Content 1")
    fs.CreateDirectory("./container_dir/subdir")
    fs.CreateFileWithContent("./container_dir/subdir/file2.txt", "Content 2")

    // Remove all contents but keep the directory
    success := fs.RemoveDirectoryContents("./container_dir")
    if success {
        fmt.Println("Directory contents removed successfully")
    } else {
        fmt.Println("Failed to remove directory contents")
    }

    // Verify the directory exists but is empty
    dirExists := fs.PathExists("./container_dir")
    isEmpty := fs.IsDirectoryEmpty("./container_dir")
    fmt.Printf("Directory exists: %t\n", dirExists)
    fmt.Printf("Directory is empty: %t\n", isEmpty)
}
```

**Expected Output:**

```
Directory contents removed successfully
Directory exists: true
Directory is empty: true
```

</details>

### RemoveDirectoryTree

Removes a directory tree structure matching the provided structure map. The structure is a map where keys are directory names and values are either nil (for empty directories) or nested maps (for subdirectories).

**Parameters:**

-   `basePath`: The base directory path where the tree structure exists
-   `structure`: A map representing the directory structure to remove

**Returns:**

-   `bool`: true if the directory tree was removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // First create a directory structure
    structure := map[string]interface{}{
        "docs": map[string]interface{}{
            "technical": nil,
            "user": map[string]interface{}{
                "guides": nil,
            },
        },
        "src": map[string]interface{}{
            "main": nil,
        },
    }

    fs.CreateDirectoryTree("./project", structure)

    // Now remove part of the structure
    removeStructure := map[string]interface{}{
        "docs": map[string]interface{}{
            "technical": nil,
        },
        "src": nil,
    }

    success := fs.RemoveDirectoryTree("./project", removeStructure)
    if success {
        fmt.Println("Directory tree structure removed successfully")
    } else {
        fmt.Println("Failed to remove directory tree structure")
    }

    // Check what remains
    fmt.Println("Remaining directories:")
    checkDirectoryExists(fs, "./project/docs")
    checkDirectoryExists(fs, "./project/docs/technical")
    checkDirectoryExists(fs, "./project/docs/user")
    checkDirectoryExists(fs, "./project/docs/user/guides")
    checkDirectoryExists(fs, "./project/src")
}

func checkDirectoryExists(fs *ufs.UFS, path string) {
    if fs.IsDirectory(path) {
        fmt.Printf("- %s exists\n", path)
    } else {
        fmt.Printf("- %s does not exist\n", path)
    }
}
```

**Expected Output:**

```
Directory tree structure removed successfully
Remaining directories:
- ./project/docs exists
- ./project/docs/technical does not exist
- ./project/docs/user exists
- ./project/docs/user/guides exists
- ./project/src does not exist
```

</details>

### RemoveAllLinks

Removes all symbolic links in the specified directory. This function does not recurse into subdirectories.

**Parameters:**

-   `dirPath`: The absolute or relative path to the directory to clean

**Returns:**

-   `bool`: true if all links were removed successfully, false if any removal failed
-   `int`: The number of links removed

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a directory with files and symlinks
    fs.CreateDirectory("./links_dir")
    fs.CreateFileWithContent("./links_dir/target1.txt", "Target 1 content")
    fs.CreateFileWithContent("./links_dir/target2.txt", "Target 2 content")

    // Create symlinks
    fs.CreateSymlink("./links_dir/target1.txt", "./links_dir/link1")
    fs.CreateSymlink("./links_dir/target2.txt", "./links_dir/link2")
    fs.CreateFile("./links_dir/regular_file.txt") // Not a symlink

    // Remove all symlinks
    success, count := fs.RemoveAllLinks("./links_dir")
    if success {
        fmt.Printf("Successfully removed %d symlinks\n", count)
    } else {
        fmt.Println("Some symlinks could not be removed")
    }

    // List remaining files
    files := fs.GetFileList("./links_dir")
    fmt.Println("Remaining files:")
    for _, file := range files {
        fmt.Printf("- %s\n", file)
    }
}
```

**Expected Output:**

```
Successfully removed 2 symlinks
Remaining files:
- target1.txt
- target2.txt
- regular_file.txt
```

</details>

### RemoveByPattern

Removes all files matching a specified pattern in the given directory. This function uses filepath.Match for pattern matching.

**Parameters:**

-   `dirPath`: The absolute or relative path to the directory to clean
-   `pattern`: The pattern to match files against (e.g., "_.tmp", "backup-_")

**Returns:**

-   `bool`: true if all matching files were removed successfully, false if any removal failed
-   `int`: The number of files removed

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Create a directory with various files
    fs.CreateDirectory("./pattern_dir")
    fs.CreateFile("./pattern_dir/temp1.tmp")
    fs.CreateFile("./pattern_dir/temp2.tmp")
    fs.CreateFile("./pattern_dir/important.doc")
    fs.CreateFile("./pattern_dir/backup-2023.zip")
    fs.CreateFile("./pattern_dir/backup-2024.zip")

    // Remove all .tmp files
    success, count := fs.RemoveByPattern("./pattern_dir", "*.tmp")
    if success {
        fmt.Printf("Successfully removed %d .tmp files\n", count)
    } else {
        fmt.Println("Some .tmp files could not be removed")
    }

    // Remove all backup files
    success, count = fs.RemoveByPattern("./pattern_dir", "backup-*.zip")
    if success {
        fmt.Printf("Successfully removed %d backup files\n", count)
    } else {
        fmt.Println("Some backup files could not be removed")
    }

    // List remaining files
    files := fs.GetFileList("./pattern_dir")
    fmt.Println("Remaining files:")
    for _, file := range files {
        fmt.Printf("- %s\n", file)
    }
}
```

**Expected Output:**

```
Successfully removed 2 .tmp files
Successfully removed 2 backup files
Remaining files:
- important.doc
```

</details>

### SafeRemoveFile

Removes a file only if it matches the expected size and/or modification time. This provides a safety check before deletion to prevent accidental removal of important files.

**Parameters:**

-   `path`: The absolute or relative path to the file to remove
-   `expectedSize`: The expected size of the file in bytes, or -1 to skip this check
-   `expectedModTime`: The expected modification time of the file, or nil to skip this check

**Returns:**

-   `bool`: true if the file was removed successfully, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "os"
    "time"
)

func main() {
    fs := ufs.New()

    // Create a test file
    testContent := "This is a test file with known content"
    fs.CreateFileWithContent("./safe_remove.txt", testContent)

    // Get file info for the expected mod time
    fileInfo, _ := os.Stat("./safe_remove.txt")

    // Try to remove with correct size but wrong mod time
    wrongTime := time.Now().Add(-24 * time.Hour) // 1 day ago
    wrongTimeInfo, _ := fileInfo.(*os.FileInfo)
    success := fs.SafeRemoveFile("./safe_remove.txt", int64(len(testContent)), &wrongTimeInfo)
    if success {
        fmt.Println("File removed with wrong mod time (unexpected)")
    } else {
        fmt.Println("File not removed due to mod time mismatch (expected)")
    }

    // Try to remove with wrong size
    success = fs.SafeRemoveFile("./safe_remove.txt", 1000, nil)
    if success {
        fmt.Println("File removed with wrong size (unexpected)")
    } else {
        fmt.Println("File not removed due to size mismatch (expected)")
    }

    // Try to remove with correct size
    success = fs.SafeRemoveFile("./safe_remove.txt", int64(len(testContent)), nil)
    if success {
        fmt.Println("File removed successfully with correct size")
    } else {
        fmt.Println("File not removed despite correct size (unexpected)")
    }

    // Verify file is gone
    exists := fs.PathExists("./safe_remove.txt")
    fmt.Printf("File still exists: %t\n", exists)
}
```

**Expected Output:**

```
File not removed due to mod time mismatch (expected)
File not removed due to size mismatch (expected)
File removed successfully with correct size
File still exists: false
```

Note: The `expectedModTime` parameter usage might vary in practice due to the complexity of comparing file times precisely.

</details>

## Error Handling and Safety

All functions in Removing.go handle errors internally and return boolean values to indicate success or failure. When an error occurs, it is logged through the UFS instance's error handling mechanism, which can be configured to display errors, log them to a file, or handle them silently.

Safety features include:

-   Path validation to ensure operations target the correct files/directories
-   Type checking to prevent misuse (e.g., trying to remove a directory with RemoveFile)
-   Empty directory verification before removal
-   Backup options for critical files
-   Pattern matching for targeted cleanup

## Performance Considerations

-   `RemoveDirectoryRecursive` can be resource-intensive for very large directories with many nested files
-   For large cleanup operations, consider using pattern-based removal instead of processing each file individually
-   When removing many files, be aware of potential I/O bottlenecks

## Platform Compatibility

These functions are designed to work across different operating systems, but some behaviors may vary:

-   Symlink handling is different on Windows vs. Unix-like systems
-   File permission requirements for deletion can vary by platform
-   Pattern matching follows the rules of Go's filepath.Match function, which has platform-specific behavior for path separators
