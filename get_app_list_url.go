package steam_integration

import (
	"github.com/arelate/gog_integration"
	"net/url"
)

// https://partner.steamgames.com/doc/webapi/ISteamApps#GetAppList
func AppListUrl() *url.URL {
	return &url.URL{
		Scheme: gog_integration.HttpsScheme,
		Host:   ApiSteamHost,
		Path:   getAppsListV2,
	}
}

func VangoghAppListUrl(_ string, _ gog_integration.Media) *url.URL {
	return AppListUrl()
}
