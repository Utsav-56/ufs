# UFS - Ultimate File System Utils

A comprehensive Go library that provides abstracted functions for handling file system operations. UFS simplifies common file operations with a clean, consistent API that works across different platforms.

## Installation

```bash
go get github.com/utsav-56/ufs
```

## Usage

```go
import "github.com/utsav-56/ufs"

```

The `ufs` package provides flexible file and directory utilities that can be used either as static functions or via an instance of the `UFS` struct. Below are detailed usage instructions for both approaches.

---

### 1. Using Static Methods

You can call utility functions directly from the `ufs` package without creating an instance. This is convenient for quick checks or scripts.

**Examples:**

#### Path Properties

```go
import "github.com/utsav-56/ufs"


func main(){

if ufs.PathExists("myfile.txt") {
    // File exists
}

if ufs.IsDirectory("myfolder") {
    // It's a directory
}


size := ufs.GetFileSize("data.txt")
fmt.Println("File size:", size)

folders := ufs.GetFolderList("./somefolder")
fmt.Println("Folders:", folders)


ufs.CompressDirectory("./mydir", "./mydir.zip")
ufs.ExtractArchive("./mydir.zip", "./outputdir")
}

```

All of the functions can be used as a static functions using ufs or by creatig a new instance of UFS using NewUFs and you can pass your customization options there

## 2. Using via instances

You can create an instance of the `UFS` struct to use all utilities with customizable options. This approach is recommended for better IDE support and advanced configuration.

**Example:**

```go
package main

import (
    "fmt"
    "io/fs"
    "github.com/utsav-56/ufs"
)

func main() {
    // Create a new UFS instance with custom options
    opts := &ufs.Options{
        ShowError:      true,
        ReturnReadable: false,
        // prettifyError is internal; set via struct if needed
    }
    u := ufs.NewUfs(opts)

    // File creation examples
    ok := u.CreateFile("example_file.txt")
    fmt.Printf("CreateFile: %v\n", ok)

    ok = u.CreateFileWithContent("example_content.txt", "Hello, World!")
    fmt.Printf("CreateFileWithContent: %v\n", ok)

    ok = u.CreateFileWithContentAndPermissions("example_content_perm.txt", "Hello, Permissions!", 0644)
    fmt.Printf("CreateFileWithContentAndPermissions: %v\n", ok)

    ok = u.CreateFileWithPermissions("example_perm.txt", 0644)
    fmt.Printf("CreateFileWithPermissions: %v\n", ok)

    // Directory creation examples
    ok = u.CreateDirectory("example_dir")
    fmt.Printf("CreateDirectory: %v\n", ok)

    ok = u.CreateDirectoryWithPermissions("example_dir_perm", 0755)
    fmt.Printf("CreateDirectoryWithPermissions: %v\n", ok)

    // Link creation examples
    ok = u.CreateSymlink("example_file.txt", "example_symlink.txt")
    fmt.Printf("CreateSymlink: %v\n", ok)

    ok = u.CreateHardLink("example_file.txt", "example_hardlink.txt")
    fmt.Printf("CreateHardLink: %v\n", ok)

    // Directory tree creation
    structure := map[string]interface{}{
        "dir1": nil,
        "dir2": map[string]interface{}{
            "subdir1": nil,
            "subdir2": map[string]interface{}{
                "subsubdir": nil,
            },
        },
    }
    ok = u.CreateDirectoryTree("example_tree", structure)
    fmt.Printf("CreateDirectoryTree: %v\n", ok)

    // Directory tree with permissions
    structurePerm := map[string]interface{}{
        "dirA": nil,
        "dirB": map[string]interface{}{
            "subdirA": nil,
        },
    }
    ok = u.CreateDirectoryTreeWithPermissions("example_tree_perm", structurePerm, fs.FileMode(0755))
    fmt.Printf("CreateDirectoryTreeWithPermissions: %v\n", ok)

    // Symlink directory tree
    ok = u.SymlinkDirectoryTree("example_tree", "example_tree_symlinks", true)
    fmt.Printf("SymlinkDirectoryTree: %v\n", ok)
}
```

You can also change options at runtime using:

```go
u.SetOptions(&ufs.Options{
    ShowError:      false,
    ReturnReadable: true,
})
```

> Note it is recommended to use via instance because of documentations IDE like Vscode or Jetbrains shows proper documentation and usage example if used instance version

## Documentation

### Modules and Utility Functions

#### `PathProperties.go`

-   Detailed documentation of pathProperties can be found here: [path-properties.go.md](documentations/path-properties.go.md)

Provides utilities for working with file and directory paths.

**Basic Path Checks:**

