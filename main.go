package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
)

func main() {

	var (
		s [18]string
		c = make(chan Result)
		t = time.After(time.Second * 10)
	)

	head := color.New(color.FgHiBlue).Add(color.Bold).SprintFunc()
	note := color.New(color.FgGreen).SprintFunc()
	log.Println(head("项目地址：github.com/wangyuling93/backtrace-jxyd"))
	log.Println(note("正在测试三网回程路由..."))

	for i := range rIp {
		go trace(c, i)
	}

loop:
	for range s {
		select {
		case o := <-c:
			s[o.i] = o.s
		case <-t:
			break loop
		}
	}

	for i, r := range s {
		if i > 0 && i%3 == 0 {
			fmt.Println()
		}
		log.Println(r)
	}
}
