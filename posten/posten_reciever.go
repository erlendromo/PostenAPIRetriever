package posten

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/erlendromo/PostenAPIRetriever/utils"
)

// Returns the json-response from Posten-API (complete response if 'complete' tag is not nil)
func NewPostenResponse(ctx context.Context, postalcode string, complete ...bool) (DataResponse, error) {
	subCtx, cancel := context.WithCancel(ctx)
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
	if err = json.NewDecoder(resp.Body).Decode(&postenResponse); err != nil {
		return nil, err
	}

	if postenResponse.Metadata.TotalHits == 0 {
		return nil, fmt.Errorf("no addresses registered")
	}

	if complete != nil {
		return postenResponse.completeData()
	}

	return postenResponse.extractedData()
}
