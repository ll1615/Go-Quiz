package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"time"
)

var files []*File

const (
	rootPath         = `F:/Code/Go` //source directory
	KB       float64 = 1024
	MB       float64 = 1024 * 1024
)

func init() {
	files = make([]*File, 0)
}

func main() {
	start := time.Now()
	defer func() {
		fmt.Printf("Elapsed Time: %v\n", time.Now().Sub(start))
	}()

	GetAllFiles(rootPath)

	sort.Sort(Files(files))
	length := len(files)
	if length > 0 {
		fmt.Println(files[length-1])
	} else {
		fmt.Println("There is no file in this directory!")
	}
}

func GetAllFiles(path string) {
	file, err := os.Stat(path)
	if err == nil {
		if !file.IsDir() {
			SaveFileInfo(file, path)
		} else {
			files, _ := ioutil.ReadDir(path)
			for _, f := range files {
				GetAllFiles(path + "/" + f.Name())
			}
		}
	}
}

func SaveFileInfo(f os.FileInfo, path string) {
	file := &File{Path: path, Name: f.Name(), Size: f.Size()}
	files = append(files, file)
}

type File struct {
	Path string
	Name string
	Size int64
}

func (f *File) String() string {
	return fmt.Sprintf("Path=%s, Name=%s, Size=%.3fMB\n", f.Path, f.Name, float64(f.Size)/MB)
}

type Files []*File

func (f Files) Len() int {
	return len(f)
}

func (f Files) Less(i, j int) bool {
	if f[i].Size < f[j].Size {
		return true
	}
	return false
}

func (f Files) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
