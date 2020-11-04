package Go_Git_DdrFs_Zip

import (
	"archive/zip"
	"fmt"
	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git"
	"github.com/go-git/go-git/storage/memory"
	"io"
	"os"
	"path/filepath"
	"strings"
)
var storer *memory.Storage
var fs billy.Filesystem

func CloneToZip(url string, dest string) error  {
	// Clones the given repository in memory, creating the remote, the local
	// branches and fetching the objects, exactly as:
	//Info("git clone https://github.com/go-git/go-billy")
	storer = memory.NewStorage()
	fs = memfs.New()
	_, err := git.Clone(storer, fs, &git.CloneOptions{
		URL: url,
	})
	if err != nil {
		return err
	}
	fmt.Println("Repository cloned")
	// ---------
	// Подготовим зипчик
	destinationFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	myZip := zip.NewWriter(destinationFile)
	defer myZip.Close()
	//--
	err = readFiles("/",fs,myZip)
	if err != nil {
		return err
	}
	return nil
}
func readFiles(dir string, reader billy.Filesystem, zipWriter *zip.Writer) (err error) {
	files, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}
	for _, file := range files {
		//fmt.Printf("Nm:%s | IsDir:%v\n", file.Name(), file.IsDir())
		if file.IsDir() {
			// Считаем эту папку
			err = readFiles(filepath.Join(dir,file.Name()),reader,zipWriter)
			if err != nil {
				return err
			}
			continue
		}
		fPath := filepath.Join(dir,file.Name())
		fsFile, err := fs.Open(fPath)
		if err != nil {
			return err
		}
		// Удалим какие то слеши, которые почемуто отоброажаются в архиве
		fPath = strings.TrimPrefix(fPath, "\\")
		fPath = strings.TrimPrefix(fPath, "/")

		zipFile, err := zipWriter.Create(fPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(zipFile, fsFile)
		if err != nil {
			return err
		}
	}
	return nil
}