# Path-properties.go

The `Path-properties.go` file provides a comprehensive set of functions for checking the properties and attributes of files and directories in the file system. These functions allow you to determine existence, type, permissions, and other characteristics of paths with a simple, unified interface.

## Overview

This module includes functionality for:

-   Checking if paths exist and determining their type (file or directory)
-   Verifying if files or directories are empty
-   Determining if paths are located in system, user, or current directories
-   Checking file and directory attributes such as hidden status
-   Verifying permissions including read, write, and execute access

All functions handle errors internally and return boolean values to indicate the result, making them easy to use in conditional statements.

## Basic Path Functions

### PathExists

Checks if a file or directory exists at the specified path.

**Parameters:**

-   `path`: The absolute or relative path to check

**Returns:**

-   `bool`: True if the path exists, false otherwise

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

    // Check if a path exists
    if fs.PathExists("/path/to/something") {
        fmt.Println("Path exists")
    } else {
        fmt.Println("Path does not exist")
    }

    // Create a file and check again
    fs.CreateFile("./test_file.txt")

    if fs.PathExists("./test_file.txt") {
        fmt.Println("Test file exists")
    } else {
        fmt.Println("Test file does not exist")
    }
}
```

**Expected Output:**

```
Path does not exist
Test file exists
```

</details>

### IsFile

Checks if the specified path points to a regular file.

**Parameters:**

-   `path`: The absolute or relative path to check

**Returns:**

-   `bool`: True if the path exists and is a regular file, false otherwise

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

    // Create a file and a directory
    fs.CreateFile("./just_a_file.txt")
    fs.CreateDirectory("./just_a_directory")

    // Check if paths are files
    fmt.Printf("./just_a_file.txt is a file: %t\n", fs.IsFile("./just_a_file.txt"))
    fmt.Printf("./just_a_directory is a file: %t\n", fs.IsFile("./just_a_directory"))
    fmt.Printf("./nonexistent_path is a file: %t\n", fs.IsFile("./nonexistent_path"))
}
```

**Expected Output:**

```
./just_a_file.txt is a file: true
./just_a_directory is a file: false
./nonexistent_path is a file: false
```

</details>

### IsDirectory

Checks if the specified path points to a directory.

**Parameters:**

-   `path`: The absolute or relative path to check

**Returns:**

-   `bool`: True if the path exists and is a directory, false otherwise

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

    // Create a file and a directory
    fs.CreateFile("./not_a_directory.txt")
    fs.CreateDirectory("./yes_a_directory")

    // Check if paths are directories
    fmt.Printf("./not_a_directory.txt is a directory: %t\n", fs.IsDirectory("./not_a_directory.txt"))
    fmt.Printf("./yes_a_directory is a directory: %t\n", fs.IsDirectory("./yes_a_directory"))
    fmt.Printf("./nonexistent_path is a directory: %t\n", fs.IsDirectory("./nonexistent_path"))
}
```

**Expected Output:**

```
./not_a_directory.txt is a directory: false
./yes_a_directory is a directory: true
./nonexistent_path is a directory: false
```

</details>

## Empty Status Functions

### IsDirectoryEmpty

Checks if the specified directory is empty.

**Parameters:**

-   `path`: The absolute or relative path to the directory

**Returns:**

-   `bool`: True if the directory exists and is empty, false otherwise

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

    // Create empty and non-empty directories
    fs.CreateDirectory("./empty_dir")
    fs.CreateDirectory("./non_empty_dir")
    fs.CreateFile("./non_empty_dir/file.txt")

    // Check if directories are empty
    fmt.Printf("./empty_dir is empty: %t\n", fs.IsDirectoryEmpty("./empty_dir"))
    fmt.Printf("./non_empty_dir is empty: %t\n", fs.IsDirectoryEmpty("./non_empty_dir"))
    fmt.Printf("./nonexistent_dir is empty: %t\n", fs.IsDirectoryEmpty("./nonexistent_dir"))
}
```

**Expected Output:**

```
./empty_dir is empty: true
./non_empty_dir is empty: false
./nonexistent_dir is empty: false
```

</details>

### IsFileEmpty

