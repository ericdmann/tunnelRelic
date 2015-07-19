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

```golang
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

```golang
	anEvent := make(map[string]string)
	anEvent["key"] = "value"
	anEvent["Dammit"] = "Yoyo"
```

### Send an event
```golang
	gopherTunnel.RegisterEvent(anEvent)
```

### Examples
Check out usage.go. 

Here's sample output if you're doing it right!

```

	tunnelRelic: Added event to send-queue. Currently  1  events in the queue

	[{"Dammit":"Yoyo","eventType":"Purchase","key":"value"}]

	tunnelRelic: Added event to send-queue. Currently  2  events in the queue

	[{"Dammit":"Yoyo","eventType":"Purchase","key":"value"}]

	tunnelRelic: Added event to send-queue. Currently  3  events in the queue

	[{"Dammit":"Yoyo","eventType":"Purchase","key":"value"}]

	tunnelRelic: Event queue buffer reached!

	tunnelRelic: Gophers will now proceed to deliver queued events to New Relic.

	tunnelRelic: Sending queued request to New Relic. Response:  {"success": true}

```


## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D


## TODO

Verbose logging on/off