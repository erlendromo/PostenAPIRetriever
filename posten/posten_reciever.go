package posten

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/erlendromo/PostenAPIRetriever/utils"
)

// Returns the complete json-response from Posten-API
func NewPostenResponse(ctx context.Context, postalcode string) (*PostenResponse, error) {
	subCtx, cancel := context.WithCancel(ctx)
	defer cancel()

	req, err := http.NewRequestWithContext(subCtx, http.MethodGet, fmt.Sprintf("%s%s", utils.BASE_URL, postalcode), nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var postenResponse PostenResponse
	if err = json.NewDecoder(resp.Body).Decode(&postenResponse); err != nil {
		return nil, err
	}

	if postenResponse.Metadata.TotalHits == 0 {
		return nil, fmt.Errorf("no addresses registered")
	}

	return &postenResponse, err
}
