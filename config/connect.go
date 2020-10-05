package config

import (
	"bytes"
	"github.com/fj-onathan/blotion/vars"
	"io/ioutil"
	"log"
	"net/http"
)

// 🧲 Connect to notion post API for reversing
func notionAPI(paramsData []byte, api string) []byte {
	uri := vars.NotionHTTP + api
	body := bytes.NewBuffer(paramsData)
	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	req.Header.Set("Content-Type", vars.ContentType)
	req.Header.Set("User-Agent", vars.UserAgent)
	req.Header.Set("Accept-Language", vars.AcceptLang)

	var rsp *http.Response
	httpClient := *http.DefaultClient
	rsp, err = httpClient.Do(req)

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return data
}
