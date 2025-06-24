# File-Reader_Writer.go

The `file-Reader_Writer.go` file provides comprehensive functionality for reading from and writing to files in the file system. Unlike other modules in the UFS package, these functions return error values in addition to their results, acknowledging the critical nature of file I/O operations and enabling more detailed error handling.

## Overview

This module includes functionality for:

-   Reading file contents as bytes or strings
-   Writing data to files (creating or overwriting)
-   Appending data to existing files
-   Copying files with or without preserving permissions
-   Moving files with permission preservation
-   Advanced file operations like assembling multiple files, splitting large files, and line-based manipulations

All functions provide detailed error information when operations fail, allowing for robust error handling in applications that work with file content.

## Basic Reading Functions

### ReadFile

Reads the content of a file and returns it as a byte slice.

**Parameters:**

-   `path`: The absolute or relative path to the file to read

**Returns:**

-   `[]byte`: The content of the file as a byte slice
-   `error`: An error if the file couldn't be read or doesn't exist

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

    // First create a file with some content
    fs.CreateFileWithContent("./sample.txt", "Hello, World!\nThis is a test file.")

    // Read the file contents
    data, err := fs.ReadFile("./sample.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    fmt.Printf("File content as bytes: %v\n", data)
    fmt.Printf("File content as string: %s\n", data)
}
```

**Expected Output:**

```
File content as bytes: [72 101 108 108 111 44 32 87 111 114 108 100 33 10 84 104 105 115 32 105 115 32 97 32 116 101 115 116 32 102 105 108 101 46]
File content as string: Hello, World!
This is a test file.
```

</details>

### ReadFileAsString

Reads the content of a file and returns it as a string.

**Parameters:**

-   `path`: The absolute or relative path to the file to read

**Returns:**

-   `string`: The content of the file as a string
-   `error`: An error if the file couldn't be read or doesn't exist

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

    // First create a file with some content
    fs.CreateFileWithContent("./text_file.txt", "Line 1\nLine 2\nLine 3")

    // Read the file contents as a string
    content, err := fs.ReadFileAsString("./text_file.txt")
    if err != nil {
        fmt.Printf("Error reading file: %v\n", err)
        return
    }

    fmt.Printf("File content: %s\n", content)
}
```

**Expected Output:**

```
File content: Line 1
Line 2
Line 3
```

</details>

## Basic Writing Functions

### WriteFile

Writes data to a file, creating it if it doesn't exist or overwriting it if it does.

**Parameters:**

-   `path`: The absolute or relative path to the file to write
-   `data`: The data to write to the file as a byte slice

**Returns:**

-   `error`: An error if the file couldn't be written to

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

    // Write binary data to a file
    data := []byte{72, 101, 108, 108, 111, 44, 32, 87, 111, 114, 108, 100, 33}

    err := fs.WriteFile("./output.bin", data)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
        return
    }

    fmt.Println("File written successfully")

    // Verify the content
    content, _ := fs.ReadFileAsString("./output.bin")
    fmt.Printf("File content: %s\n", content)
}
```

**Expected Output:**

```
File written successfully
File content: Hello, World!
```

</details>

### WriteStringToFile

Writes a string to a file, creating it if it doesn't exist or overwriting it if it does.

**Parameters:**

-   `path`: The absolute or relative path to the file to write
-   `content`: The string content to write to the file

**Returns:**

-   `error`: An error if the file couldn't be written to

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

    // Write a string to a file
    content := "This is a string\nWith multiple lines\nWritten to a file"

    err := fs.WriteStringToFile("./string_output.txt", content)
    if err != nil {
        fmt.Printf("Error writing string to file: %v\n", err)
        return
    }

    fmt.Println("String written to file successfully")

    // Verify the content
    readContent, _ := fs.ReadFileAsString("./string_output.txt")
    fmt.Printf("File content:\n%s\n", readContent)
}
```

**Expected Output:**

```
String written to file successfully
File content:
This is a string
With multiple lines
Written to a file
```

</details>

## Appending Functions

### AppendToFile

Appends data to a file, creating it if it doesn't exist.

**Parameters:**

-   `path`: The absolute or relative path to the file to append to
-   `data`: The data to append to the file as a byte slice

**Returns:**

