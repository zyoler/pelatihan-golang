package tools

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func HTTPNotif(method string, url string, jsondatastring string, arrheader map[string]string) (int, string, error) {

	var finalresult string

	var ioreader io.Reader
	if jsondatastring != "" {
		ioreader = bytes.NewBuffer([]byte(jsondatastring))
	}

	req, err := http.NewRequest(method, url, ioreader)
	if err != nil {
		log.Println("Error, NewRequest: " + err.Error())
	}

	if len(arrheader) > 0 {
		for k, v := range arrheader {
			req.Header.Set(k, v)
		}
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error, client.Do: " + err.Error())
		finalresult = ""
		return 0, finalresult, err
	}
	defer resp.Body.Close()

	respbody, _ := ioutil.ReadAll(resp.Body)
	finalresult = string(respbody)
	return resp.StatusCode, finalresult, nil
}

func Curl(method, baseUrl string, requestBody []byte, data interface{}, header map[string][]string) error {

	timeout := time.Duration(20 * time.Second)

	client := http.Client{
		Timeout: timeout,
	}
	log.Println("Isi url", baseUrl)
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
	defer response.Body.Close()
	log.Println("Response body", response.Body)
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		log.Println("Error Json Decode TO Data Interface ", err)
	}
	return nil
}
