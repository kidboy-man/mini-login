package utils

import (
	"auth-service/datatransfers"
	ctx "context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/bitly/go-simplejson"
)

var (
	once   sync.Once
	client *http.Client
)

func init() {
	once.Do(func() {
		client = &http.Client{
			Timeout: 60 * time.Second,
		}
	})
}

func PostJSONRequest(url, payload string) (result []byte, err error) {
	req, _ := http.NewRequestWithContext(ctx.Background(), http.MethodPost, url, strings.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	response, err := client.Do(req)
	fmt.Println("error response:", err)
	if err != nil {
		return
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		resJSON, err := simplejson.NewFromReader(response.Body)
		if err != nil {
			return nil, err
		}

		err = parseError(resJSON, response.StatusCode)
		return nil, err
	}

	result, err = parseBody(response.Body)
	return
}

func parseBody(body io.ReadCloser) (result []byte, err error) {

	resJSON, err := simplejson.NewFromReader(body)
	if err != nil {
		return
	}

	if !resJSON.GetPath("success").MustBool() {
		err = parseError(resJSON, http.StatusInternalServerError)
		return
	}

	JSONdata, err := resJSON.GetPath("data").MarshalJSON()
	if err != nil {
		return
	}

	result = JSONdata
	return
}

func parseError(resJSON *simplejson.Json, httpCode int) (err error) {

	JSONErr, err := resJSON.GetPath("error").MarshalJSON()
	if err != nil {
		return
	}

	var errResponse *datatransfers.CustomError
	err = json.Unmarshal(JSONErr, &errResponse)
	if err != nil {
		return
	}

	err = errResponse
	return
}
