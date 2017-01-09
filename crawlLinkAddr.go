package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ll1615/go-snippets/links"
)

type list struct {
	links []string
	depth int
}

type ulink struct {
	link  string
	depth int
}

var (
	crawlDepth = 2
	sourceLink = flag.String("l", "", "provide the source link to crawl")
)

func main() {
	worklist := make(chan list)
	unseenlinks := make(chan ulink)

	flag.Parse()
	if *sourceLink == "" {
		log.Fatalln("Please input the link to crawl with -l ")
	}

	go func() {
		worklist <- list{[]string{*sourceLink}, 0}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for unlink := range unseenlinks {
				if unlink.depth > crawlDepth {
					os.Exit(0)
				}
				foundlinks := crawl(unlink.link, unlink.depth)
				go func() { worklist <- foundlinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for wlist := range worklist {
		for _, link := range wlist.links {
			if !seen[link] {
				seen[link] = true
				unseenlinks <- ulink{link, wlist.depth}
			}
		}
	}
}

func crawl(url string, depth int) list {
	fmt.Println(url)

	linklist, err := links.Extract(url)

	if err != nil {
		log.Print(err)
	}
	return list{linklist, depth + 1}
}

