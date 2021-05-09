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
	// Create
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
		case data1 := <-ms.ExitChannel:
			fmt.Printf(" (1) After recieve from case <-ms.ExitChannel [%v] at time %v \n", strconv.FormatBool(data1), time.Since(ms.StartTime))
			data2 := <-ms.ExitChannel
			fmt.Printf(" (2) After recieve from data := <-ms.ExitChannel [%v] at time %v \n", strconv.FormatBool(data2), time.Since(ms.StartTime))
			exit = data2
		}

	}
	return nil
}

// Stop stop the services
func (ms *MainService) Stop(isStop bool) {
	if ms.ExitChannel == nil {
		return
	}
	ms.ExitChannel <- isStop
	//ms.ExitChannel <- true
}
