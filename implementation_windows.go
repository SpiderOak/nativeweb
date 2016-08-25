// +build windows

// Windows-specific implementation of the nativeweb client.

package nativeweb

import (
	"fmt"
	"net/http"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// We want winhttp to be a singleton to avoid opening up too many
// handles to the same DLL.

// TODO: if thread safety becomes an issue here, each nativeweb
// client's just going to get its own winhttp DLL handle and we'll
// move on with life.
var winhttp *windows.LazyDLL
var once sync.Once

func winHTTP() *windows.LazyDLL {
	once.Do(func() {
		winhttp = windows.NewLazySystemDLL("winhttp.dll")
	})
	return winhttp
}

type nativeWebImpl struct {
}

func New() (NativeWeb, error) {
	return &nativeWebImpl{}, nil
}

func (impl *nativeWebImpl) Get(url string) (*http.Response, error) {

	hSession, _, err := winHTTP().NewProc("WinHttpOpen").Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("WinHTTP Example/1.0"))),
		uintptr(WINHTTP_ACCESS_TYPE_DEFAULT_PROXY),
		uintptr(WINHTTP_NO_PROXY_NAME),
		uintptr(WINHTTP_NO_PROXY_BYPASS),
		0)

	if hSession != 0 {
		defer winHTTP().NewProc("WinHttpCloseHandle").Call(hSession)
	} else {
		return nil, err
	}

	hConnect, _, err := winHTTP().NewProc("WinHttpConnect").Call(hSession,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("spideroak.com"))),
		uintptr(INTERNET_DEFAULT_HTTPS_PORT), 0)

	if hConnect != 0 {
		defer winHTTP().NewProc("WinHttpCloseHandle").Call(hConnect)
	} else {
		return nil, err
	}

	hRequest, _, err := winHTTP().NewProc("WinHttpOpenRequest").Call(
		hConnect,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("GET"))), 0, 0, 0, 0,
		uintptr(WINHTTP_FLAG_SECURE))

	if hRequest != 0 {
		defer winHTTP().NewProc("WinHttpCloseHandle").Call(hRequest)
	} else {
		return nil, err
	}

	// This is actually where we call out to HTTP.
	bResults, _, err := winHTTP().NewProc("WinHttpSendRequest").Call(hRequest,
		uintptr(WINHTTP_NO_ADDITIONAL_HEADERS), 0,
		uintptr(WINHTTP_NO_REQUEST_DATA), 0, 0, 0)

	bResults, _, err = winHTTP().NewProc("WinHttpReceiveResponse").Call(hRequest, 0)

	// The following dance is beacuse the WinHTTP API is in C, so
	// we have to first make a call to determine how big the
	// results are, then we can allocate memory for the results,
	// then we can call WinHttpQueryHeaders *again* to actually
	// give us the data we're looking for.

	var dwSize uint32
	if bResults == 1 {
		_, _, err := winHTTP().NewProc("WinHttpQueryHeaders").Call(hRequest,
			uintptr(WINHTTP_QUERY_RAW_HEADERS_CRLF),
			uintptr(WINHTTP_HEADER_NAME_BY_INDEX),
			uintptr(0),
			uintptr(unsafe.Pointer(&dwSize)),
			uintptr(WINHTTP_NO_HEADER_INDEX))

		fmt.Printf("Error is %d: %v\n", err, err)

		if err == windows.ERROR_INSUFFICIENT_BUFFER {
		}
	}

	fmt.Printf("bResults %d error: %d\n", bResults, err)

	return nil, nil
}
