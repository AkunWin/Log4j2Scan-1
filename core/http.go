package core

import (
	"fmt"
	"github.com/KpLi0rn/Log4j2Scan/config"
	"github.com/KpLi0rn/Log4j2Scan/log"
	"github.com/KpLi0rn/Log4j2Scan/model"
	"net/http"
	"sync"
)

var (
	resultList []*model.Result
	lock       sync.Mutex
)

func StartHttpServer(renderChan *chan *model.Result) {
	log.Info("start result http server")
	mux := http.NewServeMux()
	mux.Handle(config.DefaultHttpPath, &resultHandler{})
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.HttpPort),
		WriteTimeout: config.DefaultHttpTimeout,
		Handler:      mux,
	}
	go listenData(renderChan)
	_ = server.ListenAndServe()
}

func listenData(renderChan *chan *model.Result) {
	for {
		select {
		case res := <-*renderChan:
			lock.Lock()
			resultList = append(resultList, res)
			lock.Unlock()
		}
	}
}

type resultHandler struct {
}

func (handler *resultHandler) ServeHTTP(w http.ResponseWriter, _ *http.Request) {
	lock.Lock()
	_, _ = w.Write(RenderHtml(resultList))
	lock.Unlock()
}
