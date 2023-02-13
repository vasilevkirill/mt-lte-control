package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetOutput(os.Stdout)
	Log.SetLevel(logrus.InfoLevel)
	loadConfig()
	R.Status = make(map[string]interface{})
	R.lastUpdate = time.Now().AddDate(-1, 0, 0)
	runWebServer()
}
