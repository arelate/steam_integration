package steam_integration

// scheme
const (
	httpsScheme = "https"
)

// hosts
const (
	SteamHost          = "steampowered.com"
	ApiSteamHost       = "api." + SteamHost
	SteamCommunityHost = "steamcommunity.com"
)

// paths
const (
	iSteamAppsPath = "/ISteamApps"
	getAppsListV2  = iSteamAppsPath + "/GetAppList/v2"
	appPath        = "/app"
)
