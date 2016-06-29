package instagram

import (
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
		URL: url,
	}

	return img, nil
}

type Image struct {
	URL *url.URL
}

func (i *Image) Clean() string {

	// delete ?ig_cache_key=*
	i.URL.RawQuery = ""

	return i.URL.String()
}

func (i *Image) Original() string {

	pathItems := strings.Split(i.URL.Path, "/")
	fileName := pathItems[len(pathItems)-1:][0]
	i.URL.Path = strings.Join([]string{"t", fileName}, "/")

	return i.URL.String()
}

func (i *Image) Thumbnail() string {

	pathItems := strings.Split(i.URL.Path, "/")
	fileName := pathItems[len(pathItems)-1:][0]
	i.URL.Path = strings.Join([]string{"t", "s320x320", fileName}, "/")

	return i.URL.String()
}

func (i *Image) Size(w, h uint) string {

	pathItems := strings.Split(i.URL.Path, "/")
	fileName := pathItems[len(pathItems)-1:][0]
	size := fmt.Sprintf("s%dx%d", w, h)

	i.URL.Path = strings.Join([]string{"t", size, fileName}, "/")

	return i.URL.String()
}
