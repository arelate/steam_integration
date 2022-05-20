package steam_integration

import (
	"github.com/arelate/gog_integration"
	"net/url"
)

func SteamCommunityUrl(appId string) *url.URL {
	return &url.URL{
		Scheme: gog_integration.HttpsScheme,
		Host:   SteamCommunityHost,
		Path:   appPath + "/" + appId,
	}
}
