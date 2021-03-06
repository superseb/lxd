package main

import (
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"

	"github.com/lxc/lxd/shared"
)

func doMemDump() {
	f, err := os.Create(*memProfile)
	if err != nil {
		shared.Debugf("Error opening memory profile file '%s': %s\n", err)
		return
	}
	pprof.WriteHeapProfile(f)
	f.Close()
}

func memProfiler() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGUSR1)
	for {
		<-ch
		doMemDump()
	}
}