-   `error`: An error if the file couldn't be appended to

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

    // Create a file with initial content
    fs.WriteStringToFile("./log.txt", "Initial log entry\n")

    // Append more data to the file
    err := fs.AppendToFile("./log.txt", []byte("Second log entry\n"))
    if err != nil {
        fmt.Printf("Error appending to file: %v\n", err)
        return
    }

    // Append even more data
    err = fs.AppendToFile("./log.txt", []byte("Third log entry\n"))
    if err != nil {
        fmt.Printf("Error appending to file: %v\n", err)
        return
    }

    fmt.Println("Data appended to file successfully")

    // Verify the content
    content, _ := fs.ReadFileAsString("./log.txt")
    fmt.Printf("Final file content:\n%s", content)
}
```

**Expected Output:**

```
Data appended to file successfully
Final file content:
Initial log entry
Second log entry
Third log entry
```

</details>

### AppendStringToFile

Appends a string to a file, creating it if it doesn't exist.

**Parameters:**

-   `path`: The absolute or relative path to the file to append to
-   `content`: The string content to append to the file

**Returns:**

-   `error`: An error if the file couldn't be appended to

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

    // Create a file with initial content
    fs.WriteStringToFile("./notes.txt", "Note 1: Remember to buy groceries\n")

    // Append more notes
    err := fs.AppendStringToFile("./notes.txt", "Note 2: Call dentist\n")
    if err != nil {
        fmt.Printf("Error appending string to file: %v\n", err)
        return
    }

    err = fs.AppendStringToFile("./notes.txt", "Note 3: Finish project report\n")
    if err != nil {
        fmt.Printf("Error appending string to file: %v\n", err)
        return
    }

    fmt.Println("Strings appended to file successfully")

    // Verify the content
    content, _ := fs.ReadFileAsString("./notes.txt")
    fmt.Printf("Notes:\n%s", content)
}
```

**Expected Output:**

```
Strings appended to file successfully
Notes:
Note 1: Remember to buy groceries
Note 2: Call dentist
Note 3: Finish project report
```

</details>

## Copying and Moving Functions

### CopyFile

Copies the content of one file to another.

**Parameters:**

-   `src`: The absolute or relative path to the source file
-   `dst`: The absolute or relative path to the destination file

**Returns:**

-   `error`: An error if the file couldn't be copied

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

    // Create a source file
    fs.WriteStringToFile("./original.txt", "This is the original file content")

    // Copy the file to a new location
    err := fs.CopyFile("./original.txt", "./copy.txt")
    if err != nil {
        fmt.Printf("Error copying file: %v\n", err)
        return
    }

    fmt.Println("File copied successfully")

    // Verify both files exist and have the same content
    originalContent, _ := fs.ReadFileAsString("./original.txt")
    copyContent, _ := fs.ReadFileAsString("./copy.txt")

    fmt.Printf("Original file content: %s\n", originalContent)
    fmt.Printf("Copied file content: %s\n", copyContent)
}
```

**Expected Output:**

```
File copied successfully
Original file content: This is the original file content
Copied file content: This is the original file content
```

</details>

### CopyFileWithPermissions

Copies a file to a new location, preserving its permissions.

**Parameters:**

-   `src`: The absolute or relative path to the source file
-   `dst`: The absolute or relative path to the destination file

**Returns:**

-   `error`: An error if the file couldn't be copied or permissions couldn't be preserved

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "os"
)

func main() {
    fs := ufs.New()

    // Create a source file with specific permissions
    sourceFile := "./executable.sh"
    fs.WriteStringToFile(sourceFile, "#!/bin/sh\necho 'Hello from script'")

    // On Unix systems, make it executable
    if os.Getuid() >= 0 { // Check if we're on a Unix-like system
        os.Chmod(sourceFile, 0755) // rwxr-xr-x
    }

    // Copy the file with permissions preserved
    err := fs.CopyFileWithPermissions(sourceFile, "./executable_copy.sh")
    if err != nil {
        fmt.Printf("Error copying file with permissions: %v\n", err)
        return
    }

    fmt.Println("File copied with permissions successfully")

    // On Unix systems, we could verify permissions are the same
    if os.Getuid() >= 0 {
        srcInfo, _ := os.Stat(sourceFile)
        dstInfo, _ := os.Stat("./executable_copy.sh")
        fmt.Printf("Source permissions: %v\n", srcInfo.Mode())
        fmt.Printf("Destination permissions: %v\n", dstInfo.Mode())
    }
}
```

