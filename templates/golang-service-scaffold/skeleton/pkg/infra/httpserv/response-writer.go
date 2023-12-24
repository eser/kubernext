// This file contains code from ginmiddlewares (https://github.com/ringsaturn/ginmiddlewares),
// which is a middleware library for gin, licensed under the MIT license.

// Copyright (c) 2023-present Eser Ozvataf
// Copyright (c) 2021 vearne

package httpserv

import (
	"bytes"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type ResponseWriter struct {
	gin.ResponseWriter
	h    http.Header
	body *bytes.Buffer

	code        int
	mu          sync.Mutex
	timedOut    bool
	wroteHeader bool
}

func NewResponseWriter(c *Context) *ResponseWriter {
	buffer := GetBuff()
	writer := &ResponseWriter{
		body:           buffer,
		ResponseWriter: c.Writer,
		h:              make(http.Header),
	}

	return writer
}

func (w *ResponseWriter) Write(b []byte) (int, error) { //nolint:varnamelen
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.timedOut {
		return 0, nil
	}

	writtenBytes, err := w.body.Write(b)
	if err != nil {
		return writtenBytes, fmt.Errorf("response writer write error: %w", err)
	}

	return writtenBytes, nil
}

func (w *ResponseWriter) WriteHeader(code int) {
	checkWriteHeaderCode(code)
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.timedOut {
		return
	}

	w.writeHeader(code)
}

func (w *ResponseWriter) writeHeader(code int) {
	w.wroteHeader = true
	w.code = code
}

func (w *ResponseWriter) WriteHeaderNow() {}

func (w *ResponseWriter) Header() http.Header {
	return w.h
}

func checkWriteHeaderCode(code int) {
	if code < 100 || code > 999 {
		panic(fmt.Sprintf("invalid WriteHeader code %v", code))
	}
}

func (w *ResponseWriter) Done(c *Context) {
	dst := w.ResponseWriter.Header()
	for k, vv := range w.Header() {
		dst[k] = vv
	}

	if !w.wroteHeader {
		w.code = http.StatusOK
	}

	w.ResponseWriter.WriteHeader(w.code)

	_, err := w.ResponseWriter.Write(w.body.Bytes())
	if err != nil {
		panic(err)
	}

	PutBuff(w.body)
}

func ResponseWriterHandler(c *gin.Context) {
	w := NewResponseWriter(c)
	c.Writer = w

	defer w.Done(c)
	c.Next()
}
