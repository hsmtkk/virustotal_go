package files

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type UploadAnalyzer interface {
	UploadAnalyze(contents []byte) (string, error)
}

const endPoint = "https://www.virustotal.com/api/v3/files"

func New(apiKey string) UploadAnalyzer {
	return &uploadAnalyzerImpl{client: http.DefaultClient, endPoint: endPoint, apiKey: apiKey}
}

func NewForTest(client *http.Client, endPoint string) UploadAnalyzer {
	return &uploadAnalyzerImpl{client: client, endPoint: endPoint}
}

type uploadAnalyzerImpl struct {
	client   *http.Client
	endPoint string
	apiKey   string
}

const applicationJson = "application/json"

// TODO mutlipart/form-data
// https://ayada.dev/posts/multipart-requests-in-go/

func (self *uploadAnalyzerImpl) UploadAnalyze(contents []byte) (string, error) {
	req, err := http.NewRequest(http.MethodPost, self.endPoint, bytes.NewReader(contents))
	if err != nil {
		return "", err
	}
	req.Header.Set(xApiKey, self.apiKey)
	resp, err := self.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}
