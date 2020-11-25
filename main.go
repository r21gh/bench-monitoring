package main

import (
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"sync"
	"strconv"
	"os"
	"time"
)

// Monitor struct consists of monitoring items
/*
	Alloc: currently allocated number of bytes on the heap,
	TotalAlloc: cumulative max bytes allocated on the heap (will not decrease),
	Sys: total memory obtained from the OS,
	Mallocs and Frees: number of allocations, deallocations, and live objects (mallocs - frees),
	PauseTotalNs: total GC pauses since the app has started,
	NumGC: number of completed GC cycles
*/
type Monitor struct {
	Alloc,
	TotalAlloc,
	Sys,
	Mallocs,
	Frees,
	LiveObjects,
	PauseTotalNs 	uint64
	NumGC			uint32
	NumGoroutine	int
}

// NewMonitor : This function is responsible for creating a new monitoring 
func NewMonitor(wg *sync.WaitGroup, duration int) {
	defer wg.Done()
	var m Monitor
	var rtm runtime.MemStats 
	var interval = time.Duration(duration) * time.Second

	for {
		<- time.After(interval)

		// Read full memory stats
		runtime.ReadMemStats(&rtm)

		// Number of goroutines
		m.NumGoroutine = runtime.NumGoroutine()

		// Misc memory stats
		m.Alloc = rtm.Alloc
		m.TotalAlloc = rtm.TotalAlloc
		m.Sys = rtm.Sys
		m.Mallocs = rtm.Mallocs
		m.Frees = rtm.Frees

		// Live objects = Mallocs - Frees
		m.LiveObjects = m.Mallocs - m.Frees

		// GC Stats 
		m.PauseTotalNs = rtm.PauseTotalNs
		m.NumGC = rtm.NumGC

		// Just encode to json and print
		b, err := json.Marshal(m)
		if err != nil {
			log.Printf("an error occured during marshaling: %v", err.Error())
		}
		
		fmt.Println(string(b))
	}
}

func main() {

	var monitorDuration string
	
	if len(os.Args) > 1 {
		monitorDuration = os.Args[1]
	} else {
		log.Fatalf("invalid/empty input monitor number")
	}

	IntNumberDuration, err :=strconv.Atoi(monitorDuration)
	if err != nil {
		log.Fatalf("an error occured during converting the monitor duration to integer: %v", err.Error())
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go NewMonitor(&wg, IntNumberDuration)
	wg.Wait()

}