package generator

import (
	"fmt"
	"time"
)

type Generator interface {
	Log(msg string) error
}

type generator struct {
	Hostname   string
	ThreadName string
}

func NewGenerator(hostname string, threadname string) Generator {
	return &generator{
		Hostname:   hostname,
		ThreadName: threadname,
	}
}

func (g generator) Log(msg string) error {
	t := time.Now()
	fmt.Printf("[%v][%v][%v]: %v \n", t.Format("2006-01-02 15:04:05"), g.Hostname, g.ThreadName, msg)
	return nil
}
