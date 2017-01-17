package steamid

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

// ResolveID tries to resolve a SteamID 64 from a search query.
// It checks for queries such as vanity url's or SteamID's.
// If an invalid apiKey is used, a SteamID 64 will still be if
// the query is not a vanity url.
// If no SteamID 64 could be resolved then 0 is returned.
func ResolveID(query, apiKey string) ID64 {
	query = strings.Replace(query, " ", "", -1)
	query = strings.Trim(query, "/")

	if strings.Contains(query, "steamcommunity.com/profiles/") {
		id64, err := strconv.ParseInt(query[strings.Index(query, "steamcommunity.com/profiles/")+len("steamcommunity.com/profiles/"):], 10, 64)
		if err != nil {
			goto id
		}

		if len(strconv.FormatInt(id64, 10)) != 17 {
			goto id
		}

		return ID64(id64)
	}

id:
	if regexp.MustCompile(`^STEAM_0:(0|1):[0-9]{1}[0-9]{0,8}$`).MatchString(query) {
		id64 := ID(query).To64()

		if len(strconv.FormatUint(uint64(id64), 10)) != 17 {
			goto id64
		}

		return id64
	}

id64:
	if regexp.MustCompile(`^\d{17}$`).MatchString(query) {
		id64, err := strconv.ParseInt(query, 10, 64)
		if err != nil {
			goto id32
		}

		return ID64(id64)
	}

id32:
	if regexp.MustCompile(`^\d{9}$`).MatchString(query) {
		id64, err := strconv.ParseInt(query, 10, 64)
		if err != nil {
			goto id3
		}

		return ID64(id64)
	}

id3:
	if regexp.MustCompile(`(\[)?U:1:\d+(\])?`).MatchString(strings.ToUpper(query)) {
		return ID3(query).To64()
	}

	if strings.Contains(query, "steamcommunity.com/id/") {
		query = query[strings.Index(query, "steamcommunity.com/id/")+len("steamcommunity.com/id/"):]
	}

	urlData := url.Values{
		"key":       {apiKey},
		"vanityurl": {query},
	}
	resp, err := http.Get("http://api.steampowered.com/ISteamUser/ResolveVanityURL/v0001/?" + urlData.Encode())
	if err != nil {
		return 0
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0
	}

	var vanityURLJSON struct {
		Response struct {
			Steamid string
			Success int
		}
	}

	err = json.Unmarshal(body, &vanityURLJSON)
	if err != nil {
		return 0
	}

	if vanityURLJSON.Response.Success != 1 {
		return 0
	}

	if len(vanityURLJSON.Response.Steamid) != 17 {
		return 0
	}

	id64, err := strconv.ParseInt(vanityURLJSON.Response.Steamid, 10, 64)
	if err != nil {
		return 0
	}

	return ID64(id64)
}
