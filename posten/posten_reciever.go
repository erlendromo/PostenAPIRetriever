package posten

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/erlendromo/PostenAPIRetriever/utils"
)

func NewPostenResponse(ctx context.Context, postalcode string) (*PostenResponse, error) {
	subCtx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	req, err := http.NewRequestWithContext(subCtx, http.MethodGet, fmt.Sprintf("%s%s", utils.BASE_URL, postalcode), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Millisecond * 500,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var postenResponse PostenResponse
	err = json.NewDecoder(resp.Body).Decode(&postenResponse)
	return &postenResponse, err
}
