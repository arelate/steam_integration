package steam_integration

import (
	"net/url"
	"strconv"
	"strings"
)

func AppReviewsUrl(appId uint32) *url.URL {
	//https://partner.steamgames.com/doc/store/getreviews
	return &url.URL{
		Scheme: httpsScheme,
		Host:   StoreHost,
		Path:   strings.Replace(getAppReviewsPathTemplate, "{appId}", strconv.Itoa(int(appId)), -1),
	}
}
