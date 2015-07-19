package main

import (
	"./tunnel_relic"
	"fmt"
	"time"
)

func main() {

	//create tunnelRelic
	myAccount := "123"
	myAPIKey := "456"
	myTunnelName := "Purchase"
	tunnel := tunnelRelic.NewTunnel(myAccount, myAPIKey, myTunnelName, 3, 2)

	//Create sample event (can be reused)
	anEvent := make(map[string]string)
	anEvent["key"] = "value"
	anEvent["Dammit"] = "Yoyo"

	//send event
	tunnel.RegisterEvent(anEvent)
}
