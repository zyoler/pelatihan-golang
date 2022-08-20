package tools

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func Curl(method, baseUrl string, requestBody []byte, data interface{}, header map[string][]string) error {

	timeout := time.Duration(20 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}
	request, err := http.NewRequest(method, baseUrl, bytes.NewBuffer(requestBody))
	request.Header = header
	if err != nil {
		log.Println("Error Create Object HTTP Request ", err)
		return err
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Error Hit Do Request ", err)
		return err
	}
	// defer response.Body.Close()
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		log.Println("Error Json Decode TO Data Interface ", err)
	}
	return nil
}
