package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

const (
	KB float64 = 1024
	MB float64 = 1024 * 1024
)

var sema = make(chan struct{}, 20)

func main() {
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var outFileCh = make(chan *File)
	var outFile = new(File)
	go server(outFile, outFileCh)

	for _, root := range roots {
		wg.Add(1)
		go SearchFile(&wg, outFileCh, root)
	}
	wg.Wait()
	close(outFileCh)

	fmt.Println(outFile)
}

// SearchFile searchs in the path
func SearchFile(wg *sync.WaitGroup, outFileCh chan<- *File, path string) {
	defer wg.Done()

	for _, entry := range dirents(path) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(path, entry.Name())
			go SearchFile(wg, outFileCh, subdir)
		} else {
			outFileCh <- &File{Path: path, Name: entry.Name(), Size: entry.Size()}
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

func server(outFile *File, outFileCh <-chan *File) {
	for {
		file, ok := <-outFileCh
		if !ok {
			break
		}
		if outFile.Size < file.Size {
			*outFile = *file
		}
	}
}

// File to store concerned message
type File struct {
	Path string
	Name string
	Size int64
}

func (f *File) String() string {
	return fmt.Sprintf("Path=%s, Name=%s, Size=%.3fMB\n", f.Path, f.Name, float64(f.Size)/MB)
}
