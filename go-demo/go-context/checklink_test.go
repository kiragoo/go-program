package gocontext

import "testing"

func TestCheckLink(t *testing.T) {
	var links = []string{
		"http://www.baidu.com",
		"http://www.jd.com",
		"http://www.taobao.com",
	}
	c := make(chan string)
	for _, link := range links {
		go CheckLink(link, c)
	}
	for {
		go CheckLink(<-c, c)
	}
}
