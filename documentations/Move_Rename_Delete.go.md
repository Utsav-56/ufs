# Move-Rename_delete.go

The `Move-Rename_delete.go` file provides comprehensive functionality for moving, renaming, and deleting files and directories in the file system. These functions offer a unified and safe approach to file system management with built-in validation and error handling.

## Overview

This module includes functionality for:

-   Moving and renaming files and directories
-   Deleting files and directories with various safety checks
-   Conditional operations that only execute if certain criteria are met
-   Operations with automatic backups

All functions handle errors internally and return boolean values to indicate success or failure, making them easy to use in conditional statements.

## Basic Move and Delete Functions

### MoveFile

Moves or renames a file from one path to another. If the destination already exists, it will be overwritten. Parent directories for the destination will be created if they don't exist.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source file
-   `destPath`: The absolute or relative path where the file should be moved to

**Returns:**

-   `bool`: true if the file was moved successfully, false otherwise

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

    // Create a test file
    fs.CreateFileWithContent("./original.txt", "This is a test file")

    // Move the file to a new location
    success := fs.MoveFile("./original.txt", "./moved/file.txt")
    if success {
        fmt.Println("File moved successfully")
    } else {
        fmt.Println("Failed to move file")
    }

    // Verify the file is in the new location
    if fs.PathExists("./moved/file.txt") {
        content, _ := fs.ReadFileAsString("./moved/file.txt")
        fmt.Printf("Content at new location: %s\n", content)
    }

    // Verify the original file no longer exists
    fmt.Printf("Original file exists: %t\n", fs.PathExists("./original.txt"))
}
```

**Expected Output:**

```
File moved successfully
Content at new location: This is a test file
Original file exists: false
```

</details>

### DeleteFile

Deletes a file at the specified path. This is a wrapper around RemoveFile for consistency with naming.

**Parameters:**

-   `path`: The absolute or relative path to the file to delete

**Returns:**

-   `bool`: true if the file was deleted successfully, false otherwise

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

    // Create a test file
    fs.CreateFile("./to_delete.txt")

    // Verify file exists
    fmt.Printf("File exists before deletion: %t\n", fs.PathExists("./to_delete.txt"))

    // Delete the file
    success := fs.DeleteFile("./to_delete.txt")
    if success {
        fmt.Println("File deleted successfully")
    } else {
        fmt.Println("Failed to delete file")
    }

    // Verify file no longer exists
    fmt.Printf("File exists after deletion: %t\n", fs.PathExists("./to_delete.txt"))
}
```

**Expected Output:**

```
File exists before deletion: true
File deleted successfully
File exists after deletion: false
```

</details>

### DeleteDirectory

Deletes a directory at the specified path, including all its contents. This is a wrapper around RemoveDirectoryRecursive for consistency with naming.

**Parameters:**

-   `path`: The absolute or relative path to the directory to delete

**Returns:**

-   `bool`: true if the directory was deleted successfully, false otherwise

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

    // Create a test directory with some contents
    fs.CreateDirectory("./test_dir")
    fs.CreateFile("./test_dir/file1.txt")
    fs.CreateDirectory("./test_dir/subdir")
    fs.CreateFile("./test_dir/subdir/file2.txt")

    // Verify directory exists
    fmt.Printf("Directory exists before deletion: %t\n", fs.PathExists("./test_dir"))

    // Delete the directory and all its contents
    success := fs.DeleteDirectory("./test_dir")
    if success {
        fmt.Println("Directory deleted successfully")
    } else {
        fmt.Println("Failed to delete directory")
    }

    // Verify directory no longer exists
    fmt.Printf("Directory exists after deletion: %t\n", fs.PathExists("./test_dir"))
}
```

**Expected Output:**

```
Directory exists before deletion: true
Directory deleted successfully
Directory exists after deletion: false
```

</details>

### MoveDirectory

Moves or renames a directory from one path to another. If the destination already exists as a directory, it will attempt to merge the contents. Parent directories for the destination will be created if they don't exist.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source directory
-   `destPath`: The absolute or relative path where the directory should be moved to

**Returns:**

-   `bool`: true if the directory was moved successfully, false otherwise

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

    // Create a test directory with some contents
    fs.CreateDirectory("./source_dir")
    fs.CreateFileWithContent("./source_dir/file1.txt", "File 1 content")
    fs.CreateDirectory("./source_dir/subdir")
    fs.CreateFileWithContent("./source_dir/subdir/file2.txt", "File 2 content")

    // Move the directory to a new location
    success := fs.MoveDirectory("./source_dir", "./destination_dir")
    if success {
        fmt.Println("Directory moved successfully")
    } else {
        fmt.Println("Failed to move directory")
    }

    // Verify the directory and its contents are in the new location
    if fs.PathExists("./destination_dir") {
        file1Exists := fs.PathExists("./destination_dir/file1.txt")
        file2Exists := fs.PathExists("./destination_dir/subdir/file2.txt")
        fmt.Printf("Destination contains file1.txt: %t\n", file1Exists)
        fmt.Printf("Destination contains subdir/file2.txt: %t\n", file2Exists)
    }

    // Verify the original directory no longer exists
    fmt.Printf("Source directory exists: %t\n", fs.PathExists("./source_dir"))
}
```

