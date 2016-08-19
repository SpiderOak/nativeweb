// Simple HTTP client implementation that makes use of native HTTP APIs.
//
// Cross-platform. We use:
// * WinHTTP on Win32
// * NSURLSesion on Darwin
// * Golang's own net/http/client on Linux (May do some gomobile hijinks
//   on Android in the future).

package nativeweb

import (
	"net/http"
)

type NativeWeb interface {
	// Returns data from an HTTP request using system-specific methods.
	GetURL(req *http.Request) (*http.Response, error)
	TestFetchURL(url string)
}
