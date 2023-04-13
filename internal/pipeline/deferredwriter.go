// Package pipeline ...
package pipeline

import (
	"net/http"
	"strings"
)

// DefferedResponesWriter ...
type DefferedResponesWriter struct {
	http.ResponseWriter
	strings.Builder
	statusCode int
}

// Write ...
func (dw *DefferedResponesWriter) Write(data []byte) (int, error) {
	return dw.Builder.Write(data)
}

// FlushData ...
func (dw *DefferedResponesWriter) FlushData() {
	if dw.statusCode == 0 {
		dw.statusCode = http.StatusOK
	}
	dw.ResponseWriter.WriteHeader(dw.statusCode)
	dw.ResponseWriter.Write([]byte(dw.Builder.String()))
}

// WriteHeader ...
func (dw *DefferedResponesWriter) WriteHeader(statusCode int) {
	dw.statusCode = statusCode
}
