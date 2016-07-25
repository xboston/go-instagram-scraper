package instagram

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
)

func NewImage(imageURL string) (img *Image, err error) {

	url, err := url.Parse(imageURL)

	if err != nil {
		return img, err
	}

	img = &Image{
		URL: *url,
	}

	img.Clean()

	return img, nil
}

func NewImageFromThumbnail(imageURL string) (img *Image, err error) {

	if !strings.Contains(imageURL, "s150x150") {
		return img, errors.New("No thumbnail url")
	}

	url, err := url.Parse(imageURL)

	if err != nil {
		return img, err
	}

	img = &Image{
		URL: *url,
	}

	img.Clean()

	return img, nil
}

type Image struct {
	URL url.URL
}

func (i Image) String() string {

	return i.Original()
}

// delete ?ig_cache_key=*
func (i *Image) Clean() string {

	i.URL.RawQuery = ""

	return i.URL.String()
}

func (i Image) Original() string {

	pathItems := strings.Split(i.URL.Path, "/")
	fileName := pathItems[len(pathItems)-1:][0]

	i.URL.Path = strings.Join([]string{"t", fileName}, "/")

	return i.URL.String()
}

func (i Image) Thumbnail() string {

	return i.Size(320, 320)
}

func (i Image) Standart() string {

	return i.Size(640, 640)
}

func (i Image) Size(w, h uint) string {

	pathItems := strings.Split(i.URL.Path, "/")
	fileName := pathItems[len(pathItems)-1:][0]
	size := fmt.Sprintf("s%dx%d", w, h)

	i.URL.Path = strings.Join([]string{"t", size, fileName}, "/")

	return i.URL.String()
}
func (i Image) ThumbnailSquare() string {

	return i.ThumbnailSquareSize(320, 320)
}

func (i Image) StandartSquare() string {

	return i.ThumbnailSquareSize(640, 640)
}

func (i Image) ThumbnailSquareSize(w, h uint) string {

	size := fmt.Sprintf("s%dx%d", w, h)
	i.URL.Path = strings.Replace(i.URL.Path, "s150x150", size, -1)

	return i.URL.String()
}
