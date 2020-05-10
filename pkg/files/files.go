package files

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/hsmtkk/virustotal_go/pkg"
	"github.com/hsmtkk/virustotal_go/pkg/files/parser"
)

type UploadAnalyzer interface {
	UploadAnalyze(fileName string, contents []byte) (string, error)
}

func New(apiKey string) UploadAnalyzer {
	endPoint := "https://www.virustotal.com/api/v3/files"
	return &uploadAnalyzerImpl{client: http.DefaultClient, endPoint: endPoint, apiKey: apiKey}
}

func NewForTest(client *http.Client, endPoint string) UploadAnalyzer {
	return &uploadAnalyzerImpl{client: client, endPoint: endPoint, apiKey: "dummy"}
}

type uploadAnalyzerImpl struct {
	client   *http.Client
	endPoint string
	apiKey   string
}

func (imp *uploadAnalyzerImpl) UploadAnalyze(fileName string, contents []byte) (string, error) {
	formData, contentType, err := makeFormData(fileName, contents)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, imp.endPoint, bytes.NewReader(formData))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set(pkg.XAPIKey, imp.apiKey)
	resp, err := imp.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	ps := parser.New()
	id, err := ps.ParseResponse(string(bs))
	if err != nil {
		return "", err
	}
	return id, nil
}

func makeFormData(fileName string, contents []byte) ([]byte, string, error) {
	var buf bytes.Buffer
	mimeWriter := multipart.NewWriter(&buf)
	partWriter, err := mimeWriter.CreateFormFile("file", fileName)
	n, err := partWriter.Write(contents)
	if err != nil {
		return nil, "", err
	}
	if n < len(contents) {
		return nil, "", fmt.Errorf("failed to write all contents")
	}
	mimeWriter.Close()
	return buf.Bytes(), mimeWriter.FormDataContentType(), nil
}
