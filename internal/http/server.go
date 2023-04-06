// Package http ...
package http

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/koha90/platform/internal/config"
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/pkg/logging"
)

type pipelineAdaptor struct {
	pipeline.RequestPipeline
}

func (p pipelineAdaptor) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	p.ProcessRequest(request, writer)
}

// Serve reads http, https parametrs from Configuration serve
func Serve(
	pl pipeline.RequestPipeline,
	cfg config.Configuration,
	logger logging.Logger,
) *sync.WaitGroup {
	wg := sync.WaitGroup{}

	adaptor := pipelineAdaptor{RequestPipeline: pl}

	enableHTTP := cfg.GetBoolDefault("http:enableHttp", true)
	if enableHTTP {
		httpPort := cfg.GetIntDefault("http:port", 5000)
		logger.Debugf("Starting HTTP server on port %v", httpPort)
		wg.Add(1)
		go func() {
			err := http.ListenAndServe(fmt.Sprintf(":%v", httpPort), adaptor)
			if err != nil {
				panic(err)
			}
		}()
	}

	enableHTTPS := cfg.GetBoolDefault("http:enableHttps", false)
	if enableHTTPS {
		httpsPort := cfg.GetIntDefault("http:httpsPort", 5500)
		certFile, cfok := cfg.GetString("http:httpsCert")
		keyFile, kfok := cfg.GetString("http:httpsKey")
		if cfok && kfok {
			logger.Debugf("Srarting HTTPS server on port %v", httpsPort)
			wg.Add(1)
			go func() {
				err := http.ListenAndServeTLS(
					fmt.Sprintf(":%v", httpsPort),
					certFile,
					keyFile,
					adaptor,
				)
				if err != nil {
					panic(err)
				}
			}()
		} else {
			panic("HTTPS certificate setting not found")
		}
	}

	return &wg
}
