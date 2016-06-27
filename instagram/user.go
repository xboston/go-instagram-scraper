package instagram

import "fmt"

// UsersService - сервис работы с юзером
type UsersService struct {
	client *Client
}

// Get - получение информации о юзере
func (s *UsersService) Get(userLogin string) (*User, error) {

	u := fmt.Sprintf("/%s/?__a=1", userLogin)

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	user := new(User)
	_, err = s.client.Do(req, user)
	return user, err
}

// User - информация о пользователе
type User struct {
	User struct {
		Biography              interface{} `json:"biography"`
		BlockedByViewer        bool        `json:"blocked_by_viewer"`
		CountryBlock           interface{} `json:"country_block"`
		ExternalURL            string      `json:"external_url"`
		ExternalURLLinkshimmed string      `json:"external_url_linkshimmed"`
		FollowedBy             struct {
			Count int `json:"count"`
		} `json:"followed_by"`
		FollowedByViewer bool `json:"followed_by_viewer"`
		Follows          struct {
			Count int `json:"count"`
		} `json:"follows"`
		FollowsViewer      bool   `json:"follows_viewer"`
		FullName           string `json:"full_name"`
		HasBlockedViewer   bool   `json:"has_blocked_viewer"`
		HasRequestedViewer bool   `json:"has_requested_viewer"`
		ID                 string `json:"id"`
		IsPrivate          bool   `json:"is_private"`
		IsVerified         bool   `json:"is_verified"`
		Media              struct {
			Count int `json:"count"`
			Nodes []struct {
				Caption  string `json:"caption"`
				Code     string `json:"code"`
				Comments struct {
					Count int `json:"count"`
				} `json:"comments"`
				Date       int `json:"date"`
				Dimensions struct {
					Height int `json:"height"`
					Width  int `json:"width"`
				} `json:"dimensions"`
				DisplaySrc string `json:"display_src"`
				ID         string `json:"id"`
				IsVideo    bool   `json:"is_video"`
				Likes      struct {
					Count int `json:"count"`
				} `json:"likes"`
				Owner struct {
					ID string `json:"id"`
				} `json:"owner"`
				ThumbnailSrc string `json:"thumbnail_src"`
			} `json:"nodes"`
			PageInfo struct {
				EndCursor       string `json:"end_cursor"`
				HasNextPage     bool   `json:"has_next_page"`
				HasPreviousPage bool   `json:"has_previous_page"`
				StartCursor     string `json:"start_cursor"`
			} `json:"page_info"`
		} `json:"media"`
		ProfilePicURL     string `json:"profile_pic_url"`
		ProfilePicURLHd   string `json:"profile_pic_url_hd"`
		RequestedByViewer bool   `json:"requested_by_viewer"`
		Username          string `json:"username"`
	} `json:"user"`
}
