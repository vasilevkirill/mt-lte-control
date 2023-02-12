package main

import (
	"fmt"
	"gopkg.in/routeros.v2"
	"time"
)

func getInfoOverApi() (interface{}, error) {
	timeout := 4 * time.Second
	connString := fmt.Sprintf("%s:%d", Config.GetString("router.ip"), Config.GetInt("router.apiport"))
	conn, err := routeros.DialTimeout(connString, Config.GetString("router.user"), Config.GetString("router.password"), timeout)
	if err != nil {
		return nil, err
	}
	r, err := conn.RunArgs(ar)

	return nil, nil
}

func sendCommand(command []string) (interface{}, error) {
	timeout := 4 * time.Second
	connString := fmt.Sprintf("%s:%d", Config.GetString("router.ip"), Config.GetInt("router.apiport"))
	conn, err := routeros.DialTimeout(connString, Config.GetString("router.user"), Config.GetString("router.password"), timeout)
	if err != nil {
		return nil, err
	}
	r, err := conn.RunArgs(command)
	if err != nil {
		return nil, err
	}
	return r.Re, nil
}
