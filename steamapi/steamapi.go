package steamapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"errors"
)

const apiBaseURL = "https://api.steampowered.com"

var invalidKeyErrorMessage = "<html><head><title>Forbidden</title></head><body><h1>Forbidden</h1>Access is denied. Retrying will not help. Please verify your <pre>key=</pre> parameter.</body></html>"
var invalidKeyError = errors.New("Invalid API key")

// Key is used to hold a Steam API key.
type Key string

// NewKey returns a type Key from the given string.
func NewKey(key string) Key {
	return Key(key)
}

func requestAPI(i, method string, version int, params url.Values, respData interface{}) error {
	resp, err := http.Get(fmt.Sprintf("%s/%s/%s/v%d?%s", apiBaseURL, i, method, version, params.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, respData)
	if err != nil {
		if string(body) == invalidKeyErrorMessage {
			return invalidKeyError
		}
		return err
	}

	return nil
}
