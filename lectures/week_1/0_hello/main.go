package main

import "github.com/sirupsen/logrus"

func main() {
	println("Hello, World!")
	logrus.SetLevel(logrus.DebugLevel)
	logrus.WithField("level", "debug").Debug("Hello, World!")
}
