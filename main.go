package main

import (
	"context"
	"fmt"
	"log"

	"github.com/erlendromo/PostenAPIRetriever/posten"
)

func main() {
	postenData, err := posten.NewPostenResponse(context.Background(), "2372")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", postenData)
}
