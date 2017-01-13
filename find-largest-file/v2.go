package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"sync"

	"github.com/ll1615/lib"
)

const (
	KB float64 = 1024
	MB float64 = 1024 * 1024
)

var outFiles = make([]*File, 0)
var sema = make(chan struct{}, 20)

func main() {
	defer lib.Trace("main")()

	var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	for _, root := range roots {
		wg.Add(1)
		go SearchFile(&wg, root)
	}
	wg.Wait()

	sort.Sort(Files(outFiles))

	length := len(outFiles)
	if length > 0 {
		fmt.Println(outFiles[length-1])
	} else {
		fmt.Println("There is no file in this directory!")
	}
}

func SearchFile(wg *sync.WaitGroup, path string) {
	defer wg.Done()

	for _, entry := range dirents(path) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(path, entry.Name())
			go SearchFile(wg, subdir)
		} else {
			SaveFileInfo(entry, path)
		}
	}
}

func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}

	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "dir error: %v\n", err)
		return nil
	}
	return entries
}

func SaveFileInfo(f os.FileInfo, path string) {
	file := &File{Path: path, Name: f.Name(), Size: f.Size()}
	outFiles = append(outFiles, file)
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
