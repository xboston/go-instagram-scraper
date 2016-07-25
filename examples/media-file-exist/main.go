package main

import (
	"log"

	instagram "github.com/xboston/go-instagram/instagram"
)

var (
	client *instagram.Client
)

func init() {
	client = instagram.NewClient(nil)
}

func main() {

	log.Println("Начали")

	file := "https://scontent-amt2-1.cdninstagram.com/t51.2885-15/s750x750/sh0.08/e35/123_321_000_n1.jpg?ig_cache_key=MTI5NjY1MzA2ODkxMzYyNDMxNw%3D%3D.2"
	res, err := client.Media.FileExist(file)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(res, file)

	file = "https://scontent-amt2-1.cdninstagram.com/t51.2885-19/s150x150/12965122_1192953944077961_1039555908_a.jpg"

	res, err = client.Media.FileExist(file)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(res, file)

	log.Println("Закончили")
}
