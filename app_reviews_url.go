package steam_integration

import (
	"net/url"
	"strconv"
)

func AppReviewsUrl(appId uint32) *url.URL {
	//https://partner.steamgames.com/doc/store/getreviews
	return &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   getAppReviewsPath + strconv.Itoa(int(appId)),
	}
}