**Expected Output (on Unix-like systems):**

```
File copied with permissions successfully
Source permissions: -rwxr-xr-x
Destination permissions: -rwxr-xr-x
```

Note: On Windows, permissions work differently, so the output would reflect Windows permission attributes.

</details>

### MoveFileWithPermissions

Moves a file to a new location, preserving its permissions.

**Parameters:**

-   `src`: The absolute or relative path to the source file
-   `dst`: The absolute or relative path to the destination file

**Returns:**

-   `error`: An error if the file couldn't be moved or permissions couldn't be preserved

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "os"
)

func main() {
    fs := ufs.New()

    // Create a source file with specific content
    sourceFile := "./source_document.txt"
    fs.WriteStringToFile(sourceFile, "This document will be moved")

    // On Unix systems, set custom permissions
    if os.Getuid() >= 0 {
        os.Chmod(sourceFile, 0640) // rw-r-----
    }

    // Move the file with permissions preserved
    err := fs.MoveFileWithPermissions(sourceFile, "./moved_document.txt")
    if err != nil {
        fmt.Printf("Error moving file with permissions: %v\n", err)
        return
    }

    fmt.Println("File moved with permissions successfully")

    // Verify the source file no longer exists
    sourceExists := fs.PathExists(sourceFile)
    fmt.Printf("Source file still exists: %t\n", sourceExists)

    // Verify the destination file exists and has the correct content
    if fs.PathExists("./moved_document.txt") {
        content, _ := fs.ReadFileAsString("./moved_document.txt")
        fmt.Printf("Moved file content: %s\n", content)

        // On Unix systems, verify permissions
        if os.Getuid() >= 0 {
            fileInfo, _ := os.Stat("./moved_document.txt")
            fmt.Printf("Moved file permissions: %v\n", fileInfo.Mode())
        }
    }
}
```

**Expected Output (on Unix-like systems):**

```
File moved with permissions successfully
Source file still exists: false
Moved file content: This document will be moved
Moved file permissions: -rw-r-----
```

</details>

## Advanced File Operations

### AssembleFiles

Combines multiple files into a single file in the order of the provided slice.

**Parameters:**

-   `srcFiles`: A slice of file paths to be combined
-   `dst`: The path to the destination file

**Returns:**

-   `error`: An error if the files couldn't be combined

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

    // Create some source files
    fs.WriteStringToFile("./part1.txt", "This is part 1 content.\n")
    fs.WriteStringToFile("./part2.txt", "This is part 2 content.\n")
    fs.WriteStringToFile("./part3.txt", "This is part 3 content.\n")

    // Combine the files
    srcFiles := []string{"./part1.txt", "./part2.txt", "./part3.txt"}
    err := fs.AssembleFiles(srcFiles, "./combined.txt")
    if err != nil {
        fmt.Printf("Error assembling files: %v\n", err)
        return
    }

    fmt.Println("Files assembled successfully")

    // Verify the combined content
    combined, _ := fs.ReadFileAsString("./combined.txt")
    fmt.Printf("Combined file content:\n%s", combined)
}
```

**Expected Output:**

```
Files assembled successfully
Combined file content:
This is part 1 content.
This is part 2 content.
This is part 3 content.
```

</details>

### SplitFile

Splits a file into multiple files based on a specified size limit.

**Parameters:**

-   `src`: The path to the source file to split
-   `chunkSize`: The maximum size in bytes of each split file

**Returns:**

-   `[]string`: A slice of paths to the created split files
-   `error`: An error if the file couldn't be split

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "strings"
)

