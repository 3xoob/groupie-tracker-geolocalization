package groupieGeo

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		errStr := fmt.Sprintf("failed to fetch data: %v", err.Error())
		log.Println(errStr)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		errStr := fmt.Sprintf("received non-OK HTTP status: %d", response.StatusCode)
		log.Println(errStr)
		return nil, fmt.Errorf(errStr)
	}
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		errStr := fmt.Sprintf("failed to read response body: %v", err.Error())
		log.Println(errStr)
		return nil, err
	}
	return dataBytes, nil
}
