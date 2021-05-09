package service

import (
	"fmt"
	"strconv"
	"time"
)

type MainService struct {
	ExitChannel chan bool
	StartTime   time.Time
}

func NewMainService(executeTime time.Time) *MainService {
	return &MainService{StartTime: executeTime}
}

func (ms *MainService) Start() error {
	//time.Sleep(40 * time.Millisecond)
	ms.ExitChannel = make(chan bool, 1)
	//fmt.Printf(" %v at time %v \n", "Start()", time.Since(ms.StartTime))
	exit := false
	for {
		fmt.Printf(" %v at time %v \n", "In for loop", time.Since(ms.StartTime))
		if exit {
			fmt.Printf(" %v at time %v \n", "exit program", time.Since(ms.StartTime))
			break
		}
		fmt.Printf(" %v at time %v \n", "Waiting .... value from ExitChannel ", time.Since(ms.StartTime))
		select {
		case <-ms.ExitChannel:
			fmt.Printf(" (1) After recieve from case <-ms.ExitChannel  %v \n", time.Since(ms.StartTime))
			data := <-ms.ExitChannel
			fmt.Printf(" (2) After recieve from data := <-ms.ExitChannel  %v at time %v \n", strconv.FormatBool(data), time.Since(ms.StartTime))
			exit = data
		}

	}
	return nil
}

// Stop stop the services
func (ms *MainService) Stop() {
	if ms.ExitChannel == nil {
		return
	}
	ms.ExitChannel <- true
}
