package gomain

import "runtime"

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func Main() {
	for f := range mainfunc {
		f()
	}
}

// queue of work to run in main thread.
var mainfunc = make(chan func(), 100)

// do runs f on the main thread.
func Do(f func()) {
	done := make(chan bool, 1)
	mainfunc <- func() {
		f()
		done <- true
	}
	<-done
}

func Exit() {
	close(mainfunc)
}
