package csa

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func (c *CSASvc) getJSON(resp *http.Response) (io.Reader, error) {
	var jsonBuffer io.Reader

	if c.debugSave {
		// Temporary save response
		f, err := ioutil.TempFile(os.TempDir(), "teams-*.json")
		if err != nil {
			return nil, fmt.Errorf("unable to create temporary file")
		}
		jsonBytes, err := ioutil.ReadAll(resp.Body)
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