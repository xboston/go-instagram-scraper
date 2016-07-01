package instagram

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "https://api.github.com/"
	baseURL        = "https://www.instagram.com"
	// Version - версия
	libraryVersion = "1.0.0"
	// UserAgent - строка UserAgent
	userAgent = "github.com/xboston/go-instagram " + libraryVersion
)

// Client - объект клиента
type Client struct {
	client *http.Client

	UserAgent string
	BaseURL   *url.URL

	Users *UsersService
	Media *MediaService
	Tag   *TagService
}

// NewClient returns a new Instagram API client.  If a nil httpClient is
// provided, http.DefaultClient will be used.
func NewClient(httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	baseURLObj, _ := url.Parse(baseURL)

	c := &Client{
		client:    httpClient,
		BaseURL:   baseURLObj,
		UserAgent: userAgent,
	}

	c.Users = &UsersService{client: c}
	c.Media = &MediaService{client: c}
	c.Tag = &TagService{client: c}

	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash. If
// specified
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	return req, nil
}

// Do sends an API request and returns the API response. The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if v != nil {
		err = json.NewDecoder(resp.Body).Decode(v)
	}
	return resp, err
}

// // ErrorResponse represents a Response which contains an error
// type ErrorResponse Response

// func getAccountPageLink(userName string) string {
// 	return strings.Replace("{username}", userName, accountPage, 1)
// }

// func getAccountJSONLink(username string) string {
// 	return strings.Replace("{username}", username, accountJSONInfo, 1)
// }

// func getAccountMediasJSONLink(userName, maxID string) string {
// 	url := strings.Replace(accountMedias, "{username}", userName, 1)
// 	return strings.Replace(url, "{max_id}", maxID, 1)
// }

// func getMediaPageLink(code string) string {
// 	return strings.Replace("{code}", code, mediaLink, 1)
// }

// func getMediaJSONLink(code string) string {
// 	return strings.Replace(mediaJSONInfo, "{code}", code, 1)
// }
