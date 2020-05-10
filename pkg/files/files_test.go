package files_test

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"testing"

	"github.com/hsmtkk/virustotal_go/pkg/apikey"
	"github.com/hsmtkk/virustotal_go/pkg/files"
	"github.com/stretchr/testify/assert"
)

func TestReal(t *testing.T) {
	apiKey, err := apikey.New().LoadAPIKey()
	assert.Nil(t, err, "should be nil")
	api := files.New(apiKey)
	id, err := api.UploadAnalyze(readFile())
	assert.Nil(t, err, "should be nil")
	assert.NotEmpty(t, id, "should not be empty")
}

func TestLocal(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bs, err := httputil.DumpRequest(r, true)
		assert.Nil(t, err, "should be nil")
		log.Println(string(bs))
		src, err := os.Open("example_response.json")
		assert.Nil(t, err, "should be nil")
		io.Copy(w, src)
	}))
	defer ts.Close()

	api := files.NewForTest(ts.Client(), ts.URL)
	_, err := api.UploadAnalyze(readFile())
	assert.Nil(t, err, "should be nil")
}

func readFile() []byte {
	bs, _ := ioutil.ReadFile("./example_response.json")
	return bs
}
