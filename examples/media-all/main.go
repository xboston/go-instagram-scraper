package main

import (
	"fmt"
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

	media2()

	log.Println("Закончили")
}

func media() {

	media, err := client.Media.GetAll("xboston")

	if err != nil {
		log.Fatal(err)
	}

	n := 0
	for _, item := range media.Items {
		n = n + 1

		img, _ := instagram.NewImage(item.Images.StandardResolution.URL)

		log.Println(n, ":", img)
	}

	log.Println("media.count", len(media.Items))
}

func media2() {

	ch := make(chan *instagram.Media)

	go client.Media.GetAll2("100tskaya", ch)

	go func(c chan *instagram.Media) {
		for {
			media := <-c

			log.Println("get", media.MoreAvailable)

			// n := 0
			// for _, item := range media.Items {
			// 	n = n + 1

			// 	img, _ := instagram.NewImage(item.Images.StandardResolution.URL)
			// 	log.Println(n, ":", img)
			// }
		}
	}(ch)

	var input string
	fmt.Scanln(&input)
}