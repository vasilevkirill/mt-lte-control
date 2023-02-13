package main

import (
	"log"
	"sync"
	"time"
)

type Router struct {
	Status     map[string]interface{}
	lastUpdate time.Time
	rw         sync.Mutex
}

func (r *Router) forceGetInformation() {
	r.getLteInfo()
	r.getSimSlot()
	r.getScript()
}

func (r *Router) getScript() {

	cmd := []string{
		"/system/script/print",
		"?=name=" + Config.GetString("router.script"),
	}
	re, err := sendCommand(cmd)
	if err != nil {
		log.Println(err)
		return
	}
	r.rw.Lock()
	defer r.rw.Unlock()
	if len(re) == 1 {
		r.Status["script"] = true
		r.lastUpdate = time.Now()
	}

	return
}
func (r *Router) getLteInfo() {
	cmd := []string{
		"/interface/lte/info",
		"=number=0",
		"=once",
	}
	re, err := sendCommand(cmd)
	if err != nil {
		log.Println(err)
		return
	}
	if len(re) == 0 {
		return
	}
	r.rw.Lock()
	defer r.rw.Unlock()
	r.Status["lteInfo"] = re[0].Map
	r.lastUpdate = time.Now()
	return
}
func (r *Router) getSimSlot() {
	cmd := []string{
		"/system/routerboard/modem/print",
	}
	re, err := sendCommand(cmd)
	if err != nil {
		log.Println(err)
		return
	}
	vl := routerRePairGetValue(re, "sim-slot")
	r.rw.Lock()
	defer r.rw.Unlock()
	r.Status["simSlot"] = vl
	r.lastUpdate = time.Now()
	return
}

func (r *Router) readStatusMapKey(key string) interface{} {
	tnow := time.Now()
	diff := tnow.Sub(r.lastUpdate)
	if diff.Seconds() > float64(Config.GetInt("router.storetime")) {
		r.forceGetInformation()
	}
	r.rw.Lock()
	defer r.rw.Unlock()
	if v, ok := r.Status[key]; ok {
		return v
	}
	return ""
}
