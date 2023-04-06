// Package placeholder ...
package placeholder

import (
	"sync"

	"github.com/koha90/platform/internal/http"
	"github.com/koha90/platform/internal/pipeline"
	"github.com/koha90/platform/internal/pipeline/basic"
	"github.com/koha90/platform/internal/services"
)

func createPipeline() pipeline.RequestPipeline {
	return pipeline.CreatePipeline(
		&basic.ServicesComponent{},
		&basic.LoggingComponent{},
		&basic.ErrorComponent{},
		&basic.StaticFileComponent{},
		&SimpleMessageComponent{},
	)
}

// Start ...
func Start() {
	results, err := services.Call(http.Serve, createPipeline())
	if err == nil {
		(results[0].(*sync.WaitGroup)).Wait()
	} else {
		panic(err)
	}
}
