package instagram

import "fmt"

// MediaService - сервис работы с медиа-данными
type MediaService struct {
	client *Client
}

// Get - полученепи медиа-данных пользователя
func (s *MediaService) Get(userLogin string) (*Media, error) {

	u := fmt.Sprintf("/%s/media", userLogin) // ?max_id={max_id}

	req, err := s.client.NewRequest("GET", u, "")
	if err != nil {
		return nil, err
	}

	media := new(Media)
	_, err = s.client.Do(req, media)
	return media, err
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

func (m *Media) getImages() {
	//  $instance->imageHighResolutionUrl = str_replace('320x320', '1080x1080', $instance->imageLowResolutionUrl);
	//     if (isset($mediaArray['caption'])) {
	//         $instance->caption = $mediaArray['caption']['text'];
	//     }
	// if ($instance->type === 'video') {

	// }
}

func (m *Media) getVideo() {
	// $instance->imageHighResolutionUrl = str_replace('320x320', '1080x1080', $instance->imageLowResolutionUrl);
	//     if (isset($mediaArray['caption'])) {
	//         $instance->caption = $mediaArray['caption']['text'];
	//     }
	// if ($instance->type !== 'video') {
	//     $instance->videoLowResolutionUrl = $mediaArray['videos']['low_resolution']['url'];
	//     $instance->videoStandardResolutionUrl = $mediaArray['videos']['standard_resolution']['url'];
	//     $instance->videoLowBandwidthUrl = $mediaArray['videos']['low_bandwidth']['url'];
	// }
}

// return strpos($imageUrl, '?ig_cache_key=') ? substr($imageUrl, 0, strpos($imageUrl, '?ig_cache_key=')) : $imageUrl;
// Максимально доступное качество фото можно получить, удалив все кроме /t/ из ссылки — https://scontent.cdninstagram.com/t/12950481_1753078061593396_874826488_n.jpg, через параметры в ссылке можно управлять размером и кропом.
// или вот так: https://instagram.com/p/9BDXa_L7bm/media/?size=l
