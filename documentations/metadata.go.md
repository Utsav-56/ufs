# Metadata.go

The `Metadata.go` file provides functions for retrieving and analyzing metadata about files and directories in the file system. These functions allow you to gather information such as file sizes, permissions, modification times, and directory contents with a simple, unified interface.

## Overview

This module includes functionality for:

-   Getting file and folder sizes
-   Retrieving detailed metadata about files and directories
-   Listing files and directories within a folder
-   Counting files and subdirectories
-   Analyzing directory structures

All functions handle errors internally and return appropriate data types or empty values in case of failure, making them robust and easy to use.

## Available Functions

### GetFileSize

Returns the size of a file in bytes.

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `int64`: The size of the file in bytes (returns 0 if the file doesn't exist or is a directory)

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

    // Create a file with some content
    fs.CreateFileWithContent("./example.txt", "Hello, this is a test file!")

    // Get the file size
    size := fs.GetFileSize("./example.txt")
    fmt.Printf("File size: %d bytes\n", size)
}
```

**Expected Output:**

```
File size: 28 bytes
```

</details>

### GetFileMetadata

Retrieves comprehensive metadata about a file or directory.

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `map[string]interface{}`: A map containing file metadata with the following keys:
    -   `Name`: The base name of the file
    -   `Size`: The size of the file in bytes
    -   `Mode`: The file's permission bits as a string
    -   `ModTime`: The last modification time
    -   `IsDir`: Boolean indicating if the path is a directory

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
    fs.CreateFileWithContent("./metadata_test.txt", "Testing metadata retrieval")

    // Get the file metadata
    metadata := fs.GetFileMetadata("./metadata_test.txt")

    // Display metadata
    fmt.Printf("File name: %s\n", metadata["Name"])
    fmt.Printf("File size: %d bytes\n", metadata["Size"])
    fmt.Printf("Permissions: %s\n", metadata["Mode"])
    fmt.Printf("Last modified: %v\n", metadata["ModTime"])
    fmt.Printf("Is directory: %t\n", metadata["IsDir"])
}
```

**Expected Output:**

```
File name: metadata_test.txt
File size: 25 bytes
Permissions: -rw-rw-rw-
Last modified: 2025-06-24 15:30:45.123456789 +0000 UTC
Is directory: false
```

</details>

### GetFileList

Returns a list of all files (excluding directories) in the specified directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory to list files from

**Returns:**

-   `[]string`: A slice containing the names of all files in the directory (empty slice if error)

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

    // Create a directory with some files
    fs.CreateDirectory("./test_dir")
    fs.CreateFile("./test_dir/file1.txt")
    fs.CreateFile("./test_dir/file2.txt")
    fs.CreateDirectory("./test_dir/subdir")
    fs.CreateFile("./test_dir/file3.txt")

    // Get list of files only (no directories)
    files := fs.GetFileList("./test_dir")

    fmt.Println("Files in directory:")
    for _, file := range files {
        fmt.Printf("- %s\n", file)
    }
}
```

**Expected Output:**

```
Files in directory:
- file1.txt
- file2.txt
- file3.txt
```

</details>

### GetFolderList

Returns a list of all directories (excluding files) in the specified directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory to list folders from

**Returns:**

-   `[]string`: A slice containing the names of all subdirectories in the directory (empty slice if error)

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

    // Create a directory structure
    fs.CreateDirectory("./parent_dir")
    fs.CreateDirectory("./parent_dir/subdir1")
    fs.CreateDirectory("./parent_dir/subdir2")
    fs.CreateFile("./parent_dir/file.txt")
    fs.CreateDirectory("./parent_dir/subdir3")

    // Get list of directories only (no files)
    folders := fs.GetFolderList("./parent_dir")

    fmt.Println("Subdirectories:")
    for _, folder := range folders {
        fmt.Printf("- %s\n", folder)
    }
}
```

**Expected Output:**

```
Subdirectories:
- subdir1
- subdir2
- subdir3
```

</details>

### GetFolderFileCount

Returns the number of files (excluding directories) in a directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory to count files in

**Returns:**

-   `int`: The number of files in the directory (0 if error)

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

    // Create a directory with some files
    fs.CreateDirectory("./count_dir")
    fs.CreateFile("./count_dir/file1.txt")
    fs.CreateFile("./count_dir/file2.txt")
    fs.CreateDirectory("./count_dir/subdir") // This shouldn't be counted
    fs.CreateFile("./count_dir/file3.txt")

    // Count just the files
    count := fs.GetFolderFileCount("./count_dir")

    fmt.Printf("Number of files in directory: %d\n", count)
}
```

**Expected Output:**

```
Number of files in directory: 3
```

</details>

### GetFolderChildCount

Returns the total number of entries (both files and directories) in a directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory to count children in

**Returns:**

-   `int`: The total number of files and directories in the directory (0 if error)

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

    // Create a directory with mixed content
    fs.CreateDirectory("./total_count_dir")
    fs.CreateFile("./total_count_dir/file1.txt")
    fs.CreateFile("./total_count_dir/file2.txt")
    fs.CreateDirectory("./total_count_dir/subdir1")
    fs.CreateDirectory("./total_count_dir/subdir2")

    // Count all entries (files and directories)
    count := fs.GetFolderChildCount("./total_count_dir")

    fmt.Printf("Total items in directory: %d\n", count)
}
```

**Expected Output:**

```
Total items in directory: 4
```

</details>

### GetChildCount

