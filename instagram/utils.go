package instagram

import (
	"net/url"
	"strings"
)

// GetProfileByUrl - получение чистого навания профиля из произвольного URL
func GetProfileByUrl(profileURL string) (profileName *string, err error) {

	u, err := url.Parse(profileURL)
	if err != nil {
		return nil, err
	}

	profileName = &strings.Split(strings.Trim(u.Path, "/"), "/")[0]

	return
}