**Expected Output:**

```
Directory moved successfully
Destination contains file1.txt: true
Destination contains subdir/file2.txt: true
Source directory exists: false
```

</details>

## Conditional Move and Delete Functions

### MoveFileIfExists

Moves a file only if it exists at the source path. If the source file doesn't exist, the function returns true without doing anything.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source file
-   `destPath`: The absolute or relative path where the file should be moved to

**Returns:**

-   `bool`: true if the file was moved successfully or doesn't exist, false otherwise

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

    // Try to move a non-existent file
    success := fs.MoveFileIfExists("./nonexistent.txt", "./destination.txt")
    if success {
        fmt.Println("Operation completed successfully (file didn't exist)")
    } else {
        fmt.Println("Failed to move file")
    }

    // Create a file and try to move it
    fs.CreateFileWithContent("./existing.txt", "This file exists")

    success = fs.MoveFileIfExists("./existing.txt", "./moved_file.txt")
    if success {
        fmt.Println("Existing file moved successfully")
    } else {
        fmt.Println("Failed to move existing file")
    }

    // Verify the moved file exists
    fmt.Printf("Moved file exists: %t\n", fs.PathExists("./moved_file.txt"))
}
```

**Expected Output:**

```
Operation completed successfully (file didn't exist)
Existing file moved successfully
Moved file exists: true
```

</details>

### MoveDirectoryIfExists

Moves a directory only if it exists at the source path. If the source directory doesn't exist, the function returns true without doing anything.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source directory
-   `destPath`: The absolute or relative path where the directory should be moved to

**Returns:**

-   `bool`: true if the directory was moved successfully or doesn't exist, false otherwise

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

    // Try to move a non-existent directory
    success := fs.MoveDirectoryIfExists("./nonexistent_dir", "./destination_dir")
    if success {
        fmt.Println("Operation completed successfully (directory didn't exist)")
    } else {
        fmt.Println("Failed to move directory")
    }

    // Create a directory and try to move it
    fs.CreateDirectory("./existing_dir")
    fs.CreateFile("./existing_dir/file.txt")

    success = fs.MoveDirectoryIfExists("./existing_dir", "./moved_dir")
    if success {
        fmt.Println("Existing directory moved successfully")
    } else {
        fmt.Println("Failed to move existing directory")
    }

    // Verify the moved directory exists
    fmt.Printf("Moved directory exists: %t\n", fs.PathExists("./moved_dir"))
    fmt.Printf("File inside moved directory exists: %t\n", fs.PathExists("./moved_dir/file.txt"))
}
```

**Expected Output:**

```
Operation completed successfully (directory didn't exist)
Existing directory moved successfully
Moved directory exists: true
File inside moved directory exists: true
```

</details>

### DeleteFileIfExists

Deletes a file only if it exists at the specified path. If the file doesn't exist, the function returns true without doing anything.

**Parameters:**

-   `path`: The absolute or relative path to the file to delete

**Returns:**

-   `bool`: true if the file was deleted successfully or doesn't exist, false otherwise

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

    // Try to delete a non-existent file
    success := fs.DeleteFileIfExists("./nonexistent.txt")
    if success {
        fmt.Println("Operation completed successfully (file didn't exist)")
    } else {
        fmt.Println("Failed to delete file")
    }

    // Create a file and try to delete it
    fs.CreateFile("./existing.txt")

    success = fs.DeleteFileIfExists("./existing.txt")
    if success {
        fmt.Println("Existing file deleted successfully")
    } else {
        fmt.Println("Failed to delete existing file")
    }

    // Verify the file no longer exists
    fmt.Printf("File still exists: %t\n", fs.PathExists("./existing.txt"))
}
```

**Expected Output:**

```
Operation completed successfully (file didn't exist)
Existing file deleted successfully
File still exists: false
```

</details>

### DeleteDirectoryIfExists

Deletes a directory only if it exists at the specified path. If the directory doesn't exist, the function returns true without doing anything.

**Parameters:**

-   `path`: The absolute or relative path to the directory to delete

**Returns:**

-   `bool`: true if the directory was deleted successfully or doesn't exist, false otherwise

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

    // Try to delete a non-existent directory
    success := fs.DeleteDirectoryIfExists("./nonexistent_dir")
    if success {
        fmt.Println("Operation completed successfully (directory didn't exist)")
    } else {
        fmt.Println("Failed to delete directory")
    }

    // Create a directory and try to delete it
    fs.CreateDirectory("./existing_dir")
    fs.CreateFile("./existing_dir/file.txt")

    success = fs.DeleteDirectoryIfExists("./existing_dir")
    if success {
        fmt.Println("Existing directory deleted successfully")
    } else {
        fmt.Println("Failed to delete existing directory")
    }

    // Verify the directory no longer exists
    fmt.Printf("Directory still exists: %t\n", fs.PathExists("./existing_dir"))
}
```

**Expected Output:**

```
Operation completed successfully (directory didn't exist)
Existing directory deleted successfully
Directory still exists: false
```

</details>

## Empty-Check Operations

### MoveDirectoryIfEmpty

Moves a directory only if it is empty. If the source directory is not empty, the function returns false without doing anything.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source directory
-   `destPath`: The absolute or relative path where the directory should be moved to

**Returns:**

-   `bool`: true if the directory was moved successfully, false if it doesn't exist or is not empty

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

    // Create a non-empty directory
    fs.CreateDirectory("./non_empty_dir")
    fs.CreateFile("./non_empty_dir/file.txt")

    // Try to move the empty directory
    success := fs.MoveDirectoryIfEmpty("./empty_dir", "./moved_empty_dir")
    if success {
        fmt.Println("Empty directory moved successfully")
    } else {
        fmt.Println("Failed to move empty directory")
    }

    // Try to move the non-empty directory
    success = fs.MoveDirectoryIfEmpty("./non_empty_dir", "./moved_non_empty_dir")
    if success {
        fmt.Println("Non-empty directory moved successfully (unexpected)")
    } else {
        fmt.Println("Did not move non-empty directory (as expected)")
    }

    // Verify results
    fmt.Printf("Empty directory moved: %t\n", fs.PathExists("./moved_empty_dir"))
    fmt.Printf("Non-empty directory still exists: %t\n", fs.PathExists("./non_empty_dir"))
}
```

**Expected Output:**

```
Empty directory moved successfully
Did not move non-empty directory (as expected)
Empty directory moved: true
Non-empty directory still exists: true
```

</details>

### MoveFileIfEmpty

Moves a file only if it is empty (zero bytes). If the source file is not empty, the function returns false without doing anything.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source file
-   `destPath`: The absolute or relative path where the file should be moved to

**Returns:**

-   `bool`: true if the file was moved successfully, false if it doesn't exist or is not empty

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

    // Create an empty file
    fs.CreateFile("./empty_file.txt")

    // Create a non-empty file
    fs.CreateFileWithContent("./non_empty_file.txt", "This file has content")

    // Try to move the empty file
    success := fs.MoveFileIfEmpty("./empty_file.txt", "./moved_empty_file.txt")
    if success {
        fmt.Println("Empty file moved successfully")
    } else {
        fmt.Println("Failed to move empty file")
    }

    // Try to move the non-empty file
    success = fs.MoveFileIfEmpty("./non_empty_file.txt", "./moved_non_empty_file.txt")
    if success {
        fmt.Println("Non-empty file moved successfully (unexpected)")
    } else {
        fmt.Println("Did not move non-empty file (as expected)")
    }

    // Verify results
    fmt.Printf("Empty file moved: %t\n", fs.PathExists("./moved_empty_file.txt"))
    fmt.Printf("Non-empty file still exists: %t\n", fs.PathExists("./non_empty_file.txt"))
}
```

