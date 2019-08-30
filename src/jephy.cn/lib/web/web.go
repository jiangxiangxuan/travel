package web

import (
	"context"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"sync"
	"time"

	"jephy.cn/lib/log"
)

const (
	HttpReadTimeout = time.Minute
)

type Web struct {
	sync.Mutex
	*gin.Engine

	httpSvrs []*http.Server
}

func NewWeb() *Web {
	gin.SetMode("release")
	web := &Web{
		Engine:   gin.New(),
		httpSvrs: make([]*http.Server, 0),
	}
	return web
}

func (w *Web) Serve(protocol, addr string) error {
	var err error

	l, err := net.Listen(protocol, addr)
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	serve := func(svr *http.Server, l net.Listener) {
		err = svr.Serve(l)
		wg.Done()
	}

	wg.Add(1)
	svrMux := http.NewServeMux()
	svrMux.Handle("/", w)
	svr := http.Server{Handler: svrMux, ReadTimeout: HttpReadTimeout}
	w.httpSvrs = append(w.httpSvrs, &svr)
	go serve(&svr, l)
	wg.Wait()

	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (w *Web) Close() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	for _, svr := range w.httpSvrs {
		if err := svr.Shutdown(ctx); err != nil {
			log.Error("Shutdown http svr error:%v", err)
		}
	}
}
