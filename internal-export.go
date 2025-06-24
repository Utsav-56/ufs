package ufs

// Metadata.go functions
var GetFileSize = dufs.GetFileSize
var GetFolderSize = dufs.GetFolderSize
var GetFolderList = dufs.GetFolderList
var GetFileList = dufs.GetFileList
var GetFolderChildCount = dufs.GetFolderChildCount
var GetFolderFileCount = dufs.GetFolderFileCount
var GetFolderMetadata = dufs.GetFolderMetadata
var Get = dufs.GetFolderMetadata
var GetFileMetadata = dufs.GetFileMetadata
var GetChildCount = dufs.GetChildCount

// Creations.go functions
var CreateFile = dufs.CreateFile
var CreateFileWithContent = dufs.CreateFileWithContent
var CreateFileWithContentAndPermissions = dufs.CreateFileWithContentAndPermissions
var CreateFileWithPermissions = dufs.CreateFileWithPermissions
var CreateDirectory = dufs.CreateDirectory
var CreateDirectoryWithPermissions = dufs.CreateDirectoryWithPermissions
var CreateSymlink = dufs.CreateSymlink
var CreateHardLink = dufs.CreateHardLink
var CreateDirectoryTree = dufs.CreateDirectoryTree
var CreateDirectoryTreeWithPermissions = dufs.CreateDirectoryTreeWithPermissions
var SymlinkDirectoryTree = dufs.SymlinkDirectoryTree
var RenameFile = dufs.RenameFile
var RenameDirectory = dufs.RenameDirectory

// Removing.go functions
var RemoveFile = dufs.RemoveFile
var RemoveDirectory = dufs.RemoveDirectory
var RemoveDirectoryRecursive = dufs.RemoveDirectoryRecursive
var RemoveSymlink = dufs.RemoveSymlink
var RemoveFileWithBackup = dufs.RemoveFileWithBackup
var RemoveEmptyFiles = dufs.RemoveEmptyFiles
var RemoveEmptyDirectories = dufs.RemoveEmptyDirectories
var RemoveDirectoryContents = dufs.RemoveDirectoryContents
var RemoveDirectoryTree = dufs.RemoveDirectoryTree
var RemoveAllLinks = dufs.RemoveAllLinks
var RemoveByPattern = dufs.RemoveByPattern
var SafeRemoveFile = dufs.SafeRemoveFile

// File-Reader_Writer.go functions
var ReadFile = dufs.ReadFile
var ReadFileAsString = dufs.ReadFileAsString
var WriteFile = dufs.WriteFile
var WriteStringToFile = dufs.WriteStringToFile
var AppendToFile = dufs.AppendToFile
var AppendStringToFile = dufs.AppendStringToFile
var CopyFile = dufs.CopyFile
var MoveFile = dufs.MoveFile
var DeleteFile = dufs.DeleteFile
var CopyFileWithPermissions = dufs.CopyFileWithPermissions
var MoveFileWithPermissions = dufs.MoveFileWithPermissions
var AssembleFiles = dufs.AssembleFiles
var SplitFile = dufs.SplitFile
var CleanUpFiles = dufs.CleanUpFiles
var ReadFileWithLines = dufs.ReadFileWithLines
var AppendToLastLine = dufs.AppendToLastLine
var AppendToFirstLine = dufs.AppendToFirstLine

// Path-properties.go functions
var PathExists = dufs.PathExists
var IsFile = dufs.IsFile
var IsDirectory = dufs.IsDirectory
var IsDirectoryEmpty = dufs.IsDirectoryEmpty
var IsFileEmpty = dufs.IsFileEmpty
var IsInSystemPath = dufs.IsInSystemPath
var IsInUserPath = dufs.IsInUserPath
var IsInCurrentPath = dufs.IsInCurrentPath
var IsFileHidden = dufs.IsFileHidden
var IsFileExecutable = dufs.IsFileExecutable
var IsFileReadable = dufs.IsFileReadable
var IsFileWritable = dufs.IsFileWritable
var IsDirectoryHidden = dufs.IsDirectoryHidden
var IsDirectoryReadable = dufs.IsDirectoryReadable

// Compress-Extract.go functions
var CompressDirectory = dufs.CompressDirectory
var ExtractArchive = dufs.ExtractArchive
var CompressFile = dufs.CompressFile
var CompressHere = dufs.CompressHere
var ExtractHere = dufs.ExtractHere
var CompressFileHere = dufs.CompressFileHere
var CompressAndRemove = dufs.CompressAndRemove
var ExtractAndRemove = dufs.ExtractAndRemove
var CompressAndExtract = dufs.CompressAndExtract
var ExtractAndCompress = dufs.ExtractAndCompress
var CompressWithSystemCommand = dufs.CompressWithSystemCommand
var ExtractWithSystemCommand = dufs.ExtractWithSystemCommand

var MoveDirectory = dufs.MoveDirectory