Checks if the specified file is empty (zero bytes).

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `bool`: True if the file exists and is empty, false otherwise

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

    // Create empty and non-empty files
    fs.CreateFile("./empty_file.txt")
    fs.CreateFileWithContent("./non_empty_file.txt", "This file has content")

    // Check if files are empty
    fmt.Printf("./empty_file.txt is empty: %t\n", fs.IsFileEmpty("./empty_file.txt"))
    fmt.Printf("./non_empty_file.txt is empty: %t\n", fs.IsFileEmpty("./non_empty_file.txt"))
    fmt.Printf("./nonexistent_file.txt is empty: %t\n", fs.IsFileEmpty("./nonexistent_file.txt"))
}
```

**Expected Output:**

```
./empty_file.txt is empty: true
./non_empty_file.txt is empty: false
./nonexistent_file.txt is empty: false
```

</details>

## Path Location Functions

### IsInSystemPath

Checks if the specified path is in one of the system directories.

**Parameters:**

-   `path`: The absolute or relative path to check

**Returns:**

-   `bool`: True if the path is in a system directory, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "runtime"
)

func main() {
    fs := ufs.New()

    // Check system paths specific to the operating system
    var systemPath string

    if runtime.GOOS == "windows" {
        systemPath = "C:\\Windows\\System32\\notepad.exe"
    } else {
        systemPath = "/usr/bin/bash"
    }

    userPath := "./my_file.txt"

    fmt.Printf("%s is in system path: %t\n", systemPath, fs.IsInSystemPath(systemPath))
    fmt.Printf("%s is in system path: %t\n", userPath, fs.IsInSystemPath(userPath))
}
```

**Expected Output (on Windows):**

```
C:\Windows\System32\notepad.exe is in system path: true
./my_file.txt is in system path: false
```

**Expected Output (on Linux/macOS):**

```
/usr/bin/bash is in system path: true
./my_file.txt is in system path: false
```

</details>

### IsInUserPath

Checks if the specified path is in the user's home directory.

**Parameters:**

-   `path`: The absolute or relative path to check

**Returns:**

-   `bool`: True if the path is in the user's home directory, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "os/user"
    "path/filepath"
)

func main() {
    fs := ufs.New()

    // Get the current user's home directory
    currentUser, _ := user.Current()
    homeDir := currentUser.HomeDir

    // Check if paths are in the user's home directory
    docPath := filepath.Join(homeDir, "Documents", "file.txt")
    systemPath := "/etc/hosts"

    fmt.Printf("%s is in user path: %t\n", docPath, fs.IsInUserPath(docPath))
    fmt.Printf("%s is in user path: %t\n", systemPath, fs.IsInUserPath(systemPath))
}
```

**Expected Output (varies by system):**

```
/home/username/Documents/file.txt is in user path: true
/etc/hosts is in user path: false
```

</details>

### IsInCurrentPath

Checks if the specified path is in the current working directory.

**Parameters:**

-   `path`: The absolute or relative path to check

**Returns:**

-   `bool`: True if the path is in the current working directory, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "os"
    "path/filepath"
)

func main() {
    fs := ufs.New()

    // Get current working directory
    cwd, _ := os.Getwd()

    // Check if paths are in the current directory
    localPath := "./local_file.txt"
    absoluteLocalPath := filepath.Join(cwd, "local_file.txt")
    absoluteOtherPath := "/tmp/other_file.txt"

    fmt.Printf("%s is in current path: %t\n", localPath, fs.IsInCurrentPath(localPath))
    fmt.Printf("%s is in current path: %t\n", absoluteLocalPath, fs.IsInCurrentPath(absoluteLocalPath))
    fmt.Printf("%s is in current path: %t\n", absoluteOtherPath, fs.IsInCurrentPath(absoluteOtherPath))
}
```

**Expected Output:**

```
./local_file.txt is in current path: true
/home/username/project/local_file.txt is in current path: true
/tmp/other_file.txt is in current path: false
```

</details>

## File Attribute Functions

### IsFileHidden

Checks if a file is hidden according to the OS conventions.

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `bool`: True if the file exists and is hidden, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "runtime"
)

