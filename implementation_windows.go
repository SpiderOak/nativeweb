// +build windows

// Windows-specific implementation of the nativeweb client.

package nativeweb

import (
	"fmt"
	"net/http"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

type nativeWebImpl struct {
}

func New() (NativeWeb, error) {
	return &nativeWebImpl{}, nil
}

func (impl *nativeWebImpl) Get(url string) (*http.Response, error) {
	var winhttp = windows.NewLazySystemDLL("winhttp.dll")

	var winHttpOpen = winhttp.NewProc("WinHttpOpen")
	hSession, _, err := winHttpOpen.Call(
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("WinHTTP Example/1.0"))),
		uintptr(WINHTTP_ACCESS_TYPE_DEFAULT_PROXY),
		uintptr(WINHTTP_NO_PROXY_NAME),
		uintptr(WINHTTP_NO_PROXY_BYPASS),
		0)

	fmt.Printf("hSession error: %d\n", err)

	var winHttpConnect = winhttp.NewProc("WinHttpConnect")
	hConnect, _, err := winHttpConnect.Call(hSession,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("spideroak.com"))),
		uintptr(INTERNET_DEFAULT_HTTPS_PORT), 0)

	fmt.Printf("hConnect error: %d\n", err)

	var winHttpOpenRequest = winhttp.NewProc("WinHttpOpenRequest")
	hRequest, _, err := winHttpOpenRequest.Call(
		hConnect,
		uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("GET"))), uintptr(0),
		uintptr(0), uintptr(0),
		uintptr(0),
		uintptr(WINHTTP_FLAG_SECURE))

	fmt.Printf("Return: %d %d\n", hRequest, err)

	var winHttpSendRequest = winhttp.NewProc("WinHttpSendRequest")
	bResults, _, err := winHttpSendRequest.Call(hRequest,
		uintptr(WINHTTP_NO_ADDITIONAL_HEADERS), 0,
		uintptr(WINHTTP_NO_REQUEST_DATA), 0,
		0, 0)

	fmt.Printf("bResults %d error: %d\n", bResults, err)
	return nil, nil
}
