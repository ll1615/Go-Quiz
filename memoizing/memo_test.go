package memo

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"

	v1 "github.com/ll1615/go-snippets/memoizing/v1"
	v2 "github.com/ll1615/go-snippets/memoizing/v2"
)

var incomingURLs = []string{
	"http://baidu.com",
	"http://qq.com",
	"http://www.sina.com",
	"http://youdao.com",
	"http://qq.com",
	"http://youdao.com",
	"http://www.sina.com",
	"http://baidu.com",
}

func TestMemoV1(t *testing.T) {
	m := v1.New(httpGetBody)
	var wg sync.WaitGroup

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func TestMemoV2(t *testing.T) {
	m := v2.New(httpGetBody)
	var wg sync.WaitGroup

	for _, url := range incomingURLs {
		wg.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
			wg.Done()
		}(url)
	}
	wg.Wait()
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