func main() {
    fs := ufs.New()

    // Create regular and hidden files based on OS
    fs.CreateFile("./regular_file.txt")

    var hiddenFile string
    if runtime.GOOS == "windows" {
        // On Windows, need to use attrib command to set the hidden attribute
        hiddenFile = "./hidden_file.txt"
        fs.CreateFile(hiddenFile)
        // This would require running: attrib +h hidden_file.txt
    } else {
        // On Unix-like systems, files starting with . are hidden
        hiddenFile = "./.hidden_file"
        fs.CreateFile(hiddenFile)
    }

    fmt.Printf("./regular_file.txt is hidden: %t\n", fs.IsFileHidden("./regular_file.txt"))
    fmt.Printf("%s is hidden: %t\n", hiddenFile, fs.IsFileHidden(hiddenFile))
}
```

**Expected Output (Unix-like systems):**

```
./regular_file.txt is hidden: false
./.hidden_file is hidden: true
```

Note: On Windows, you would need to manually set the hidden attribute using the system's `attrib` command for this example to show a true result for the hidden file.

</details>

### IsFileExecutable

Checks if a file is executable by the current user.

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `bool`: True if the file exists and is executable, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "os"
    "runtime"
)

func main() {
    fs := ufs.New()

    // Create files with different executability
    regularFile := "./regular_file.txt"
    fs.CreateFile(regularFile)

    var executableFile string
    if runtime.GOOS == "windows" {
        executableFile = "./executable.bat"
    } else {
        executableFile = "./executable.sh"
    }

    fs.CreateFile(executableFile)

    // On Unix-like systems, make the file executable
    if runtime.GOOS != "windows" {
        os.Chmod(executableFile, 0755)
    }

    fmt.Printf("%s is executable: %t\n", regularFile, fs.IsFileExecutable(regularFile))
    fmt.Printf("%s is executable: %t\n", executableFile, fs.IsFileExecutable(executableFile))
}
```

**Expected Output (varies by system):**

On Unix-like systems:

```
./regular_file.txt is executable: false
./executable.sh is executable: true
```

On Windows:

```
./regular_file.txt is executable: false
./executable.bat is executable: true
```

</details>

### IsFileReadable

Checks if a file is readable by the current user.

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `bool`: True if the file exists and is readable, false otherwise

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

    // Create a readable file
    readablePath := "./readable_file.txt"
    fs.CreateFile(readablePath)

    // Check if the file is readable
    fmt.Printf("%s is readable: %t\n", readablePath, fs.IsFileReadable(readablePath))

    // On Unix-like systems, we could create a non-readable file
    if os.Getuid() == 0 { // Only if running as root/admin
        nonReadablePath := "./non_readable_file.txt"
        fs.CreateFile(nonReadablePath)
        os.Chmod(nonReadablePath, 0200) // Write-only permission

        fmt.Printf("%s is readable: %t\n", nonReadablePath, fs.IsFileReadable(nonReadablePath))
    }
}
```

**Expected Output:**

```
./readable_file.txt is readable: true
```

If running as root/admin on Unix-like systems with the second check:

```
./readable_file.txt is readable: true
./non_readable_file.txt is readable: false
```

</details>

### IsFileWritable

Checks if a file is writable by the current user.

**Parameters:**

-   `path`: The absolute or relative path to the file

**Returns:**

-   `bool`: True if the file exists and is writable, false otherwise

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

    // Create a writable file
    writablePath := "./writable_file.txt"
    fs.CreateFile(writablePath)

    // Check if the file is writable
    fmt.Printf("%s is writable: %t\n", writablePath, fs.IsFileWritable(writablePath))

    // On Unix-like systems, we could create a read-only file
    if os.Getuid() == 0 { // Only if running as root/admin
        readOnlyPath := "./read_only_file.txt"
        fs.CreateFile(readOnlyPath)
        os.Chmod(readOnlyPath, 0444) // Read-only permission

        fmt.Printf("%s is writable: %t\n", readOnlyPath, fs.IsFileWritable(readOnlyPath))
    }
}
```

**Expected Output:**

```
./writable_file.txt is writable: true
```

If running as root/admin on Unix-like systems with the second check:

```
./writable_file.txt is writable: true
./read_only_file.txt is writable: false
```

</details>

## Directory Attribute Functions

### IsDirectoryHidden

Checks if a directory is hidden according to the OS conventions.

**Parameters:**

-   `path`: The absolute or relative path to the directory

