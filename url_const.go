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
	// Steam Web API paths
	iSteamAppsInterfacePath = "/ISteamApps"
	getAppsList             = iSteamAppsInterfacePath + "/GetAppList"
	getAppsListV2           = getAppsList + "/v2"
	iSteamNewsInterfacePath = "/ISteamNews"
	getNewsForApp           = iSteamNewsInterfacePath + "/GetNewsForApp"
	getNewsForAppV2         = getNewsForApp + "/v2"

	// Steam Web site paths
	appPath = "/app"
)
