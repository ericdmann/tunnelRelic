package tunnelRelic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var tunnelLock sync.Mutex // Controls access to ruleset

type Tunnel struct {
	SendInterval    int
	SendBuffer      int
	InsightsAPI     string
	InsightsAccount string
	InsightsEvent   string
	InsightsURL     string
	Silent          bool
	SendQueue       []string
	RunningByteSize int
}

func NewTunnel(Account string, APIKey string, EventName string, send int, sendBuff int) *Tunnel {

	url := strings.Join([]string{"https://insights-collector.newrelic.com/v1/accounts/", Account, "/events"}, "")
	relic := &Tunnel{

		SendInterval:    send,
		SendBuffer:      sendBuff,
		InsightsAPI:     APIKey,
		InsightsAccount: Account,
		InsightsURL:     url,
		Silent:          false,
		InsightsEvent:   EventName,
		RunningByteSize: 0,
	}

	go relic.MaintainQueue()
	return relic

}

func NewTransaction() map[string]interface{} {
	newRelicTransaction := make(map[string]interface{})

	if hostname, err := os.Hostname(); err == nil {
		newRelicTransaction["host"] = hostname
	} else {
		newRelicTransaction["host"] = "default"
	}

	return newRelicTransaction
}

func (relic *Tunnel) MaintainQueue() {

	for true {
		time.Sleep(time.Second * time.Duration(int64(relic.SendInterval)))
		go relic.EmptyQueue()
	}

}

func (relic *Tunnel) RegisterEvent(event map[string]interface{}) {
	event["eventType"] = relic.InsightsEvent
	eventJson, err := json.Marshal(event)
	if err != nil && relic.Silent != true {
		fmt.Println("tunnelRelic: Error receiving event", err)
	}

	objectString := string(eventJson[:])
	tunnelLock.Lock()
	defer tunnelLock.Unlock()

	relic.SendQueue = append(relic.SendQueue, objectString)
	relic.RunningByteSize = relic.RunningByteSize + len(objectString)

	if relic.Silent != true {
		fmt.Println("tunnelRelic: Added event to send-queue. Currently ", len(relic.SendQueue), " events in the queue")
	}

	// If we have gone over the specified item limit
	// or we are about to go over the 5MB payload size.
	// then send to Insights
	if len(relic.SendQueue) > relic.SendBuffer || relic.RunningByteSize > 4000000 {
		if relic.Silent != true {
			fmt.Println("tunnelRelic: Event queue buffer reached!")
		}

		go relic.EmptyQueue()
	}
}

func (relic *Tunnel) EmptyQueue() {

	if len(relic.SendQueue) < 1 {
		return
	}
	if relic.Silent != true {
		fmt.Println("tunnelRelic: Gophers will now proceed to deliver queued events to New Relic.")
	}

	requestStr := "[" + strings.Join(relic.SendQueue, ",") + "]"

	var eventJson = []byte(requestStr)
	req, err := http.NewRequest("POST", relic.InsightsURL, bytes.NewBuffer(eventJson))
	if err != nil {
		relic.SendQueue = nil
		return
	}
	req.Header.Set("X-Insert-Key", relic.InsightsAPI)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	relic.SendQueue = nil
	relic.RunningByteSize = 0

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	if relic.Silent != true {
		fmt.Println("tunnelRelic: Sending queued request to New Relic. Response: ", string(body))
	}

}
