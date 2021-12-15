package main

import (
	"flag"
	"fmt"
	"github.com/KpLi0rn/Log4j2Scan/config"
	"github.com/KpLi0rn/Log4j2Scan/core"
	"github.com/KpLi0rn/Log4j2Scan/log"
	"github.com/KpLi0rn/Log4j2Scan/model"
	"github.com/KpLi0rn/Log4j2Scan/util"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	ResultChan chan *model.Result
	RenderChan chan *model.Result
)

func main() {
	core.PrintLogo(config.GetAuthors())
	parserInput()
	ResultChan = make(chan *model.Result, config.DefaultChannelSize)
	RenderChan = make(chan *model.Result, config.DefaultChannelSize)
	go core.StartFakeServer(&ResultChan)
	go core.StartHttpServer(&RenderChan)
	go func() {
		for {
			select {
			case res := <-ResultChan:
				info := fmt.Sprintf("%s->%s", res.Name, res.Host)
				log.Info("log4j2 detected")
				log.Info(info)
				data := &model.Result{
					Host:   res.Host,
					Name:   res.Name,
					Finger: res.Finger,
				}
				RenderChan <- data
			}
		}
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("|------------------------------------|")
	fmt.Println("|--Payload: ldap/rmi://your-ip:port--|")
	fmt.Println("|------------------------------------|")
	wait()
}

func parserInput() {
	var (
		port     int
		httpPort int
		help     bool
	)
	flag.IntVar(&port, "p", 8001, "server port")
	flag.IntVar(&httpPort, "http-port", 8888, "http port")
	flag.BoolVar(&help, "help", false, "help info")
	flag.Parse()
	if help {
		flag.PrintDefaults()
		return
	}
	if !util.CheckPort(port) || !util.CheckPort(httpPort) {
		os.Exit(-1)
	}
	log.Info("use port: %d", port)
	log.Info("use http port: %d", httpPort)
	config.Port = port
	config.HttpPort = httpPort
}

func wait() {
	sign := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sign
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()
	<-done
}
