package worker

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/wgeorgecook/channels/channeling"
)

type Names int

//go:generate stringer -type=Names
const (
	Unknown Names = iota
	Bashful
	Doc
	Dopey
	Grumpy
	Happy
	Sleepy
	Sneezy
)

type Dwarf struct {
	Name     Names
	Error    error
	WorkChan channeling.WorkChan
}

func New(name Names, wc channeling.WorkChan) *Dwarf {
	d := &Dwarf{
		Name:     name,
		WorkChan: wc,
	}

	var err error
	switch name {
	case Grumpy:
		err = fmt.Errorf("%s is unhappy because the pickaxe is too heavy", d.Name)
	case Sleepy:
		err = fmt.Errorf("%s could not mine today because he was too tired", d.Name)
	case Happy:
		err = fmt.Errorf("%s can't find joy in mining without his friends", d.Name)
	case Doc:
		err = fmt.Errorf("%s is confused about the right tools for the job", d.Name)
	case Bashful:
		err = fmt.Errorf("%s is too shy to ask for help with the mining", d.Name)
	case Sneezy:
		err = fmt.Errorf("%s is sneezing too much to work", d.Name)
	case Dopey:
		err = fmt.Errorf("%s forgot where he left the pickaxe and can't start mining", d.Name)
	case Unknown:
	default:
		err = fmt.Errorf("%s no one showed up to mine", d.Name)
	}

	d.Error = err

	return d
}

func (d Dwarf) logStep(s string) {
	log.Printf("%s: %s\n", d.Name, s)
}

func (d Dwarf) sendDone() {
	d.WorkChan <- true
}

func (d Dwarf) DoWork() error {
	d.logStep("doing some work!")
	defer d.sendDone()
	randTime := rand.Int() % 100
	time.Sleep(time.Duration(randTime/5) * time.Second)
	if randTime%2 == 0 {
		return d.Error
	}
	d.logStep("success!")
	return nil
}
