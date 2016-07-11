package instagram

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
)

// MediaService - сервис работы с медиа-данными
type MediaService struct {
	client *Client
}

// GetByLoginAndMaxID - получение медиа-элементов пользователя по логину и max_id
func (s *MediaService) GetByLoginAndMaxID(userLogin, maxID string) (media *Media, err error) {

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

	resp, err := s.client.Do(req, &media)

	if http.StatusOK != resp.StatusCode {
		return nil, errors.New("Media not found")
	}

	return media, err
}

// Get - получение медиа-элементов пользователя
func (s *MediaService) Get(userLogin string) (media *Media, err error) {

	media, err = s.GetByLoginAndMaxID(userLogin, "")
	return
}

// GetAll - получение полного списка медиа-элементов пользователя
func (s *MediaService) GetAll(userLogin string) (media *Media, err error) {

	media, err = s.GetByLoginAndMaxID(userLogin, "")
	moreAvailable := media.MoreAvailable

	for moreAvailable {

		maxID := media.Items[len(media.Items)-1].ID

		moreMedia, err := s.GetByLoginAndMaxID(userLogin, maxID)

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
func (s *MediaService) GetAllWithCallback(userLogin string, mediaFunc func(*Media)) {

	var wg sync.WaitGroup
	wg.Add(1)

	channelMedia := make(chan *Media)
	defer close(channelMedia)

	go func() {

		defer wg.Done()

		var (
			media, moreMedia *Media
			err              error
			moreAvailable    bool
		)

		media, err = s.GetByLoginAndMaxID(userLogin, "")
		if err != nil {
			log.Fatal(err.Error())
		}

		wg.Add(1)
		channelMedia <- media

		moreAvailable = media.MoreAvailable
		maxID := media.Items[len(media.Items)-1].ID

		for moreAvailable {

			moreMedia, err = s.GetByLoginAndMaxID(userLogin, maxID)

			if err != nil {
				log.Fatal(err.Error())
			}

			wg.Add(1)
			channelMedia <- moreMedia

			moreAvailable = moreMedia.MoreAvailable
			maxID = moreMedia.Items[len(moreMedia.Items)-1].ID
		}

	}()

	go func() {
		for {
			media := <-channelMedia
			mediaFunc(media)
			wg.Done()
		}
	}()

	wg.Wait()
}

// Exist - проверка существования медиа
func (s *MediaService) Exist(mediaID string) (bool, error) {

	mediaID = strings.TrimSpace(mediaID)

	if mediaID == "" {
		return false, errors.New("mediaID empty")
	}

	u := fmt.Sprintf("/p/%s/?__a=1", mediaID)

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return false, err
	}

	resp, err := s.client.Do(req, nil)

	log.Println(resp.Request.URL, resp.StatusCode)

	if http.StatusOK != resp.StatusCode {
		return false, errors.New("Media not exist")
	}

	return true, nil
}

// Media - инфомрация о медаи-данных пользователя
type Media struct {
	Items []struct {
		AltMediaURL       string `json:"alt_media_url"`
		CanDeleteComments bool   `json:"can_delete_comments"`
		CanViewComments   bool   `json:"can_view_comments"`
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
			Count int `json:"count"`
			Data  []struct {
				CreatedTime string `json:"created_time"`
				From        struct {
					FullName       string `json:"full_name"`
					ID             string `json:"id"`
					ProfilePicture string `json:"profile_picture"`
					Username       string `json:"username"`
				} `json:"from"`
				ID   string `json:"id"`
				Text string `json:"text"`
			} `json:"data"`
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
		Link     string      `json:"link"`
		Location interface{} `json:"location"`
		Type     string      `json:"type"`
		User     struct {
			FullName       string `json:"full_name"`
			ID             string `json:"id"`
			ProfilePicture string `json:"profile_picture"`
			Username       string `json:"username"`
		} `json:"user"`
		UserHasLiked bool `json:"user_has_liked"`
		VideoViews   int  `json:"video_views"`
		Videos       struct {
			LowBandwidth struct {
				Height int    `json:"height"`
				URL    string `json:"url"`
				Width  int    `json:"width"`
			} `json:"low_bandwidth"`
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
		} `json:"videos"`
	} `json:"items"`
	MoreAvailable bool   `json:"more_available"`
	Status        string `json:"status"`
}
