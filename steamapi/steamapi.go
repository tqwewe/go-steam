package steamapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const apiBaseURL = "https://api.steampowered.com"

// Key is used to hold a Steam API key.
type Key string

// NewKey returns a type Key from the given string.
func NewKey(key string) Key {
	return Key(key)
}

func requestAPI(i, method string, version int, params url.Values) ([]byte, error) {
	requestURL := fmt.Sprintf("%v/%v/%v/v%v?", apiBaseURL, i, method, version)

	resp, err := http.Get(requestURL + params.Encode())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
