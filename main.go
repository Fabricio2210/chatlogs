package main

import (
	"fmt"
	"github.com/Fabricio2210/wrapper"
	"github.com/robfig/cron"
)

func main() {
	c := cron.New()
	c.AddFunc("0 10 11 * *", func() {
		wrapper.Wrapper("DSP")
	})
	c.AddFunc("0 15 11 * *", func() {
		wrapper.Wrapper("DDM")
	})
	c.AddFunc("0 20 11 * *", func() {
		wrapper.Wrapper("RAW")
	})
	c.AddFunc("0 25 11 * *", func() {
		wrapper.Wrapper("POP")
	})
	c.AddFunc("0 30 11 * *", func() {
		wrapper.Wrapper("SHINKO")
	})
	c.AddFunc("0 35 11 * *", func() {
		wrapper.Wrapper("AQUA")
	})
	c.AddFunc("0 40 11 * *", func() {
		wrapper.Wrapper("BEAM")
	})
	c.AddFunc("0 45 11 * *", func() {
		wrapper.Wrapper("DECEPTICRON")
	})
	c.AddFunc("0 50 11 * *", func() {
	 	wrapper.Wrapper("DOODY")
	})
	c.AddFunc("0 55 11 * *", func() {
		wrapper.Wrapper("MEERKAT")
	})
	c.AddFunc("0 00 12 * *", func() {
		wrapper.Wrapper("PROPER")
	})
	c.AddFunc("0 05 12 * *", func() {
		wrapper.Wrapper("REACTS")
	})
	c.AddFunc("0 10 12 * *", func() {
		wrapper.Wrapper("TBS")
	})
	c.AddFunc("0 15 12 * *", func() {
		wrapper.Wrapper("WPIG")
	})
	c.Start()
	fmt.Println("Running!!!!!!!")
	<-make(chan struct{})
}

