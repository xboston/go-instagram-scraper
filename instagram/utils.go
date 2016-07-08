package instagram

import (
	"errors"
	"net/url"
	"strings"
)

// GetProfileByURL - получение чистого навания профиля из произвольного URL
func GetProfileByURL(profileURL string) (profileName *string, err error) {

	if strings.Index(profileURL, "//") == 0 {
		profileURL = "https:" + profileURL
	}
	if strings.Index(profileURL, "://") == -1 {
		profileURL = "https://" + profileURL
	}

	u, err := url.Parse(profileURL)
	if err != nil {
		return nil, err
	}

	if "instagram.com" != u.Host && "www.instagram.com" != u.Host {
		return nil, errors.New("Host error")
	}

	var profile string

	profile = strings.Split(strings.Trim(u.Path, "/"), "/")[0]
	profile = strings.TrimSpace(profile)

	if profile == "" {
		return nil, errors.New("Empty profile")

	}

	return &profile, nil
}
