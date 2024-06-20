package main

import (
	"context"
	"fmt"
	"log"

	"github.com/erlendromo/PostenAPIRetriever/posten"
)

func main() {
	postenData, err := posten.NewPostenResponse(context.Background(), "0010")
	if err != nil {
		log.Fatal(err)
	}

	data, err := postenData.ExtractValuableData()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)
}