-   **PathExists:** Check if a path exists.
-   **IsFile:** Check if a path is a file.
-   **IsDirectory:** Check if a path is a directory.
-   **IsDirectoryEmpty:** Check if a directory is empty.
-   **IsFileEmpty:** Check if a file is empty.

**Location Checks:**

-   **IsInSystemPath:** Check if a path is in a system directory.
-   **IsInUserPath:** Check if a path is in the user's home directory.
-   **IsInCurrentPath:** Check if a path is in the current working directory.

**File Permission and Attribute Checks:**

-   **IsFileHidden:** Check if a file is hidden.
-   **IsFileExecutable:** Check if a file is executable.
-   **IsFileReadable:** Check if a file is readable.
-   **IsFileWritable:** Check if a file is writable.

**Directory Permission and Attribute Checks:**

-   **IsDirectoryHidden:** Check if a directory is hidden.
-   **IsDirectoryReadable:** Check if a directory is readable.
-   **IsDirectoryWritable:** Check if a directory is writable.

Each function includes comprehensive documentation with descriptions, parameters, return values, and usage examples. The implementation handles Windows and Unix-like systems differently where necessary, especially for hidden file detection and executable status checking.

#### `Creations.go`

-   Detailed documentation of creations module can be found here: [creations.md](documentations/creations.md)

Provides utilities for creating files, directories, and links.

**Basic File Creation:**

-   **CreateFile:** Creates an empty file.
-   **CreateFileWithContent:** Creates a file with specified content.
-   **CreateFileWithPermissions:** Creates an empty file with specified permissions.
-   **CreateFileWithContentAndPermissions:** Creates a file with specified content and permissions.

**Directory Creation:**

-   **CreateDirectory:** Creates a directory with default permissions.
-   **CreateDirectoryWithPermissions:** Creates a directory with specified permissions.

**Link Creation:**

-   **CreateSymlink:** Creates a symbolic link.
-   **CreateHardLink:** Creates a hard link.

**Advanced Directory Tree Operations:**

-   **CreateDirectoryTree:** Creates a hierarchical directory structure.
-   **CreateDirectoryTreeWithPermissions:** Creates a hierarchical directory structure with specified permissions.
-   **SymlinkDirectoryTree:** Creates symbolic links for an entire directory tree.

#### `Removing.go`

-   Detailed documentation of Removing module can be found here: [removing.go.md](documentations/removing.go.md)

Provides utilities for removing files, directories, and links with safety and flexibility.

**Basic Removal Operations:**

-   **RemoveFile:** Removes a single file.
-   **RemoveDirectory:** Removes an empty directory.
-   **RemoveDirectoryRecursive:** Removes a directory and all its contents.
-   **RemoveSymlink:** Removes a symbolic link.

**Enhanced Removal Operations:**

-   **RemoveFileWithBackup:** Creates a backup before removing a file.
-   **RemoveEmptyFiles:** Removes all empty files in a directory.
-   **RemoveEmptyDirectories:** Removes all empty directories in a directory.
-   **RemoveDirectoryContents:** Empties a directory without removing it.
-   **RemoveDirectoryTree:** Removes a directory structure matching a provided map.
-   **RemoveAllLinks:** Removes all symbolic links in a directory.
-   **RemoveByPattern:** Removes files matching a pattern.
-   **SafeRemoveFile:** Removes a file only if it matches expected criteria.

Each function includes comprehensive documentation with:

-   Description of the function
-   Parameter descriptions
-   Return value descriptions
-   Usage examples

All functions follow consistent error handling and validation patterns to prevent accidental data loss. For example, `RemoveFile` checks that the path is a file before removal, and `RemoveDirectory` ensures the directory is empty. Error handling is consistent with the rest of the library, using `handleError` and `handleMistakeWarning` methods.

#### `File-Reader_writer.go`

-   Detailed documentation of `File Reader writer` module can be found here: [file-reader_writer.go.md](documentations/file-Reader_writer.go.md)

Provides utilities for reading, writing, appending, and transferring files.

**Basic File Operations:**

-   **ReadFile:** Reads file content as bytes.
-   **ReadFileAsString:** Reads file content as a string.
-   **WriteFile:** Writes bytes to a file.
-   **WriteStringToFile:** Writes a string to a file.
-   **AppendToFile:** Appends bytes to a file.
-   **AppendStringToFile:** Appends a string to a file.
-   **DeleteFile:** Removes a file.

**File Transfer Operations:**

-   **CopyFile:** Copies a file to another location.
-   **MoveFile:** Moves a file to another location.
-   **CopyFileWithPermissions:** Copies a file while preserving permissions.
-   **MoveFileWithPermissions:** Moves a file while preserving permissions.

**Advanced File Operations:**

