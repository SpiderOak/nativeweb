package nativeweb

import (
	"io"
	"net/http"
	"net/url"
)

func (impl *nativeWebImpl) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := impl.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (impl *nativeWebImpl) Head(url string) (*http.Response, error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := impl.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (impl *nativeWebImpl) Post(url string, bodyType string, body io.Reader) (*http.Response, error) {
	return nil, nil
}

func (impl *nativeWebImpl) PostForm(url string, data url.Values) (*http.Response, error) {
	return nil, nil
}
