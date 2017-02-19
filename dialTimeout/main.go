package main

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	trans := &http.Transport{
		Dial: TimeoutDialer(time.Second*1, time.Second*1),
	}
	client := &http.Client{
		Transport: trans,
	}
	resp, err := client.Get("http://api.bilibili.com/archive_stat/stat?aid=1")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(data))
}

func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		err = conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, err
	}
}
