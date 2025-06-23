package main

import (
	"fmt"
	"os"

	"github.com/utsav-56/ufs" // Adjust the import path as necessary
)

func main() {

	// You dont need to create instance else
	// you can use them directly using ufs.PathExists, ufs.IsFile, etc.
	ufsInstance := &ufs.UFS{}

	// Prepare test paths (adjust as needed)
	testFile := "testfile.txt"
	testDir := "testdir"
	hiddenFile := ".hiddenfile"
	hiddenDir := ".hiddendir"
	execFile := "testexec.sh"

	// Create test files and directories for demonstration
	os.WriteFile(testFile, []byte("hello"), 0644)
	os.Mkdir(testDir, 0755)
	os.WriteFile(hiddenFile, []byte("hidden"), 0644)
	os.Mkdir(hiddenDir, 0755)
	os.WriteFile(execFile, []byte("#!/bin/bash\necho hi"), 0755)

	fmt.Println("PathExists:", ufsInstance.PathExists(testFile))
	fmt.Println("IsFile:", ufsInstance.IsFile(testFile))
	fmt.Println("IsDirectory:", ufsInstance.IsDirectory(testDir))
	fmt.Println("IsDirectoryEmpty:", ufsInstance.IsDirectoryEmpty(testDir))
	fmt.Println("IsFileEmpty:", ufsInstance.IsFileEmpty(testFile))
	fmt.Println("IsInSystemPath:", ufsInstance.IsInSystemPath(testFile))
	fmt.Println("IsInUserPath:", ufsInstance.IsInUserPath(testFile))
	fmt.Println("IsInCurrentPath:", ufsInstance.IsInCurrentPath(testFile))
	fmt.Println("IsFileHidden:", ufsInstance.IsFileHidden(hiddenFile))
	fmt.Println("IsFileExecutable:", ufsInstance.IsFileExecutable(execFile))
	fmt.Println("IsFileReadable:", ufsInstance.IsFileReadable(testFile))
	fmt.Println("IsFileWritable:", ufsInstance.IsFileWritable(testFile))
	fmt.Println("IsDirectoryHidden:", ufsInstance.IsDirectoryHidden(hiddenDir))
	fmt.Println("IsDirectoryReadable:", ufsInstance.IsDirectoryReadable(testDir))

	// Clean up
	os.Remove(testFile)
	os.Remove(execFile)
	os.Remove(hiddenFile)
	os.RemoveAll(testDir)
	os.RemoveAll(hiddenDir)
}
