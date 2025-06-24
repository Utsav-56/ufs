# Compress-Extract.go

The `Compress-Extract.go` file provides comprehensive functionality for compressing and extracting files and directories. It offers a unified interface for working with ZIP archives across different platforms, making it easy to package files for storage or transfer and to extract their contents when needed.

## Overview

This module includes functionality for:

-   Compressing directories and files into ZIP archives
-   Extracting the contents of ZIP archives
-   Performing convenience operations like compressing to the current directory
-   Using system-specific compression tools when available

All operations are designed to be platform-independent where possible, with fallbacks to system-specific tools when necessary. Error handling is comprehensive, ensuring reliable operation even in challenging scenarios.

## Basic Functions

### CompressDirectory

Compresses a directory into a ZIP file, preserving the directory structure within the archive.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the directory to compress
-   `destPath`: The absolute or relative path where the ZIP file will be created

**Returns:**

-   `error`: An error if the compression failed, nil otherwise

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

    // Create a directory with some files to compress
    fs.CreateDirectory("./test_dir")
    fs.CreateFileWithContent("./test_dir/file1.txt", "Hello, World!")
    fs.CreateFileWithContent("./test_dir/file2.txt", "This is another test file")
    fs.CreateDirectory("./test_dir/subdir")
    fs.CreateFileWithContent("./test_dir/subdir/file3.txt", "This is in a subdirectory")

    // Compress the directory
    err := fs.CompressDirectory("./test_dir", "./test_dir.zip")
    if err != nil {
        fmt.Printf("Error compressing directory: %v\n", err)
        return
    }

    fmt.Println("Directory compressed successfully")
}
```

**Expected Output:**

```
Directory compressed successfully
```

The resulting ZIP file will contain:

-   file1.txt
-   file2.txt
-   subdir/
-   subdir/file3.txt
</details>

### ExtractArchive

Extracts the contents of a ZIP file to a specified directory.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the ZIP file
-   `destPath`: The absolute or relative path where the contents will be extracted

**Returns:**

-   `error`: An error if the extraction failed, nil otherwise

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

    // First compress a directory to create a test ZIP file
    // (assuming test_dir exists with some files)
    fs.CompressDirectory("./test_dir", "./test_archive.zip")

    // Extract the ZIP file to a new location
    err := fs.ExtractArchive("./test_archive.zip", "./extracted_files")
    if err != nil {
        fmt.Printf("Error extracting archive: %v\n", err)
        return
    }

    fmt.Println("Archive extracted successfully")

    // List the extracted files to verify
    files, _ := fs.GetFileList("./extracted_files")
    for _, file := range files {
        fmt.Printf("Extracted: %s\n", file)
    }
}
```

**Expected Output:**

```
Archive extracted successfully
Extracted: file1.txt
Extracted: file2.txt
```

Note: The actual files listed will depend on the contents of the original archive.

</details>

### CompressFile

Compresses a single file into a ZIP archive.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the file to compress
-   `destPath`: The absolute or relative path where the ZIP file will be created

**Returns:**

-   `error`: An error if the compression failed, nil otherwise

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

    // Create a test file to compress
    fs.CreateFileWithContent("./important.txt", "This is important content that needs to be zipped")

    // Compress the file
    err := fs.CompressFile("./important.txt", "./important.zip")
    if err != nil {
        fmt.Printf("Error compressing file: %v\n", err)
        return
    }

    fmt.Println("File compressed successfully")
}
```

**Expected Output:**

```
File compressed successfully
```

The resulting ZIP file will contain only the specified file (important.txt).

</details>

## Convenience Functions

### CompressHere

Compresses a directory into a ZIP file and outputs it in the current working directory.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the directory to compress

**Returns:**

-   `string`: The path to the created ZIP file
-   `error`: An error if the compression failed, nil otherwise

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

    // Create a test directory
    fs.CreateDirectory("./documents")
    fs.CreateFileWithContent("./documents/doc1.txt", "Document 1 content")
    fs.CreateFileWithContent("./documents/doc2.txt", "Document 2 content")

    // Compress the directory to the current working directory
    zipPath, err := fs.CompressHere("./documents")
    if err != nil {
        fmt.Printf("Error compressing directory: %v\n", err)
        return
    }

    fmt.Printf("Directory compressed to: %s\n", zipPath)
}
```

**Expected Output:**

```
Directory compressed to: /current/working/directory/documents.zip
```

The actual path will reflect your current working directory.

</details>

### ExtractHere

Extracts the contents of a ZIP file in the current working directory.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the ZIP file

**Returns:**

