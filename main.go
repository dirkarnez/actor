package main

import (
	"github.com/dirkarnez/wait2die"
	"github.com/play175/wifiNotifier"
	"log"
)

func main() {
	var previousSSID = ""

	wifiNotifier.SetWifiNotifier(func(ssid string) {
		if  previousSSID == ssid {
			return
		}
		log.Println("onWifiChanged, current ssid:" + ssid)
		previousSSID = ssid
	})

	log.Println("current ssid:" + wifiNotifier.GetCurrentSSID())

	wait2die.WaitToDie(nil)
}