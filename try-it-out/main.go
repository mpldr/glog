package main

import (
	"fmt"
	"io"
	"math/rand"

	"git.sr.ht/~poldi1405/glog"
	"git.sr.ht/~poldi1405/glog/logrotation"
)

func main() {
	defer glog.PanicHandler()
	fmt.Println("######################### normal")
	glog.Trace("this is a Trace message")
	glog.Debug("this is a Debug message")
	glog.Info("this is a Info message")
	glog.Warn("this is a Warning message")
	glog.Error("this is an Error message")
	glog.Fatal("this is a Fatal message")
	fmt.Println("######################### to file")
	f1, _ := glog.AddLogFile("dbg.log", glog.DEBUG, glog.TRACE)
	defer f1.Close()
	f2, _ := glog.AddLogFile("err.log", glog.ERROR, glog.FATAL)
	defer f2.Close()
	glog.Trace("this is a Trace message")
	glog.Debug("this is a Debug message")
	glog.Info("this is a Info message")
	glog.Warn("this is a Warning message")
	glog.Error("this is an Error message")
	glog.Fatal("this is a Fatal message")
	fmt.Println("######################### drop warning and above")
	glog.SetOutputAbove(glog.WARNING, io.Discard)
	glog.Trace("this is a Trace message")
	glog.Debug("this is a Debug message")
	glog.Info("this is a Info message")
	glog.Warn("this is a Warning message")
	glog.Error("this is an Error message")
	glog.Fatal("this is a Fatal message")
	glog.Debug("this caller only shows the function")
	glog.ShowCallerLine = true
	glog.Debug("this caller shows the line")

	glog.Info("10.000.000 lines will now be logged to a file and rotated on occasion")
	r := logrotation.NewRotor("big.log")
	r.Open()
	defer r.Close()
	glog.SetOutput(glog.INFO, r)

	messages := []string{
		"a user has logged in",
		"please consider supporting on LiberaPay",
		"place your ad here",
		"i can't code UIs to save my life",
		"no other platform allows this much freedom in window alignment",
		"Windows 10 = Windows 7, but more Spyware",
		"Windows 11 = macOS, but more Spyware?",
	}

	for i := 0; i < 10000000; i++ {
		fmt.Printf("%8d/%8d\r", i, 10000000)
		glog.Info(messages[rand.Intn(len(messages))])
	}
	fmt.Println()
	teenie()
}

func teenie() {
	panic("O!M!G! is that the newest buyPhone?! AAAAAHHH!")
}
