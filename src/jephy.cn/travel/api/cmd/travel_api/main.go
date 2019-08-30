package main

import (
	"jephy.cn/lib/log"
	"jephy.cn/lib/wait"
	"jephy.cn/lib/web"
	"jephy.cn/travel/api/global"
	"jephy.cn/travel/api/restful"
	"syscall"
)

func main() {
	global.LoadConfig("/etc/travel_api.toml")

	log.SetLevelS(global.Config().LogLevel)

	webServer := web.NewWeb()
	restful.RegisterURLs(webServer)

	exit := func() {
		log.Fatal("receive exit signal")
		webServer.Close()
	}

	wait.RegSignalFunc(syscall.SIGTERM, exit)
	wait.RegSignalFunc(syscall.SIGQUIT, exit)
	wait.RegSignalFunc(syscall.SIGINT, exit)
	go wait.Wait()

	if err := webServer.Serve("tcp", global.Config().HTTPAddr); err != nil {
		log.Fatal(err.Error())
	}
}