**Expected Output:**

```
Empty file moved successfully
Did not move non-empty file (as expected)
Empty file moved: true
Non-empty file still exists: true
```

</details>

### DeleteFileIfEmpty

Deletes a file only if it is empty (zero bytes). If the file is not empty, the function returns false without doing anything.

**Parameters:**

-   `path`: The absolute or relative path to the file to delete

**Returns:**

-   `bool`: true if the file was deleted successfully, false if it doesn't exist or is not empty

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

    // Create an empty file
    fs.CreateFile("./empty_file.txt")

    // Create a non-empty file
    fs.CreateFileWithContent("./non_empty_file.txt", "This file has content")

    // Try to delete the empty file
    success := fs.DeleteFileIfEmpty("./empty_file.txt")
    if success {
        fmt.Println("Empty file deleted successfully")
    } else {
        fmt.Println("Failed to delete empty file")
    }

    // Try to delete the non-empty file
    success = fs.DeleteFileIfEmpty("./non_empty_file.txt")
    if success {
        fmt.Println("Non-empty file deleted successfully (unexpected)")
    } else {
        fmt.Println("Did not delete non-empty file (as expected)")
    }

    // Verify results
    fmt.Printf("Empty file still exists: %t\n", fs.PathExists("./empty_file.txt"))
    fmt.Printf("Non-empty file still exists: %t\n", fs.PathExists("./non_empty_file.txt"))
}
```

**Expected Output:**

```
Empty file deleted successfully
Did not delete non-empty file (as expected)
Empty file still exists: false
Non-empty file still exists: true
```

</details>

### DeleteDirectoryIfEmpty

Deletes a directory only if it is empty. If the directory is not empty, the function returns false without doing anything.

**Parameters:**

-   `path`: The absolute or relative path to the directory to delete

**Returns:**

-   `bool`: true if the directory was deleted successfully, false if it doesn't exist or is not empty

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

    // Create a non-empty directory
    fs.CreateDirectory("./non_empty_dir")
    fs.CreateFile("./non_empty_dir/file.txt")

    // Try to delete the empty directory
    success := fs.DeleteDirectoryIfEmpty("./empty_dir")
    if success {
        fmt.Println("Empty directory deleted successfully")
    } else {
        fmt.Println("Failed to delete empty directory")
    }

    // Try to delete the non-empty directory
    success = fs.DeleteDirectoryIfEmpty("./non_empty_dir")
    if success {
        fmt.Println("Non-empty directory deleted successfully (unexpected)")
    } else {
        fmt.Println("Did not delete non-empty directory (as expected)")
    }

    // Verify results
    fmt.Printf("Empty directory still exists: %t\n", fs.PathExists("./empty_dir"))
    fmt.Printf("Non-empty directory still exists: %t\n", fs.PathExists("./non_empty_dir"))
}
```

