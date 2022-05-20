package helpers

import (
	"io"
	"net/http"
)

func HttpGet(url string) []byte {
	resp, getErr := http.Get(url)

	if getErr != nil {
		panic(getErr.Error())
	}
	defer func(body io.ReadCloser) {
		if err := body.Close(); err != nil {
			panic(err.Error())
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}
	return body
}
