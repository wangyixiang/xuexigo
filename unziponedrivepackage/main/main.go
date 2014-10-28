package main

import (
	"log"
	"path"
	unzip "xuexigo/unziponedrivepackage"
)

func main() {
	oneDriveFiles := []string{
		"OneDrive-2014-07-05 (1).zip",
		"OneDrive-2014-07-05 (2).zip",
		"OneDrive-2014-07-05 (3).zip",
		"OneDrive-2014-07-05 (4).zip",
		"OneDrive-2014-07-05 (5).zip",
		"OneDrive-2014-07-05.zip",
		"OneDrive-2014-07-06 (1).zip",
		"OneDrive-2014-07-06 (2).zip",
		"OneDrive-2014-07-06 (3).zip",
		"OneDrive-2014-07-06 (4).zip",
		"OneDrive-2014-07-06 (5).zip",
		"OneDrive-2014-07-06.zip",
		"OneDrive-2014-07-07 (1).zip",
		"OneDrive-2014-07-07 (2).zip",
		"OneDrive-2014-07-07.zip"}
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	oneDriveDir := "/mnt/hgfs/E/Download"
	for _, filename := range oneDriveFiles {
		err := unzip.UnzipPackageTo(path.Join(oneDriveDir, filename), path.Join(oneDriveDir, "oneDrive"), true)
		if err != nil {
			log.Fatal(err)
		}
	}
}
