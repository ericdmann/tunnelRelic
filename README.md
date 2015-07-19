# tunnelRelic

tunnelRelic allows you to register/create/send custom events and custom attributes to New Relic Insights. Multiple tunnels (events) can be created to send events to different Insights events. 

## Requirements

You'll need: 
	New Relic Account # 
	Insights API key

Both pieces of information are found here -> https://insights.newrelic.com/accounts/YOUR_ACCOUNT_NUM/manage/api_keys


## Setup/Usage

Usage flow is drop-dead simple. Create tunnel -> send events to a tunnel -> we'll send them along to New Relic
	

### Create a tunnel

```go
	gopherTunnel := tunnelRelic.NewTunnel(myAccount, myAPIKey, myTunnelName, flushInterval, bufferLimit)
```

### Tunnel properties
	flushInterval   int 		//How often should we send queued events?
	bufferLimit     int 		//Ignore interval if we reach this many events in the queue
	myAPIKey     	string 		//An insert key from the API keys tab of Insights
	myAccount	 	string 		//Your numberical New Relic account
	myTunnelName   	string 		//Name of the event you're registering

### Create an event
Events can be resused/modified, and do not need to be allocated each time you want to register an event. 

```go
	anEvent := make(map[string]string)
	anEvent["key"] = "value"
	anEvent["Dammit"] = "Yoyo"
```

### Send an event
```go
	gopherTunnel.RegisterEvent(anEvent)
```

### Example

Events sent to Insights are fully queryable/pullable through the Insights product.


Queryable<br>
<img src="http://d.pr/i/1gYOX+" style="width: 400px;"/><br>


Explorable<br>
<img src="http://d.pr/i/1lnO4+" style="width: 400px;"/><br>
 

```go

	package main

	import (
		"./src"
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


```


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D


## TODO

Verbose logging on/off
Multiple events in single post