-   `string`: The path to the directory where the archive was extracted
-   `error`: An error if the extraction failed, nil otherwise

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

    // Extract an archive to the current directory
    // (assuming backup.zip exists)
    extractPath, err := fs.ExtractHere("./backup.zip")
    if err != nil {
        fmt.Printf("Error extracting archive: %v\n", err)
        return
    }

    fmt.Printf("Archive extracted to: %s\n", extractPath)
}
```

**Expected Output:**

```
Archive extracted to: /current/working/directory/backup
```

The extraction directory will have the same name as the archive without the extension.

</details>

### CompressFileHere

Compresses a single file into a ZIP file in the current working directory.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the file to compress

**Returns:**

-   `string`: The path to the created ZIP file
-   `error`: An error if the compression failed, nil otherwise

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
    fs.CreateFileWithContent("./report.pdf", "This would be PDF content in a real scenario")

    // Compress the file to the current working directory
    zipPath, err := fs.CompressFileHere("./report.pdf")
    if err != nil {
        fmt.Printf("Error compressing file: %v\n", err)
        return
    }

    fmt.Printf("File compressed to: %s\n", zipPath)
}
```

**Expected Output:**

```
File compressed to: /current/working/directory/report.pdf.zip
```

The ZIP file will be created in the current working directory with the original filename plus .zip extension.

</details>

## Advanced Operations (Use with Caution)

### CompressAndRemove

Compresses a directory into a ZIP file and removes the original directory.

**WARNING: This permanently deletes the source directory. Use with extreme caution.**

**Parameters:**

-   `sourcePath`: The absolute or relative path to the directory to compress and remove
-   `destPath`: The absolute or relative path where the ZIP file will be created

**Returns:**

-   `error`: An error if the operation failed, nil otherwise

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

    // Create a test directory
    fs.CreateDirectory("./temp_files")
    fs.CreateFileWithContent("./temp_files/temp1.txt", "Temporary file 1")
    fs.CreateFileWithContent("./temp_files/temp2.txt", "Temporary file 2")

    // Compress the directory and remove the original
    err := fs.CompressAndRemove("./temp_files", "./temp_archive.zip")
    if err != nil {
        fmt.Printf("Error compressing and removing directory: %v\n", err)
        return
    }

    fmt.Println("Directory compressed and removed successfully")

    // Verify the directory no longer exists
    exists := fs.PathExists("./temp_files")
    fmt.Printf("Original directory still exists: %t\n", exists)
}
```

**Expected Output:**

```
Directory compressed and removed successfully
Original directory still exists: false
```

</details>

### ExtractAndRemove

Extracts a ZIP file and removes the original ZIP file.

**WARNING: This permanently deletes the source ZIP file. Use with extreme caution.**

**Parameters:**

-   `sourcePath`: The absolute or relative path to the ZIP file to extract and remove
-   `destPath`: The absolute or relative path where the contents will be extracted

**Returns:**

-   `error`: An error if the operation failed, nil otherwise

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

    // First create a test archive
    fs.CreateDirectory("./original")
    fs.CreateFileWithContent("./original/file.txt", "Test content")
    fs.CompressDirectory("./original", "./original.zip")

    // Extract the archive and remove the ZIP file
    err := fs.ExtractAndRemove("./original.zip", "./extracted_content")
    if err != nil {
        fmt.Printf("Error extracting and removing archive: %v\n", err)
        return
    }

    fmt.Println("Archive extracted and removed successfully")

    // Verify the ZIP file no longer exists
    exists := fs.PathExists("./original.zip")
    fmt.Printf("Original ZIP file still exists: %t\n", exists)
}
```

**Expected Output:**

```
Archive extracted and removed successfully
Original ZIP file still exists: false
```

</details>

### CompressAndExtract

Compresses a directory and extracts it to a specified location.

**WARNING: This operation is inefficient for most use cases as it creates a temporary ZIP file.**

**Parameters:**

-   `sourcePath`: The absolute or relative path to the directory to compress
-   `tempPath`: The absolute or relative path where the temporary ZIP file will be created
-   `finalPath`: The absolute or relative path where the contents will be extracted

**Returns:**

