package main

import (
	"fmt"
	"io"

	"git.sr.ht/~poldi1405/glog"
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
	teenie()
}

func teenie() {
	panic("O!M!G! is that the newest buyPhone?! AAAAAHHH!")
}
