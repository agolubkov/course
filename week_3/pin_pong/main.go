package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
)

func comm(id int, ch chan string) {
	for {
		select {
		case mes := <-ch:
			switch mes {
			case "PING":
				ch <- "PONG"
				logrus.WithFields(logrus.Fields{"uid": id, "got": "PING", "sent": "PONG"}).Info("New Round")
			case "PONG":
				logrus.WithFields(logrus.Fields{"uid": id, "got": "PONG"}).Info("New Round")
			}
		case <-time.After(time.Millisecond * time.Duration(1000+rand.Intn(1000))):
			ch <- "PING"
			logrus.WithFields(logrus.Fields{"uid": id, "sent": "PING"}).Info("New Round")
		}
	}
}

func main() {
	ch := make(chan string, 1)
	go comm(1, ch)
	go comm(2, ch)
	fmt.Scanln()
}
