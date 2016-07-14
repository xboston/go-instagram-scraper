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

	getUsernameByID()

	log.Println("Закончили")

}

func getUsernameByID() {

	checkUserID := uint(27105452)

	username, err := client.Query.GetUsernameByID(checkUserID)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(checkUserID, "=>", username)
}
