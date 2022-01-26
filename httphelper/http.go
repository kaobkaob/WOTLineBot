package httphelper

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Do(inMethod, inURL string, inParam url.Values, inHeaders map[string]string) ([]byte, error) {

	var err error

	var body = strings.NewReader(inParam.Encode())

	if inMethod == "GET" {
		inURL += inParam.Encode()
	}

	req, _ := http.NewRequest(inMethod, inURL, body)

	for k, v := range inHeaders {
		req.Header.Set(k, v)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer resp.Body.Close()

	readbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}

	return readbody, err
}