Returns separate counts for the number of files and directories in a directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory to count children in

**Returns:**

-   `int`: The number of directories (first return value)
-   `int`: The number of files (second return value)

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

    // Create a directory with mixed content
    fs.CreateDirectory("./detailed_count_dir")
    fs.CreateFile("./detailed_count_dir/file1.txt")
    fs.CreateFile("./detailed_count_dir/file2.txt")
    fs.CreateDirectory("./detailed_count_dir/subdir1")
    fs.CreateDirectory("./detailed_count_dir/subdir2")
    fs.CreateFile("./detailed_count_dir/file3.txt")

    // Get separate counts for folders and files
    folderCount, fileCount := fs.GetChildCount("./detailed_count_dir")

    fmt.Printf("Directory contains %d folders and %d files\n", folderCount, fileCount)
}
```

**Expected Output:**

```
Directory contains 2 folders and 3 files
```

</details>

### GetFolderMetadata

Retrieves comprehensive metadata about a directory.

**Parameters:**

-   `path`: The absolute or relative path to the directory

**Returns:**

-   `map[string]interface{}`: A map containing directory metadata with the following keys:
    -   `Name`: The base name of the directory
    -   `Size`: The size of the directory entry itself (not contents)
    -   `Mode`: The directory's permission bits as a string
    -   `ModTime`: The last modification time
    -   `IsDir`: Boolean indicating if the path is a directory (always true)
    -   `ChildNum`: The number of children in the directory

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

    // Create a test directory with some content
    fs.CreateDirectory("./folder_metadata_test")
    fs.CreateFile("./folder_metadata_test/file1.txt")
    fs.CreateDirectory("./folder_metadata_test/subdir")

    // Get the directory metadata
    metadata := fs.GetFolderMetadata("./folder_metadata_test")

    // Display metadata
    fmt.Printf("Folder name: %s\n", metadata["Name"])
    fmt.Printf("Permissions: %s\n", metadata["Mode"])
    fmt.Printf("Last modified: %v\n", metadata["ModTime"])
    fmt.Printf("Is directory: %t\n", metadata["IsDir"])
    fmt.Printf("Child count: %d\n", metadata["ChildNum"])
}
```

**Expected Output:**

```
Folder name: folder_metadata_test
Permissions: drwxrwxrwx
Last modified: 2025-06-24 15:35:12.987654321 +0000 UTC
Is directory: true
Child count: 2
```

</details>

### GetFolderSize

Recursively calculates the total size of a directory and all its contents.

**Parameters:**

-   `path`: The absolute or relative path to the directory to calculate size for

**Returns:**

-   `int64`: The total size of all files in the directory tree in bytes (0 if error)

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

    // Create a directory structure with files of different sizes
    fs.CreateDirectory("./size_test_dir")
    fs.CreateFileWithContent("./size_test_dir/file1.txt", "This is a small file")

    // Create a nested directory with more files
    fs.CreateDirectory("./size_test_dir/nested")
    fs.CreateFileWithContent("./size_test_dir/nested/file2.txt",
        "This is a slightly larger file with more content than the first one")

    // Get the total size of the directory and all its contents
    totalSize := fs.GetFolderSize("./size_test_dir")

    fmt.Printf("Total folder size: %d bytes\n", totalSize)
}
```

**Expected Output:**

```
Total folder size: 89 bytes
```

(The exact size will depend on the content of the files)

</details>

## Advanced Usage

### Combining Functions for Enhanced Analysis

You can combine these metadata functions to perform more complex analyses of your file system.

<details>
<summary>Example: Analyzing Directory Contents</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "path/filepath"
)

func main() {
    fs := ufs.New()
    directoryPath := "./analysis_dir"

    // Get basic counts
    folderCount, fileCount := fs.GetChildCount(directoryPath)
    fmt.Printf("Directory contains %d folders and %d files\n", folderCount, fileCount)

    // Calculate average file size
    files := fs.GetFileList(directoryPath)
    var totalSize int64
    for _, fileName := range files {
        filePath := filepath.Join(directoryPath, fileName)
        totalSize += fs.GetFileSize(filePath)
    }

    var avgSize float64
    if fileCount > 0 {
        avgSize = float64(totalSize) / float64(fileCount)
    }

    fmt.Printf("Total size of all files: %d bytes\n", totalSize)
    fmt.Printf("Average file size: %.2f bytes\n", avgSize)

    // List subdirectories and their item counts
    folders := fs.GetFolderList(directoryPath)
    fmt.Println("\nSubdirectory analysis:")
    for _, folderName := range folders {
        folderPath := filepath.Join(directoryPath, folderName)
        subFolderCount, subFileCount := fs.GetChildCount(folderPath)
        folderSize := fs.GetFolderSize(folderPath)

        fmt.Printf("- %s: %d folders, %d files, %d bytes total\n",
            folderName, subFolderCount, subFileCount, folderSize)
    }
}
```

</details>

## Performance Considerations

-   `GetFolderSize` traverses the entire directory tree recursively, which can be slow for very large directories.
-   When analyzing large file systems, consider using these functions selectively to avoid performance issues.
-   For frequent metadata queries on the same files, consider caching the results rather than repeatedly calling these functions.

## Error Handling

All functions in Metadata.go handle errors internally and return appropriate default values when errors occur:

-   Size functions return 0
-   List functions return empty slices
-   Metadata functions return nil

Errors are logged through the UFS instance's error handling mechanism, which can be configured to display errors, log them to a file, or handle them silently.
