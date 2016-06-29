package main

import (
	"log"

	instagram "github.com/xboston/go-instagram/instagram"
)

func main() {

	imgURL := "https://scontent-ams3-1.cdninstagram.com/t51.2885-15/e35/13151005_514712292048857_1289305583_n.jpg?ig_cache_key=MTI1Mzc2Mjg0MzAwNzE5NDk4MA%3D%3D.2"

	img, err := instagram.NewImage(imgURL)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(img.Clean())
	log.Println(img.Original())
	log.Println(img.Thumbnail())
	log.Println(img.Size(640, 640))
}
