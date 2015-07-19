package main

import (
	"fmt"
	"github.com/ericdmann/tunnelRelic"
	"time"
)

func main() {

	//create tunnelRelic
	myAccount := "123"
	myAPIKey := "456"
	myTunnelName := "Purchase"
	flushInterval := 60
	bufferLimit := 15
	tunnel := tunnelRelic.NewTunnel(myAccount, myAPIKey, myTunnelName, flushInterval, bufferLimit)

	//Create sample event (can be reused)
	anEvent := make(map[string]string)
	anEvent["key"] = "value"
	anEvent["Dammit"] = "Yoyo"

	//send event
	tunnel.RegisterEvent(anEvent)
}
