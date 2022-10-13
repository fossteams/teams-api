package util

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func GetJSON(resp *http.Response, debugSave bool) (io.Reader, error) {
	var jsonBuffer io.Reader

	if debugSave {
		// Temporary save response
		f, err := ioutil.TempFile(os.TempDir(), "teams-*.json")
		if err != nil {
			return nil, fmt.Errorf("unable to create temporary file")
		}
		jsonBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, fmt.Errorf("unable to read response body: %v", err)
		}

		_, _ = f.Write(jsonBytes)
		fmt.Printf("saved temporary json to %v\n", f.Name())
		jsonBuffer = bytes.NewReader(jsonBytes)
	} else {
		jsonBuffer = resp.Body
	}
	return jsonBuffer, nil
}
