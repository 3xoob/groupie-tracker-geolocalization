package groupieGeo

import (
	"encoding/json"
	"io"
	"net/http"
)

func ExtractToStruct(link string, v interface{}) error {
	resp, err := http.Get(link)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bodyBytes, v)
}
