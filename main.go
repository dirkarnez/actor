package main

import (
	"bytes"
	"fmt"
	"github.com/dirkarnez/wait2die"
	"github.com/play175/wifiNotifier"
	"log"
	"os/exec"
)

var previousSSID = ""

func main() {

	wifiNotify(func(ssid string) {
		if ssid == "Default"{
			Chrome()
		}
	})

	wait2die.WaitToDie(nil)
}

type subscriber func(ssid string)

func wifiNotify(subscribers... subscriber) {
	var dispatcher = func(ssid string) {
		for _, s := range subscribers {
			s(ssid)
		}
	}
	wifiNotifier.SetWifiNotifier(func(ssid string) {
		if  previousSSID == ssid {
			return
		}

		dispatcher(ssid)

		log.Println("onWifiChanged, current ssid:" + ssid)
		previousSSID = ssid
	})

	dispatcher(wifiNotifier.GetCurrentSSID())
}
func Chrome() {
	cmd := exec.Command("C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe","--chrome-frame", "--app=https://google.com")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
	fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
}