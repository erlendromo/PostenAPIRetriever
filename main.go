package main

import (
	"context"
	"fmt"

	"github.com/erlendromo/PostenAPIRetriever/posten"
)

func main() {
	postenData, _ := posten.NewPostenResponse(context.Background(), "2822")

	data, _ := postenData.ExtractValuableData()

	fmt.Printf("%+v\n", data)
}
