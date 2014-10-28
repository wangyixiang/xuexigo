package unziponedrivepackage

import (
	zip "archive/zip"
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

const (
	indexFileName = "Encoding Errors.txt"
)

func getIndexMap(r io.Reader) (indexMap map[string]string, err error) {
	sStarter := "Original File Name  ->  New File Name"
	sMarker := "->"
	inputReader := bufio.NewReader(r)
	indexMap = make(map[string]string)

	for {
		line, err := inputReader.ReadString('\n')
		line = strings.TrimSpace(line)
		if err == io.EOF {
			return nil, err
		}
		if line == sStarter {
			break
		}
	}

	for {
		line, err := inputReader.ReadString('\n')
		if line == "" && err == nil {
			continue
		}
		iCount := strings.Count(line, sMarker)
		switch iCount {
		case 0:
			fmt.Println("Strange Data0:" + line)
		case 1:
			sSlice := strings.Split(line, sMarker)
			indexMap[strings.TrimSpace(sSlice[1])] = strings.TrimSpace(sSlice[0])
		default:
			fmt.Println("Strange Data2:" + line)
		}
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
	}
	return indexMap, err
}

func UnzipPackageTo(packageName string, dstDirectory string, overWrite bool) error {
	rc, err := zip.OpenReader(packageName)
	dstDirectory = path.Clean(dstDirectory)
	if err != nil {
		return err
	}
	defer rc.Close()

	_, err = os.Stat(dstDirectory)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(dstDirectory, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if dstDirectory != "." && !overWrite {
			return errors.New(dstDirectory + " already exists.")
		}
	}

	var nameMap map[string]string = nil

	for _, f := range rc.File {
		if f.Name == indexFileName {
			r, err := f.Open()
			if err != nil {
				return err
			}
			defer r.Close()
			nameMap, err = getIndexMap(r)
			if err != nil {
				return err
			}
			break
		}
	}

	for _, f := range rc.File {
		if f.Name == indexFileName {
			continue
		}
		if f.Mode().IsDir() {
			err := os.MkdirAll(dstDirectory+string(os.PathSeparator)+f.Name, 0755)
			if err != nil {
				return err
			}
			continue
		}
		if nameMap != nil {
			if originalName, ok := nameMap[f.Name]; ok {
				err := unzipFileFromPackage(f, dstDirectory+string(os.PathSeparator)+originalName)
				if err != nil {
					return err
				}
				continue
			}
		}
		unzipFileFromPackage(f, dstDirectory+string(os.PathSeparator)+f.Name)
	}
	return nil
}

func unzipFileFromPackage(f *zip.File, filename string) error {
	writer, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer writer.Close()
	reader, err := f.Open()
	if err != nil {
		return err
	}
	defer reader.Close()
	if _, err := io.Copy(writer, reader); err != nil {
		return err
	}
	fmt.Println("extract " + f.Name + " to " + filename)
	return nil
}

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
	oneDriveDir := "/mnt/hgfs/E/Download"
	for _, filename := range oneDriveFiles {
		ok := UnzipPackageTo(path.Join(oneDriveDir, filename), path.Join(oneDriveDir, "oneDrive"), true)
		if ok != nil {
			log.Fatal(ok)
		}
	}
}
