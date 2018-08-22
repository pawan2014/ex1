package streamer

import (
	"log"
	"time"

	"github.com/ex1/streamer/config"
	"github.com/ex1/streamer/model"
)

var doneChannels = make(map[string]chan bool)

func StreamMachine(machine model.Machine) {

	log.Println(machine.Machineid, "-", machine.MachineName)
	for _, p := range machine.Tags {
		log.Printf("\t %v-%v", p.TagName, p.Frequency)
	}

	// timeChan := time.After(5 * time.Second)
	// go func() {
	// 	<-timeChan
	// 	for k, v := range doneChannels {
	// 		log.Printf("Signal Done for %v", k)
	// 		close(v)
	// 	}
	// }()

	StartStreaming(machine, doneChannels)

}

func StartStreaming(machine model.Machine, doneChannels map[string]chan bool) {
	log.Println(machine.Machineid)
	for _, p := range machine.Tags {
		log.Println("\t", p.TagName)
		doneChannels[p.TagName] = make(chan bool, 1)
		go SendData(p.TagName, time.Duration(p.Frequency)*time.Second, doneChannels[p.TagName])

	}
}

func StopStreaming(machine string) {
	stop(machine, doneChannels)
}

func stop(machine string, doneChannels map[string]chan bool) {
	log.Printf("send done signal to all the channels")
	for k, v := range doneChannels {
		log.Printf("Signal for %v", k)
		close(v)
		delete(doneChannels, k)
	}
}

// Send to MQTT
func SendData(tageName string, sleep time.Duration, done chan bool) {
	log.Println("inside go routine", tageName)
	for {
		select {
		case <-done:
			log.Println("Stopping ", tageName)
			return
		default:
			log.Println("timeseries for ", tageName)
			config.Publish(tageName + ": timeseries")
			time.Sleep(sleep)
		}
	}
}
