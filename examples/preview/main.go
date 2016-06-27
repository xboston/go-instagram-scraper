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

	user()
	media()
	tag()

	log.Println("Закончили")

}

func user() {

	user, err := client.Users.Get("xboston")

	if err != nil {
		log.Fatal(err)
	}

	log.Println(user.User.Username, user.User.FullName, user.User.FollowedBy.Count)
}

func media() {

	media, err := client.Media.Get("xboston")

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range media.Items {

		log.Println(item.Caption.Text, item.Images.StandardResolution.URL)
	}
}

func tag() {

	tagMedia, err := client.Tag.Search("boobs")

	if err != nil {
		log.Fatal(err)
	}

	for _, item := range tagMedia.Tag.Media.Nodes {

		log.Println("Media", item.DisplaySrc)
	}

	for _, item := range tagMedia.Tag.TopPosts.Nodes {

		log.Println("TopPosts", item.DisplaySrc)
	}
}
