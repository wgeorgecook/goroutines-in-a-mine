package main

import (
	"log"

	"github.com/wgeorgecook/channels/channeling"
	"github.com/wgeorgecook/channels/worker"
)

var (
	inChan   channeling.InputChan
	errChan  channeling.ErrChan
	doneChan channeling.DoneChan
	workChan channeling.WorkChan
)

func init() {
	inChan = make(channeling.InputChan)
	errChan = make(channeling.ErrChan)
	doneChan = make(chan bool, 1)
	workChan = make(chan bool, 1)
}

func main() {
	log.Println("Hello world!")

	log.Println("starting to listen for work")
	go channeling.Process(inChan, errChan, doneChan, workChan)

	log.Println("starting to listen for errors")
	go channeling.ListenForErrors(errChan)

	log.Println("starting work with 7 workers")
	var names = []worker.Names{
		worker.Bashful,
		worker.Doc,
		worker.Dopey,
		worker.Grumpy,
		worker.Happy,
		worker.Sleepy,
		worker.Sneezy,
	}
	for _, name := range names {
		inChan <- worker.New(name, workChan)
	}

	// block until we're done
	log.Println("Blocking until work completes...")
	<-doneChan
	defer log.Println("Goodbye!")
}
