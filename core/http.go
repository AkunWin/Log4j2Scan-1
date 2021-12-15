package core

import (
	"github.com/KpLi0rn/Log4j2Scan/log"
	module "github.com/KpLi0rn/Log4j2Scan/model"
	"net/http"
	"sync"
	"time"
)

var (
	resultList []*module.Result
	lock       sync.Mutex
)

func StartHttpServer(renderChan *chan *module.Result) {
	log.Info("start result http server")
	mux := http.NewServeMux()
	mux.Handle("/", &resultHandler{})
	server := &http.Server{
		Addr:         ":8888",
		WriteTimeout: time.Second * 3, //设置3秒的写超时
		Handler:      mux,
	}
	go func() {
		for {
			select {
			case res := <-*renderChan:
				lock.Lock()
				resultList = append(resultList, res)
				lock.Unlock()
			}
		}
	}()
	_ = server.ListenAndServe()
}

type resultHandler struct {
}

func (handler *resultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	_, _ = w.Write(RenderHtml(resultList))
	lock.Unlock()
}
