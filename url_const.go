package steam_integration

import "net/url"

// scheme
const (
	httpsScheme = "https"
)

// hosts
const (
	SteamPoweredHost   = "steampowered.com"
	StoreHost          = "store." + SteamPoweredHost
	ApiHost            = "api." + SteamPoweredHost
	SteamCommunityHost = "steamcommunity.com"
)

// paths
const (
	// Steam Web API paths
	iSteamAppsInterfacePath   = "/ISteamApps"
	getAppsList               = iSteamAppsInterfacePath + "/GetAppList"
	getAppsListV2             = getAppsList + "/v2"
	iSteamNewsInterfacePath   = "/ISteamNews"
	getNewsForApp             = iSteamNewsInterfacePath + "/GetNewsForApp"
	getNewsForAppV2           = getNewsForApp + "/v2"
	getAppReviewsPathTemplate = "/appreviews/{appId}?json=1"

	// Steam Website paths
	appPath = "/app"
)

type SteamUrlFunc func(appId uint32) *url.URL
