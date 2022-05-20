package steam_integration

const (
	httpsScheme = "https"
)

// hosts
const (
	SteamHost    = "steampowered.com"
	ApiSteamHost = "api." + SteamHost
)

// paths
const (
	iSteamAppsPath = "/ISteamApps"
	getAppsListV2  = iSteamAppsPath + "/GetAppList/v2"
)
