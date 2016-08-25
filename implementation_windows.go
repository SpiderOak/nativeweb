// +build windows

// Windows-specific implementation of the nativeweb client.

package nativeweb

import (
	"bufio"
	"bytes"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// We want winhttp to be a singleton to avoid opening up too many
// handles to the same DLL.
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

func (impl *nativeWebImpl) queryRawHeaders(hRequest uintptr) (string, error) {
	var dwSize uint32
	var lpOutBuffer []uint16

	bResults, _, err := winHTTP().NewProc("WinHttpQueryHeaders").Call(hRequest,
		uintptr(WINHTTP_QUERY_RAW_HEADERS_CRLF),
		uintptr(WINHTTP_HEADER_NAME_BY_INDEX),
		0,
		uintptr(unsafe.Pointer(&dwSize)),
		uintptr(WINHTTP_NO_HEADER_INDEX))

	if bResults != 1 {
		if err == windows.ERROR_INSUFFICIENT_BUFFER {
			lpOutBuffer = make([]uint16, int(dwSize))

			bResults, _, err = winHTTP().NewProc("WinHttpQueryHeaders").Call(hRequest,
				uintptr(WINHTTP_QUERY_RAW_HEADERS_CRLF),
				uintptr(WINHTTP_HEADER_NAME_BY_INDEX),
				uintptr(unsafe.Pointer(&lpOutBuffer[0])),
				uintptr(unsafe.Pointer(&dwSize)),
				uintptr(WINHTTP_NO_HEADER_INDEX))

			if bResults == 1 {
				return syscall.UTF16ToString(lpOutBuffer), nil
			} else {
				return "", err
			}
		} else {
			return "", err
		}
	} else {
		panic("bResults == 1 with null pointer to header string")
	}
}

func (impl *nativeWebImpl) makeSession(req *http.Request) (uintptr, error) {
	hSession, _, err := winHTTP().NewProc("WinHttpOpen").Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("NativeWeb WinHTTP"))),
		uintptr(WINHTTP_ACCESS_TYPE_DEFAULT_PROXY),
		uintptr(WINHTTP_NO_PROXY_NAME),
		uintptr(WINHTTP_NO_PROXY_BYPASS),
		0)

	return hSession, err
}

func (impl *nativeWebImpl) makeConnection(req *http.Request, hSession uintptr) (uintptr, error) {
	// Golang considers the port part of the hostname, whereas WinHTTP
	// doesn't. We need to split apart the host in the request, as well
	// as set default ports if none are specified.
	var port int
	var host string
	var err error

	res := strings.Split(req.URL.Host, ":")
	switch {
	case len(res) != 1:
		port, err = strconv.Atoi(res[1])
		if err != nil {
			return 0, err
		}
	case req.URL.Scheme == "https":
		port = 443
	case req.URL.Scheme == "http":
		port = 80
	case true:
		port = 0 // WinHTTP will pick automagically depending on if
		// we request a secure connection or not.
	}

	host = res[0]

	hConnect, _, err := winHTTP().NewProc("WinHttpConnect").Call(hSession,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(host))),
		uintptr(port), 0)
	return hConnect, err
}

func (impl *nativeWebImpl) makeRequest(req *http.Request, hConnect uintptr) (uintptr, error) {

	hRequest, _, err := winHTTP().NewProc("WinHttpOpenRequest").Call(
		hConnect,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(req.Method))),
		0, // Object Name
		0, // HTTP version (default 1.1)
		0, // Referrer
		0, // Accept Types
		uintptr(WINHTTP_FLAG_SECURE))

	return hRequest, err
}

func (impl *nativeWebImpl) runRequest(req *http.Request) (*string, *string, error) {
	hSession, err := impl.makeSession(req)
	if hSession != 0 {
		defer winHTTP().NewProc("WinHttpCloseHandle").Call(hSession)
	} else {
		return nil, nil, err
	}

	hConnect, err := impl.makeConnection(req, hSession)
	if hConnect != 0 {
		defer winHTTP().NewProc("WinHttpCloseHandle").Call(hConnect)
	} else {
		return nil, nil, err
	}

	hRequest, err := impl.makeRequest(req, hConnect)
	if hRequest != 0 {
		defer winHTTP().NewProc("WinHttpCloseHandle").Call(hRequest)
	} else {
		return nil, nil, err
	}

	// This is actually where we call out to HTTP.
	bResults, _, err := winHTTP().NewProc("WinHttpSendRequest").Call(hRequest,
		uintptr(WINHTTP_NO_ADDITIONAL_HEADERS), 0,
		uintptr(WINHTTP_NO_REQUEST_DATA), 0, 0, 0)

	bResults, _, err = winHTTP().NewProc("WinHttpReceiveResponse").Call(hRequest, 0)

	if bResults != 1 {
		return nil, nil, err
	}

	headers, err := impl.queryRawHeaders(hRequest)

	if err != nil {
		return nil, nil, err
	}

	return &headers, nil, nil
}

func (impl *nativeWebImpl) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)

	headers, _, err := impl.runRequest(req)

	if err != nil {
		return nil, err
	}

	goResp, err := http.ReadResponse(
		bufio.NewReader(bytes.NewBufferString(*headers)), nil)
	if err != nil {
		return nil, err
	}

	return goResp, nil
}
