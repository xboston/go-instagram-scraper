package main

import (
	"log"

	instagram "github.com/xboston/go-instagram/instagram"
)

func main() {

	log.Println("Начали")

	user()
	media()
	tag()

	log.Println("Закончили")

}

func user() {

	client := instagram.NewClient(nil)
	user, err := client.Users.Get("xboston")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user.User.Username, user.User.FullName, user.User.FollowedBy.Count)
}

func media() {

	client := instagram.NewClient(nil)
	media, err := client.Media.Get("xboston")

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range media.Items {

		log.Println(item.Caption.Text, item.Images.StandardResolution.URL)
	}
}

func tag() {

	client := instagram.NewClient(nil)
	tagMedia, err := client.Tag.Search("imiss")

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range tagMedia.Tag.Media.Nodes {

		log.Println(item.Caption, item.DisplaySrc)
	}
}
