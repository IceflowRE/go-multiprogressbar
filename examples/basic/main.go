package main

import (
	"fmt"
	"github.com/IceflowRE/go-multiprogressbar"
	"github.com/schollz/progressbar/v3"
	"time"
)

func main() {
	mpb := multiprogressbar.New()
	for _, pBar := range []*progressbar.ProgressBar{
		progressbar.New(150),
		progressbar.New(200),
		progressbar.New(250),
	} {
		mpb.Add(pBar)
	}

	mpb.Get(0).Describe("Bar Zero")
	mpb.Get(1).Describe("Bar One")
	mpb.Get(2).Describe("Bar Two")
	
	for val := 0; val < 300; val++ {
		time.Sleep(10 * time.Millisecond)
		barId := val % 3
		mpb.Get(barId).Add(1)
		if val == 100 {
			mpb.Get(0).Describe("CHANGED")
		}
	}
	mpb.Finish()
	fmt.Printf("TADA")
}
