package steamapi

import "net/url"

import "strconv"
import "encoding/json"
import "errors"

// PlayerSummaries contains basic profile information for a Steam user.
type PlayerSummaries struct {
	Avatar                   string `json:"avatar"`
	AvatarFull               string `json:"avatarfull"`
	AvatarMedium             string `json:"avatarmedium"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	LastLogOff               int    `json:"lastlogoff"`
	LocCountryCode           string `json:"loccountrycode"`
	PersonaName              string `json:"personaname"`
	PersonaState             int    `json:"personastate"`
	PersonaStateFlags        int    `json:"personastateflags"`
	PrimaryClanID            string `json:"primaryclanid"`
	ProfileState             int    `json:"profilestate"`
	ProfileURL               string `json:"profileurl"`
	RealName                 string `json:"realname"`
	SteamID                  string `json:"steamid"`
	TimeCreated              int    `json:"timecreated"`
}

// GetPlayerSummaries returns basic profile information for a list of 64-bit Steam IDs.
func (k Key) GetPlayerSummaries(id64s ...uint64) ([]PlayerSummaries, error) {
	var (
		err         error
		params      = url.Values{}
		steamIDList string
	)

	for i, id64 := range id64s {
		steamIDList += strconv.FormatUint(id64, 10)
		if i < len(id64s)-1 {
			steamIDList += ","
		}
	}

	params.Add("key", string(k))
	params.Add("steamids", steamIDList)

	body, err := requestAPI("ISteamUser", "GetPlayerSummaries", 2, params)
	if err != nil {
		return []PlayerSummaries{}, err
	}

	var respData struct {
		Response struct {
			Players []PlayerSummaries `json:"players"`
		} `json:"response"`
	}

	err = json.Unmarshal(body, &respData)
	if err != nil {
		return []PlayerSummaries{}, err
	}

	return respData.Response.Players, nil
}

// GetSinglePlayerSummaries returns basic profile information for a single 64-bit Steam ID.
func (k Key) GetSinglePlayerSummaries(id64 uint64) (PlayerSummaries, error) {
	summaries, err := k.GetPlayerSummaries(id64)
	if err != nil {
		return PlayerSummaries{}, err
	}

	if len(summaries) == 0 {
		return PlayerSummaries{}, errors.New("No player summaries were responded")
	}

	return summaries[0], nil
}