func main() {
    fs := ufs.New()

    // Create a larger file to split
    content := strings.Repeat("This is some content that will be repeated many times.\n", 100)
    fs.WriteStringToFile("./large_file.txt", content)

    // Split the file into chunks of roughly 500 bytes each
    chunkSize := int64(500)
    splitFiles, err := fs.SplitFile("./large_file.txt", chunkSize)
    if err != nil {
        fmt.Printf("Error splitting file: %v\n", err)
        return
    }

    fmt.Printf("File split into %d parts:\n", len(splitFiles))
    for i, file := range splitFiles {
        fileSize, _ := fs.GetFileSize(file)
        fmt.Printf("Part %d: %s (size: %d bytes)\n", i+1, file, fileSize)
    }

    // Verify we can reassemble the files to get the original content
    err = fs.AssembleFiles(splitFiles, "./reassembled.txt")
    if err != nil {
        fmt.Printf("Error reassembling files: %v\n", err)
        return
    }

    originalSize, _ := fs.GetFileSize("./large_file.txt")
    reassembledSize, _ := fs.GetFileSize("./reassembled.txt")
    fmt.Printf("Original size: %d bytes, Reassembled size: %d bytes\n", originalSize, reassembledSize)

    if originalSize == reassembledSize {
        fmt.Println("Files match in size - split and reassembly successful")
    }
}
```

**Expected Output:**

```
File split into 10 parts:
Part 1: ./large_file_1.txt (size: 500 bytes)
Part 2: ./large_file_2.txt (size: 500 bytes)
...
Part 10: ./large_file_10.txt (size: 500 bytes)
Original size: 5000 bytes, Reassembled size: 5000 bytes
Files match in size - split and reassembly successful
```

Note: The actual number of parts and their sizes will depend on the content and chunk size.

</details>

### CleanUpFiles

Removes empty files from the given slice of file paths.

**Parameters:**

-   `files`: A slice of file paths to check and potentially remove

**Returns:**

-   `[]string`: A slice of paths to files that were removed
-   `error`: An error if any file couldn't be checked or removed

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

    // Create a mix of empty and non-empty files
    fs.CreateFile("./empty1.txt")
    fs.CreateFile("./empty2.txt")
    fs.WriteStringToFile("./nonempty1.txt", "This file has content")
    fs.CreateFile("./empty3.txt")
    fs.WriteStringToFile("./nonempty2.txt", "This file also has content")

    // Clean up empty files
    filesToCheck := []string{
        "./empty1.txt",
        "./empty2.txt",
        "./nonempty1.txt",
        "./empty3.txt",
        "./nonempty2.txt",
    }

    removedFiles, err := fs.CleanUpFiles(filesToCheck)
    if err != nil {
        fmt.Printf("Error during cleanup: %v\n", err)
        // Continue anyway to see what was removed
    }

    fmt.Printf("Removed %d empty files:\n", len(removedFiles))
    for i, file := range removedFiles {
        fmt.Printf("%d. %s\n", i+1, file)
    }

    // Verify which files still exist
    fmt.Println("\nChecking remaining files:")
    for _, file := range filesToCheck {
        exists := fs.PathExists(file)
        fmt.Printf("%s exists: %t\n", file, exists)
    }
}
```

**Expected Output:**

```
Removed 3 empty files:
1. ./empty1.txt
2. ./empty2.txt
3. ./empty3.txt

Checking remaining files:
./empty1.txt exists: false
./empty2.txt exists: false
./nonempty1.txt exists: true
./empty3.txt exists: false
./nonempty2.txt exists: true
```

</details>

### ReadFileWithLines

Reads a file and returns its content as a slice of strings, each representing a line in the file.

**Parameters:**

-   `path`: The path to the file to read

**Returns:**

