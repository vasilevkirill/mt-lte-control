package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	loadConfig()
	runWebServer()
}
