package ufs

import (
	"log"
	"os"
	"path/filepath"
)

// GetFileSize returns the size of the given file in bytes.
// Logs errors internally and returns 0 on failure.
/*
Usage:
size := ufs.GetFileSize("/path/to/file.txt")
*/

func GetFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		log.Println("GetFileSize error:", err)
		return 0
	}
	if info.IsDir() {
		log.Println("GetFileSize: path is a directory, not a file")
		return 0
	}
	return info.Size()
}

// GetFolderSize recursively calculates the total size of a folder.
// Logs errors internally and returns 0 on failure.
/*
Usage:
size := ufs.GetFolderSize("/path/to/directory")
*/
func GetFolderSize(path string) int64 {
	var size int64
	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			log.Println("Walk error:", err)
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				log.Println("Info error:", err)
				return nil
			}
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		log.Println("GetFolderSize error:", err)
		return 0
	}
	return size
}

// GetFolderList returns a list of folder names under the given path.
// Logs errors internally and returns an empty slice on failure.
/*
Usage:
folders := ufs.GetFolderList("/path/to/directory")
*/
func GetFolderList(path string) []string {
	var folders []string
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Println("GetFolderList error:", err)
		return []string{}
	}
	for _, entry := range entries {
		if entry.IsDir() {
			folders = append(folders, entry.Name())
		}
	}
	return folders
}

// GetFileList returns a list of file names under the given path (non-recursive).
// Logs errors internally and returns an empty slice on failure.
/*
Usage:
files := ufs.GetFileList("/path/to/directory")
*/

func GetFileList(path string) []string {
	var files []string
	entries, err := os.ReadDir(path)
	if err != nil {
		log.Println("GetFileList error:", err)
		return []string{}
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files
}