-   `error`: An error if the operation failed, nil otherwise

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

    // Create a test directory
    fs.CreateDirectory("./source_dir")
    fs.CreateFileWithContent("./source_dir/file1.txt", "Content 1")
    fs.CreateFileWithContent("./source_dir/file2.txt", "Content 2")

    // Compress and extract in one operation
    err := fs.CompressAndExtract("./source_dir", "./temp.zip", "./destination_dir")
    if err != nil {
        fmt.Printf("Error in compress and extract operation: %v\n", err)
        return
    }

    fmt.Println("Directory compressed and extracted successfully")

    // Verify the temporary ZIP file was removed
    exists := fs.PathExists("./temp.zip")
    fmt.Printf("Temporary ZIP file still exists: %t\n", exists)
}
```

**Expected Output:**

```
Directory compressed and extracted successfully
Temporary ZIP file still exists: false
```

</details>

### ExtractAndCompress

Extracts a ZIP file and compresses it to a specified location.

**WARNING: This operation is inefficient for most use cases as it creates a temporary directory.**

**Parameters:**

-   `sourcePath`: The absolute or relative path to the ZIP file to extract
-   `tempPath`: The absolute or relative path where the contents will be temporarily extracted
-   `finalPath`: The absolute or relative path where the new ZIP file will be created

**Returns:**

-   `error`: An error if the operation failed, nil otherwise

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

    // First create a test archive
    fs.CreateDirectory("./original_data")
    fs.CreateFileWithContent("./original_data/data.txt", "Important data")
    fs.CompressDirectory("./original_data", "./original.zip")

    // Extract and recompress in one operation
    err := fs.ExtractAndCompress("./original.zip", "./temp_dir", "./repackaged.zip")
    if err != nil {
        fmt.Printf("Error in extract and compress operation: %v\n", err)
        return
    }

    fmt.Println("Archive extracted and recompressed successfully")

    // Verify the temporary directory was removed
    exists := fs.PathExists("./temp_dir")
    fmt.Printf("Temporary directory still exists: %t\n", exists)
}
```

**Expected Output:**

```
Archive extracted and recompressed successfully
Temporary directory still exists: false
```

</details>

## System Command Operations

### CompressWithSystemCommand

Uses the system's native compression tools to create an archive with various compression formats.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the directory to compress
-   `destPath`: The absolute or relative path where the archive will be created
-   `format`: The compression format to use (e.g., "gzip", "bzip2", "xz")

**Returns:**

-   `error`: An error if the compression failed, nil otherwise

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

    // Create a test directory
    fs.CreateDirectory("./system_test")
    fs.CreateFileWithContent("./system_test/file1.txt", "Test content 1")
    fs.CreateFileWithContent("./system_test/file2.txt", "Test content 2")

    // Compress using system command with gzip format
    err := fs.CompressWithSystemCommand("./system_test", "./system_test.tar.gz", "gzip")
    if err != nil {
        fmt.Printf("Error compressing with system command: %v\n", err)
        return
    }

    fmt.Println("Directory compressed successfully using system command")
}
```

**Expected Output:**

```
Directory compressed successfully using system command
```

Note: This function requires the appropriate system tools to be installed. On Windows, it needs tar.exe which is available in Windows 10 and later.

</details>

### ExtractWithSystemCommand

Uses the system's native extraction tools to extract an archive.

**Parameters:**

-   `sourcePath`: The absolute or relative path to the archive to extract
-   `destPath`: The absolute or relative path where the contents will be extracted

**Returns:**

-   `error`: An error if the extraction failed, nil otherwise

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

    // Extract an archive using system command
    // (assuming system_archive.tar.gz exists)
    err := fs.ExtractWithSystemCommand("./system_archive.tar.gz", "./system_extracted")
    if err != nil {
        fmt.Printf("Error extracting with system command: %v\n", err)
        return
    }

    fmt.Println("Archive extracted successfully using system command")
}
```

**Expected Output:**

```
Archive extracted successfully using system command
```

Note: This function requires the appropriate system tools to be installed. On Windows, it needs tar.exe which is available in Windows 10 and later.

</details>

## Security Considerations

The compression and extraction functions include protection against common security issues:

1. **Zip Slip Vulnerability**: The extraction code includes checks to prevent directory traversal attacks where a malicious archive might try to write files outside the intended extraction directory.

2. **Path Validation**: All paths are validated and converted to absolute paths to ensure consistent behavior.

3. **Resource Management**: Files are properly closed using defer statements to prevent resource leaks.

## Performance Considerations

-   For large directories, compression and extraction operations can be memory and CPU intensive.
-   Consider using streaming approaches for very large files if memory usage is a concern.
-   The system command operations may be more efficient for large archives as they leverage optimized native tools.

## Platform Compatibility

-   All basic ZIP operations work across platforms (Windows, macOS, Linux).
-   The system command operations (`CompressWithSystemCommand` and `ExtractWithSystemCommand`) require:
    -   On Windows: Windows 10 or later with tar.exe available
    -   On Unix-like systems: The tar command must be installed
-   Different compression formats (gzip, bzip2, xz) may have different system requirements.
