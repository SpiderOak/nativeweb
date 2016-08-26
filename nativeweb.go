// Simple HTTP client implementation that makes use of native HTTP APIs.
//
// Cross-platform. We use:
// * WinHTTP on Win32
// * NSURLSesion on Darwin
// * Golang's own net/http/client on Linux (May do some gomobile hijinks
//   on Android in the future).

// The goal is to be able to provide a system by which we mimic as much
// of the behavior of net/http.Client as much as possible. We will document
// where we deviate from assumed behavior.

package nativeweb

import (
	"io"
	"net/http"
	"net/url"
)

type NativeWeb interface {
	// Do sends an HTTP request and returns an HTTP response,
	// following policy as configured on the client. An error is
	// returned if caused by client policy (such as
	// CheckRedirect), or failure to speak HTTP (such as a
	// connectivity problem). A non-2xx status code doesn't cause
	// an error.
	//
	// If the returned error is nil, the Response will contain a
	// non-nill Body which the user is expected to close. If the
	// Body is not closed, the NativeWeb's underlying transport
	// mechanism may not be able to re-use a persistent TCP
	// connection to the server for a subsequent "keep-alive"
	// request. The request Body, if non-nil, nill be closed by
	// NativeWeb, even on errors. On error, any Response can be
	// ignored.
	//
	// Generally, Get, Post, or PostForm will be used instead of Do.
	Do(req *http.Request) (*http.Response, error)

	// Issues a GET to the specified URL. Will follow redirects.
	// As in http.Client, when err is nil the response will contain
	// a non-nil Body. Caller should close resp.Body when done reading
	// from it.
	//
	// To make a request with custom headers, use http.NewRequest and
	// NativeWeb.Do.
	Get(url string) (*http.Response, error)

	// Issues a HEAD to the specified URL. Will follow redirects.
	Head(url string) (*http.Response, error)

	// Issues a POST to the specified URL. Caller should close Body from
	// the response when done reading from it. If provided body is an
	// io.Closer, it is closed after the request.
	//
	// To make a request with custom headers, use http.NewRequest and
	// NativeWeb.Do.
	Post(url string, bodyType string, body io.Reader) (*http.Response, error)

	// Issues a POST to the specified URL, with the data's keys
	// and values URL-encoded as the request body. The
	// Content-Type header is set to
	// application/x-www-form-urlencoded. To use other headers,
	// use http.NewRequest and DefaultClient.Do.
	PostForm(url string, data url.Values) (*http.Response, error)
}
