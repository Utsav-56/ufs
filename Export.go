package ufs

/*
Export exports the UFS functions for external use.

It divides functionalities into different categories,
such as metadata operations, path properties, and file operations.

can be used as follows:
import "github.com/utsav-56/ufs"

and then the categories can be accessed like this:

ufs.metadata.GetFileSize("path/to/file")
ufs.archive.CompressDirectory("path/to/dir", "path/to/archive.zip")
ufs.filefunctions.CreateFile("path/to/file.txt")
ufs.dirfunctions.CreateDirectory("path/to/dir")

This allows for a clean and organized way to access the various functionalities provided by the UFS package.

The static way of using the package is also there where all the functions can be accessed directly like this:
ufs.GetFileSize("path/to/file")
ufs.CompressDirectory("path/to/dir", "path/to/archive.zip")
ufs.CreateFile("path/to/file.txt")
ufs.CreateDirectory("path/to/dir")

or by creating an instance of UFS:
ufsInstance := &ufs.UFS{}

// ufsInstance.GetFileSize("path/to/file")

This package provides a unified file system interface for various operations such as metadata retrieval, path properties, file and directory creation, compression, and extraction.
It is designed to be easy to use and flexible, allowing users to perform common file system tasks with minimal code.

*/

type metadata struct{}
type archive struct{}
type fileFunctions struct{}
type dirFunctions struct{}

var Metadata = metadata{}
var Archive = archive{}
var FileFunctions = fileFunctions{}
var DirFunctions = dirFunctions{}

// Exported metadata methods

func (metadata) GetFileSize(path string) int64 {
	return GetFileSize(path)
}

func (metadata) GetFileMetadata(path string) map[string]interface{} {
	return GetFileMetadata(path)
}

func (metadata) GetFileList(path string) []string {
	return GetFileList(path)
}

func (metadata) GetFolderList(path string) []string {
	return GetFolderList(path)
}

func (metadata) GetFolderFileCount(path string) int {
	return GetFolderChildCount(path)
}

func (metadata) GetFolderChildCount(path string) int {
	return GetFolderChildCount(path)
}

func (metadata) GetChildCount(path string) (int, int) {
	return GetChildCount(path)
}

func (metadata) GetFolderMetadata(path string) map[string]interface{} {
	return GetFolderMetadata(path)
}

func (metadata) GetFolderSize(path string) int64 {
	return GetFolderSize(path)
}

// Exported archive methods
func (archive) CompressDirectory(sourcePath, destPath string) error {
	return CompressDirectory(sourcePath, destPath)
}

func (archive) ExtractArchive(sourcePath, destPath string) error {
	return ExtractArchive(sourcePath, destPath)
}

func (archive) CompressFile(sourcePath, destPath string) error {
	return CompressFile(sourcePath, destPath)
}

func (archive) CompressHere(sourcePath string) (string, error) {
	return CompressHere(sourcePath)
}

func (archive) ExtractHere(sourcePath string) (string, error) {
	return ExtractHere(sourcePath)
}

func (archive) CompressFileHere(sourcePath string) (string, error) {
	return CompressFileHere(sourcePath)
}

func (archive) CompressAndRemove(sourcePath, destPath string) error {
	return CompressAndRemove(sourcePath, destPath)
}

func (archive) ExtractAndRemove(sourcePath, destPath string) error {
	return ExtractAndRemove(sourcePath, destPath)
}

func (archive) CompressAndExtract(sourcePath, tempPath, finalPath string) error {
	return CompressAndExtract(sourcePath, tempPath, finalPath)
}

func (archive) ExtractAndCompress(sourcePath, tempPath, finalPath string) error {
	return ExtractAndCompress(sourcePath, tempPath, finalPath)
}

func (archive) CompressWithSystemCommand(sourcePath, destPath, format string) error {
	return CompressWithSystemCommand(sourcePath, destPath, format)
}

func (archive) ExtractWithSystemCommand(sourcePath, destPath string) error {
	return ExtractWithSystemCommand(sourcePath, destPath)

}

// Exported file functions methods
func (fileFunctions) ReadFile(path string) ([]byte, error) {
	return ReadFile(path)
}

func (fileFunctions) ReadFileAsString(path string) (string, error) {
	return ReadFileAsString(path)
}

func (fileFunctions) WriteFile(path string, data []byte) error {
	return WriteFile(path, data)
}

func (fileFunctions) WriteStringToFile(path string, content string) error {
	return WriteStringToFile(path, content)
}

func (fileFunctions) AppendToFile(path string, data []byte) error {
	return AppendToFile(path, data)
}

func (fileFunctions) AppendStringToFile(path string, content string) error {
	return AppendStringToFile(path, content)
}

func (fileFunctions) CopyFile(src, dst string) error {
	return CopyFile(src, dst)
}

func (fileFunctions) MoveFile(src, dst string) bool {
	return dufs.MoveFile(src, dst)
}

func (fileFunctions) DeleteFile(path string) bool {
	return DeleteFile(path)
}

func (fileFunctions) CopyFileWithPermissions(src, dst string) error {
	return CopyFileWithPermissions(src, dst)
}

func (fileFunctions) MoveFileWithPermissions(src, dst string) error {
	return MoveFileWithPermissions(src, dst)
}

func (fileFunctions) AssembleFiles(srcFiles []string, dst string) error {
	return AssembleFiles(srcFiles, dst)
}

func (fileFunctions) SplitFile(src string, chunkSize int64) ([]string, error) {
	return SplitFile(src, chunkSize)
}

func (fileFunctions) CleanUpFiles(files []string) ([]string, error) {
	return CleanUpFiles(files)
}

func (fileFunctions) ReadFileWithLines(path string) ([]string, error) {
	return ReadFileWithLines(path)
}

func (fileFunctions) AppendToLastLine(path string, content string) error {
	return AppendToLastLine(path, content)
}

func (fileFunctions) AppendToFirstLine(path string, content string) error {
	return AppendToFirstLine(path, content)
}

// Exported directory functions methods
func (dirFunctions) CreateFile(path string) bool {
	return CreateFile(path)
}

func (dirFunctions) CreateDirectory(path string) bool {
	return CreateDirectory(path)
}

func (dirFunctions) DeleteDirectory(path string) bool {
	return RemoveDirectory(path)
}

func (dirFunctions) RenameDirectory(oldPath, newPath string) bool {
	return RenameDirectory(oldPath, newPath)
}

func (dirFunctions) CopyDirectory(src, dst string) bool {
	return dufs.copyDirectoryRecursive(src, dst)
}

func (dirFunctions) MoveDirectory(src, dst string) bool {
	return dufs.MoveDirectory(src, dst)
}

func (dirFunctions) IsDirectory(path string) bool {
	return IsDirectory(path)
}

func (dirFunctions) IsDirectoryEmpty(path string) bool {
	return IsDirectoryEmpty(path)
}

func (dirFunctions) IsDirectoryHidden(path string) bool {
	return IsDirectoryHidden(path)
}

func (dirFunctions) IsDirectoryReadable(path string) bool {
	return IsDirectoryReadable(path)
}
