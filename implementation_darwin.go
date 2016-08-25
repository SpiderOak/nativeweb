// +build darwin

// Darwin-specific implementation of nativeweb client.

package nativeweb

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "mechanics_darwin.h"
import "C"

import (
	"net/http"
	"unsafe"
)

type nativeWebImpl struct {
}

// New returns an entity that implements the NativeWeb interface.
func New() (NativeWeb, error) {
	return &nativeWebImpl{}, nil
}

func (impl *nativeWebImpl) Get(url string) (*http.Response, error) {
	var results unsafe.Pointer

	urlString := C.CString(url)
	defer C.free(unsafe.Pointer(urlString))

	results = C.FetchURL(urlString)

	goResp := http.Response{
		Status:        C.GoString(C.StatusText(results)),
		StatusCode:    int(C.StatusCode(results)),
		Proto:         "HTTP/1.1", // This should be a bit more... dynamical.
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(C.ContentLength(results)),
	}

	return &goResp, nil
}
