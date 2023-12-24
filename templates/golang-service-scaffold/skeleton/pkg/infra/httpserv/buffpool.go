// This file contains code from gin-timeout (https://github.com/vearne/gin-timeout),
// which is a timeout middleware for gin, licensed under the MIT license.

// Copyright (c) 2023-present Eser Ozvataf
// Copyright (c) 2021 vearne

package httpserv

import (
	"bytes"
	"sync"
)

const BuffSize = 10 * 1024

var buffPool sync.Pool

func GetBuff() *bytes.Buffer {
	item, ok := buffPool.Get().(*bytes.Buffer)

	if ok && item != nil {
		return item
	}

	// var byteSlice []byte
	byteSlice := make([]byte, 0, BuffSize)

	return bytes.NewBuffer(byteSlice)
}

func PutBuff(buffer *bytes.Buffer) {
	buffer.Reset()
	buffPool.Put(buffer)
}
