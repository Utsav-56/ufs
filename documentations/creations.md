# Creations.go

The `Creations.go` file provides functions to create files, directories, and links in the UFS (Ultimate File System). These functions make it easy to create and manage file system resources with a simple, unified interface.

## Overview

This module includes functionality for:

-   Creating empty files and files with content
-   Creating directories with custom permissions
-   Creating symbolic and hard links
-   Building complex directory structures

All functions handle errors internally and return boolean values to indicate success or failure, making them easy to use in conditional statements.

## Available Functions

### CreateFile

Creates a new empty file at the specified path. If the file already exists, it will be truncated to zero length.

**Parameters:**

-   `path`: The absolute or relative path to the file to create

**Returns:**

-   `bool`: true if the file was created successfully, false otherwise

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
    success := fs.CreateFile("./example.txt")
    if success {
        fmt.Println("File created successfully")
    } else {
        fmt.Println("Failed to create file")
    }
}
```

**Expected Output:**

```
File created successfully
```

</details>

### CreateFileWithContent

Creates a new file at the specified path with the given content. If the file already exists, it will be overwritten.

**Parameters:**

-   `path`: The absolute or relative path to the file to create
-   `content`: The content to write to the file as a string

**Returns:**

-   `bool`: true if the file was created and written successfully, false otherwise

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
    content := "Hello, World!\nThis is a test file."
    success := fs.CreateFileWithContent("./hello.txt", content)
    if success {
        fmt.Println("File created with content successfully")
    } else {
        fmt.Println("Failed to create file with content")
    }
}
```

**Expected Output:**

```
File created with content successfully
```

If you check the file contents:

```
Hello, World!
This is a test file.
```

</details>

### CreateFileWithContentAndPermissions

Creates a new file at the specified path with the given content and permissions. If the file already exists, it will be overwritten.

**Parameters:**

-   `path`: The absolute or relative path to the file to create
-   `content`: The content to write to the file as a string
-   `perm`: The file permissions (e.g., 0644 for read/write by owner, read-only for others)

**Returns:**

-   `bool`: true if the file was created and written successfully, false otherwise

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
    content := "This is a secure file with custom permissions."
    success := fs.CreateFileWithContentAndPermissions("./secure.txt", content, 0600)
    if success {
        fmt.Println("Secure file created successfully")
    } else {
        fmt.Println("Failed to create secure file")
    }
}
```

**Expected Output:**

```
Secure file created successfully
```

If you check the file permissions (on Unix-like systems):

```
-rw------- 1 user user 43 Jun 24 14:30 secure.txt
```

</details>

### CreateFileWithPermissions

Creates a new empty file at the specified path with the given permissions. If the file already exists, it will be truncated to zero length.

**Parameters:**

-   `path`: The absolute or relative path to the file to create
-   `perm`: The file permissions (e.g., 0644 for read/write by owner, read-only for others)

**Returns:**

-   `bool`: true if the file was created successfully, false otherwise

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
    success := fs.CreateFileWithPermissions("./executable.sh", 0755)
    if success {
        fs.WriteStringToFile("./executable.sh", "#!/bin/bash\necho 'Hello from bash script'")
        fmt.Println("Executable file created successfully")
    } else {
        fmt.Println("Failed to create executable file")
    }
}
```

**Expected Output:**

```
Executable file created successfully
```

If you check the file permissions (on Unix-like systems):

```
-rwxr-xr-x 1 user user 40 Jun 24 14:32 executable.sh
```

</details>

### CreateDirectory

Creates a new directory at the specified path. If the directory already exists, no error is returned.

**Parameters:**

-   `path`: The absolute or relative path to the directory to create

**Returns:**

-   `bool`: true if the directory was created successfully or already exists, false otherwise

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
    success := fs.CreateDirectory("./my_directory")
    if success {
        fmt.Println("Directory created successfully")
    } else {
        fmt.Println("Failed to create directory")
    }
}
```

**Expected Output:**

```
Directory created successfully
```

</details>

### CreateDirectoryWithPermissions

Creates a new directory at the specified path with the given permissions. If the directory already exists, no error is returned but permissions won't be changed.

**Parameters:**

-   `path`: The absolute or relative path to the directory to create
-   `perm`: The directory permissions (e.g., 0755 for read/write/execute by owner, read/execute for others)

**Returns:**

-   `bool`: true if the directory was created successfully, false otherwise

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
    success := fs.CreateDirectoryWithPermissions("./private_dir", 0700)
    if success {
        fmt.Println("Private directory created successfully")
    } else {
        fmt.Println("Failed to create private directory")
    }
}
```

**Expected Output:**

```
Private directory created successfully
```

If you check the directory permissions (on Unix-like systems):

```
drwx------ 2 user user 4096 Jun 24 14:35 private_dir
```

</details>

### CreateSymlink

Creates a symbolic link at the specified path pointing to the target.

**Parameters:**

-   `target`: The file or directory that the symlink will point to
-   `symlink`: The path where the symlink will be created

**Returns:**