**Expected Output:**

```
Empty directory deleted successfully
Did not delete non-empty directory (as expected)
Empty directory still exists: false
Non-empty directory still exists: true
```

</details>

## Rename Functions

### RenameFile

Renames a file without moving it to a different directory. This is a convenience wrapper around MoveFile for cases where only the name changes.

**Parameters:**

-   `path`: The absolute or relative path to the file to rename
-   `newName`: The new name for the file (not a path, just the filename)

**Returns:**

-   `bool`: true if the file was renamed successfully, false otherwise

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

    // Create a test file
    fs.CreateFileWithContent("./test_folder/original.txt", "This is a test file")

    // Rename the file
    success := fs.RenameFile("./test_folder/original.txt", "renamed.txt")
    if success {
        fmt.Println("File renamed successfully")
    } else {
        fmt.Println("Failed to rename file")
    }

    // Verify the file is renamed but still in the same directory
    fmt.Printf("Original file exists: %t\n", fs.PathExists("./test_folder/original.txt"))
    fmt.Printf("Renamed file exists: %t\n", fs.PathExists("./test_folder/renamed.txt"))

    // Try with an invalid name (containing path separators)
    success = fs.RenameFile("./test_folder/renamed.txt", "invalid/name.txt")
    if success {
        fmt.Println("File renamed with path separator (unexpected)")
    } else {
        fmt.Println("Failed to rename file with path separator (as expected)")
    }
}
```

**Expected Output:**

```
File renamed successfully
Original file exists: false
Renamed file exists: true
Failed to rename file with path separator (as expected)
```

</details>

### RenameDirectory

Renames a directory without moving it to a different location. This is a convenience wrapper around MoveDirectory for cases where only the name changes.

**Parameters:**

-   `path`: The absolute or relative path to the directory to rename
-   `newName`: The new name for the directory (not a path, just the directory name)

**Returns:**

-   `bool`: true if the directory was renamed successfully, false otherwise

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

    // Create a test directory with a file
    fs.CreateDirectory("./parent/original_dir")
    fs.CreateFile("./parent/original_dir/file.txt")

    // Rename the directory
    success := fs.RenameDirectory("./parent/original_dir", "renamed_dir")
    if success {
        fmt.Println("Directory renamed successfully")
    } else {
        fmt.Println("Failed to rename directory")
    }

    // Verify the directory is renamed but still in the same parent directory
    fmt.Printf("Original directory exists: %t\n", fs.PathExists("./parent/original_dir"))
    fmt.Printf("Renamed directory exists: %t\n", fs.PathExists("./parent/renamed_dir"))
    fmt.Printf("File in renamed directory exists: %t\n", fs.PathExists("./parent/renamed_dir/file.txt"))

    // Try with an invalid name (containing path separators)
    success = fs.RenameDirectory("./parent/renamed_dir", "invalid/name")
    if success {
        fmt.Println("Directory renamed with path separator (unexpected)")
    } else {
        fmt.Println("Failed to rename directory with path separator (as expected)")
    }
}
```

