package steam_integration

import "net/url"

// https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppList
func AppListUrl() *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   ApiSteamHost,
		Path:   getAppsListV2,
	}
}
