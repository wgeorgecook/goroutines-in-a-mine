package channeling

import (
	"log"
)

type Input interface {
	DoWork() error
}
type InputChan chan (Input)
type ErrChan chan (error)
type DoneChan chan (bool)
type WorkChan chan (bool)

func Process(inChan InputChan, errChan ErrChan, done DoneChan, workChan WorkChan) {
	doWork := func(i Input) {
		if err := i.DoWork(); err != nil {
			errChan <- err
		}
	}

	log.Println("Process is listening...")
	var doneCounter = 0
	for {
		select {
		case i := <-inChan:
			go doWork(i)
		case <-workChan:
			doneCounter++
			if doneCounter == 7 {
				done <- true
			}
		}
	}

}

func ListenForErrors(errChan ErrChan) {
	logErr := func(err error) {
		log.Println("error in processing: ", err.Error())
	}

	log.Println("ListenForErrors is listening...")
	for err := range errChan {
		go logErr(err)
	}
}
