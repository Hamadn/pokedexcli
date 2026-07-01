package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchJSON(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch %s: %s", url, res.Status)
	}

	return io.ReadAll(res.Body)
}