**Expected Output:**

```
Directory renamed successfully
Original directory exists: false
Renamed directory exists: true
File in renamed directory exists: true
Failed to rename directory with path separator (as expected)
```

</details>

## Backup Operations

### MoveWithBackup

Moves a file or directory after creating a backup of the destination if it exists. The backup will have the same name with ".bak" appended.

**Parameters:**

-   `srcPath`: The absolute or relative path to the source file or directory
-   `destPath`: The absolute or relative path where the file or directory should be moved to

**Returns:**

-   `bool`: true if the move and backup were successful, false otherwise
-   `string`: The path to the backup if one was created, or an empty string

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

    // Create source and destination files
    fs.CreateFileWithContent("./source.txt", "Source file content")
    fs.CreateFileWithContent("./destination.txt", "Original destination content")

    // Move with backup
    success, backupPath := fs.MoveWithBackup("./source.txt", "./destination.txt")
    if success {
        fmt.Println("File moved with backup successfully")
        fmt.Printf("Backup created at: %s\n", backupPath)
    } else {
        fmt.Println("Failed to move file with backup")
    }

    // Verify source, destination, and backup
    fmt.Printf("Source exists: %t\n", fs.PathExists("./source.txt"))

    if fs.PathExists("./destination.txt") {
        content, _ := fs.ReadFileAsString("./destination.txt")
        fmt.Printf("Destination content: %s\n", content)
    }

    if fs.PathExists(backupPath) {
        content, _ := fs.ReadFileAsString(backupPath)
        fmt.Printf("Backup content: %s\n", content)
    }
}
```

**Expected Output:**

```
File moved with backup successfully
Backup created at: ./destination.txt.bak
Source exists: false
Destination content: Source file content
Backup content: Original destination content
```

</details>

### DeleteWithBackup

Deletes a file or directory after creating a backup. The backup will have the same name with ".bak" appended.

**Parameters:**

-   `path`: The absolute or relative path to the file or directory to delete

**Returns:**

-   `bool`: true if the deletion and backup were successful, false otherwise
-   `string`: The path to the backup that was created

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

    // Create a file to delete
    fs.CreateFileWithContent("./to_delete.txt", "Important content that should be backed up")

    // Delete with backup
    success, backupPath := fs.DeleteWithBackup("./to_delete.txt")
    if success {
        fmt.Println("File deleted with backup successfully")
        fmt.Printf("Backup created at: %s\n", backupPath)
    } else {
        fmt.Println("Failed to delete file with backup")
    }

    // Verify original file is gone and backup exists
    fmt.Printf("Original file exists: %t\n", fs.PathExists("./to_delete.txt"))
    fmt.Printf("Backup exists: %t\n", fs.PathExists(backupPath))

    // Check backup content
    if fs.PathExists(backupPath) {
        content, _ := fs.ReadFileAsString(backupPath)
        fmt.Printf("Backup content: %s\n", content)
    }
}
```

