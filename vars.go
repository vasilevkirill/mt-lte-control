package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	Log         = logrus.New()
	Config      = viper.New()
	configPaths = []string{".", "/etc", "/etc/defaults", "c:\\", "c:\\windows\\"}
	R           Router
)
