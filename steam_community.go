package steam_integration

import (
	"net/url"
)

func SteamCommunityUrl(appId string) *url.URL {
	return &url.URL{
		Scheme: httpsScheme,
		Host:   SteamCommunityHost,
		Path:   appPath + "/" + appId,
	}
}
