package lib

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"testing"
)

func TestThumbnail(t *testing.T) {
	defer Trace()()

	relativeDir := "../images"
	runtime.GOMAXPROCS(runtime.NumCPU())

	imgPath := filepath.Join(os.Getenv("GOPATH"), `src/github.com/ll1615/go-snippets/images/*.thumb.*`)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", "del", imgPath)
	} else {
		cmd = exec.Command("bash", "-c", "rm", imgPath)
	}
	if err := cmd.Run(); err != nil {
		t.Errorf("error delete thumb files: %v", err)
	}

	pics := getFiles(relativeDir)

	pics = Map(pics, func(s string) string {
		if strings.HasSuffix(s, ".jpg") {
			return relativeDir + "/" + s
		}
		return ""
	})

	chPics := make(chan string, len(pics))

	for _, p := range pics {
		if p != "" {
			chPics <- p
		}
	}
	close(chPics)

	if got, want := makeThumbnails6(chPics), int64(106580); got != want {
		t.Errorf("tatal file size error: %d, want %d", got, want)
	}
}

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f)
	}
}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	for range filenames {
		<-ch
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)

	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}

	return nil
}

func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}

	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
		}(f)
	}

	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbfiles = append(thumbfiles, it.thumbfile)
	}

	return thumbfiles, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup

	for f := range filenames {
		wg.Add(1)

		go func(f string) {
			defer wg.Done()
			thumb, err := ImageFile(f)
			if err != nil {
				log.Println("thumb error:", err)
				return
			}
			info, _ := os.Stat(thumb)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		total += size
	}
	return total
}

func getFiles(folder string) []string {
	folder, _ = filepath.Abs(folder)
	pics := make([]string, 0)

	files, _ := ioutil.ReadDir(folder)
	for _, f := range files {
		pics = append(pics, f.Name())
	}
	return pics
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, s := range vs {
		vsm[i] = f(s)
	}
	return vsm
}
