package f

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

type reqHeaders struct {
	Key   string
	Value string
}

func doRequest(method string, url string, body *[]byte, headers []reqHeaders) ([]byte, error) {
	var (
		req *http.Request
		err error
	)
	if body != nil {
		r := bytes.NewReader(*body)
		req, err = http.NewRequest(method, url, r)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
	}

	for _, rh := range headers {
		req.Header.Set(rh.Key, rh.Value)
	}

	req.Header.Set("content-type", "application/json")

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("err resp.Body.Close(). err: %+v, resp: %+v", err, resp)
		}
	}()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, err
}