-   `[]string`: A slice of strings, each representing a line in the file
-   `error`: An error if the file couldn't be read

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

    // Create a file with multiple lines
    content := "Line 1: First line of the file\nLine 2: Second line\nLine 3: Third line\nLine 4: Last line"
    fs.WriteStringToFile("./multiline.txt", content)

    // Read the file line by line
    lines, err := fs.ReadFileWithLines("./multiline.txt")
    if err != nil {
        fmt.Printf("Error reading file lines: %v\n", err)
        return
    }

    fmt.Printf("File contains %d lines:\n", len(lines))
    for i, line := range lines {
        fmt.Printf("%d: %s\n", i+1, line)
    }
}
```

**Expected Output:**

```
File contains 4 lines:
1: Line 1: First line of the file
2: Line 2: Second line
3: Line 3: Third line
4: Line 4: Last line
```

</details>

### AppendToLastLine

Appends a string to the last line of a file, or creates a new line if the file is empty or doesn't exist.

**Parameters:**

-   `path`: The path to the file
-   `content`: The string to append

**Returns:**

-   `error`: An error if the file couldn't be read or written to

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

    // Create a file with some initial content
    fs.WriteStringToFile("./log_file.txt", "Log entry 1\nLog entry 2\nLog entry 3")

    // Append to the last line (adds a new line)
    err := fs.AppendToLastLine("./log_file.txt", "Log entry 4")
    if err != nil {
        fmt.Printf("Error appending to last line: %v\n", err)
        return
    }

    fmt.Println("Content appended successfully")

    // Verify the file content
    content, _ := fs.ReadFileAsString("./log_file.txt")
    fmt.Printf("File content after append:\n%s\n", content)

    // Try with a non-existent file
    err = fs.AppendToLastLine("./new_log.txt", "First log entry in new file")
    if err != nil {
        fmt.Printf("Error appending to new file: %v\n", err)
        return
    }

    newContent, _ := fs.ReadFileAsString("./new_log.txt")
    fmt.Printf("\nNew file content:\n%s\n", newContent)
}
```

**Expected Output:**

```
Content appended successfully
File content after append:
Log entry 1
Log entry 2
Log entry 3
Log entry 4

New file content:
First log entry in new file
```

</details>

### AppendToFirstLine

Appends a string to the first line of a file, gracefully shifting the current first line to the second line.

**Parameters:**

-   `path`: The path to the file
-   `content`: The string to add as the new first line

**Returns:**

-   `error`: An error if the file couldn't be read or written to

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

    // Create a file with some initial content
    fs.WriteStringToFile("./document.txt", "Original first line\nSecond line\nThird line")

    // Add a new first line
    err := fs.AppendToFirstLine("./document.txt", "New first line")
    if err != nil {
        fmt.Printf("Error appending to first line: %v\n", err)
        return
    }

    fmt.Println("Content prepended successfully")

    // Verify the file content
    content, _ := fs.ReadFileAsString("./document.txt")
    fmt.Printf("File content after prepend:\n%s\n", content)

    // Try with a non-existent file
    err = fs.AppendToFirstLine("./new_document.txt", "First line in new file")
    if err != nil {
        fmt.Printf("Error appending to new file: %v\n", err)
        return
    }

    newContent, _ := fs.ReadFileAsString("./new_document.txt")
    fmt.Printf("\nNew file content:\n%s\n", newContent)
}
```

**Expected Output:**

```
Content prepended successfully
File content after prepend:
New first line
Original first line
Second line
Third line

New file content:
First line in new file
```

</details>

## Error Handling

Unlike other modules in the UFS package, the file-Reader_Writer functions return detailed error information when operations fail. This design decision acknowledges the critical nature of file I/O operations and enables more detailed error handling in applications.

When an error occurs during a file operation, the error is wrapped with contextual information about which function encountered the error, making debugging easier.

<details>
<summary>Error Handling Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Try to read a non-existent file
    _, err := fs.ReadFile("./nonexistent_file.txt")
    if err != nil {
        fmt.Printf("Error details: %v\n", err)

        // You can handle different types of errors
        if os.IsNotExist(err) {
            fmt.Println("The file does not exist - creating it...")
            fs.CreateFile("./nonexistent_file.txt")
        } else if os.IsPermission(err) {
            fmt.Println("Permission denied - check file permissions")
        } else {
            fmt.Println("Unknown error occurred")
        }
    }
}
```

</details>

## Performance Considerations

-   For large files, `ReadFile` and `WriteFile` load the entire file into memory, which can be inefficient
-   When working with very large files, consider using `SplitFile` to break them into manageable chunks
-   For line-by-line processing of large files, `ReadFileWithLines` is more memory-efficient than reading the entire file
-   When combining multiple files, `AssembleFiles` uses buffered I/O for better performance

## Platform Compatibility

These functions are designed to work across different operating systems with consistent behavior:

-   Path separators (backslashes on Windows, forward slashes on Unix) are handled automatically
-   File permissions work differently between Windows and Unix-like systems
-   Directory creation uses appropriate permissions for the platform (0755 on Unix, equivalent on Windows)
-   Line endings (CRLF on Windows, LF on Unix) are preserved when reading and writing text files