-   `bool`: true if the symlink was created successfully, false otherwise

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

    // First create a target file
    fs.CreateFileWithContent("./original.txt", "This is the original file")

    // Create a symlink to the target
    success := fs.CreateSymlink("./original.txt", "./link_to_original.txt")
    if success {
        fmt.Println("Symlink created successfully")
    } else {
        fmt.Println("Failed to create symlink")
    }
}
```

**Expected Output:**

```
Symlink created successfully
```

If you check the link (on Unix-like systems):

```
link_to_original.txt -> original.txt
```

</details>

### CreateHardLink

Creates a hard link at the specified path pointing to the target. Both the target and link paths must be on the same file system.

**Parameters:**

-   `target`: The file that the hard link will refer to
-   `link`: The path where the hard link will be created

**Returns:**

-   `bool`: true if the hard link was created successfully, false otherwise

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

    // First create a target file
    fs.CreateFileWithContent("./source.txt", "This is the source file for hard linking")

    // Create a hard link to the target
    success := fs.CreateHardLink("./source.txt", "./hardlink_to_source.txt")
    if success {
        fmt.Println("Hard link created successfully")
    } else {
        fmt.Println("Failed to create hard link")
    }
}
```

**Expected Output:**

```
Hard link created successfully
```

Both files will appear as regular files, but they share the same inode and data blocks.

</details>

### CreateDirectoryTree

Creates a directory tree based on the provided structure. The structure is a map where keys are directory names and values are either nil (for empty directories) or nested maps (for subdirectories).

**Parameters:**

-   `basePath`: The base directory path where the tree will be created
-   `structure`: A map representing the directory structure to create

**Returns:**

-   `bool`: true if the directory tree was created successfully, false otherwise

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

    // Define a directory structure
    structure := map[string]interface{}{
        "docs": map[string]interface{}{
            "technical": nil,
            "user": map[string]interface{}{
                "guides": nil,
                "tutorials": nil,
            },
        },
        "src": map[string]interface{}{
            "main": nil,
            "test": nil,
        },
    }

    success := fs.CreateDirectoryTree("./project", structure)
    if success {
        fmt.Println("Directory tree created successfully")
    } else {
        fmt.Println("Failed to create directory tree")
    }
}
```

**Expected Output:**

```
Directory tree created successfully
```

This will create the following directory structure:

```
project/
├── docs/
│   ├── technical/
│   └── user/
│       ├── guides/
│       └── tutorials/
└── src/
    ├── main/
    └── test/
```

</details>

### CreateDirectoryTreeWithPermissions

Creates a directory tree with the specified permissions. The structure is a map where keys are directory names and values are either nil (for empty directories) or nested maps (for subdirectories).

**Parameters:**

-   `basePath`: The base directory path where the tree will be created
-   `structure`: A map representing the directory structure to create
-   `perm`: The permissions to apply to all directories in the tree

**Returns:**

-   `bool`: true if the directory tree was created successfully, false otherwise

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

    // Define a directory structure
    structure := map[string]interface{}{
        "config": map[string]interface{}{
            "secure": nil,
        },
        "data": map[string]interface{}{
            "private": nil,
            "public": nil,
        },
    }

    // Create with restricted permissions (owner access only)
    success := fs.CreateDirectoryTreeWithPermissions("./secure_app", structure, 0700)
    if success {
        fmt.Println("Secure directory tree created successfully")
    } else {
        fmt.Println("Failed to create secure directory tree")
    }
}
```

**Expected Output:**

```
Secure directory tree created successfully
```

This will create the following directory structure, all with 0700 permissions:

```
secure_app/
├── config/
│   └── secure/
└── data/
    ├── private/
    └── public/
```

</details>

### SymlinkDirectoryTree

Creates symbolic links for an entire directory tree. This function walks through the source directory tree and creates corresponding symbolic links in the destination directory.

**Parameters:**

-   `sourceDir`: The source directory tree to be symlinked
-   `destDir`: The destination directory where symlinks will be created
-   `recursive`: If true, symlinks subdirectories recursively; otherwise, only symlinks files in the top-level directory

**Returns:**

-   `bool`: true if the directory tree was symlinked successfully, false otherwise

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

    // First create a source directory structure with some files
    fs.CreateDirectory("./original/subdir1")
    fs.CreateDirectory("./original/subdir2")
    fs.CreateFileWithContent("./original/file1.txt", "Content of file 1")
    fs.CreateFileWithContent("./original/subdir1/file2.txt", "Content of file 2")

    // Now create symlinks to the entire directory structure
    success := fs.SymlinkDirectoryTree("./original", "./linked_copy", true)
    if success {
        fmt.Println("Directory tree symlinked successfully")
    } else {
        fmt.Println("Failed to symlink directory tree")
    }
}
```

**Expected Output:**

```
Directory tree symlinked successfully
```

This will create the following structure:

```
linked_copy/
├── file1.txt -> ../original/file1.txt
├── subdir1/
│   └── file2.txt -> ../../original/subdir1/file2.txt
└── subdir2/
```

All files in the linked_copy directory are symbolic links to the corresponding files in the original directory.

</details>

## Error Handling

All functions in Creations.go handle errors internally and return boolean values to indicate success or failure. When an error occurs, it is logged through the UFS instance's error handling mechanism, which can be configured to display errors, log them to a file, or handle them silently.

## Performance Considerations

-   When creating large directory trees or symlinking entire directory structures, be aware that these operations can be resource-intensive.
-   Creating files with large content using `CreateFileWithContent` or `CreateFileWithContentAndPermissions` loads the entire content into memory. For very large files, consider using streaming operations instead.

## Platform Compatibility

These functions are designed to work across different operating systems, but some functionality may have platform-specific behavior:

-   Permissions have different meanings on Windows vs. Unix-like systems
-   Symbolic links may require administrative privileges on Windows
-   Hard links have platform-specific limitations (e.g., they can't span filesystems and can't point to directories on most systems)
