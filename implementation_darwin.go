// +build darwin

// Darwin-specific implementation of nativeweb client.

package nativeweb

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Cocoa
// #include "darwin_mechanics.h"
import "C"

import (
	"fmt"
	"net/http"
	"unsafe"
)

type nativeWebImpl struct {
}

// New returns an entity that implements the NativeWeb interface.
func New() (NativeWeb, error) {
	return &nativeWebImpl{}, nil
}

func (impl *nativeWebImpl) GetURL(req *http.Request) (*http.Response, error) {
	return nil, nil
}

func (impl *nativeWebImpl) TestFetchURL(url string) {
	urlString := C.CString(url)
	defer C.free(unsafe.Pointer(urlString))

	data := C.GoString(C.FetchURL(urlString))

	fmt.Printf("%s\n", data)
}
