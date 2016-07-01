package instagram

import (
	"fmt"
	"log"
)

// MediaService - сервис работы с медиа-данными
type MediaService struct {
	client *Client
}

func (s *MediaService) GetByLoginAndMaxId(userLogin, maxID string) (media *Media, err error) {

	var u string

	if maxID != "" {
		u = fmt.Sprintf("/%s/media?max_id=%s", userLogin, maxID)
	} else {
		u = fmt.Sprintf("/%s/media", userLogin)
	}

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &media)
	return media, err
}

// Get - получение медиа-элементов пользователя
func (s *MediaService) Get(userLogin string) (media *Media, err error) {

	media, err = s.GetByLoginAndMaxId(userLogin, "")
	return
}

// GetAll - получение полного списка медиа-элементов пользователя
func (s *MediaService) GetAll(userLogin string) (media *Media, err error) {

	media, err = s.GetByLoginAndMaxId(userLogin, "")
	moreAvailable := media.MoreAvailable

	for moreAvailable {

		maxID := media.Items[len(media.Items)-1].ID

		moreMedia, err := s.GetByLoginAndMaxId(userLogin, maxID)

		if err != nil {
			return media, err
		}

		media.Items = append(media.Items, moreMedia.Items...)

		moreAvailable = moreMedia.MoreAvailable
		maxID = moreMedia.Items[len(moreMedia.Items)-1].ID
	}

	return
}

// GetAllWithCallback - получение полного списка медиа-элементов пользователя с передачей его в пользовательскую функцию
func (s *MediaService) GetAllWithCallback(userLogin string, m func(*Media)) {

	ch := make(chan *Media)

	go func(s *MediaService, userLogin string, ch chan *Media) {
		var (
			media, moreMedia *Media
			err              error
		)

		media, err = s.GetByLoginAndMaxId(userLogin, "")
		moreAvailable := media.MoreAvailable

		ch <- media

		maxID := media.Items[len(media.Items)-1].ID

		for moreAvailable {

			moreMedia, err = s.GetByLoginAndMaxId(userLogin, maxID)

			if err != nil {
				log.Fatal(err.Error())
			}

			ch <- moreMedia

			moreAvailable = moreMedia.MoreAvailable
			maxID = moreMedia.Items[len(moreMedia.Items)-1].ID
		}
	}(s, userLogin, ch)

	go func(c chan *Media) {
		for {
			media := <-c
			m(media)
		}
	}(ch)
}

// Media - инфомрация о медаи-данных пользователя
type Media struct {
	Items []struct {
		AltMediaURL       interface{} `json:"alt_media_url"`
		CanDeleteComments bool        `json:"can_delete_comments"`
		CanViewComments   bool        `json:"can_view_comments"`
		Caption           struct {
			CreatedTime string `json:"created_time"`
			From        struct {
				FullName       string `json:"full_name"`
				ID             string `json:"id"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"from"`
			ID   string `json:"id"`
			Text string `json:"text"`
		} `json:"caption"`
		Code     string `json:"code"`
		Comments struct {
			Count int           `json:"count"`
			Data  []interface{} `json:"data"`
		} `json:"comments"`
		CreatedTime string `json:"created_time"`
		ID          string `json:"id"`
		Images      struct {
			LowResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_resolution"`
			StandardResolution struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"standard_resolution"`
			Thumbnail struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"thumbnail"`
		} `json:"images"`
		Likes struct {
			Count int `json:"count"`
			Data  []struct {
				FullName       string `json:"full_name"`
				ID             string `json:"id"`
				ProfilePicture string `json:"profile_picture"`
				Username       string `json:"username"`
			} `json:"data"`
		} `json:"likes"`
		Link     string `json:"link"`
		Location struct {
			Name string `json:"name"`
		} `json:"location"`
		Type string `json:"type"`
		User struct {
			FullName       string `json:"full_name"`
			ID             string `json:"id"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"user"`
		UserHasLiked bool `json:"user_has_liked"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	Status        string `json:"status"`
}
