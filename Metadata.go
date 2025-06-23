package ufs

import (
	"os"
	"path/filepath"
)

// GetFileSize returns the size of the given file in bytes.
// Logs errors internally and returns 0 on failure.
/*
Usage:
size := ufs.GetFileSize("/path/to/file.txt")
*/
func (ufs *UFS) GetFileSize(path string) int64 {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "GetFileSize")
		return 0
	}
	if info.IsDir() {
		ufs.handleMistakeWarning("GetFileSize called on a directory, returning 0")
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
func (ufs *UFS) GetFolderSize(path string) int64 {
	var size int64
	err := filepath.WalkDir(path, func(p string, d os.DirEntry, err error) error {
		if err != nil {
			ufs.handleError(err, "GetFolderSize")
			return nil
		}
		if !d.IsDir() {
			info, err := d.Info()
			if err != nil {
				ufs.handleError(err, "GetFolderSize")
				return nil
			}
			size += info.Size()
		}
		return nil
	})
	if err != nil {
		ufs.handleError(err, "GetFolderSize")
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
func (ufs *UFS) GetFolderList(path string) []string {
	var folders []string
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetFolderList")
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
func (ufs *UFS) GetFileList(path string) []string {
	var files []string
	entries, err := os.ReadDir(path)
	if err != nil {
		ufs.handleError(err, "GetFileList")
		return []string{}
	}
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}
	return files
}

func (ufs *UFS) GetFileMetadata(path string) map[string]interface{} {
	info, err := os.Stat(path)
	if err != nil {
		ufs.handleError(err, "GetFileMetadata")
		return nil
	}
	metadata := map[string]interface{}{
		"Name":    info.Name(),
		"Size":    info.Size(),
		"Mode":    info.Mode().String(),
		"ModTime": info.ModTime(),
		"IsDir":   info.IsDir(),
	}

	return metadata

}
