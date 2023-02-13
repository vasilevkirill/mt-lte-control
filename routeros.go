package main

import (
	"fmt"
	"gopkg.in/routeros.v2"
	"gopkg.in/routeros.v2/proto"
	"time"
)

func routerRePairGetValue(re []*proto.Sentence, key string) string {
	if len(re) == 0 {
		return ""
	}
	if key == "" {
		return ""
	}
	for _, v := range re {
		for _, vv := range v.List {
			if vv.Key == key {
				return vv.Value
			}
		}
	}
	return ""
}
func sendCommand(command []string) ([]*proto.Sentence, error) {
	timeout := 10 * time.Second
	connString := fmt.Sprintf("%s:%d", Config.GetString("router.address"), Config.GetInt("router.apiport"))
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
