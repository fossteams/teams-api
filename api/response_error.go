package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func InvalidResponseError(resp *http.Response) error {
	bodyString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid status code %d", resp.StatusCode)
	}
	return fmt.Errorf("invalid status code %d: resp = %s", resp.StatusCode, string(bodyString))
}
