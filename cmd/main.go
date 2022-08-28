package main

import (
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/puizeabix/log-generator/pkg/generator"
)

func main() {
	concurrent := flag.Int("concurrent", 10, "Number threads to run concurrently, default is 10")
	sleepus := flag.Int("sleepus", 500, "Sleep duration between each message in us (microseconds) default is 500")
	msg := flag.String("msg", "This is a sample log message", "The log message that want to print")

	flag.Parse()
	hostname, err := os.Hostname()
	if err != nil {
		panic(err.Error())
	}

	var wg sync.WaitGroup

	for i := 0; i < *concurrent; i++ {
		wg.Add(1)
		go func(h string, idx int, d int, m string) {
			tn := fmt.Sprintf("thread-%d", idx)
			g := generator.NewGenerator(h, tn)
			c := 1
			for true {
				g.Log(fmt.Sprintf("[%d]%s", c, m))
				c++
				time.Sleep(time.Microsecond * time.Duration(d))
			}
			wg.Done()
		}(hostname, i, *sleepus, *msg)
	}

	wg.Wait()
}
