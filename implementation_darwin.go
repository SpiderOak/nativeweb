// +build darwin

// Darwin-specific implementation of nativeweb client.

package nativeweb

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "mechanics_darwin.h"
import "C"

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"runtime"
	"unsafe"
)

type nativeWebImpl struct {
	urlSession unsafe.Pointer
}

// New creates a fresh instance of the nativeWebImpl type, implementing the NativeWeb API.
func New() (NativeWeb, error) {
	impl := nativeWebImpl{
		urlSession: C.OpenSession(),
	}

	// Make sure we release the urlSession if the NativeWeb client gets GC'ed.
	runtime.SetFinalizer(&impl, func(n *nativeWebImpl) {
		C.Release(n.urlSession)
	})

	return &impl, nil
}

func (impl *nativeWebImpl) Do(req *http.Request) (*http.Response, error) {
	var results unsafe.Pointer

	urlString := C.CString(req.URL.String())
	defer C.free(unsafe.Pointer(urlString))

	results = C.FetchURL(impl.urlSession, urlString)
	bLen := C.DataBytesSize(results)

	b := unsafe.Pointer(C.DataBytes(results))
	bBytes := bytes.NewBuffer(C.GoBytes(b, C.int(bLen)))

	goResp := http.Response{
		Status:        C.GoString(C.StatusText(results)),
		StatusCode:    int(C.StatusCode(results)),
		Proto:         "HTTP/1.1", // This should be a bit more... dynamical.
		ProtoMajor:    1,
		ProtoMinor:    1,
		ContentLength: int64(C.ContentLength(results)),
		Body:          ioutil.NopCloser(bufio.NewReader(bBytes)),
	}

	if int64(bLen) != goResp.ContentLength {
		return nil, errors.New("Incomplete download")
	}

	return &goResp, nil
}
