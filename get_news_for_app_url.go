package steam_integration

import (
	"net/url"
	"strconv"
)

//https://partner.steamgames.com/doc/webapi/ISteamNews#GetNewsForApp
func NewsForApp(appId uint32) *url.URL {
	u := &url.URL{
		Scheme: httpsScheme,
		Host:   ApiSteamHost,
		Path:   getNewsForAppV2,
	}

	q := u.Query()
	q.Add("appid", strconv.Itoa(int(appId)))
	u.RawQuery = q.Encode()

	return u
}
