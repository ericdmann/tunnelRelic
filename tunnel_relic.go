package tunnelRelic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type Tunnel struct {
	SendInterval    int
	SendBuffer      int
	InsightsAPI     string
	InsightsAccount string
	InsightsEvent   string
	InsightsURL     string
	SendQueue       []string
}

func NewTunnel(Account string, APIKey string, EventName string, send int, sendBuff int) *Tunnel {

	url := strings.Join([]string{"https://insights-collector.newrelic.com/v1/accounts/", Account, "/events"}, "")
	relic := &Tunnel{

		SendInterval:    send,
		SendBuffer:      sendBuff,
		InsightsAPI:     APIKey,
		InsightsAccount: Account,
		InsightsURL:     url,
		InsightsEvent:   EventName,
	}

	go relic.MaintainQueue()
	return relic

}

func (relic *Tunnel) MaintainQueue() {

	for true {
		time.Sleep(time.Second * time.Duration(int64(relic.SendInterval)))
		relic.EmptyQueue()
	}

}

func (relic *Tunnel) RegisterEvent(event map[string]interface{}) {
	event["eventType"] = relic.InsightsEvent
	eventJson, err := json.Marshal(event)
	if err != nil {
		fmt.Println("tunnelRelic: Error receiving event", err)
	}

	objectString := string(eventJson[:])
	relic.SendQueue = append(relic.SendQueue, objectString)

	fmt.Println("tunnelRelic: Added event to send-queue. Currently ", len(relic.SendQueue), " events in the queue")
	//fmt.Println(objectString)
	if len(relic.SendQueue) > relic.SendBuffer {
		fmt.Println("tunnelRelic: Event queue buffer reached!")
		relic.EmptyQueue()
	}
}

func (relic *Tunnel) EmptyQueue() {

	if len(relic.SendQueue) < 1 {
		return
	}
	fmt.Println("tunnelRelic: Gophers will now proceed to deliver queued events to New Relic.")

	requestStr := "[" + strings.Join(relic.SendQueue, ",") + "]"

	fmt.Println(requestStr)
	var eventJson = []byte(requestStr)
	req, err := http.NewRequest("POST", relic.InsightsURL, bytes.NewBuffer(eventJson))
	req.Header.Set("X-Insert-Key", relic.InsightsAPI)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("tunnelRelic: Sending queued request to New Relic. Response: ", string(body))

	relic.SendQueue = nil
}
