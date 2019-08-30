package wait

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"jephy.cn/lib/log"
)

var (
	_sigs       = make(chan os.Signal, 1)
	_signalFunc = make(map[os.Signal]func())
)

func RegSignalFunc(sig os.Signal, f func()) {
	log.Info("add process signal:[%v] func=%v", sig, f)
	signal.Notify(_sigs, sig)
	_signalFunc[sig] = f
}

func Wait() {
	signal.Notify(_sigs, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT)

	callfunc := func(sig os.Signal) {
		if fu, ok := _signalFunc[sig]; ok {
			log.Info("signal recevied[%v] and call handle func %v", sig, fu)
			fu()
		} else {
			log.Warn("dafault case signal[%v=%d], do nothing.", sig, sig)
		}
	}

	for {
		msg := <-_sigs
		log.Warn("Recevied signal: %v", msg)
		switch msg {
		default:
			callfunc(msg)
		case os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGQUIT:
			log.Warn("signal [%v] recevied, app will exit at 1 second after sleep.", msg)
			callfunc(msg)
			time.Sleep(time.Second)
			return
		}
	}
}