**Returns:**

-   `bool`: True if the directory exists and is hidden, false otherwise

<details>
<summary>Usage Example</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
    "runtime"
)

func main() {
    fs := ufs.New()

    // Create regular and hidden directories based on OS
    fs.CreateDirectory("./regular_dir")

    var hiddenDir string
    if runtime.GOOS == "windows" {
        // On Windows, need to use attrib command to set the hidden attribute
        hiddenDir = "./hidden_dir"
        fs.CreateDirectory(hiddenDir)
        // This would require running: attrib +h hidden_dir
    } else {
        // On Unix-like systems, directories starting with . are hidden
        hiddenDir = "./.hidden_dir"
        fs.CreateDirectory(hiddenDir)
    }

    fmt.Printf("./regular_dir is hidden: %t\n", fs.IsDirectoryHidden("./regular_dir"))
    fmt.Printf("%s is hidden: %t\n", hiddenDir, fs.IsDirectoryHidden(hiddenDir))
}
```

**Expected Output (Unix-like systems):**

```
./regular_dir is hidden: false
./.hidden_dir is hidden: true
```

Note: On Windows, you would need to manually set the hidden attribute using the system's `attrib` command for this example to show a true result for the hidden directory.

</details>

### IsDirectoryReadable

Checks if a directory is readable by the current user.

**Parameters:**

-   `path`: The absolute or relative path to the directory

**Returns:**

-   `bool`: True if the directory exists and is readable, false otherwise

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

    // Create a readable directory
    readablePath := "./readable_dir"
    fs.CreateDirectory(readablePath)

    // Check if the directory is readable
    fmt.Printf("%s is readable: %t\n", readablePath, fs.IsDirectoryReadable(readablePath))

    // On Unix-like systems, we could create a non-readable directory
    if os.Getuid() == 0 { // Only if running as root/admin
        nonReadablePath := "./non_readable_dir"
        fs.CreateDirectory(nonReadablePath)
        os.Chmod(nonReadablePath, 0300) // Write and execute permission, but no read

        fmt.Printf("%s is readable: %t\n", nonReadablePath, fs.IsDirectoryReadable(nonReadablePath))
    }
}
```

**Expected Output:**

```
./readable_dir is readable: true
```

If running as root/admin on Unix-like systems with the second check:

```
./readable_dir is readable: true
./non_readable_dir is readable: false
```

</details>

## Path Safety Functions

These functions help verify path properties before performing potentially destructive operations:

<details>
<summary>Example: Combining Path Property Checks for Safe Operations</summary>

```go
package main

import (
    "fmt"
    "github.com/yourusername/ufs"
)

func main() {
    fs := ufs.New()

    // Safely delete a file only if it's not in system path and is writable
    filePath := "./temp_file.txt"
    fs.CreateFile(filePath)

    if !fs.IsInSystemPath(filePath) && fs.IsFileWritable(filePath) {
        fs.RemoveFile(filePath)
        fmt.Println("File safely deleted")
    } else {
        fmt.Println("File not deleted - safety checks failed")
    }

    // Safely create a file only if the parent directory is writable
    newFilePath := "./output/data.txt"
    dirPath := "./output"

    fs.CreateDirectory(dirPath)

    if fs.IsDirectory(dirPath) && fs.IsDirectoryReadable(dirPath) {
        fs.CreateFile(newFilePath)
        fmt.Println("File safely created")
    } else {
        fmt.Println("File not created - safety checks failed")
    }
}
```

</details>

## Performance Considerations

-   Path property functions are designed to be efficient, but they do require system calls which can add up if called repeatedly
-   For operations that check multiple properties on the same path, consider caching the results rather than calling each function separately
-   The `os.Stat` operation, which many of these functions use internally, can be a bottleneck in high-volume file operations

## Platform Compatibility

These functions are designed to work across different operating systems, but with some important differences:

-   Hidden file/directory detection is handled differently on Windows vs. Unix-like systems
-   Executable file detection on Windows is based on file extensions, while on Unix it's based on permission bits
-   Permission checking may behave differently across platforms, especially with complex permission systems like ACLs
-   Windows paths may use backslashes (`\`) while Unix uses forward slashes (`/`), but these functions handle the conversion internally
