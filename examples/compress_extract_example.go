package main

import "github.com/utsav-56/ufs"

func main() {

	ufsInstance := &ufs.UFS{}

	ufsInstance.CompressDirectory(".", "./compressed.zip")

	ufsInstance.ExtractArchive("./compressed.zip", "./extracted")


	ufs.GetFolderList("./extracted")

}
