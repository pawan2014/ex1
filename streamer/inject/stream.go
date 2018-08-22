package inject

import (
	"log"
	"time"

	"github.com/ex1/streamer/config"
	"github.com/ex1/streamer/model"
)

var doneChannels = make(map[string]chan bool)
var reportingChannel = make(map[string]chan string)

func StreamMachine(machine model.Machine) {

	// log.Println(machine.Machineid, "-", machine.MachineName)
	// for _, p := range machine.Tags {
	// 	log.Printf("\t %v-%v", p.TagName, p.Frequency)
	// }

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
	reportingChannel[machine.Machineid] = make(chan string, 100)
	for _, p := range machine.Tags {
		log.Println("\t", p.TagName)
		concatName := machine.Machineid + p.TagID
		doneChannels[concatName] = make(chan bool, 1)
		go SendData(p.TagID, time.Duration(p.Frequency)*time.Second, doneChannels[concatName], reportingChannel[machine.Machineid])

	}
	go func(rptChan chan string) {
		log.Println("Started listing to error  channel")
	loop:
		for {
			select {
			case mydata, ok := <-rptChan:
				if !ok {
					break loop
				} else {
					log.Printf("Recieved error" + mydata)
				}
			}
		}
		log.Println("Closing Error channel")
	}(reportingChannel[machine.Machineid])

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
	if _, present := reportingChannel[machine]; present != false {
		close(reportingChannel[machine])
		delete(reportingChannel, machine)
	}

}

// Send to MQTT
func SendData(tagId string, sleep time.Duration, done chan bool, rptChan chan string) {
	log.Println("Inside Goroutine ", tagId)
loop:
	for {
		select {
		case <-done:
			log.Println("Stopped.. ", tagId)
			break loop
		default:
			log.Println("Pushing sensor value for  ", tagId)
			err := config.Publish(tagId + ": timeseries")
			if err != nil {
				log.Println("Got Error:Rabbit pusblish fail : log Error ", tagId)
				rptChan <- "Tag ID" + tagId + " : " + err.Error()
				log.Println("Done:Got Error:Rabbit pusblish fail : log Error ", tagId)
			}
			time.Sleep(sleep)
		}
	}
	log.Println("Closing Goroutine ", tagId)
}

func CleanAllChannels() {

	log.Printf("Cleaing all...")
	for k, v := range doneChannels {
		close(v)
		delete(doneChannels, k)
	}
	for k1, v1 := range reportingChannel {
		close(v1)
		delete(doneChannels, k1)
	}
	log.Printf("Channel Length is %v,%v", len(doneChannels), len(reportingChannel))
}