-   **AssembleFiles:** Combines multiple files into one.
-   **SplitFile:** Splits a file into multiple smaller files.
-   **CleanUpFiles:** Removes empty files from a provided list.
-   **ReadFileWithLines:** Reads a file as an array of lines.
-   **AppendToLastLine:** Adds content to the end of a file.
-   **AppendToFirstLine:** Adds content to the beginning of a file.

Each function includes:

-   Preconditions checks (e.g., file existence)
-   Detailed documentation (parameters, return values, usage examples)
-   Consistent error handling with a common error wrapping function
-   Automatic creation of parent directories when needed
-   Use of appropriate file permissions (`0644` for files, `0755` for directories)

#### `Compress-Extract.go`

-   Detailed documentation of `compressing` module can be found here: [compressing.go.md](documentations/compressing.go.md)

Provides utilities for compressing and extracting files and directories, supporting both pure Go and system command approaches.

**Core Functions:**

-   **CompressDirectory:** Compresses a directory into a ZIP file.
-   **ExtractArchive:** Extracts the contents of a ZIP file to a specified directory.
-   **CompressFile:** Compresses a single file into a ZIP file.

**Convenience Functions:**

-   **CompressHere:** Compresses a directory and outputs the ZIP file to the current working directory.
-   **ExtractHere:** Extracts an archive to the current working directory.
-   **CompressFileHere:** Compresses a file and outputs the ZIP file to the current working directory.

**Advanced (Dangerous) Functions:**

-   **CompressAndRemove:** Compresses a directory and removes the original.
-   **ExtractAndRemove:** Extracts an archive and removes the original.
-   **CompressAndExtract:** Compresses a directory and then extracts it elsewhere.
-   **ExtractAndCompress:** Extracts an archive and then compresses it elsewhere.

**System Command Functions:**

-   **CompressWithSystemCommand:** Uses system tools for compression with formats like gzip, bzip2, and xz.
-   **ExtractWithSystemCommand:** Uses system tools for extraction.

Each function includes:

-   Detailed documentation explaining its purpose
-   Parameter and return value descriptions
-   Example usage code
-   Proper error handling

The implementation uses Go's standard `archive/zip` package for pure Go operations and system commands for additional formats. Platform differences between Windows and Unix-like systems are handled, and safety checks (such as protection against zip slip vulnerabilities) are included.

### `Move-Rename_Delete.go`

-   Detailed documentation of pathProperties can be found here: [Move_Rename_Delete.go.md](documentations/Move_Rename_Delete.go.md)

Provides utilities for moving, renaming, and deleting files and directories, with robust handling of edge cases and consistent error management.

**Basic Operations:**

-   **MoveFile:** Moves or renames a file, creating the destination directory if needed.
-   **DeleteFile:** Removes a file (wrapper around `RemoveFile` for consistent naming).
-   **DeleteDirectory:** Removes a directory and its contents (wrapper around `RemoveDirectoryRecursive`).
-   **MoveDirectory:** Moves or renames a directory, merging contents if the destination exists.

**Conditional Operations:**

-   **MoveFileIfExists:** Moves a file only if the source exists.
-   **MoveDirectoryIfExists:** Moves a directory only if the source exists.
-   **DeleteFileIfExists:** Deletes a file only if it exists.
-   **DeleteDirectoryIfExists:** Deletes a directory only if it exists.

**Empty File/Directory Operations:**

-   **MoveFileIfEmpty:** Moves a file only if it is empty (zero bytes).
-   **MoveDirectoryIfEmpty:** Moves a directory only if it is empty.
-   **DeleteFileIfEmpty:** Deletes a file only if it is empty.
-   **DeleteDirectoryIfEmpty:** Deletes a directory only if it is empty.

**Renaming Operations:**

-   **RenameFile:** Renames a file within the same directory.
-   **RenameDirectory:** Renames a directory in place.

**Backup Operations:**

-   **MoveWithBackup:** Creates a backup of the destination before moving.
-   **DeleteWithBackup:** Creates a backup before deletion.

**Helper Functions:**

-   **copyThenDelete:** Copies a file and deletes the source (for cross-filesystem moves).
-   **mergeDirectories:** Merges the contents of two directories.
-   **copyDirectoryRecursive:** Recursively copies directory structures.

Each function includes:

-   Clear documentation explaining its purpose
-   Parameter and return value descriptions
-   Example usage code
-   Proper error handling using the existing `handleError` method

The implementation handles edge cases such as:

-   Cross-filesystem moves
-   Directory merging
-   Path validation
-   Handling empty vs. non-empty files/directories
-   Creating parent directories as needed
-   Backup creation and restoration if operations fail

All functions return boolean values to indicate success or failure, with error logging managed by the UFS instance.

> Each function includes detailed documentation with:
>
> -   Description of the function
> -   Parameter descriptions
> -   Return value descriptions
> -   Usage examples
>     ` `

## License

MIT

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