**Expected Output:**

```
File deleted with backup successfully
Backup created at: ./to_delete.txt.bak
Original file exists: false
Backup exists: true
Backup content: Important content that should be backed up
```

</details>

## Error Handling and Safety

All functions in Move-Rename_delete.go handle errors internally and return boolean values to indicate success or failure. When an error occurs, it is logged through the UFS instance's error handling mechanism, which can be configured to display errors, log them to a file, or handle them silently.

Safety features include:

-   Type checking to ensure operations are applied to the correct type of path (file or directory)
-   Path validation to prevent invalid operations
-   Conditional operations that only execute when specific criteria are met
-   Automatic backup creation for potentially destructive operations

## Advanced Usage

<details>
<summary>Example: Complex File Organization with Safety Checks</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "path/filepath"
)

func main() {
    fs := ufs.New()

    // Organize files by type with safety checks
    organizeFolder := func(sourcePath string) {
        // Create organization folders
        fs.CreateDirectory(filepath.Join(sourcePath, "documents"))
        fs.CreateDirectory(filepath.Join(sourcePath, "images"))
        fs.CreateDirectory(filepath.Join(sourcePath, "other"))

        // Get list of files in the source directory
        files := fs.GetFileList(sourcePath)

        for _, file := range files {
            fullPath := filepath.Join(sourcePath, file)
            ext := filepath.Ext(file)

            var destFolder string
            switch ext {
            case ".txt", ".doc", ".pdf":
                destFolder = "documents"
            case ".jpg", ".png", ".gif":
                destFolder = "images"
            default:
                destFolder = "other"
            }

            destPath := filepath.Join(sourcePath, destFolder, file)

            // Move with backup for safety
            success, backupPath := fs.MoveWithBackup(fullPath, destPath)
            if success {
                fmt.Printf("Moved %s to %s\n", file, destFolder)
                if backupPath != "" {
                    fmt.Printf("  (Created backup at %s)\n", backupPath)
                }
            } else {
                fmt.Printf("Failed to move %s\n", file)
            }
        }

        // Clean up empty directories
        dirs := fs.GetFolderList(sourcePath)
        for _, dir := range dirs {
            dirPath := filepath.Join(sourcePath, dir)
            if fs.IsDirectoryEmpty(dirPath) {
                fs.DeleteDirectoryIfEmpty(dirPath)
                fmt.Printf("Removed empty directory: %s\n", dir)
            }
        }
    }

    // Use the function
    organizeFolder("./my_files")
}
```

</details>

## Platform Compatibility

These functions are designed to work across different operating systems with consistent behavior:

-   Path separators (backslashes on Windows, forward slashes on Unix) are handled automatically
-   When moving files across different filesystems (where simple rename operations fail), the functions automatically fall back to a copy-and-delete approach
-   Directory structures are preserved during move operations
-   Backups are created with platform-specific path handling
