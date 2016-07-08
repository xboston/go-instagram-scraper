package main

import (
	"log"

	instagram "github.com/xboston/go-instagram/instagram"
)

func main() {

	imgURL := "https://scontent-ams3-1.cdninstagram.com/t51.2885-15/e35/13151005_514712292048857_1289305583_n.jpg?ig_cache_key=MTI1Mzc2Mjg0MzAwNzE5NDk4MA%3D%3D.2"
	imgURL = "https://scontent-ams3-1.cdninstagram.com/t51.2885-15/e35/13557105_695516713922456_777411316_n.jpg?ig_cache_key=MTI4OTgzNzg2ODA2NzQyNjE4Ng%3D%3D.2"
	imgURL = "https://scontent-ams3-1.cdninstagram.com/t51.2885-15/s150x150/e35/c135.0.810.810/12424604_182025218829349_1270015183_n.jpg?ig_cache_key=MTE3NDgwMzQ4NDkxOTY2NTgzMQ%3D%3D.2.c"

	img, err := instagram.NewImageFromThumbnail(imgURL)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(img.Clean())
	//log.Println(img.Original())
	log.Println(img.Thumbnail())
	log.Println(img.Standart())
	log.Println(img.Size(150, 150))
	log.Println(img.Size(1080, 1080))
	log.Println(img.ThumbnailSquare())
	log.Println(img.StandartSquare())
}
