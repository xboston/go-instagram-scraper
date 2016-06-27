package instagram

import (
	"fmt"
	"log"
)

const (
	//tagSearch = "https://www.instagram.com/explore/tags/%s/?__a=1&max_id={max_id}"
	tagSearch = "https://www.instagram.com/explore/tags/%s/?__a=1"
)

// TagService - сервис работы с тегами
type TagService struct {
	client *Client
}

// Search - поиск по тегу
func (s *TagService) Search(tag string) (*Tag, error) {
	u := fmt.Sprintf(tagSearch, tag)

	log.Println(u)

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	tagObj := new(Tag)
	_, err = s.client.Do(req, tagObj)
	return tagObj, err
}

// Tag - медиа-данные по тегу
type Tag struct {
	Tag struct {
		Media struct {
			Count    int `json:"count"`
			PageInfo struct {
				HasPreviousPage bool   `json:"has_previous_page"`
				StartCursor     string `json:"start_cursor"`
				EndCursor       string `json:"end_cursor"`
				HasNextPage     bool   `json:"has_next_page"`
			} `json:"page_info"`
			Nodes []struct {
				Code       string `json:"code"`
				Dimensions struct {
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"dimensions"`
				Owner struct {
					ID string `json:"id"`
				} `json:"owner"`
				Comments struct {
					Count int `json:"count"`
				} `json:"comments"`
				Caption string `json:"caption"`
				Likes   struct {
					Count int `json:"count"`
				} `json:"likes"`
				Date         int    `json:"date"`
				ThumbnailSrc string `json:"thumbnail_src"`
				IsVideo      bool   `json:"is_video"`
				ID           string `json:"id"`
				DisplaySrc   string `json:"display_src"`
			} `json:"nodes"`
		} `json:"media"`
		ContentAdvisory interface{} `json:"content_advisory"`
		TopPosts        struct {
			Nodes []struct {
				Code       string `json:"code"`
				Dimensions struct {
					Width  int `json:"width"`
					Height int `json:"height"`
				} `json:"dimensions"`
				Owner struct {
					ID string `json:"id"`
				} `json:"owner"`
				Comments struct {
					Count int `json:"count"`
				} `json:"comments"`
				Caption string `json:"caption"`
				Likes   struct {
					Count int `json:"count"`
				} `json:"likes"`
				Date         int    `json:"date"`
				ThumbnailSrc string `json:"thumbnail_src"`
				IsVideo      bool   `json:"is_video"`
				ID           string `json:"id"`
				DisplaySrc   string `json:"display_src"`
				VideoViews   int    `json:"video_views,omitempty"`
			} `json:"nodes"`
		} `json:"top_posts"`
		Name string `json:"name"`
	} `json:"tag"`
}
